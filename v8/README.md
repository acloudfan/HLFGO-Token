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
+ If in dev mode then run the CC
+ Launch the event listener chain events
+ Execute invoke/set on chaincode
+ Observe events on the tool everytime set is invoked



