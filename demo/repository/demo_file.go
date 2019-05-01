package repository

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/polaris1119/go-web-layout/model"
)

type fileRepository struct {
	filename string
}

func NewFileRepository(filename string) *fileRepository {
	return &fileRepository{
		filename: filename,
	}
}

func (f *fileRepository) Create(ctx context.Context, demo *model.Demo) error {
	b, err := json.Marshal(demo)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(f.filename, b, 0664)
}

func (f *fileRepository) FindById(ctx context.Context, id int) (*model.Demo, error) {
	file, err := os.Open(f.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := &model.Demo{}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(d)

	return d, err
}

func (f *fileRepository) Update(ctx context.Context, demo *model.Demo) error {
	return nil
}

func (f *fileRepository) Delete(ctx context.Context, id int) error {
	return nil
}
