package main

import "testing"
import "fmt"
import "gopkg.in/mgo.v2/bson"

func TestMongoDBConnect(t *testing.T){
	err := mdbOpenSession("localhost");
	if(err != nil){
		t.Error("Error connecting, cause: ", err);
	}
}

func TestMongoDBInsert(t *testing.T){
	todo := Todo{Name: "Write presentation", Desc: "Write IOT presentation"};
	
	err := mdbInsertTodo(&todo);
	if(err != nil){
		t.Error("Error inserting, cause: ", err);
	}
}

func TestMongoDBFind(t *testing.T){

	todos, err := mdbFindTodos(bson.M{"name": "Write presentation"});
	if(err != nil){	
		t.Error("Error finding, cause: ", err);
	}
	fmt.Println(todos);
}

func TestMongoDBRemove(t *testing.T){
	err := mdbRemoveTodo(bson.M{"name": "Write presentation"});
	if(err != nil){	
		t.Error("Error removing, cause: ", err);
	}
}

func TestCloseMongoDBManager(t *testing.T){
	mdbCloseSession();
}
