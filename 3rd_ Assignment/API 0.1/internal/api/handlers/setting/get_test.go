package setting

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSetting(t *testing.T) {
	req := httptest.NewRequest("GET", "/setting", nil)
	rw := httptest.NewRecorder()

	GetSetting(rw, req)

	resp := rw.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}

	var data map[string]string
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	expected := map[string]string{
		"Critical Level": "2.0",
		"Alert Level": "1.0",
	}

	for k, v := range expected {
		if data[k] != v {
			t.Errorf("expected %s to be %s, got %s", k, v, data[k])
		}
	}
}
