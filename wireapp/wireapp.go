//go:build wireinject
// +build wireinject

package wireapp

import (
	"github.com/google/wire"
	"github.com/khanakia/mangocms/mango/cli"
	"github.com/khanakia/mangocms/mango/configmgr"
	"github.com/khanakia/mangocms/mango/dbconn"
	"github.com/khanakia/mangocms/mango/gormdb"
	"github.com/khanakia/mangocms/mango/interfaces"
	"github.com/khanakia/mangocms/mango/natso"
	"github.com/ubgo/gofm/logger"
)

func NewGormConfig(dbConn dbconn.DbConn) gormdb.Config {
	return gormdb.Config{
		DB: dbConn.SqlDb,
	}
}

func NewConfigMgrConfig() configmgr.Config {
	return configmgr.Config{}
}

// func NewConfigMgr(config configmgr.Config) *configmgr.ConfigMgr {
// 	return &configmgr.ConfigMgr{}
// }

func Init() Plugin {
	wire.Build(
		NewConfigMgrConfig,
		// NewConfigMgr,
		configmgr.New,
		wire.Bind(new(interfaces.IConfig), new(*configmgr.ConfigMgr)),
		cli.New,
		logger.New,
		wire.Struct(new(dbconn.Config), "*"),
		dbconn.New,
		NewGormConfig,
		gormdb.New,
		wire.Struct(new(natso.Config), "*"),
		natso.New,
		// wire.Struct(new(appcore.Config), "*"),
		// appcore.New,
		wire.Struct(new(Plugin), "*"),
	)
	return Plugin{}
}
