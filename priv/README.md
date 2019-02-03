# Private data
https://hyperledger-fabric.readthedocs.io/en/release-1.3/private_data_tutorial.html


Install & Instantiate
=====================
dev-init.sh -e

. set-env.sh acme

reset-chain-env.sh
set-chain-env.sh  -n priv -v 1.0 -p token/priv -c '{"Args": ["init"]}' -C airlinechannel
set-chain-env.sh -R pcollection.json

Exercise
=========
Install & Instantiate the token/priv chaincode using "peer chaincode instantiate .." command

<Solution>
Setup the environment variables
chain.sh install
. cc.env.sh
peer chaincode instantiate -C "$CC_CHANNEL_ID" -n "$CC_NAME"  -v "$CC_VERSION" -c "$CC_CONSTRUCTOR" -o "$ORDERER_ADDRESS"  --collections-config "$GOPATH/src/token/priv/pcollection.json"


Test the setup
==============
1. Start the Environment

# Start the environment in net mode
dev-init.sh -e

reset-chain-env.sh

set-chain-env.sh  -n priv -v 1.0 -p token/priv -c '{"Args": ["init"]}' -C airlinechannel
# Use the -R option to set the PDC
# At instantiation chain.sh will specify the full path to PDC collection
set-chain-env.sh -R pcollection.0.json

Install & Instantiate
. set-env.sh acme
chain.sh install
chain.sh instantiate

. set-env.sh budget
chain.sh install

2. Invoke the Set as ACME & Query

<Terminal#1>

# Invoke to set the value for 2 tokens
. set-env.sh acme
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

set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Budget has set the OPEN data"]}'
chain.sh invoke

set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Budget has set the SECRET data"]}'
chain.sh invoke

# Get the value for 2 tokens - Budget will NOT seet the value for protected token
chain.sh query         

4. Query as ACME
<Terminal#1>
. set-env.sh acme
chain.sh query  

. set-env.sh acme

Exercise
========
Extend the priv chaincode - add a function to delete the key in specific collection

. set-env.sh acme
set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Acme has set the OPEN data"]}'
chain.sh invoke
set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Acme has set the SECRET data"]}'
chain.sh invoke


chain.sh query
set-chain-env.sh -i '{"Args": ["Del", "AcmeBudgetOpen"]}'
chain.sh invoke

chain.sh query


Experimental
============
set-chain-env.sh -i '{"Args": ["Del", "MemberOnlyTest"]}'

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
