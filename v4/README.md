# Demonstrates the use of Argument functions on the stub

To try out:

<<Terminal#1>>
dev-init.sh dev


<<Terminal#2>>
set-chain-env.sh   -n token   -v 1.0   -p token/v4
cc-run.sh

<<Terminal#1>>
chain.sh   install

set-chain-env.sh   -c   '{"Args":["init"]}'
chain.sh   instantiate

# Empty array is fine
# "args" instead of "Args" is fine too
set-chain-env.sh -i '{"args":[]}'
chain.sh   invoke

# 1
set-chain-env.sh   -i    '{"args":["func-name"]}'
chain.sh   invoke

# 2
set-chain-env.sh   -i    '{"args":["func-name","param-1","param-2"]}'
chain.sh   invoke


