package container

import (
	"context"
	"learn/pkg/http/handlers"
	"learn/pkg/http/routes"
	"learn/services/implementations"
	"learn/services/interfaces"
	"net/http"

	"go.uber.org/fx"
)

// BuildContainer sets up the dependency injection container
func BuildContainer() *fx.App {
	return fx.New(
		fx.Provide(
			NewHTTPServer,
			implementations.NewUserService,
			// handlers.NewUserHandler,
			handlers.NewHealthcheckHandler,
		),
		fx.Provide(
			func(healthCheckHandler *handlers.HealthcheckHandler) *http.ServeMux {
				return routes.SetupRoutes(healthCheckHandler)
			},
		),
		fx.Invoke(func(*http.Server, interfaces.IUserService) {}),
	)
}

func NewHTTPServer(lifecycle fx.Lifecycle, mux *http.ServeMux) *http.Server {
	server := &http.Server{Addr: ":8080", Handler: mux}
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil {
					panic("error occured")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
	return server
}
