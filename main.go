package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/dev-crusader/movie-fetcher/client"
	grpsv "github.com/dev-crusader/movie-fetcher/grpc/server"
	mv "github.com/dev-crusader/movie-fetcher/internal"
	sv "github.com/dev-crusader/movie-fetcher/restapi"
	"github.com/dev-crusader/movie-fetcher/startup"
	"github.com/dev-crusader/movie-fetcher/startup/middleware"
)

var (
	method          = middleware.MethodType
	auth            = middleware.AuthMiddleware
	logger          = middleware.Logger
	fetcherHTTPFunc = sv.Fetcher
	grpcAddr        = flag.String("grpc", ":5001", "listen address of the grpc transport")
)

func main() {

	startup.Load()
	flag.Parse()
	fmt.Printf("\nGRPC server running on port %s\n", *grpcAddr)
	go grpsv.RunGRPCServer(*grpcAddr)
	m := &mv.Movie{Client: client.GetClient()}
	http.HandleFunc("/", method(auth(logger(fetcherHTTPFunc(m, sv.GetMovieHandler))), http.MethodPost))
	fmt.Println("Http server running on port :8080")
	http.ListenAndServe(":8080", nil)
}
