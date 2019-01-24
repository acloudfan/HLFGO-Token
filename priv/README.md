# Private data
https://hyperledger-fabric.readthedocs.io/en/release-1.3/private_data_tutorial.html

1. Start the Environment

# Start the environment in net mode
dev-init.sh -e

reset-chain-env.sh

set-chain-env.sh  -n priv -v 1.0 -p token/priv -c '{"Args": ["init"]}' -C airlinechannel
# Use the -R option to set the PDC
# At instantiation chain.sh will specify the full path to PDC collection
set-chain-env.sh -R pcollection.json

2. Invoke the Set as ACME & Query

<Terminal#1>
. set-env.sh acme
chain.sh install
chain.sh instantiate-priv
# Invoke to set the value for 2 tokens
set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Acme has set the OPEN data"]}'
chain.sh invoke
set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Acme has set the SECRET data"]}'
chain.sh invoke
# Get the value for 2 tokens
set-chain-env.sh -q '{"Args": ["Get"]}'
chain.sh query

3. Invoke the Set as BUDGET & Query

<Terminal#2>
. set-env.sh budget
chain.sh install
# Get the value for 2 tokens
chain.sh query         

set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Budget has set the OPEN data"]}'
chain.sh invoke

set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Budget has set the SECRET data"]}'
chain.sh invoke

# Get the value for 2 tokens - Budget will NOT seet the value for protected token
chain.sh query         

4. Query as ACME
<Terminal#1>
chain.sh query  


Testing in Dev Mode
====================
Use the instructions below to test the PDC in in DEV mode

Install & Instantiate
======================
Regular install process
Instantiate requires the collection.json to be specified
--collections-config

# Start the environment in Dev mode
dev-init.sh dev
set-chain-env.sh  -n priv -v 1.0 -p token/priv -c '{"Args": ["init"]}' 
# Use the -R option to set the PDC
# At instantiation chain.sh will specify the full path to PDC collection
set-chain-env.sh -R pcollection.json

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
set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Acme has set the OPEN data"]}'
chain.sh invoke
set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Acme has set the SECRET data"]}'
chain.sh invoke

2. Acme can get both public and secret data
set-chain-env.sh -q '{"Args": ["Get"]}'
chain.sh query

3. Budget can the public & the private data
# Switch context to budget
. set-env.sh budget

# Change the parameters for invoke
set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Budget has set the OPEN data"]}'
chain.sh invoke

set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Budget has set the SECRET data"]}'
chain.sh invoke

4. Budget can get only the public data
chain.sh query
