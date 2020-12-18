package ar

import (
	"github.com/models"
	"golang.org/x/net/context"
)

type Repository interface {
	Fetch(ctx context.Context, limit, offset int) ([]*models.Ar, error)
	GetCount(ctx context.Context) (int, error)
	Insert(ctx context.Context, a *models.Ar) (*int, error)
}
