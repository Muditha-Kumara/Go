package setting

import (
	"encoding/json"
	"net/http"
)

// * The GET method retrieves the settings for critical and alert levels *
// curl -s -u admin123:Testing@123 -H "Content-Type: application/json" http://127.0.0.1:8080/setting
func GetSetting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"Critical Level": "2.0",
		"Alert Level": "1.0",
	}
	json.NewEncoder(w).Encode(response)
}
