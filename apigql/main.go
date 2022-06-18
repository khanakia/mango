package apigql

import (
	"github.com/khanakia/mangocms/apigql/graph/resolverfn"
	"github.com/khanakia/mangocms/apigql/pkg/gql"
	"github.com/khanakia/mangocms/apigql/pkg/middleware"
	"github.com/khanakia/mangocms/apigql/pkg/server_handler"
	"github.com/ubgo/gofm/ginserver"
	"github.com/ubgo/gqlgenfn"

	"github.com/khanakia/mangocms/wireapp"
)

func Boot(plugin wireapp.Plugin) {
	serverServer := ginserver.New(ginserver.Config{
		BeforeHandler: middleware.CORSMiddleware(),
	})

	serverServer.Router.Use(middleware.BlockHostsMiddleware())
	serverServer.Router.Use(gqlgenfn.GinContextToContextMiddleware())

	resolver := &resolverfn.Resolver{
		GormDB: plugin.GormDB,
		Logger: plugin.Logger,
	}
	gqlConfig := gql.Config{
		GormDB:   plugin.GormDB,
		Server:   serverServer,
		Resolver: resolver,
	}
	gql.New(gqlConfig)

	server_handler.New(server_handler.Config{
		GormDB: plugin.GormDB,
		Logger: plugin.Logger,
		Server: serverServer,
	})

	serverServer.Start()
}
