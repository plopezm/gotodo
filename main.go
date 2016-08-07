package main

import (
    "log"
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
	"TodoRemove",
	"DELETE",
	"/api/v1/todos/{todoId}",
	TodoRemove,
    },
}



func main() {
    router := NewRouter(routes)

    log.Fatal(http.ListenAndServe(":8080", router))
}
