package logic

import (
	"context"
	"encoding/json"
	"github.com/polaris1119/go-web-layout/demo"
	"github.com/polaris1119/go-web-layout/model"
)

type demoLogic struct {
	repository demo.Repository
}

func NewDemoLogic(repository demo.Repository) *demoLogic {
	return &demoLogic{repository: repository}
}

func (dl *demoLogic) Fetch(ctx context.Context, id int) (*model.Demo, error) {
	return dl.repository.FindById(ctx, id)
}

func (dl *demoLogic) Create(ctx context.Context, data string) error {
	d := &model.Demo{}
	err := json.Unmarshal([]byte(data), d)
	if err != nil {
		return err
	}

	return dl.repository.Create(ctx, d)
}
