# Shows the use of Logging Interface from chaincode

<<Terminal #1>>
# Launch env in net mode
dev-init.sh

# Set up chaincode environment
. set-env.sh  acme
set-chain-env.sh -n token  -v 1.0  -p token/v2 -c '{"Args":["init"]}'
set-chain-env.sh -L ERROR


# Install the chaincode
chain.sh install
chain.sh instantiate

<<Terminal #2>>

