package cmd

import (
	globalrouter "ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Hudai)

	mux := http.NewServeMux()

	initRoutes(mux, manager)

	fmt.Println("Server Running on: 8080")

	globalRouter := globalrouter.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
