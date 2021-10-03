#!/bin/bash
set -ev

# install chaincode for channelseller
docker exec cli peer chaincode install -n ato-cc -v 1.0 -p chaincode/go
sleep 1

# instantiate chaincode for channelseller
docker exec cli peer chaincode instantiate -o orderer.ato.com:7050 -C channelseller -n ato-cc -v 1.0 -c '{"Args":[""]}' -P "OR ('SellerOrg.member','BuyerOrg.member')"
sleep 10

# invoke chaincode for channelseller
docker exec cli peer chaincode invoke -o orderer.ato.com:7050 -C channelseller -n ato-cc -c '{"function":"initWallet","Args":[""]}'
docker exec cli peer chaincode invoke -o orderer.ato.com:7050 -C channelseller -n ato-cc -c '{"function":"setGoods","Args":["Fabric", "state", "Hyper", "cate", "20", "content", "1Q2W3E4R"]}'
sleep 3

docker exec cli peer chaincode invoke -o orderer.ato.com:7050 -C channelseller -n ato-cc -c '{"function":"purchaseGoods","Args":["5T6Y7U8I", "GD0"]}'
sleep 3

docker exec cli peer chaincode query -o orderer.ato.com: 7050 -C channelseller -n ato-cc -c '{"function":"getWallet","Args":["1Q2W3E4R"]}'
docker exec cli peer chaincode query -o orderer.ato.com: 7050 -C channelseller -n ato-cc -c '{"function":"getWallet","Args":["5T6Y7U8I"]}'
sleep 1

# query chaincode for channelseller
docker exec cli peer chaincode query -o orderer.ato.com:7050 -C channelseller -n ato-cc -c '{"function":"getAllGoods","Args":[""]}'