package main

import "github.com/adolphlwq/tinyurl/cmd"

func main() {
	/*
		var (
			configPath string
			host       string
			port       string
		)
		flag.StringVar(&configPath, "config", "default.properties", "config path")
		flag.StringVar(&host, "host", "0.0.0.0", "tinyurl server bind host")
		flag.StringVar(&port, "port", "8877", "tinyurl server bind port")
		flag.Parse()

		mysqlClient := mysql.NewMySQLClient(configPath)
		appService := &server.ServiceProvider{
			MysqlClient: mysqlClient,
			UriUUID:     uriuuid.BasicURIUUID{},
		}
		addr := net.JoinHostPort(host, port)

		server.Start(addr, appService)
	*/

	cmd.Execute()
}
