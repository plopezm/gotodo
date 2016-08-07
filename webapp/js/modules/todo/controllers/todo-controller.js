app.controller('todoController', ['$scope', 'todoService', function($scope, todoService){

	this.todoList = [];

	this.init = function(){
		todoService.getTodoList().then(function(response){
			$scope.todoController.todoList = response.data;
		});
	}	
	
	this.init();
}]);

