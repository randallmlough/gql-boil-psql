package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/randallmloug/gqlboil/auth"
	"github.com/randallmloug/gqlboil/graphql_models"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	tb "github.com/didip/tollbooth"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	initLogger()
	db := getDatabase()

	port := os.Getenv("PORT")

	// Create a limiter
	lmt := tb.NewLimiter(100, nil)
	lmt.SetOnLimitReached(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response, err := json.Marshal(map[string]interface{}{
			"Message": "Rate exceeded",
		})
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(`{"code": 500, "message": "Could not write response"}`))
			return
		}
		w.WriteHeader(429)
		w.Write(response)
		return
	})

	c := graphql_models.Config{
		Resolvers: &Resolver{},
	}

	c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		if exist := auth.ExistsInContext((ctx)); !exist {
			return nil, fmt.Errorf("Access denied")
		}
		return next(ctx)
	}

	srv := handler.New(graphql_models.NewExecutableSchema(c))
	srv.Use(extension.Introspection{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	// srv.Use(apollotracing.Tracer{})

	r := chi.NewRouter()

	r.Use(auth.Middleware(db))
	r.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	r.Handle("/graphql", srv)

	log.Info().Str("host", "localhost").Str("port", port).Msg("Connect GraphQL playground")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal().Err(err).Str("Could not listen to port", port).Send()
	}
}

func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	level, levelError := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	if levelError != nil {
		// Default level for this example is info, unless debug flag is present
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(level)
	}
}
