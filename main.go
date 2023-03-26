package main

import (
	"flag"
	"os"

	"github.com/tongsq/go-lib/logger"
	"proxy-collect/bootstrap"
	"proxy-collect/bootstrap/servers"
	"proxy-collect/config"
)

var (
	configFile string
	serverList bootstrap.StringList
)

func main() {
	flag.Var(&serverList, "S", "start server type")
	flag.StringVar(&configFile, "C", "conf.yaml", "")
	flag.Parse()
	if flag.NFlag() == 0 || len(serverList) == 0 {
		serverList = append(serverList, "job", "api", "tunnel")
	}
	config.YamlPath = configFile
	//init
	bootstrap.Bootstrap()

	for _, server := range serverList {
		if server == bootstrap.ServerALl {
			go servers.StartApiServer()
			go servers.StartJobServer()
			go servers.StartTunnelServer()
			break
		} else if server == bootstrap.ServerJob {
			go servers.StartJobServer()
		} else if server == bootstrap.ServerApi {
			go servers.StartApiServer()
		} else if server == bootstrap.ServerTunnel {
			go servers.StartTunnelServer()
		} else {
			logger.Error("unknown server", nil)
			os.Exit(1)
		}
	}
	select {}
}
