#!/bin/bash
set -ev

# Shut down the Docker containers that might be currently running.
docker-compose -f docker-compose.yaml stop
# docker-compose -f docker-compose-ca.yaml stop
docker-compose stop
docker-compose -f docker-compose.yaml kill && docker-compose -f docker-compose.yaml down --volumes --remove-orphans
docker-compose -f kill && docker-compose down --volumes --remove-orphans

# remove the local state
# rm -rf $GOPATH/src/ato/application/wallet/

# Your system is now clean
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)

# remove chaincode docker images
docker rmi -f $(docker images dev-* -q)
