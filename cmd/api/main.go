package main

import (
	"context"
	"flag"
	"i3/config"
	"i3/internal/http_handler"
	"i3/internal/repository"
	"i3/pkg/datasource"
	"i3/pkg/logger"
	"i3/pkg/router"
)

var (
	ctx = context.Background()
)

func InitConfig() {
	config.InitConfig()
	logger.InitLogger()
}

func main() {
	flag.Parse()

	pg := datasource.NewPgx(ctx)

	repo := repository.New(pg)
	httpHandler := http_handler.New(repo)
	r := router.NewGin(httpHandler)

	r.Run(":" + config.ReadConfig().Port)

	// Graceful shutdown implementation

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
