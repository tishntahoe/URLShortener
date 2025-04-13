package storageInit

import (
	"database/sql"
	"github.com/tishntahoe/UrlShortener/internal/storage"
)

func StorageInit(connstr string) (*storage.StorageStr, error) {
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		// logger
		return nil, err
	}
	if err := db.Ping(); err != nil {
		// logger
		return nil, err
	}
	return &storage.StorageStr{Db: db}, nil
}
