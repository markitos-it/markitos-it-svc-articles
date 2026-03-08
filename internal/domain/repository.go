package domain

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Article, error)
	GetByID(ctx context.Context, id string) (*Article, error)
	Create(ctx context.Context, doc *Article) error
	Update(ctx context.Context, doc *Article) error
	Delete(ctx context.Context, id string) error
}
