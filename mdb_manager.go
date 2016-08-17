package main

import (
        //"fmt"
	//"log"
        "gopkg.in/mgo.v2"
	"errors"
        "gopkg.in/mgo.v2/bson"
)

var session *mgo.Session;

var mongoDBSessionNotCreated = errors.New("MongoDBManager: Session not created");
var mongoDBItemExists = errors.New("MongoDBManager: TODO exists");
var mongoDBItemNotFound = errors.New("MongoDBManager: TODO not found");

func connect(url string) error{	
	if(session == nil){
		s, err := mgo.Dial(url);
		session = s;
		return err;
	}
	return nil;
}

func closeSession() error{
	if(session == nil){
		return mongoDBSessionNotCreated;
	}
	session.Close();
	session = nil;
	return nil;
}

func insertTodo(todo *Todo) error{
	if(session == nil){
		return mongoDBSessionNotCreated;
	}	
	c := session.DB("todoApp").C("todos");
    	result, err := c.Find(bson.M{"id": todo.Id}).Count()
	if(err != nil){
		return err;
	}
	if(result == 0){
		return c.Insert(todo);
	}else{
		return mongoDBItemExists;
	}
}

func findTodos(pattern bson.M) (Todos, error){
	var result Todos;

	if(session == nil){
		return nil, mongoDBSessionNotCreated;
	}

	c := session.DB("todoApp").C("todos");

	//err := c.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0}).All(&result)
	err := c.Find(pattern).All(&result)
	return result, err;
} 

func removeTodo(pattern bson.M) error{
	if(session == nil){
		return mongoDBSessionNotCreated;
	}
	c := session.DB("todoApp").C("todos");
	info, err := c.RemoveAll(pattern);
	
	if(err != nil){
		return err;
	}
	if(info.Removed == 0){
		return mongoDBItemNotFound;
	}
	return err;
}



