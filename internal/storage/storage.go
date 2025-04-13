package storage

import (
	"context"
	"database/sql"
)

type StorageStr struct {
	Db *sql.DB
}

type Storage interface {
	Save(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
}

func Save() {

}

func Get() {

}
