#!/bin/bash
set -ev

docker-compose -f docker-compose-ca.yaml up -d ca.seller.ato.com

sleep 1
cd $GOPATH/src/ATO_Blockchain-master/application/sdk
node enrollAdmin.js
sleep 1
node registUsers.js