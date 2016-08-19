package main

import (
    "log"
    "fmt"
    "net/http"
)

var routes = Routes{
    Route{
        "TodoIndex",
        "GET",
        "/api/v1/todos",
        TodoIndex,
    },
    Route{
        "TodoShow",
        "GET",
        "/api/v1/todos/{todoId}",
        TodoShow,
    },
    Route{
	"TodoCreate",
	"PUT",
	"/api/v1/todos",
	TodoCreate,
    },
    Route{
	"TodoComplete",
	"POST",
	"/api/v1/todos/complete/{todoId}",
	TodoComplete,
    },
    Route{
	"TodoRemove",
	"DELETE",
	"/api/v1/todos/{todoId}",
	TodoRemove,
    },
}

func main() {
    port := ":8080";
    fmt.Println("====================================");
    fmt.Println("Starting server at port "+port);
    fmt.Println("====================================");

    router := NewRouter(routes);

    //Adding path as web-page server
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("webapp")));

    //Open new mongodb session
    mdbOpenSession("localhost");

    fmt.Println("====================================");
    log.Fatal(http.ListenAndServe(port, router));
}
