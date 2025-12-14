package data

import (
	"context"
	"goapi/internal/api/repository/models"
	"time"
)

// * Implementation of DataService for SQLite database *
type DataServiceSQLite struct {
	repo models.DataRepository
}

func NewDataServiceSQLite(repo models.DataRepository) *DataServiceSQLite {
	return &DataServiceSQLite{
		 repo: repo,
	}
}

func (ds *DataServiceSQLite) Create(data *models.Data, ctx context.Context) error {
	if err := ds.ValidateData(data); err != nil {
		 return DataError{Message: "Invalid data."}
	}
	return ds.repo.Create(data, ctx)
}

func (ds *DataServiceSQLite) ReadOne(deviceID string, vehicalID string, ctx context.Context) (*models.Data, error) {
	// Implement lookup by deviceID and vehicalID if needed, or adapt repo method
	// Placeholder: return nil, nil
	return nil, nil
}

func (ds *DataServiceSQLite) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Data, error) {
	return ds.repo.ReadMany(page, rowsPerPage, ctx)
}

func (ds *DataServiceSQLite) ReadByVehicalID(vehicalID string, ctx context.Context) ([]*models.Data, error) {
	return ds.repo.ReadByVehicalID(vehicalID, ctx)
}

func (ds *DataServiceSQLite) Update(data *models.Data, ctx context.Context) (int64, error) {

	if err := ds.ValidateData(data); err != nil {
		return 0, DataError{Message: "Invalid data: " + err.Error()}
	}
	return ds.repo.Update(data, ctx)
}

func (ds *DataServiceSQLite) Delete(deviceID string, vehicalID string, ctx context.Context) (int64, error) {
	// Implement delete by deviceID and vehicalID if needed, or adapt repo method
	// Placeholder: return 0, nil
	return 0, nil
}

func (ds *DataServiceSQLite) ValidateData(data *models.Data) error {
       var errMsg string
       if data.DeviceID == "" || len(data.DeviceID) > 50 {
	       errMsg += "DeviceID is required and must be less than 50 characters. "
       }
       if len(data.VehicalID) > 50 {
	       errMsg += "VehicalID must be less than 50 characters. "
       }
       if len(data.AlertType) > 50 {
	       errMsg += "AlertType must be less than 50 characters. "
       }
       if len(data.Location) > 100 {
	       errMsg += "Location must be less than 100 characters. "
       }
       _, err := time.Parse("2006-01-02T15:04:05Z", data.DateTime)
       if err != nil {
	       errMsg += "DateTime must be in the format: 2021-01-01T12:00:00Z. "
       }
       if errMsg != "" {
	       return DataError{Message: errMsg}
       }
       return nil
}
