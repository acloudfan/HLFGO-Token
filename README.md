# HLFGO-Token
Multiple versions of the chaincode

# Cloning
Clone it under gocc/src/token
git clone https://github.com/acloudfan/HLFGO-Token.git token

# Versions

# Setting up the environment
> . set-env.sh   acme  *OR*   > . set-env.sh   budget 
> set-chain-env.sh  -l golang  -p token/v1  -n token -v 1.0 -c '{"Args":["init"]}' -C airlinechannel 

{v1}
Both query & invoke execute the invoke function in chaincode
> set-chain-env.sh  -q '{"Args":["get"]}'
> set-chain-env.sh  -i '{"Args":["set"]}'

{v2}
Using the Logger API to log messages

{v3}
Getting the information in chaincode
- Transaction ID
- Timestamp
- Channel ID

{v4}
Getting the arguments in a chaincode
> set-chain-env.sh -q '{"Args":["FunctionName","Arg-1", "Arg-2"]}'
> chain.sh  query

{v5}
Shows the use of GetState | PutState | DelState
> set-chain-env.sh -q '{"Args":["get"]}'
> set-chain-env.sh -i '{"Args":["set"]}'
> set-chain-env.sh -i '{"Args":["delete"]}'

{v6}
Solution to the exercise. Create a new function - call it AddState

{v7}
Multiple tokens - functions will apply to the token - so pass an additional parameter
Solution to the exercise








# ERC20

> . set-env.sh   acme  *OR*   > . set-env.sh   budget 

Init
====
export CONSTRUCTOR='{"Args":["init","ACFT","A Cloud Fan Token!!!","raj"]}'
> set-chain-env.sh  -l golang  -p token/ERC20  -n erc20 -v 1.0 -c '{"Args":["init","ACFT","1000","A Cloud Fan Token!!!","raj"]}' -C airlinechannel 

set-chain-env.sh -c '$CONSTRUCTOR'

set-chain-env.sh -c '{"Args":["init","100","ACFT","A Cloud Fan Token!!!","raj"]}'

Queries
=======
> set-chain-env.sh -q '{"Args":["totalSupply"]}'
> set-chain-env.sh -q '{"Args":["balanceOf","sam"]}'

Invoke
======
> set-chain-env.sh -i '{"Args":["transfer","raj","sam","100"]}'
> set-chain-env.sh -i '{"Args":["transfer","sam","peter","50"]}'

