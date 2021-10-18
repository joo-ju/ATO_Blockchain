'use strict';
var app = angular.module('application', []);
app.controller('AppCtrl', function($scope, appFactory) {
  $scope("#success_getwallet").hide();
  $scope("#success_getallgoods").hide();
  $scope("#success_setgoods").hide();
  $scope("#success_getgoods").hide();
  $scope("#success_changegoodsprice").hide();
  $scope("#success_deletegoods").hide();
  
  $scope.getWallet = function() {
    appFactory.getWallet($scope.walletid, function(data) {
      $scope.search_wallet = data;
      $("#success_getwallet").show();
    });
  }

  $scope.getAllWallet = function() {
    appFactory.getAllWallet(function(data) {
      var array = [];
      
      for (var i=0; i<data.length; i++) {
        parseInt(data[i].key);
        data[i].Record.Key = data[i].Key;
        array.push(data[i].Record);
        $("#success_getallgoods").hide();
      }
      
      array.sort(function(a, b) {
        return parseFloat(a.Key) - parseFloat(b.Key);
      });

      $scope.allGoods = array;
    });
  }

  $scope.getGoods = function() {
    appFactory.getGoods($scope.musickey, function(data) {
      $("#success_getgoods").show();
    });
  }

  $scope.setGoods = function() {
    appFactory.setGoods($scopte.goods, function(data) {
      $scope.create_goods = data;
      $("#success_setgoods").show();
    });
  }

  $scope.purchaseGoods = function(key) {
    appFactory.purchaseGoods(key, function(data) {
      var array = [];
      
      for (var i=0; i<data.length; i++) {
        parseInt(data[i].Key);
        data[i].Record.Key = data[i].Key;
        array.push(data[i].Record);
        $("#success_getallgoods").hide();
      }
      
      array.sort(function(a, b) {
        return parseFloat(a.Key) - parseFloat(b.Key);
      });

      $scope.allGoods = array;
    });
  }

  $scope.changeGoodsPrice = function() {
    appFactory.changeGoodsPrice($scope.change, function(data) {
      $scope.change_goods_price = data;
      $("#success_changegoodsprice").show();
    });
  }

  $scope.deleteGoods = function() {
    appFactory.deleteGoods($scope.goodskeydelete, function(data) {
      $scope.delete_goods = data;
      $("#success_deletegoods").show();
    });
  }
  
});

app.factory('appFactory', function($http) {
  var factory = {};

  factory.getWallet = function(key, callback) {
    $http.get('/api/getWallet?walletid='+key).success(function(output) {
      callback(output)
    });
  }

  factory.getAllGoods = function(callback) {
    $http.get('/api/getAllGoods/').success(function(output) {
      callback(output)
    });
  }
  
  factory.getGoods = function(key, callback) {
    $http.get('/api/getGoods?goodskey='+key).success(function(output) {
      callback(output)
    });
  }

  factory.setGoods = function(data, callback) {
    $http.get('/api/setGoods?title='+data.title+'&state='+data.state+'&seller='+data.seller+
              '&category='+data.category+'&price='+data.price+'&content='+data.content+
              '&walletid'+data.walletid).success(function(output) {
      callback(output)
    });
  }

  factory.purchaseGoods = function(key, callback) {
    $http.get('/api/getAllGoods/').success(function(output) {
      callback(output)
    });
  }

  factory.changeGoodsPrice = function(data, callback) {
    $http.get('/api/deleteGoods?goodskey'+data.goodskey+'&price='+data.price).success(function(output) {
      callback(output)
    });
  }

  factory.deleteGoods = function(key, callback) {
    $http.get('/api/deleteGoods?goodskey='+key).success(functon(output) {
      callback(output)
    });
  }

  return factory;
});
