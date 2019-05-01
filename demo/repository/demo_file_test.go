package repository

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/polaris1119/go-web-layout/model"
)

func Test_fileRepository_Create(t *testing.T) {
	type fields struct {
		filename string
	}
	type args struct {
		ctx  context.Context
		demo *model.Demo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"first",
			fields{"tmp.json"},
			args{context.Background(), &model.Demo{ID: 1, Name: "polaris"}},
			false,
		},
	}
	defer func() {
		os.Remove("tmp.json")
	}()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileRepository{
				filename: tt.fields.filename,
			}
			if err := f.Create(tt.args.ctx, tt.args.demo); (err != nil) != tt.wantErr {
				t.Errorf("fileRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileRepository_FindById(t *testing.T) {
	type fields struct {
		filename string
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
			f := &fileRepository{
				filename: tt.fields.filename,
			}
			got, err := f.FindById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileRepository.FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileRepository.FindById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileRepository_Update(t *testing.T) {
	type fields struct {
		filename string
	}
	type args struct {
		ctx  context.Context
		demo *model.Demo
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
			f := &fileRepository{
				filename: tt.fields.filename,
			}
			if err := f.Update(tt.args.ctx, tt.args.demo); (err != nil) != tt.wantErr {
				t.Errorf("fileRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileRepository_Delete(t *testing.T) {
	type fields struct {
		filename string
	}
	type args struct {
		ctx context.Context
		id  int
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
			f := &fileRepository{
				filename: tt.fields.filename,
			}
			if err := f.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("fileRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
