package demo

import (
	"context"
	"github.com/polaris1119/go-web-layout/model"
)

type Logic interface {
	Fetch(ctx context.Context, id int) (*model.Demo, error)
	Create(ctx context.Context, data string) error
}
