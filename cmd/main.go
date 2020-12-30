package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"Ayudaap.org/src/routes"

	"Ayudaap.org/pkg/config"
	"github.com/gorilla/handlers"
	"github.com/spf13/viper"
)

func main() {

	config.Initconfig()
	puerto := viper.GetInt("API.PUERTO")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	//Si esta en modo desarrollo se habilita para todos los origenes, en productivo se debe restringir
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	r := routes.GetRouter()

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, handlers.CompressHandler(handlers.CORS(originsOk, headersOk, methodsOk)(r))),
		Addr:         fmt.Sprintf(":%d", puerto),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	fmt.Printf("Ejecutando en :%d\n", puerto)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
