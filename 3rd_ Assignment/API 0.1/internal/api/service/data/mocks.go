package data

import (
	"context"
	"goapi/internal/api/repository/models"
)

// * Mock implementation of DataService for testing purposes, always returns a successful response and Data object(s) *
type MockDataServiceSuccessful struct{}
// GetLatestByDeviceID returns mock latest data for a given deviceID (simulate found data)
func (m *MockDataServiceSuccessful) GetLatestByDeviceID(deviceID string, ctx context.Context) (*models.Data, error) {
       return &models.Data{
	       DeviceID:   deviceID,
	       VehicalID:  "vehical1",
	       Data:       "latest alert data",
	       AlertType:  "type1",
	       DateTime:   "2021-01-01T00:01:00Z",
	       Location:   "location1",
       }, nil
}
// ReadByVehicalID returns mock data for a given vehicalID (simulate found data)
func (m *MockDataServiceSuccessful) ReadByVehicalID(vehicalID string, ctx context.Context) ([]*models.Data, error) {
       return []*models.Data{
	       {
		       DeviceID:   "device1",
		       VehicalID:  vehicalID,
		       Data:       "alert data 1",
		       AlertType:  "type1",
		       DateTime:   "2021-01-01T00:00:00Z",
		       Location:   "location1",
	       },
       }, nil
}

func (m *MockDataServiceSuccessful) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Data, error) {
       return []*models.Data{
	       {
		       DeviceID:   "device1",
		       VehicalID:  "vehical1",
		       Data:       "alert data 1",
		       AlertType:  "type1",
		       DateTime:   "2021-01-01T00:00:00Z",
		       Location:   "location1",
	       },
	       {
		       DeviceID:   "device2",
		       VehicalID:  "vehical2",
		       Data:       "alert data 2",
		       AlertType:  "type2",
		       DateTime:   "2021-01-01T00:00:00Z",
		       Location:   "location2",
	       },
       }, nil
}

func (m *MockDataServiceSuccessful) ReadOne(deviceID string, vehicalID string, ctx context.Context) (*models.Data, error) {
       return &models.Data{
	       DeviceID:   deviceID,
	       VehicalID:  vehicalID,
	       Data:       "alert data",
	       AlertType:  "type1",
	       DateTime:   "2021-01-01T00:00:00Z",
	       Location:   "location1",
       }, nil
}

func (m *MockDataServiceSuccessful) Create(data *models.Data, ctx context.Context) error {
	return nil
}

func (m *MockDataServiceSuccessful) Update(data *models.Data, ctx context.Context) (int64, error) {
	return 1, nil
}

func (m *MockDataServiceSuccessful) Delete(deviceID string, vehicalID string, ctx context.Context) (int64, error) {
	return 1, nil
}

func (m *MockDataServiceSuccessful) ValidateData(data *models.Data) error {
	return nil
}

// * Mock implementation of DataService for testing purposes, always returns empty data *

type MockDataServiceNotFound struct{}
// GetLatestByDeviceID returns nil (simulate not found)
func (m *MockDataServiceNotFound) GetLatestByDeviceID(deviceID string, ctx context.Context) (*models.Data, error) {
	return nil, nil
}
// ReadByVehicalID returns empty slice (simulate not found)
func (m *MockDataServiceNotFound) ReadByVehicalID(vehicalID string, ctx context.Context) ([]*models.Data, error) {
	return []*models.Data{}, nil
}

func (m *MockDataServiceNotFound) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Data, error) {
	return []*models.Data{}, nil
}

func (m *MockDataServiceNotFound) ReadOne(deviceID string, vehicalID string, ctx context.Context) (*models.Data, error) {
	return nil, nil
}

func (m *MockDataServiceNotFound) Create(data *models.Data, ctx context.Context) error {
	return nil
}

func (m *MockDataServiceNotFound) Update(data *models.Data, ctx context.Context) (int64, error) {
	return 0, nil
}

func (m *MockDataServiceNotFound) Delete(deviceID string, vehicalID string, ctx context.Context) (int64, error) {
	return 0, nil
}

func (m *MockDataServiceNotFound) ValidateData(data *models.Data) error {
	return nil
}

// * Mock implementation of DataService for testing purposes, always returns an error *
type MockDataServiceError struct{}
// GetLatestByDeviceID returns an error (simulate DB/service error)
func (m *MockDataServiceError) GetLatestByDeviceID(deviceID string, ctx context.Context) (*models.Data, error) {
	return nil, DataError{Message: "Error reading latest data."}
}
// ReadByVehicalID returns an error (simulate DB/service error)
func (m *MockDataServiceError) ReadByVehicalID(vehicalID string, ctx context.Context) ([]*models.Data, error) {
	return nil, DataError{Message: "Error reading data."}
}

func (m *MockDataServiceError) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Data, error) {
	return nil, DataError{Message: "Error reading data."}
}

func (m *MockDataServiceError) ReadOne(deviceID string, vehicalID string, ctx context.Context) (*models.Data, error) {
	return nil, DataError{Message: "Error reading data."}
}

func (m *MockDataServiceError) Create(data *models.Data, ctx context.Context) error {
	return DataError{Message: "Error creating data."}
}

func (m *MockDataServiceError) Update(data *models.Data, ctx context.Context) (int64, error) {
	return 0, DataError{Message: "Error updating data."}
}

func (m *MockDataServiceError) Delete(deviceID string, vehicalID string, ctx context.Context) (int64, error) {
	return 0, DataError{Message: "Error deleting data."}
}

func (m *MockDataServiceError) ValidateData(data *models.Data) error {
	return nil
}
