package demo

import (
	"context"
	"github.com/polaris1119/go-web-layout/model"
)

type Repository interface {
	Create(ctx context.Context, demo *model.Demo) error
	FindById(ctx context.Context, id int) (*model.Demo, error)
	Update(ctx context.Context, demo *model.Demo) error
	Delete(ctx context.Context, id int) error
}
