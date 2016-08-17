app.controller('todoController', ['$scope', 'todoService', function($scope, todoService){

	this.todoList = [];
	this.newTodo = {
		name: "",
		desc: ""
	}

	this.init = function(){
		todoService.getTodoList().then(function(response){
			$scope.todoController.todoList = response.data;
		});
	}	

	this.saveTodo = function(){
		todoService.saveTodo(this.newTodo).then(function(response){
			console.log(response);
		});
	}

	this.removeTodo = function(todo){
		todoService.removeTodo(todo).then(function(response){
			console.log(response);
		});
	}
	
	this.init();
}]);

