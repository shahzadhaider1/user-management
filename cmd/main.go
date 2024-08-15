package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"

	"user-management/graph"
	"user-management/internal/routes"
)

func main() {
	dbURL := "postgres://user:password@localhost:5432/user_management" // Adjust this as needed
	dbPool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	r := graph.NewResolver(dbPool)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: r})))

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	routes.InitRoutes(router)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to run server, error: %v\n", err)
		}
	}()
	log.Println("server running on port 8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("failed to shutdown the server, error:", err)
	}

	log.Println("server stopped successfully")
}
