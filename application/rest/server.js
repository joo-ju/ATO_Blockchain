const express = require("express");
const app = express();

var path = require("path");
var sdk = require("./sdk");

const PORT = 3001;
const HOST = "localhost";

app.get("/api/getWallet", function (req, res) {
  var walletid = req.query.walletid;

  let args = [walletid];

  sdk.send(false, "getWallet", args, res);
});
app.get("/api/setGoods", function (req, res) {
  var title = req.query.title;
  var content = req.query.content;
  var price = req.query.price;
  var category = req.query.category;
  var walletid = req.query.walletid;

  let args = [title, content, price, category, walletid];
  sdk.send(true, "setGoods", args, res);
});
app.get("/api/getAllgoods", function (req, res) {
  let args = [];
  sdk.send(false, "getAllGoods", args, res);
});
app.get("/api/purchaseGoods", function (req, res) {
  var walletid = req.query.walletid;
  var key = req.query.goodskey;

  let args = [walletid, key];
  sdk.send(true, "purchaseGoods", args, res);
});
app.use(express.static(path.join(__dirname, "./client")));

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
