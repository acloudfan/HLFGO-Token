InvokeChaincode
===============
Shows the use of Invocation of a chaincode from another chaincode

Testing in Dev mode
===================

#Terminal 1
- Set env context
  .  set-env.sh   acme
- Launch the environment
  dev-init.sh   dev
- Set the chaincode env for 'token'
  set-chain-env.sh  -n token   -v 1.0  -p token/v5   -c '{"args":[]}'
- Launch the 'token' chaincode
  cc-run.sh

#Terminal 2
- Set env context
  .  set-env.sh   acme
- Install & Instantiate the 'token' chaincode
  chain.sh install
  chain.sh instantiate

- Set the chaincode env for 'caller'
  set-chain-env.sh  -n caller   -v 1.0  -p token/v6   -c '{"args":[]}'
- Launch the 'caller' chaincode
  cc-run.sh

#Terminal 3
- Set env context
  .  set-env.sh   acme
- Install & Instantiate the 'caller' chaincode
  chain.sh install
  chain.sh instantiate

Now we are ready to test
- Setup the query & invoke parameters
  set-chain-env.sh   -i '{"args":["setOnCaller"]}'
  set-chain-env.sh   -q '{"args":["getOnCaller"]}'

- To test execute invoke | query on caller
  chain.sh query
  chain.sh invoke