package data

import (
	"context"
	service "goapi/internal/api/service/data"
	"log"
	"net/http"
	"time"
)

// * The DELETE method removes a resource identified by a URI *
// * curl -X DELETE http://127.0.0.1:8080/data/1 -i -u admin123:Testing@123 -H "Content-Type: application/json"
func DeleteHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.DataService) {
       deviceID := r.PathValue("device_id")
       vehicalID := r.PathValue("vehical_id")
       if deviceID == "" || vehicalID == "" {
	       w.WriteHeader(http.StatusBadRequest)
	       w.Write([]byte(`{"error": "Missing device_id or vehical_id."}`))
	       return
       }

       ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
       defer cancel()

       aff, err := ds.Delete(deviceID, vehicalID, ctx)
       if err != nil {
	       logger.Println("Could not delete data:", err, deviceID, vehicalID)
	       http.Error(w, "Internal Server error", http.StatusInternalServerError)
	       return
       }

       if aff == 0 {
	       w.WriteHeader(http.StatusNotFound)
	       w.Write([]byte(`{"error": "Resource not found."}`))
	       return
       }

	// * This is a Success, response in JSON and with a 204 status code when data was successfully deleted
	w.WriteHeader(http.StatusNoContent)
}
 