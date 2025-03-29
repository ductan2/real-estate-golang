package initialize

import (
	"bytes"
	"encoding/json"
	"ecommerce/global"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Custom hook để gửi log lên Elasticsearch
type ElasticHook struct {
	ElasticsearchURL string
	IndexName        string
	Username         string
	Password         string
}

// Gửi log đến Elasticsearch
func (hook *ElasticHook) Fire(entry *logrus.Entry) error {
	logData, err := json.Marshal(entry.Data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/_doc", hook.ElasticsearchURL, hook.IndexName), bytes.NewBuffer(logData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Nếu có username/password, thêm Basic Auth
	if hook.Username != "" && hook.Password != "" {
		req.SetBasicAuth(hook.Username, hook.Password)
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("failed to send log to Elasticsearch, status code: %d", resp.StatusCode)
	}
	return nil
}

// Bắt buộc phải có để logrus nhận Hook
func (hook *ElasticHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func InitELK() {
	if !global.Config.ELK.Enabled {
		return
	}

	// Khởi tạo Elasticsearch client
	cfg := elasticsearch.Config{
		Addresses: []string{global.Config.ELK.ElasticsearchURL},
		Username:  global.Config.ELK.Username,
		Password:  global.Config.ELK.Password,
	}
	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	// Kiểm tra kết nối Elasticsearch
	res, err := esClient.Info()
	if err != nil {
		log.Fatalf("Error getting Elasticsearch info: %s", err)
	}
	defer res.Body.Close()
	fmt.Println("Connected to Elasticsearch:", res)

	// Kiểm tra và tạo index nếu chưa tồn tại
	indexName := global.Config.ELK.IndexName
	exists, err := esClient.Indices.Exists([]string{indexName})
	if err != nil {
		log.Fatalf("Error checking if index exists: %s", err)
	}
	defer exists.Body.Close()

	if exists.StatusCode == 404 {
		// Tạo index với mapping
		createIndex, err := esClient.Indices.Create(
			indexName,
			esClient.Indices.Create.WithBody(strings.NewReader(`{
				"mappings": {
					"properties": {
						"timestamp": { "type": "date" },
						"level": { "type": "keyword" },
						"message": { "type": "text" },
						"service": { "type": "keyword" },
						"trace_id": { "type": "keyword" }
					}
				}
			}`)),
		)
		if err != nil {
			log.Fatalf("Error creating index: %s", err)
		}
		defer createIndex.Body.Close()
		fmt.Println("Index created successfully:", indexName)
	}

	// Cấu hình Logrus
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	// Ghi log vào file
	logger.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    100, // MB
		MaxBackups: 3,
		MaxAge:     28,   // Ngày
		Compress:   true, // Nén file log
	})

	// Cấu hình mức log
	level, err := logrus.ParseLevel("info")
	if err != nil {
		log.Fatalf("Error parsing log level: %s", err)
	}
	logger.SetLevel(level)

	// Thêm hook gửi log đến Elasticsearch
	elasticHook := &ElasticHook{
		ElasticsearchURL: global.Config.ELK.ElasticsearchURL,
		IndexName:        global.Config.ELK.IndexName,
		Username:         global.Config.ELK.Username,
		Password:         global.Config.ELK.Password,
	}
	logger.AddHook(elasticHook)

	global.Logger = logger
	fmt.Println("ELK logging initialized successfully")
}
