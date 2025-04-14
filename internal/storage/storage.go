package storage

import (
	"context"
)

type StorageInterface interface {
	Save(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
}

func (s StorageStuct) Save(ctx context.Context, key, value string) error {

	return nil
}

func (s StorageStuct) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}
