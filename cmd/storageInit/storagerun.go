package storageInit

import (
	"context"
	"database/sql"
	"github.com/redis/go-redis/v9"
	"github.com/tishntahoe/UrlShortener/internal/storage"
	"github.com/tishntahoe/UrlShortener/pkg/cfg"
	"github.com/tishntahoe/UrlShortener/pkg/logger"
	"time"
)

func StorageInit(conf *cfg.Cfg) error {

	// PG для добавления авторизации при получении ссылки (возможность - не обязательность).

	//db, err := PostgresLaunch(conf.ConnectionPgString)
	//if err != nil {
	// logger
	//	return err
	//}
	rds, err := RedisLaunch(conf.ConnectioRedisIP, conf.ConnectioRedisPass)
	if err != nil {
		logger.ErrorHandler(err, logger.GetWorkDir())
		return err
	}
	storage.Storage = &storage.StorageStuct{Db: nil, RedisConn: rds}
	return nil
}

func PostgresLaunch(connstr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		logger.ErrorHandler(err, logger.GetWorkDir())
		return nil, err
	}
	if err := db.Ping(); err != nil {
		logger.ErrorHandler(err, logger.GetWorkDir())
		return nil, err
	}
	return db, nil
}

func RedisLaunch(ip, pass string) (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	rds := redis.NewClient(&redis.Options{
		Addr:     ip,
		Password: pass,
	})
	err := rds.Ping(ctx).Err()
	if err != nil {
		logger.ErrorHandler(err, logger.GetWorkDir())
		return nil, err
	}

	return rds, nil
}
