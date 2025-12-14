package models

import "context"

type Data struct {
	DeviceID    string  `json:"device_id"`
	VehicalID   string  `json:"vehical_id"`
	Data        string  `json:"data"`
	AlertType   string  `json:"alert_type"`
	DateTime    string  `json:"date_time"`
	Location    string  `json:"location"`
}

type DataRepository interface {
	Create(Data *Data, ctx context.Context) error
	ReadOne(id int, ctx context.Context) (*Data, error)
	ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*Data, error)
	ReadByVehicalID(vehicalID string, ctx context.Context) ([]*Data, error)
	Update(data *Data, ctx context.Context) (int64, error)
	Delete(data *Data, ctx context.Context) (int64, error)
}
