package main

import (
	"fmt"
	"github.com/KenerChang/api-server/example"
	"github.com/KenerChang/api-server/middleware"
	"github.com/KenerChang/api-server/util/logger"
	"github.com/KenerChang/api-server/util/route"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var modules = []route.Module{
	example.ModuleInfo,
}

func main() {
	// setup routes
	r := setRoutes()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

func setRoutes() http.Handler {
	moduleRoutes := mux.NewRouter()
	for _, module := range modules {

		for _, entrypoint := range module.Entrypoints {
			path := fmt.Sprintf("/api/%s/v%d/%s", module.Name, entrypoint.Version, entrypoint.Path)
			logger.Info.Printf(nil, "set path: %s", path)

			moduleRoutes.HandleFunc(path, entrypoint.Callback).Methods(entrypoint.Method)
		}
	}
	moduleRoutes.Use(middleware.RequestIDMiddleware)

	return moduleRoutes
}
