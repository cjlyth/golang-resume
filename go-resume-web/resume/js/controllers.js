'use strict';
/* Controllers */
angular.module('resumeApp.controllers', ['ngCookies', 'resumeApp.services'])
  .controller('GridCtrl', [function() {}])
  .controller('ResumeCtrl', ['$scope', '$cookieStore', 'resumeData', 'companies',
  	function($scope, $cookieStore, resumeData, companies) {
			$scope.selectedRole = $scope.$eval("$cookieStore.get('selectedRole')");
			$scope.selectRole = function(role){
				$scope.$evalAsync('selectedRole = "' + role + '"');
			}
			$scope.$watch('selectedRole', "selectedRole && $cookieStore.put('selectedRole', selectedRole)");
    	$scope.resume = { experience: resumeData.getExperience(companies) };
  	}])


  .controller('ProjectController', ['$scope', 'filterFilter', function($scope, filterFilter) {
    this.data = $scope.project;
    this.longName = $scope.project.longName;
    this.shortName = $scope.project.shortName;
    this.dates =  $scope.project.dates;
    // this.array = [
    //   {name: 'Tobias'},
    //   {name: 'Jeff'},
    //   {name: 'Brian'},
    //   {name: 'Igor'},
    //   {name: 'James'},
    //   {name: 'Brad'}
    // ];
    // this.filteredArray = filterFilter(this.array, 'a');
  }])
;
