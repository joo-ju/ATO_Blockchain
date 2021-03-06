#!/bin/bash
set -ev

# install chaincode for channelseller
docker exec cli peer chaincode install -n atocc -v 1.0 -p chaincode/go
sleep 1
# instantiate chaincode for channelseller
docker exec cli peer chaincode instantiate -o orderer.ato.com:7050 -C channelseller -n atocc -v 1.0 -c '{"Args":[""]}' -P "OR ('SellerOrg.member','BuyerOrg.member')"
sleep 10
# invoke chaincode for channelseller
docker exec cli peer chaincode invoke -o orderer.ato.com:7050 -C channelseller -n atocc -c '{"function":"initWallet","Args":[""]}'
docker exec cli peer chaincode invoke -o orderer.ato.com:7050 -C channelseller -n atocc -c '{"function":"setGoods","Args":["제목", "내용", "20", "카테고리", "1Q2W3E4R"]}'
sleep 3
# query chaincode for channelseller
docker exec cli peer chaincode query -o orderer.ato.com:7050 -C channelseller -n atocc -c '{"function":"getAllGoods","Args":[""]}'