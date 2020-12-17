package repository

import (
	"context"
	"database/sql"

	"github.com/sales_order"

	"github.com/models"
	"github.com/sirupsen/logrus"
)

type SalesOrderRepository struct {
	Conn *sql.DB
}

func NewSalesOrderRepository(Conn *sql.DB) sales_order.Repository {
	return &SalesOrderRepository{Conn: Conn}
}

func (f SalesOrderRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.SalesOrder, error) {
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

	result := make([]*models.SalesOrder, 0)
	for rows.Next() {
		t := new(models.SalesOrder)
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
			&t.SPRAS,
			&t.KTOPL,
			&t.COA,
			&t.TXT20,
			&t.TXT50,
			&t.MCOD1,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *SalesOrderRepository) Fetch(ctx context.Context, limit, offset int) ([]*models.SalesOrder, error) {
	if limit != 0 {
		query := `SELECT * FROM sales_orders where is_deleted = 0 AND is_active = 1`

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
		query := `SELECT * FROM sales_orders where is_deleted = 0 AND is_active = 1`

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

func (m *SalesOrderRepository) GetCount(ctx context.Context) (int, error) {
	query := `SELECT count(*) AS count FROM sales_orders WHERE is_deleted = 0 and is_active = 1`

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

func (m *SalesOrderRepository) Insert(ctx context.Context, a *models.SalesOrder) (*int, error) {
	query := `INSERT sales_orders SET created_by=? , created_date=? , modified_by=?, modified_date=? , 	
				deleted_by=? , deleted_date=? , is_deleted=? , is_active=? , spras=?, ktopl=?, coa=?,
				txt20=?, txt50=?, mcod1=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	res, err := stmt.ExecContext(ctx, a.CreatedBy, a.CreatedDate, nil, nil, nil, nil, 0, 1, a.SPRAS, a.KTOPL,
		a.COA, a.TXT20, a.TXT50, a.MCOD1)
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
