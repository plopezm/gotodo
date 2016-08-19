package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "strconv"
    "gopkg.in/mgo.v2/bson"
    "github.com/gorilla/mux"
)

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos, err := mdbFindTodos(nil);
    if(err != nil){
	w.WriteHeader(http.StatusBadRequest);
	fmt.Fprintf(w, "error finding: %s",err)
	return;
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8");
    w.WriteHeader(http.StatusOK);

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]

    todos, err := mdbFindTodos(bson.M{"id": todoId});
    if(err != nil){
	w.WriteHeader(http.StatusBadRequest);
	fmt.Fprintf(w, "error showing: %s",err)
	return;
    }

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }

}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &todo); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
	return;
    }

    err = mdbInsertTodo(&todo);
    if(err != nil){
	w.WriteHeader(http.StatusBadRequest);
	fmt.Fprintf(w, "error inserting: %s",err)
	return;
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(todo); err != nil {
        panic(err)
    }
}

func TodoComplete(w http.ResponseWriter, r *http.Request){
    /*
    vars := mux.Vars(r);
    todoId := vars["todoId"];

    i, err := strconv.Atoi(todoId);
    if err != nil {
    	w.WriteHeader(http.StatusBadRequest);
	fmt.Fprintf(w, "<todoId> must be integer")
	return;
    }
    err = RepoCompleteTodo(i);
    if err != nil {
    	w.WriteHeader(http.StatusNotFound);
	fmt.Fprintf(w, "Todo object not found")
	return;
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK);
    */
}

func TodoRemove(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r);
    todoId := vars["todoId"];

    i, err := strconv.Atoi(todoId);
    if err != nil {
    	w.WriteHeader(http.StatusBadRequest);
	fmt.Fprintf(w, "<todoId> must be integer")
	return;
    }
    err = mdbRemoveTodo(bson.M{"id": i});
    if err != nil {
    	w.WriteHeader(http.StatusNotFound);
	fmt.Fprintf(w, "Todo object not found")
	return;
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    
}



