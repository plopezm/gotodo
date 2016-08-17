package main

import "testing"
import "gopkg.in/mgo.v2/bson"

func TestMongoDBConnect(t *testing.T){
	err := connect("localhost");
	if(err != nil){
		t.Error("Error connecting, cause: ", err);
	}
}

func TestMongoDBInsert(t *testing.T){
	todo := Todo{Name: "Write presentation", Desc: "Write IOT presentation"};
	
	err := insertTodo(&todo);
	if(err != nil){
		t.Error("Error inserting, cause: ", err);
	}
}

func TestMongoDBFind(t *testing.T){

	_, err := findTodos(bson.M{"name": "Write presentation"});
	if(err != nil){	
		t.Error("Error finding, cause: ", err);
	}
}

func TestMongoDBRemove(t *testing.T){
	err := removeTodo(bson.M{"name": "Write presentation"});
	if(err != nil){	
		t.Error("Error removing, cause: ", err);
	}
}

func TestCloseMongoDBManager(t *testing.T){
	closeSession();
}
