'use strict';
/* Directives */
angular.module('resumeApp.directives', [])

.directive('dateRange', [function(){
  var ddo = {
    scope: {
      range: "=dateRange",
      format: "@dateFormat"
    },
    templateUrl: 'partials/dateRange.html'
  };
  return ddo;
}])

.directive('cols', [function(){
  var ddo = {
    replace: false,
    scope: false,
    templateUrl: 'partials/columnGuide.html'
  };
  return ddo;
}])

.directive('project', [function(){
  var ddo = {
    templateUrl: 'partials/project.html'
  };
  return ddo;
}]);


// ,
// link: function postLink(scope, attr, ele) {
//   console.log('link',arguments);
// }
