app.service('todoService',['$http', function($http){
	
	this.getTodoList = function(){
		return $http.get('/api/v1/todos');
	}	

}]);
