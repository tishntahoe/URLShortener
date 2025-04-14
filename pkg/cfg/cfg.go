package cfg

import "os"

type Cfg struct {
	ConnectionPgString,
	ConnectioRedisIP,
	ConnectioRedisPass string
}

func CfgLaunch() *Cfg {
	return &Cfg{
		ConnectionPgString: os.Getenv("CONNECTION_STRING"),
		ConnectioRedisIP:   os.Getenv("CONNECTION_REDIS_IP"),
		ConnectioRedisPass: os.Getenv("CONNECTION_REDIS_PASS"),
	}
}
