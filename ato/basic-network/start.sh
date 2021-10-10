#!/bin/bash
set -ev

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

docker-compose -f docker-compose.yaml down
docker-compose -f docker-compose.yaml up -d orderer.ato.com peer0.seller.ato.com peer1.seller.ato.com peer0.buyer.ato.com peer1.buyer.ato.com cli

# wait for Hyperledger Fabric to start
# incase of errors when running later commands, issue export FABRIC_START_TIMEOUT=<larger number>
export FABRIC_START_TIMEOUT=10
sleep ${FABRIC_START_TIMEOUT}

# Create the channel
docker exec cli peer channel create -o orderer.ato.com:7050 -c channelseller -f /etc/hyperledger/configtx/channel1.tx

# Join peer0.seller.ato.com to the channel and Update the Anchor Peers in Channel1
docker exec cli peer channel join -b channelseller.block
docker exec cli peer channel update -o orderer.ato.com:7050 -c channelseller -f /etc/hyperledger/configtx/SellerOrganchors.tx

# Join peer1.seller.ato.com to the channel
docker exec -e "CORE_PEER_ADDRESS=peer1.seller.ato.com:7051" cli peer channel join -b channelseller.block

# Join peer0.buyer.ato.com to the channel and update the Anchor Peers in Channel1
docker exec -e "CORE_PEER_LOCALMSPID=BuyerOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/buyer.ato.com/users/Admin@buyer.ato.com/msp" -e "CORE_PEER_ADDRESS=peer0.buyer.ato.com:7051" cli peer channel join -b channelseller.block
docker exec -e "CORE_PEER_LOCALMSPID=BuyerOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/buyer.ato.com/users/Admin@buyer.ato.com/msp" -e "CORE_PEER_ADDRESS=peer0.buyer.ato.com:7051" cli peer channel update -o orderer.ato.com:7050 -c channelseller -f /etc/hyperledger/configtx/BuyerOrganchors.tx

# Join peer0.buyer.ato.com to the channel
docker exec -e "CORE_PEER_LOCALMSPID=BuyerOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/buyer.ato.com/users/Admin@buyer.ato.com/msp" -e "CORE_PEER_ADDRESS=peer1.buyer.ato.com:7051" cli peer channel join -b channelseller.block
