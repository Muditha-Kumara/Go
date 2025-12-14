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
	// result variable removed (no longer needed)
	if err == nil && lastData != nil {
		// Parse date_time fields
		newTime, err1 := time.Parse(time.RFC3339, data.DateTime)
		lastTime, err2 := time.Parse(time.RFC3339, lastData.DateTime)
	       // The time gap check is still performed, but the result is not returned in the response anymore
	       // (If you want to log or use this info internally, you can do so here)
	       // if newTime.Sub(lastTime) <= time.Minute && newTime.Sub(lastTime) >= 0 {
	       //     // Do something if needed
	       // }
	       // No assignment to 'result' needed
	       if err1 == nil && err2 == nil {
		       _ = (newTime.Sub(lastTime) <= time.Minute && newTime.Sub(lastTime) >= 0) // evaluated but not used
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

	// Respond with the full data object as JSON (to match test expectation)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
