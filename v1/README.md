Launch the environment
======================
Dev  mode:     dev-init.sh   dev        (Suggested)
Net  mode:     dev-init.sh   net


Try out v1/token
================

> Setup the chaincode environment
set-chain-env.sh   -n  token   -p token/v1    -i '{"Args": ["init"]}'  

<<Terminal #1>>
> Setup the organization context
. set-env.sh acme

<<Terminal #2>>
> Start the chaincode - ONLY if env started in dev mode
cc-run.sh

<<Terminal #1>>
> Install the chaincode
chain.sh    install 
chain.sh    instantiate 

> Setup the invoke & query parameters
set-chain-env.sh  -i '{"Args": ["invoke"] }'
set-chain-env.sh  -q '{"Args": ["invoke"] }' 