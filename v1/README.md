Launch the environment
======================
Dev  mode:     dev-init.sh   dev  

Try out v1/token
================
#1  

<<Terminal #1>>  

- Setup the chaincode environment
set-chain-env.sh   -n  token   -p token/v1    -i '{"Args": ["init"]}'  


- Setup the organization context to acme
. set-env.sh acme

#2

<<Terminal #2>>

- Start the chaincode 

cc-run.sh

#3

<<Terminal #1>>

- Install & Instantiate the chaincode
chain.sh    install 
chain.sh    instantiate                             <<Observe terminal#2>>

Checkout explorer - you should see 1 transaction against the chaicode 'token'

#4

<<Terminal #1>>

- Setup and execute the invoke & query API
set-chain-env.sh  -i '{"Args": ["invoke"] }'
chain.sh  invoke                                    <<Observe terminal#2>>

set-chain-env.sh  -q '{"Args": ["invoke"] }' 
chain.sh  query

Checkout explorer - you should see 2 transaction against the chaicode 'token'