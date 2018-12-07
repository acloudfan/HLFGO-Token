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
> set-chain-env.sh  -q '{"Args":["get"]}'
> set-chain-env.sh  -i '{"Args":["set"]}'

{v2}
> 

# ERC20

> . set-env.sh   acme  *OR*   > . set-env.sh   budget 

Init
====
export CONSTRUCTOR='{"Args":["init","ACFT","A Cloud Fan Token!!!","raj"]}'
> set-chain-env.sh  -l golang  -p token/ERC20  -n erc20 -v 1.0 -c '{"Args":["init","ACFT","1000","A Cloud Fan Token!!!","raj"]}' -C airlinechannel 

 set-chain-env.sh -c '$CONSTRUCTOR'

  set-chain-env.sh -c '{"Args":["init","100","ACFT","ACloudFanToken!!!","raj"]}'

Queries
=======
> set-chain-env.sh -q '{"Args":["totalSupply"]}'
> set-chain-env.sh -q '{"Args":["balanceOf","sam"]}'

Invoke
======
> set-chain-env.sh -i '{"Args":["transfer","raj","sam","100"]}'


