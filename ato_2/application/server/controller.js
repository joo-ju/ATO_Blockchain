var sdk = require('./sdk.js');
module.exports = function(app) {

  app.get('/api/getWallet', function (req, res) {
    var walletid = req.query.walletid;
    let args = [walletid];
    sdk.send(false, 'getWallet', args, res);
  });
  
  app.get('/api/setWallet', function (req, res) {
    var name = req.query.name;
    var id = req.query.id;
    var coin = req.query.coin;
    let args = [name, id, coin];
    sdk.send(true, 'setWallet', args, res);
  });
  
  app.get('/api/getGoods', function (req, res) {
    var goodskey = req.query.goodskey;
    let args = [goodskey];
    sdk.send(true, 'getGoods', args, res);
  });
  
  app.get('/api/setGoods', function (req, res) {
    var title = req.query.title;
    var state = req.query.state;
    var seller = req.query.seller;
    var category = req.query.category;
    var price = req.query.price;
    var content = req.query.content;
    var walletid = req.query.price;
    let args = [title, state, seller, category, price, content, walletid];
    sdk.send(true, 'setGoods', args, res);
  });
  
  app.get('/api/getAllGoods', function (req, res) {
    let args = [];
    sdk.send(true, 'getAllGoods', args, res);
  });
  
  app.get('/api/purchaseGoods', function (req, res) {
    var walletid = req.query.walletid;
    var goodskey = req.query.goodskey
    let args = [walletid, goodskey];
    sdk.send(true, 'purchaseGoods', args, res);
  });
  
  app.get('/api/changeGoodsPrice', function (req, res) {
    var goodskey = req.query.goodskey;
    var price = req.query.price;
    let args = [goodskey, price];
    sdk.send(true, 'changeGoodsPrice', args, res);
  });
  
  app.get('/api/deleteGoods', function (req, res) {
    var goodskey = req.query.goodskey;
    let args = [goodskey];
    sdk.send(true, 'deleteGoods', args, res);
  });
}