# Shows the use of PutState | GetState | DelState

# Test the original token.go 
1. Launch the env in dev mode       dev-init.sh dev
2. Setup the env context            .  set-env.sh acme
3. Setup the chain env              set-chain-env.sh  -n token -v 1.0 -p token/v5  -c '{"args":[]}' 
4. Run the chaincode                cc-run.sh
5. In a new terminal set the env context     .  set-env.sh acme
6. Install                          chain.sh install
7. Instantiate                      chain.sh instantiate
8. Setup invoke & query args        set-chain-env.sh   -i   '{"args":["set"]}' -q   '{"args":["get"]}' 

To Test execute invoke/query .... everytime invoke is called value will increment by 10
9. Execute invoke                   chain.sh  invoke
10. Execute query                   chain.sh  query


# Using the test.sh file
# Steps to be followed only after the 'Token delete func added'
The test may be run net & dev mode. 
Prior to running the test install & instantiate the chaincode
1. Launch the setup in net mode
   dev-init.sh net

2. Set the org context
   . set-env.sh acme

3. Set the chaincode env
   set-chain-env.sh -n token -v 1.0 -p token/v5

4. Install & Instantiate
   chain.sh  install

   set-chain-env.sh  -c  '{"args":[]}'
   chain.sh  instantiate

5. Launch the test
   cc-test.sh

