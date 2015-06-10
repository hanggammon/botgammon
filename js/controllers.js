// defines app and exposes underscore.js "_" everywhere within app
var botgammonApp = angular.module('botgammonApp', [])
  .constant('_', window._)
  .run(function ($rootScope) {
     $rootScope._ = window._;
  });

botgammonApp.controller('GameCtrl', function ($scope, $http) {   
  $scope.model = {
    // board state according to https://github.com/hanggammon/botgammon/blob/master/doc/protocol.md 
    game: undefined,
  };
  $scope.loadData = function() {
    $http.get('api/v1/games/0/').success(function(data) {
      $scope.model.game = data;
    });
  };

  // initial load
  $scope.loadData();
});
