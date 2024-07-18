package main

import (
	"context"
	"i3/config"
	apicontroller "i3/internal/api/controller"
	apirepository "i3/internal/api/repository"
	apiusecase "i3/internal/api/usecase"
	"i3/internal/meteo"
	"i3/internal/scheduler"
	"i3/pkg/datasource"
	"i3/pkg/logger"
	"i3/pkg/router"
	"time"
)

var (
	ctx = context.Background()
)

func InitConfig() {
	config.InitConfig()
	logger.InitLogger()
}

func main() {
	pg := datasource.NewPgx(ctx)
	repo := apirepository.New(pg)
	uc := apiusecase.New(repo)
	ctrl := apicontroller.New(uc)

	// Scheduler
	redis := datasource.NewRedis()
	meteo := meteo.New(redis)

	s := scheduler.New()
	s.ScheduleJob(1*time.Minute, scheduler.NewWeatherJob(meteo, repo))
	s.Start()

	// Router
	r := router.NewGin(ctrl)
	r.Run(":" + config.ReadConfig().Port)

	s.Stop()

	// Graceful implementation

	// srv := &http.Server{
	// 	Addr:    ":" + config.ReadConfig().Port,
	// 	Handler: r,
	// }

	// go func() {
	// 	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()

	// log.Printf("Server is running on port :: %v", config.ReadConfig().Port)

	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// <-sigChan

	// shutdown, release := context.WithTimeout(context.Background(), 10*time.Second)
	// defer release()

	// if err := srv.Shutdown(shutdown); err != nil {
	// 	log.Fatal(err)
	// }
}
