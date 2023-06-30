// export PATH="$PATH:$(go env GOPATH)/bin"
// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative schema.proto

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"opentaxi/database/auth/xauth"
	"opentaxi/graph"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func conexion() *sql.DB {
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbhost := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	loc := "America%2FLa_Paz"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=%s", dbuser, dbpass, dbhost, dbname, loc)
	db, err := sql.Open("mysql", dsn)

	er := db.Ping()
	if er != nil {
		panic(er.Error())
	}

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	db := conexion()
	router := chi.NewRouter()
	router.Use(xauth.AuthMiddleware(db))
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	resolver := &graph.Resolver{DB: db}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	show_playground := os.Getenv("PLAYGROUND")
	if show_playground == "1" {
		router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
