package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/zainul/ark/xlog"
	ownhttp "github.com/zainul/arkana-kit/internal/delivery/http"
	"github.com/zainul/arkana-kit/internal/middleware"
	"github.com/zainul/arkana-kit/internal/pkg/error/usecaseerror"
	"github.com/zainul/arkana-kit/internal/pkg/initial"
	"github.com/zainul/arkana-kit/internal/repository/store"
	"github.com/zainul/arkana-kit/internal/usecase"
	"github.com/zainul/nux"
)

// LogInit is initial log
func LogInit() {
	xlog.NewXLog("some server", "pass", "user", "queue name", "user-svc", 5672)
}

func main() {

	var wait time.Duration
	flag.DurationVar(&wait,
		"graceful-timeout", time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	db := initial.GetDB()
	defer db.Close()

	// mgb := initial.GetMongoDB()
	// mq := initial.GetMQ()

	// init logger
	LogInit()

	nux.NewError(usecaseerror.GetErrors())

	r := mux.NewRouter()
	store := store.NewStore(db)
	usecaseUser := usecase.NewUser(store)
	ownhttp.NewUserHandler(r, usecaseUser)

	r.Use(middleware.ContentType)

	http.Handle("/", r)

	// config the http server
	srv := &http.Server{
		Addr: "0.0.0.0:7000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		// WriteTimeout: time.Second * 1,
		ReadTimeout: time.Second * 5, // -> it will be in config file
		// IdleTimeout: time.Second * 1,
		Handler: r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Server up :7000")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
