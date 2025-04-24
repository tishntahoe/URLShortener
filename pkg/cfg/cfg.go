package cfg

import "os"

type Cfg struct {
	ConnectionPgString,
	ConnectioRedisIP,
	ConnectioRedisPass,
	ConnectionIpServer string
}

func CfgLaunch() *Cfg {
	return &Cfg{
		ConnectionPgString: os.Getenv("CONNECTION_STRING"),
		ConnectioRedisIP:   os.Getenv("CONNECTION_REDIS_IP"),
		ConnectioRedisPass: os.Getenv("CONNECTION_REDIS_PASS"),
		ConnectionIpServer: os.Getenv("CONNECTION_IP_SERVER_ADDRESS"),
	}
}
