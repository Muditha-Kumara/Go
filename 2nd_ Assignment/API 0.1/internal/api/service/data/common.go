package data

import (
	"context"
	"goapi/internal/api/repository/models"
)
type DataService interface {
	Create(data *models.Data, ctx context.Context) error
	ReadOne(deviceID string, vehicalID string, ctx context.Context) (*models.Data, error)
	ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Data, error)
	ReadByVehicalID(vehicalID string, ctx context.Context) ([]*models.Data, error)
	Update(data *models.Data, ctx context.Context) (int64, error)
	Delete(deviceID string, vehicalID string, ctx context.Context) (int64, error)
	ValidateData(data *models.Data) error
	GetLatestByDeviceID(deviceID string, ctx context.Context) (*models.Data, error)
}

type DataError struct {
	Message string
}

func (de DataError) Error() string {
	return de.Message
}
