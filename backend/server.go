package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/zenith110/portfilo/graph"
	"github.com/zenith110/portfilo/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("GRAPHQLPORT")
	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"https://abrahannevarez.dev/", "https://www.abrahannevarez.de/v", "https://graphql.abrahannevarez.dev/"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(&transport.Websocket{
        Upgrader: websocket.Upgrader{
            CheckOrigin: func(r *http.Request) bool {
                // Check against your desired domains here
                 return r.Host == allowedOrigin
            },
            ReadBufferSize:  1024,
            WriteBufferSize: 1024,
        },
    })
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, router))
}
