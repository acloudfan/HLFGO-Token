# Private data
https://hyperledger-fabric.readthedocs.io/en/release-1.3/private_data_tutorial.html


Install & Instantiate
======================
Regular install process
Instantiate requires the collection.json to be specified
--collections-config

# Start the environment in Dev mode
dev-init.sh dev

# Launch the chaincode instance on Acme Peer
<Terminal#1>
. set-env.sh acme
cc-run.sh

# Launch the chaincode instance on Budget Peer
<Terminal#1>
. set-env.sh budget
cc-run.sh

<Terminal#3>
# Install the chaincode on Acme & Budget peers
. set-env.sh acme
chain.sh install
chain.sh instantiate-priv

. set-env.sh budget
chain.sh install

Test
====
Invalid collection name will lead to error



1. Acme can set both the public & private data
. set-env.sh acme
set-chain-env.sh -i '{"Args": ["Set","airlineOpen", "Acme has set the OPEN data"]}'
chain.sh invoke
set-chain-env.sh -i '{"Args": ["Set","acmePrivCollection", "Acme has set the SECRET data"]}'
chain.sh invoke

2. Acme can get both public and secret data
set-chain-env.sh -q '{"Args": ["Get"]}'
chain.sh query

3. Budget can set only the public but not the private
# Switch context to budget
. set-env.sh budget

# Change the parameters for invoke
set-chain-env.sh -i '{"Args": ["Set","airlineOpen", "Budget has set the OPEN data"]}'
chain.sh invoke

set-chain-env.sh -i '{"Args": ["Set","acmePrivCollection", "Budget has set the SECRET data"]}'
chain.sh invoke

4. Budget can get only the public data
chain.sh query
