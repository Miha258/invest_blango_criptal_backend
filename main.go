package main

import (
	"context"
	"invest_blango_criptal_backend/handlers"
	"invest_blango_criptal_backend/repository"
	"invest_blango_criptal_backend/service"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)


func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	
	if err := initConfig(); err != nil {
		logrus.Fatalf("Init config error: %s", err.Error())
	}

	mongoClient, ctx := connectMongoDB()
	repos := repository.NewRepository(mongoClient, ctx) // work with db
	services := service.NewService(repos) // buisnes logic
	handlers := handlers.NewHandler(services) //request handlers
	
	srv := new(Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error while running server: %s", err.Error())
	}
}


func connectMongoDB() (*mongo.Client, *context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongodb_uri"))); if err != nil {
		logrus.Fatalf("Init config error: %s", err.Error())
	}
	ctx := context.Background()
    err = client.Connect(ctx)
    if err != nil {
        logrus.Fatal(err)
    }
	return client, &ctx
}


func initConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("configs")
	viper.AddConfigPath("configs")
	return viper.ReadInConfig()
}


type Server struct {
	httpServer *http.Server
}


func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
		Handler: handler,
		MaxHeaderBytes: 1 << 20, //1 M,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}


func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}