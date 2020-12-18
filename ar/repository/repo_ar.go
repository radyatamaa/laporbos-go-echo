package repository

import (
	"context"
	"database/sql"

	"github.com/ar"

	"github.com/models"
	"github.com/sirupsen/logrus"
)

type ArRepository struct {
	Conn *sql.DB
}

func NewArRepository(Conn *sql.DB) ar.Repository {
	return &ArRepository{Conn: Conn}
}

func (f ArRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Ar, error) {
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

	result := make([]*models.Ar, 0)
	for rows.Next() {
		t := new(models.Ar)
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
			&t.DocNumber,
			&t.DocDate,
			&t.PostingDate,
			&t.TglJatuhTempo,
			&t.Customer,
			&t.DC,
			&t.Amount,
			&t.DocCurrency,
			&t.DocHeader,
			&t.Assignment,
			&t.SalesDocument,
			&t.BillingDocument,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *ArRepository) Fetch(ctx context.Context, limit, offset int) ([]*models.Ar, error) {
	if limit != 0 {
		query := `SELECT * FROM ars where is_deleted = 0 AND is_active = 1`

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
		query := `SELECT * FROM ars where is_deleted = 0 AND is_active = 1`

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

func (m *ArRepository) GetCount(ctx context.Context) (int, error) {
	query := `SELECT count(*) AS count FROM ars WHERE is_deleted = 0 and is_active = 1`

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

func (m *ArRepository) Insert(ctx context.Context, a *models.Ar) (*int, error) {
	query := `INSERT ars SET created_by=? , created_date=? , modified_by=?, modified_date=? , 	
				deleted_by=? , deleted_date=? , is_deleted=? , is_active=? , doc_number=?, doc_date=?, posting_date=?,
				tgl_jatuh_tempo=?, customer=?, dc=?, amount=?, doc_currency=?, doc_header=?,assignment=?, sales_document=?, 
				billing_document=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	res, err := stmt.ExecContext(ctx, a.CreatedBy, a.CreatedDate, nil, nil, nil, nil, 0, 1,
		a.DocNumber,
		a.DocDate,
		a.PostingDate,
		a.TglJatuhTempo,
		a.Customer,
		a.DC,
		a.Amount,
		a.DocCurrency,
		a.DocHeader,
		a.Assignment,
		a.SalesDocument,
		a.BillingDocument,
	)
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
