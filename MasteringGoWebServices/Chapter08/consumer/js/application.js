var SocialNetworkApp = angular.module('SocialNetwork', ['ngRoute']);

SocialNetworkApp.config(function($routeProvider) {
	$routeProvider
	.when('/login', 
		{
			controller: 'Authentication',
			templateUrl: '/views/auth.html'
		}
	).when('/users', 
		{
			controller: 'Users',
			templateUrl: '/views/users.html'
		}
	).when('/statuses', 
		{
			controller: 'Statuses',
			templateUrl: '/views/statuses.html'
		}
	);
});

SocialNetworkApp.controller('Users',['$scope', '$http', '$location', '$routeParams', function($scope,$http,$location,$routeParams) {
	
	$scope.users = [];
	$http.get('https://www.mastergoco.com/api/users').success(function(data) {
		$scope.users = data.users;
	});

}]);

SocialNetworkApp.controller('Statuses',['$scope', '$http', '$location', '$routeParams', function($scope,$http,$location,$routeParams) {
	
	$scope.statuses = [];
	
	$scope.getStatuses = function() {
		$http.get('https://www.mastergoco.com/api/statuses').success(function(data) {
		
		});
	};
	
	$scope.postStatus = function() {
	
	}
	
	$scope.getStatuses();

}]);
 
SocialNetworkApp.controller('Authentication',['$scope', '$http', '$location', '$routeParams', function($scope,$http,$location,$routeParams) {

	$scope.login = function() {

		postData = { email: $scope.loginEmail, password: $scope.loginPassword };
		$http({
			url: 'https://www.mastergoco.com/api/user',
			method: 'POST',
			data: JSON.stringify(postData),
			headers: {'Content-Type': 'application/json'}
		}).success(function(data) {
			$location.path('/users');
		}).error(function(data,status,headers,config) {
			alert ("Error: " + status);
		});

	};

	$scope.register = function() {
		postData = { user: $scope.registerUser, email: $scope.registerEmail, first: $scope.registerFirst, last: $scope.registerLast, password: $scope.registerPassword };	
		$http.post('https://localhost:443/api/users', postData).success(function(data) {
		
			$location.path('/users');
		
		}).error(function(data,status,headers,config) {
			alert ("Error: " + status);			
		});		
	};

}]);