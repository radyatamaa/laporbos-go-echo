package master_customer

import (
	"github.com/models"
	"golang.org/x/net/context"
)

type Repository interface {
	Fetch(ctx context.Context, limit, offset int) ([]*models.MasterCustomer, error)
	GetCount(ctx context.Context) (int, error)
	Insert(ctx context.Context, a *models.MasterCustomer) (*int, error)
}
