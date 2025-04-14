package storage

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
)

type StorageStuct struct {
	Db        *sql.DB
	RedisConn *redis.Client
}

var Storage *StorageStuct
