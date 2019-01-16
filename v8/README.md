Shows the use of event....
Test tool for event

Start with v5
Update it such that whenever the value of token is changed it emits the event.
+ Event =  TokenValueChanged
+ Payload should be in JSON format
  { "value": NEW_VALUE }

Testing
=======
+ Launch the env - either mode fine
+ Install & Instantiate
    . set-env.sh   acme
    set-chain-env.sh  -n token -v 1.0 -p token/v8 -c '{"Args": ["init"]}'
    chain.sh install
    chain.sh instantiate
+ If in dev mode then run the CC
+ <Terminal#1> Launch the event listener chain events
    events.sh -t chaincode -n token -e TokenValueChanged -c airlinechannel 
+ <Terminal#2> Execute invoke/set on chaincode
    set-chain-env.sh   -i   '{"args":["set"]}' -q   '{"args":["get"]}'
    chain.sh invoke
+ Observe events on the tool everytime set is invoked




