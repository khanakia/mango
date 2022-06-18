package wireapp

import (
	"github.com/khanakia/mangocms/mango/cli"
	"github.com/khanakia/mangocms/mango/dbconn"
	"github.com/khanakia/mangocms/mango/gormdb"
	"github.com/khanakia/mangocms/mango/interfaces"
	"github.com/khanakia/mangocms/mango/natso"
	"github.com/ubgo/gofm/logger"
)

type Plugin struct {
	ConfigMgr interfaces.IConfig
	Cli       cli.Cli
	Logger    logger.Logger
	DbConn    dbconn.DbConn
	GormDB    gormdb.GormDB
	Natso     natso.Natso
	// Auth      auth_app.Auth
}
