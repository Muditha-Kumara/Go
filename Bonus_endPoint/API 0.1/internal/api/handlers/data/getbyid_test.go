package data_test

import (
	"context"
	"encoding/json"
	"goapi/internal/api/handlers/data"
	service "goapi/internal/api/service/data"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetByIDInvalidID(t *testing.T) {
	mockDataService := &service.MockDataServiceSuccessful{}
	req, err := http.NewRequest("GET", "/data/invalid", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("vehical_id", "") // Simulate missing vehical_id
	rr := httptest.NewRecorder()

	data.GetByIDHandler(rr, req, log.Default(), mockDataService)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

       if strings.TrimSpace(rr.Body.String()) != `{"error": "Missing vehical_id."}` {
	       t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), `{"error": "Missing vehical_id."}`)
       }
}

func TestGetByIdInternalError(t *testing.T) {
	mockDataService := &service.MockDataServiceError{}
	req, err := http.NewRequest("GET", "/data/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("vehical_id", "1")

	rr := httptest.NewRecorder()

	data.GetByIDHandler(rr, req, log.Default(), mockDataService)
       if status := rr.Code; status != http.StatusInternalServerError {
	       t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
       }
       if strings.TrimSpace(rr.Body.String()) != `Internal server error.` {
	       t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), `Internal server error.`)
       }
}

func TestGetByIdNotFound(t *testing.T) {
	mockDataService := &service.MockDataServiceNotFound{}
	req, err := http.NewRequest("GET", "/data/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("vehical_id", "1")

	rr := httptest.NewRecorder()

	data.GetByIDHandler(rr, req, log.Default(), mockDataService)
       if status := rr.Code; status != http.StatusNotFound {
	       t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
       }
       if strings.TrimSpace(rr.Body.String()) != `{"error": "Resource not found."}` {
	       t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), `{"error": "Resource not found."}`)
       }
}

func TestGetByIdSuccessful(t *testing.T) {
	mockDataService := &service.MockDataServiceSuccessful{}
	req, err := http.NewRequest("GET", "/data/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("vehical_id", "1")

	rr := httptest.NewRecorder()

	data.GetByIDHandler(rr, req, log.Default(), mockDataService)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

       data, _ := mockDataService.ReadByVehicalID("1", context.Background())
       expected, _ := json.Marshal(data)

       if strings.TrimSpace(rr.Body.String()) != string(expected) {
	       t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
       }
}
