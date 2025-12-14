package data

import (
	"context"
	"encoding/json"
	"goapi/internal/api/repository/models"
	service "goapi/internal/api/service/data"
	"log"
	"net/http"
	"time"
)

// * User sends a POST request to /data with a JSON payload in the request body *
// * curl -X POST http://127.0.0.1:8080/data -i -u admin123:Testing@123 -H "Content-Type: application/json" -d '{"device_id": "dev1", "vehical_id": "veh1", "data": "some data", "alert_type": "warning", "date_time": "2025-12-14T12:00:00Z", "location": "lab"}'
func PostHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.DataService) {
	var data models.Data

	// Decode the JSON payload from the request body into the data struct
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid request data. Please check your input."}`))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// Get the latest record for the same device_id
	var lastData *models.Data
	lastData, err := ds.GetLatestByDeviceID(data.DeviceID, ctx)
	result := "ok"
	if err == nil && lastData != nil {
		// Parse date_time fields
		newTime, err1 := time.Parse(time.RFC3339, data.DateTime)
		lastTime, err2 := time.Parse(time.RFC3339, lastData.DateTime)
		if err1 == nil && err2 == nil {
			if newTime.Sub(lastTime) <= time.Minute && newTime.Sub(lastTime) >= 0 {
				result = "wrong"
			}
		}
	}

	// Always insert the new data
	if err := ds.Create(&data, ctx); err != nil {
		switch err.(type) {
		case service.DataError:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "` + err.Error() + `"}`))
			return
		default:
			logger.Println("Error creating data:", err, data)
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
			return
		}
	}

	// Respond with result (ok or wrong)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}
