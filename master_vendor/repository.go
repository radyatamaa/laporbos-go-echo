package master_vendor

import (
	"github.com/models"
	"golang.org/x/net/context"
)

type Repository interface {
	Fetch(ctx context.Context, limit, offset int) ([]*models.MasterVendor, error)
	GetCount(ctx context.Context) (int, error)
	Insert(ctx context.Context, a *models.MasterVendor) (*int, error)
}
