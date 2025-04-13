package cfg

import "os"

type Cfg struct {
	ConnectionDbString string
}

func CfgLaunch() *Cfg {
	return &Cfg{
		ConnectionDbString: os.Getenv("CONNECTION_STRING"),
	}
}
