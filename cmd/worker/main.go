package main

import (
	"context"
	"i3/config"
	apirepository "i3/internal/api/repository"
	"i3/internal/meteo"
	"i3/internal/scheduler"
	"i3/pkg/datasource"
	"i3/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	ctx = context.Background()
)

func init() {
	config.InitConfig()
	logger.InitLogger()
}

func main() {
	pg := datasource.NewPgx(ctx)
	redis := datasource.NewRedis()
	repo := apirepository.New(pg)

	meteo := meteo.New(redis)
	s := scheduler.New()

	s.ScheduleJob(1*time.Minute, scheduler.NewWeatherJob(meteo, repo))

	s.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	s.Stop()
}
