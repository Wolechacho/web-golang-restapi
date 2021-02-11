package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"web-golang-restapi/helpers"
	"web-golang-restapi/routers"
	"syscall"
	"time"
)

const (
	port = 6060
)

func main() {
	//connect to the db
	go func() {
		helpers.CreateDbConnection()
	}()

	defer helpers.DB.Close()

	//create channel to the start the server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	router := routers.RegisterRoutes()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	go func() {
		//create listner and block other process from executing
		err := server.ListenAndServe()

		//maybe invalid port related issues
		if err != nil {
			panic(err)
		}
	}()
	fmt.Println("Http server started listening on the port ", port)
	<-done

	//Receive Shutdown signal
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := server.Shutdown(ctx)
	if err == nil {
		fmt.Println("Server gracefully shutdown")
	}
}
