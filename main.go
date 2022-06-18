package main

import (
	"github.com/khanakia/mangocms/apigql"
	"github.com/khanakia/mangocms/mango/plug"
	"github.com/khanakia/mangocms/wireapp"
	"github.com/spf13/cobra"
)

// type Plugin struct {
// 	ConfigMgr interfaces.IConfig
// 	Cli       cli.Cli
// 	DbConn    dbconn.DbConn
// 	GormDB    gormdb.GormDB
// 	Auth      auth_app.Auth
// 	Natso     natso.Natso
// }

// func NewGormConfig(db2 dbconn.DbConn) gormdb.Config {
// 	return gormdb.Config{
// 		DB: db2.SqlDb,
// 	}
// }

func main() {
	// configMgr := configmgr.New(configmgr.Config{})
	// cliCli := cli.New()

	// dbDB := dbconn.New(dbconn.Config{
	// 	ConfigMgr: configMgr,
	// })
	// config := NewGormConfig(dbDB)
	// gormDB := gormdb.New(config)

	// natsoNatso := natso.New(natso.Config{
	// 	Cli: cliCli,
	// })

	// authAuth := auth_app.New(auth_app.Config{
	// 	Cli:    cliCli,
	// 	GormDB: gormDB,
	// 	Natso:  natsoNatso,
	// })

	// plugin := Plugin{
	// 	ConfigMgr: configMgr,
	// 	Cli:       cliCli,
	// 	DbConn:    dbDB,
	// 	GormDB:    gormDB,
	// 	Natso:     natsoNatso,
	// 	Auth:      authAuth,
	// }

	// plug.InitPlugins(cliCli.RootCmd, plugin)
	// plugin.Cli.Execute()

	// USING WIRE
	plugin := wireapp.Init()
	plug.InitPlugins(plugin.Cli.RootCmd, plugin)

	var ServerCommand = &cobra.Command{
		Use:   "server",
		Short: "starts the HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			apigql.Boot(plugin)
		},
	}

	plugin.Cli.RootCmd.AddCommand(ServerCommand)

	plugin.Cli.Execute()
}
