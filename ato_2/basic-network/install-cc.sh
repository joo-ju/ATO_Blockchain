#!/bin/bash
set -ev

# install chaincode for channelseller
docker exec cli peer chaincode install -n ato-cc-2 -v 1.0 -p chaincode/go
sleep 1

# instantiate chaincode for channelseller
docker exec cli peer chaincode instantiate -o orderer1.ato.com:7050 -C channelseller -n ato-cc-2 -v 1.0 -c '{"Args":[""]}' -P "OR ('SellerOrg.member','BuyerOrg.member')"
sleep 10

# invoke chaincode for channelseller
docker exec cli peer chaincode invoke -o orderer1.ato.com:7050 -C channelseller -n ato-cc-2 -c '{"function":"initWallet","Args":[""]}'
docker exec cli peer chaincode invoke -o orderer1.ato.com:7050 -C channelseller -n ato-cc-2 -c '{"function":"setGoods","Args":["Fabric", "state", "Hyper", "cate", "20", "content", "1Q2W3E4R"]}'
sleep 3

# check GD0
docker exec cli peer chaincode query -o orderer1.ato.com:7050 -C channelseller -n ato-cc-2 -c '{"function":"getGoods", "Args":["GD0"]}'
sleep 1

# purchase Goods
docker exec cli peer chaincode invoke -o orderer1.ato.com:7050 -C channelseller -n ato-cc-2 -c '{"function":"purchaseGoods","Args":["5T6Y7U8I", "GD0"]}'
sleep 3

docker exec cli peer chaincode query -o orderer1.ato.com: 7050 -C channelseller -n ato-cc-2 -c '{"function":"getWallet","Args":["1Q2W3E4R"]}'
docker exec cli peer chaincode query -o orderer1.ato.com: 7050 -C channelseller -n ato-cc-2 -c '{"function":"getWallet","Args":["5T6Y7U8I"]}'
sleep 1

# check GD0
docker exec cli peer chaincode query -o orderer1.ato.com:7050 -C channelseller -n ato-cc-2 -c '{"function":"getGoods", "Args":["GD0"]}'
sleep 1

# query chaincode for channelseller
docker exec cli peer chaincode query -o orderer.ato.com:7050 -C channelseller -n ato-cc-2 -c '{"function":"getAllGoods","Args":[""]}'
