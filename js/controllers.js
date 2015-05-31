var botgammonApp = angular.module('botgammonApp', [])
  // allow DI for use in controllers, unit tests
  .constant('_', window._)
  // use in views, ng-repeat="x in _.range(3)"
  .run(function ($rootScope) {
     $rootScope._ = window._;
  });

botgammonApp.controller('GameCtrl', function ($scope, $http) {
  $scope.loadData = function() {
    $http.get('api/v1/games/0/').success(function(data) {
      $scope.game = data;
    });
  };

  // initial load
  $scope.loadData();
});
