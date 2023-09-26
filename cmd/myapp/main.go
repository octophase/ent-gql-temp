package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"myapp"
	"myapp/ent"
	"myapp/ent/privacy"
	_ "myapp/ent/runtime"
	"myapp/ent/user"
	"myapp/viewer"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/99designs/gqlgen/graphql/handler"
	//"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	//"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	// Create ent.Client and run the schema migration
	client := Open("postgresql://postgres:postgres@127.0.0.1/myapp-ent")

	// Context
	ctx := context.Background()

	// Atlas
	err := client.Schema.Create(
		ctx,
		schema.WithGlobalUniqueID(true),
		schema.WithAtlas(true),
	)
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	/*
		if err := Do(ctx, client); err != nil {
			log.Fatal(err)
		}
	*/

	router := chi.NewRouter()

	// Add CORS middleware around every request
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*", "http://localhost:3003"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(myapp.NewSchema(client))
	/*
		srv.AddTransport(&transport.Websocket{
			Upgrader: websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					// Check against your desired domains here
					return r.Host == "localhost"
				},
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
			},
		})
	*/

	router.Handle("/", playground.Handler("MyApp", "/query"))
	router.Handle("/query", srv)

	err = http.ListenAndServe(":8081", router)
	if err != nil {
		panic(err)
	}
}

func Do(ctx context.Context, client *ent.Client) error {
	// Expect operation to fail, because viewer-context
	// is missing (first mutation rule check).
	if _, err := client.User.Query().Where(user.UsernameEQ("johnsmith")).Only(ctx); !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operation to fail, but got %w", err)
	}

	// Apply the same operation with "Admin" role.
	admin := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.Admin})
	if err := client.User.Create().Exec(admin); err != nil {
		return fmt.Errorf("expect operation to pass, but got %w", err)
	}

	// Apply the same operation with "ViewOnly" role.
	viewOnly := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.View})
	if err := client.User.Create().Exec(viewOnly); !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operation to fail, but got %w", err)
	}

	// Allow all viewers to query users.
	for _, ctx := range []context.Context{ctx, viewOnly, admin} {
		// Operation should pass for all viewers.
		count := client.User.Query().CountX(ctx)
		fmt.Println(count)
	}
	return nil
}
