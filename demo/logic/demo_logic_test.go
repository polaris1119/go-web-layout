package logic

import (
	"context"
	"reflect"
	"testing"

	"github.com/polaris1119/go-web-layout/demo"
	"github.com/polaris1119/go-web-layout/model"
)

func Test_demoLogic_Fetch(t *testing.T) {
	type fields struct {
		repository demo.Repository
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Demo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl := &demoLogic{
				repository: tt.fields.repository,
			}
			got, err := dl.Fetch(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("demoLogic.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("demoLogic.Fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_demoLogic_Create(t *testing.T) {
	type fields struct {
		repository demo.Repository
	}
	type args struct {
		ctx  context.Context
		data string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl := &demoLogic{
				repository: tt.fields.repository,
			}
			if err := dl.Create(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("demoLogic.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
