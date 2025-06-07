package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mssola/useragent"
)

func GetLocationFromIP(ip string) (string) {
    url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
    resp, err := http.Get(url)
    if err != nil {
        return ""
    }
    defer resp.Body.Close()

    var result struct {
        Status      string `json:"status"`
        Country     string `json:"country"`
        RegionName  string `json:"regionName"`
        City        string `json:"city"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return ""
    }

    if result.Status != "success" {
        return ""
    }

    location := fmt.Sprintf("%s, %s, %s", result.City, result.RegionName, result.Country)
    return location
}

func GetUserAgentDetails(userAgent string) (*useragent.UserAgent) {
	ua := useragent.New(userAgent)
	if ua == nil {
		return nil
	}
	return ua
}
