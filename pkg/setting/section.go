package setting

type Config struct {
	Database  Database  `mapstructure:"database"`
	Server    Server    `mapstructure:"server"`
	Redis     Redis     `mapstructure:"redis"`
	SMTP      SMTP      `mapstructure:"smtp"`
	JWT       JWT       `mapstructure:"jwt"`
	APMServer APMServer `mapstructure:"apm_server"`
	ELK       ELK       `mapstructure:"elk"`
	Logger    Logger    `mapstructure:"logger"`
}

type Server struct {
	Port        string `mapstructure:"port"`
	Mode        string `mapstructure:"mode"`
	FrontendUrl string `mapstructure:"frontend_url"`
}

type Database struct {
	Host                  string `mapstructure:"host"`
	Port                  int    `mapstructure:"port"`
	Username              string `mapstructure:"user"`
	Password              string `mapstructure:"password"`
	Dbname                string `mapstructure:"dbname"`
	MaxOpenConnections    int    `mapstructure:"maxOpenConnections"`
	MaxIdleConnections    int    `mapstructure:"maxIdleConnections"`
	MaxLifetimeConnection int    `mapstructure:"maxLifetimeConnection"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type SMTP struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type JWT struct {
	TokenSecret         string `mapstructure:"tokenSecret"`
	TokenExpirationTime string `mapstructure:"tokenExpirationTime"`
	TokenHoursToExpire  int    `mapstructure:"tokenHoursToExpire"`
}

type APMServer struct {
	ServerURL   string `mapstructure:"server_url"`
	ServiceName string `mapstructure:"service_name"`
	Environment string `mapstructure:"environment"`
	Enabled     bool   `mapstructure:"enabled"`
}

type ELK struct {
	ElasticsearchURL string `mapstructure:"elasticsearch_url"`
	LogstashURL      string `mapstructure:"logstash_url"`
	KibanaURL        string `mapstructure:"kibana_url"`
	IndexName        string `mapstructure:"index_name"`
	Enabled          bool   `mapstructure:"enabled"`
	Username         string `mapstructure:"username"`
	Password         string `mapstructure:"password"`
}

type Logger struct {
	Level       string `mapstructure:"level"`
	Format      string `mapstructure:"format"`
	FileLogName string `mapstructure:"file_log_name"`
}
