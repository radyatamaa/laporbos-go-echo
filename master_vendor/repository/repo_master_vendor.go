package repository

import (
	"context"
	"database/sql"

	"github.com/master_vendor"

	"github.com/models"
	"github.com/sirupsen/logrus"
)

type MasterVendorRepository struct {
	Conn *sql.DB
}

func NewMasterVendorRepository(Conn *sql.DB) master_vendor.Repository {
	return &MasterVendorRepository{Conn: Conn}
}

func (f MasterVendorRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.MasterVendor, error) {
	rows, err := f.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*models.MasterVendor, 0)
	for rows.Next() {
		t := new(models.MasterVendor)
		err = rows.Scan(
			&t.Id,
			&t.CreatedBy,
			&t.CreatedDate,
			&t.ModifiedBy,
			&t.ModifiedDate,
			&t.DeletedBy,
			&t.DeletedDate,
			&t.IsDeleted,
			&t.IsActive,
			&t.KodeVendor,
			&t.NAME1 ,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *MasterVendorRepository) Fetch(ctx context.Context, limit, offset int) ([]*models.MasterVendor, error) {
	if limit != 0 {
		query := `SELECT * FROM master_vendors where is_deleted = 0 AND is_active = 1`

		//if search != ""{
		//	query = query + `AND (promo_name LIKE '%` + search + `%'` +
		//		`OR promo_desc LIKE '%` + search + `%' ` +
		//		`OR start_date LIKE '%` + search + `%' ` +
		//		`OR end_date LIKE '%` + search + `%' ` +
		//		`OR promo_code LIKE '%` + search + `%' ` +
		//		`OR max_usage LIKE '%` + search + `%' ` + `) `
		//}
		query = query + ` ORDER BY created_date desc LIMIT ? OFFSET ? `
		res, err := m.fetch(ctx, query, limit, offset)
		if err != nil {
			return nil, err
		}
		return res, err

	} else {
		query := `SELECT * FROM master_vendors where is_deleted = 0 AND is_active = 1`

		//if search != ""{
		//	query = query + `AND (promo_name LIKE '%` + search + `%'` +
		//		`OR promo_desc LIKE '%` + search + `%' ` +
		//		`OR start_date LIKE '%` + search + `%' ` +
		//		`OR end_date LIKE '%` + search + `%' ` +
		//		`OR promo_code LIKE '%` + search + `%' ` +
		//		`OR max_usage LIKE '%` + search + `%' ` + `) `
		//}
		query = query + ` ORDER BY created_date desc `
		res, err := m.fetch(ctx, query)
		if err != nil {
			return nil, err
		}
		return res, err
	}
}

func (m *MasterVendorRepository) GetCount(ctx context.Context) (int, error) {
	query := `SELECT count(*) AS count FROM master_vendors WHERE is_deleted = 0 and is_active = 1`

	rows, err := m.Conn.QueryContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	count, err := checkCount(rows)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return count, nil
}

func (m *MasterVendorRepository) Insert(ctx context.Context, a *models.MasterVendor) (*int, error) {
	query := `INSERT master_vendors SET created_by=? , created_date=? , modified_by=?, modified_date=? , 	
				deleted_by=? , deleted_date=? , is_deleted=? , is_active=? , kode_vendor=?, name1=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	res, err := stmt.ExecContext(ctx, a.CreatedBy, a.CreatedDate, nil, nil, nil, nil, 0, 1, a.KodeVendor, a.NAME1)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	a.Id = int(lastID)
	return &a.Id, nil
}

func checkCount(rows *sql.Rows) (count int, err error) {
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}
	return count, nil
}
