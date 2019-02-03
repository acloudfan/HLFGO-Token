Demonstrates Query by Range
===========================

Testing Range functions (token/qry/v1)
======================================
1. Launch dev environment in either mode
dev-init.sh

2. Setup the  env
. set-env.sh acme

3. Setup chaincode env. The init function takes args for setting sample data
reset-chain-env.sh
set-chain-env.sh   -n qry  -v 1.0 -p token/qry/v1 -c '{"Args": ["init","1","10"]}'
set-chain-env.sh   -q '{"Args": ["GetTokenByRange", "", ""]}' 

4. Install and instantiate
chain.sh install
chain.sh instantiate

5. Try out the queries now
set-chain-env.sh -q '{"Args": ["GetTokenByRange", "key1", "key3"]}' 
chain.sh query

# All data in chunks of 5 / page
set-chain-env.sh -q '{"Args": ["GetTokenByRangeWithPagination", "", "","5"]}' 


Query Result Interface
======================
https://godoc.org/github.com/hyperledger/fabric/protos/ledger/queryresult

