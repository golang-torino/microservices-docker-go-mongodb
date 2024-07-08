package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/mmorejon/microservices-docker-go-mongodb/users/pkg/models/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/log/global"
)

type application struct {
	log   *slog.Logger
	users *mongodb.UserModel
}

func main() {
	if err := run(); err != nil {
		logFatal(fmt.Sprintf("exited with error: %s", err))
	} else {
		fmt.Println("I'm done")
	}
}

func run() error {
	// Define command-line flags
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4000, "HTTP server network port")
	mongoURI := flag.String("mongoURI", "mongodb://localhost:27017", "Database hostname url")
	mongoDatabase := flag.String("mongoDatabase", "users", "Database name")
	enableCredentials := flag.Bool("enableCredentials", false, "Enable the use of credentials for mongo connection")
	flag.Parse()

	shutdown, err := setupOTelSDK(context.Background())
	if err != nil {
		return fmt.Errorf("cannot setup otel sdk: %w", err)
	}
	defer func() {
		// TODO: add a timeout?
		if err = shutdown(context.Background()); err != nil {

			logFatal(fmt.Sprintf("failed shutting down tracer provider: %s\n", err))
		}
	}()

	l := otelslog.NewLogger("website", otelslog.WithLoggerProvider(global.GetLoggerProvider()))
	slog.SetDefault(l)

	// Create mongo client configuration
	co := options.Client().ApplyURI(*mongoURI)
	if *enableCredentials {
		co.Auth = &options.Credential{
			Username: os.Getenv("MONGODB_USERNAME"),
			Password: os.Getenv("MONGODB_PASSWORD"),
		}
	}

	// Establish database connection
	client, err := mongo.NewClient(co)
	if err != nil {
		logFatal(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logFatal(err.Error())
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	l.Info("Database connection established")

	// Initialize a new instance of application containing the dependencies.
	app := &application{
		log: l,
		users: &mongodb.UserModel{
			C: client.Database(*mongoDatabase).Collection("users"),
		},
	}

	// Initialize a new http.Server struct.
	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	srv := &http.Server{
		Addr:         serverURI,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	l.Info("Starting server", "uri", serverURI)
	return srv.ListenAndServe()
}

func logFatal(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
