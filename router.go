package main

import (
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
)

func NewRouter(routes Routes) *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    
    //Looping between routes from routes.go

    fmt.Println("Adding routes:\n");
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)
        fmt.Println("Adding route: "+route.Name+" -> "+route.Method+" "+route.Pattern);
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    //Adding path as web-page server
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("webapp")))

    //Return router created
    return router
}
