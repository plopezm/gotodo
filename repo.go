package main

import "fmt"

var currentId int

var todos = Todos{};

// Give us some seed data
func init() {
    RepoCreateTodo(Todo{Name: "Write presentation", Desc: "Write IOT presentation"})
    RepoCreateTodo(Todo{Name: "Host meetup", Desc: "Be care for telephone call at 9:00 am"})
}

func RepoFindTodo(id int) Todo {
    for _, t := range todos {
        if t.Id == id {
            return t
        }
    }
    // return empty Todo if not found
    return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
    currentId += 1
    t.Id = currentId
    todos = append(todos, t)
    return t
}

func RepoCompleteTodo(id int) error {
    for _, t := range todos {
        if t.Id == id {
	    t.Completed = true;
            return nil
        }
    }
    return fmt.Errorf("Could not find Todo with id of %d to delete", id);
}

func RepoDestroyTodo(id int) error {
    for i, t := range todos {
        if t.Id == id {
            todos = append(todos[:i], todos[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}

