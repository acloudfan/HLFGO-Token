. set-env.sh acme

reset-chain-env.sh
set-chain-env.sh   -n qry  -v 1.0 -p token/qry/v2 -c '{"Args": ["init","1","10"]}'


set-chain-env.sh -q '{"Args": ["getStateOnKey", "", "",""]}' 
chain.sh query

set-chain-env.sh -q '{"Args": ["getStateRangeOnKey", "john"]}' 
chain.sh query

set-chain-env.sh -q '{"Args": ["getStateRangeOnKey", "john", "USA"]}' 
chain.sh query

set-chain-env.sh -q '{"Args": ["getStateRangeOnKey", "john", "USA","BTC"]}' 
chain.sh query

set-chain-env.sh -q '{"Args": ["getStateRangeOnKey",""]}' 
chain.sh query