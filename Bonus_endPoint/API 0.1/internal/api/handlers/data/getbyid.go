package data

import (
	"context"
	"encoding/json"
	service "goapi/internal/api/service/data"
	"log"
	"net/http"
	"time"
)

// * The GET method retrieves a resource identified by a URI *
// * curl -X GET http://127.0.0.1:8080/data/veh1 -i -u admin123:Testing@123 -H "Content-Type: application/json"
func GetByIDHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.DataService) {

	vehicalID := r.PathValue("vehical_id")
	if vehicalID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Missing vehical_id."}`))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// You may want to implement a new service method for this, e.g., ReadByVehicalID
	dataList, err := ds.ReadByVehicalID(vehicalID, ctx)
	if err != nil {
		logger.Println("Could not read by vehical_id:", err, vehicalID)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}
	if dataList == nil || len(dataList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Resource not found."}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dataList); err != nil {
		logger.Println("Error encoding data:", err, dataList)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}
}
