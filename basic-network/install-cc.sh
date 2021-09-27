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
docker exec cli peer chaincode invoke -o orderer.ato.com:7050 -C channelseller -n ato-cc -c '{"function":"setGoods","Args":[ "Fabric", "Hyper", "20", "cate", "1Q2W3E4R"]}'
sleep 3
# query chaincode for channelseller
docker exec cli peer chaincode query -o orderer.ato.com:7050 -C channelseller -n ato-cc -c '{"function":"getAllGoods","Args":[""]}'