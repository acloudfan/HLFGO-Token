#!/bin/bash

# Initialize
echo "Initializing ..."
govendor init
govendor add +external

# Add the dependencies
echo "Adding dependencies ... please wait"
govendor add github.com/hyperledger/fabric/core/chaincode/shim
govendor add github.com/hyperledger/fabric/protos/peer
govendor add github.com/hyperledger/fabric/protos/ledger/queryresult

