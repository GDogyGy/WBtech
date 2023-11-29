package main

import (
	"WbTech0/internal/config"
	"WbTech0/internal/http-server/routes"
	"WbTech0/internal/lib/json"
	"WbTech0/internal/model"
	"WbTech0/internal/storage/postgres"
	"github.com/go-chi/chi/v5"
	stan "github.com/nats-io/stan.go"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()
	// TODO: init logger: slog
	log := setupLogger(cfg.Env)
	var data []model.Order

	//TODO: init storage: PostgreSQL
	storage, err := postgres.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to init storage", slog.StringValue(err.Error()))
		os.Exit(1)
	}

	// TODO: get cache data DB
	data = storage.GetOrderModels()

	// TODO: init route: chi
	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		router.HandleFunc("/", routes.OrderRoutes(log, storage, &data))
	})

	// TODO: init server, run server
	go StartServer(cfg, log, router)
	// TODO: start subscribe nats
	Subscriber(&data, log, storage)

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}

func Subscriber(data *[]model.Order, log *slog.Logger, storage *postgres.Storage) {
	sc, err := stan.Connect("test_cluster", "stan-server", stan.NatsURL("http://localhost:4222"))
	if err != nil {
		log.Error("Failed nats connect ", slog.StringValue(err.Error()))
	}
	defer sc.Close()

	sc.Subscribe("updates", func(m *stan.Msg) {
		jsn, err := json.ParseToModel(m.Data)
		if err != nil {
			return
		}
		*data = append(*data, jsn)
		order, err := storage.InsertOrder(jsn)
		if err != nil {
			log.Error("Failed inser order in BD", slog.StringValue(err.Error()), order)
		}
	}, stan.DurableName("my-durable"), stan.MaxInflight(1), stan.AckWait(20*time.Second))
	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}

func StartServer(cfg *config.Config, log *slog.Logger, router *chi.Mux) {
	log.Info("server starting", slog.String("address", cfg.Address))
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Error("Failed to start server")
	}
}
