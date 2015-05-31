// defines app and exposes underscore.js "_" everywhere within app
var botgammonApp = angular.module('botgammonApp', [])
  .constant('_', window._)
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
