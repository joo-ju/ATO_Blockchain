#!/bin/bash
set -ev

docker-compose -f docker-compose-ca.yaml up -d ca.seller.ato.com

sleep 1
cd $GOPATH/src/ato/application/sdk
node enrollAdmin.js
sleep 1
node registUsers.js
sleep 1

node setGoods.js Fabric1 Sale Hyper Photocard 10 Photocard 1Q2W3E4R
sleep 3
node getAllGoods.js
sleep 3

node purchaseGoods.js 5T6Y7U8I GD1
sleep 5

node getWallet.js 1Q2W3E4R
node getWallet.js 5T6Y7U8I