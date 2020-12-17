package sales_order

import (
	"github.com/models"
	"golang.org/x/net/context"
)

type Usecase interface {
	GetAll(ctx context.Context, page, limit, offset int) (*models.SalesOrderDtoWithPagination, error)
	Import(ctx context.Context, fileLocation string) error
}
