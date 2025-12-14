package SQLite

import (
	"context"
	"database/sql"
	"goapi/internal/api/repository/DAL"
	"goapi/internal/api/repository/models"
)

type DataRepository struct {
	sqlDB *sql.DB
	createStmt      *sql.Stmt
	readByVehicalStmt *sql.Stmt
	readStmt        *sql.Stmt
	readManyStmt    *sql.Stmt
	updateStmt      *sql.Stmt
	deleteStmt      *sql.Stmt
	ctx             context.Context
}

// GetLatestByDeviceID returns the most recent data record for a given device_id
func (r *DataRepository) GetLatestByDeviceID(deviceID string, ctx context.Context) (*models.Data, error) {
	row := r.sqlDB.QueryRowContext(ctx, `SELECT device_id, vehical_id, data, alert_type, date_time, location FROM data WHERE device_id = ? ORDER BY date_time DESC LIMIT 1`, deviceID)
	var d models.Data
	err := row.Scan(&d.DeviceID, &d.VehicalID, &d.Data, &d.AlertType, &d.DateTime, &d.Location)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}

func NewDataRepository(sqlDB DAL.SQLDatabase, ctx context.Context) (models.DataRepository, error) {

	repo := &DataRepository{
		sqlDB: sqlDB.Connection(),
		ctx:   ctx,
	}

	   // Create the data table if it doesn't exist
	   if _, err := repo.sqlDB.Exec(`CREATE TABLE  IF NOT EXISTS data (
		   device_id VARCHAR(50) NOT NULL,
		   vehical_id VARCHAR(50),
		   data TEXT,
		   alert_type VARCHAR(50),
		   date_time TIMESTAMP,
		   location VARCHAR(100)
	   );`); err != nil {
		   repo.sqlDB.Close()
		   return nil, err
	   }

	// * Create needed Prepared SQL statements, this is more efficient than running each query individually
		createStmt, err := repo.sqlDB.Prepare(`INSERT INTO data (device_id, vehical_id, data, alert_type, date_time, location) VALUES (?, ?, ?, ?, ?, ?)`)
	if err != nil {
		repo.sqlDB.Close() // Close the database connection if statement preparation fails
		return nil, err
	}
	repo.createStmt = createStmt

		readStmt, err := repo.sqlDB.Prepare("SELECT device_id, vehical_id, data, alert_type, date_time, location FROM data WHERE device_id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readStmt = readStmt

		readManyStmt, err := repo.sqlDB.Prepare("SELECT device_id, vehical_id, data, alert_type, date_time, location FROM data LIMIT ? OFFSET ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readManyStmt = readManyStmt

		updateStmt, err := repo.sqlDB.Prepare("UPDATE data SET vehical_id = ?, data = ?, alert_type = ?, date_time = ?, location = ? WHERE device_id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.updateStmt = updateStmt

		deleteStmt, err := repo.sqlDB.Prepare("DELETE FROM data WHERE device_id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.deleteStmt = deleteStmt

		// prepare statement to read by vehical_id
		readByVehicalStmt, err := repo.sqlDB.Prepare("SELECT device_id, vehical_id, data, alert_type, date_time, location FROM data WHERE vehical_id = ?")
		if err != nil {
			repo.sqlDB.Close()
			return nil, err
		}
		repo.readByVehicalStmt = readByVehicalStmt

	go Close(ctx, repo)

	return repo, nil
}

func Close(ctx context.Context, r *DataRepository) {

	<-ctx.Done()
	r.createStmt.Close()
	r.readStmt.Close()
	r.readByVehicalStmt.Close()
	r.updateStmt.Close()
	r.deleteStmt.Close()
	r.readManyStmt.Close()
	r.sqlDB.Close()
}

func (r *DataRepository) Create(data *models.Data, ctx context.Context) error {

	_, err := r.createStmt.ExecContext(ctx, data.DeviceID, data.VehicalID, data.Data, data.AlertType, data.DateTime, data.Location)
	return err
}

func (r *DataRepository) ReadOne(id int, ctx context.Context) (*models.Data, error) {
	   row := r.readStmt.QueryRowContext(ctx, id)
	   var data models.Data
	   err := row.Scan(&data.DeviceID, &data.VehicalID, &data.Data, &data.AlertType, &data.DateTime, &data.Location)
	   if err != nil {
		   if err == sql.ErrNoRows {
			   return nil, nil
		   }
		   return nil, err
	   }
	   return &data, nil
}

func (r *DataRepository) ReadByVehicalID(vehicalID string, ctx context.Context) ([]*models.Data, error) {
	rows, err := r.readByVehicalStmt.QueryContext(ctx, vehicalID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []*models.Data
	for rows.Next() {
		var d models.Data
		err := rows.Scan(&d.DeviceID, &d.VehicalID, &d.Data, &d.AlertType, &d.DateTime, &d.Location)
		if err != nil {
			return nil, err
		}
		data = append(data, &d)
	}
	return data, nil
}

func (r *DataRepository) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Data, error) {

	if page < 1 {
		return r.ReadAll()
	}

	offset := rowsPerPage * (page - 1)
	rows, err := r.readManyStmt.QueryContext(ctx, rowsPerPage, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []*models.Data
	   for rows.Next() {
		   var d models.Data
		   err := rows.Scan(&d.DeviceID, &d.VehicalID, &d.Data, &d.AlertType, &d.DateTime, &d.Location)
		   if err != nil {
			   return nil, err
		   }
		   data = append(data, &d)
	   }
	return data, nil
}

func (r *DataRepository) ReadAll() ([]*models.Data, error) {
		rows, err := r.sqlDB.QueryContext(context.Background(), "SELECT device_id, vehical_id, data, alert_type, date_time, location FROM data")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []*models.Data
	   for rows.Next() {
		   var d models.Data
		   err := rows.Scan(&d.DeviceID, &d.VehicalID, &d.Data, &d.AlertType, &d.DateTime, &d.Location)
		   if err != nil {
			   return nil, err
		   }
		   data = append(data, &d)
	   }
	return data, nil
}

func (r *DataRepository) Update(data *models.Data, ctx context.Context) (int64, error) {
		res, err := r.updateStmt.ExecContext(ctx, data.VehicalID, data.Data, data.AlertType, data.DateTime, data.Location, data.DeviceID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (r *DataRepository) Delete(data *models.Data, ctx context.Context) (int64, error) {
		res, err := r.deleteStmt.ExecContext(ctx, data.DeviceID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
