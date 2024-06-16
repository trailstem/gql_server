package graph

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/trailstem/graphql-server/graph/services"
	"github.com/trailstem/graphql-server/internal"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Srv services.Services
}

func NewExecutableSchema() graphql.ExecutableSchema {
	return internal.NewExecutableSchema(
		internal.Config{
			Resolvers:  &Resolver{},
			Directives: internal.DirectiveRoot{IsAuthenticated: IsAuthenticated},
		})
}

func NewServer() *handler.Server {
	srv := handler.NewDefaultServer(NewExecutableSchema())

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		KeepAlivePingInterval: 10 * time.Second,
	})

	return srv
}
