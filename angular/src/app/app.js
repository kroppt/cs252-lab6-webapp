var myApp = angular.module('myApp',[]);
myApp.controller('reloadCtrl','$scope', function($scope){
$scope.count = 0;
$scope.myFunction = function(){
$scope.count++;
};
});
