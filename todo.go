package main

import "time"

type Todo struct {
    Id        int       `json:"id"`
    Name      string    `json:"name"`
    Desc      string    `json:"desc"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
    Deadline  time.Time `json:"deadline"`
}

type Todos []Todo
