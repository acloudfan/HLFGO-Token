Client Identity Library
=======================
https://github.com/hyperledger/fabric/tree/release-1.1/core/chaincode/lib/cid
https://golang.org/pkg/crypto/x509/

Users in acme
=============
- mary
"app.accounting.role=tradeapprover:ecert","department=accounting:ecert"
- john
"app.accounting.role=accountant:ecert","department=accounting:ecert"
- anil
"department=logistics:ecert"

# Make sure the context is admin otherwise the instantiate will Fail
# with endrosement error
. set-env.sh acme
. set-ca-msp.sh admin

set-chain-env.sh -n cid -v 1.0  -p token/cid -c '{"Args":["init"]}'

chain.sh install
chain.sh instantiate

set-chain-env.sh  -q '{"Args": ["ReadAttributesOfCaller"]}'


Query
=====

set-chain-env.sh  -i '{"Args": ["AsssertOnCallersDepartment"]}'

. set-ca-msp.sh  mary
chain.sh invoke

. set-ca-msp.sh  john
chain.sh query

. set-ca-msp.sh  anil
chain.sh query

Invoke
======
. set-ca-msp.sh  mary
chain.sh invoke

. set-ca-msp.sh  john
chain.sh query

. set-ca-msp.sh  anil
chain.sh query

Exercise
========

+ Copy the cid/solution/ApproveTrade.go to folder cid/
+ Uncomment the call to ApproveTrade in the invoke function

If you already have the chaincode installed then upgrade:
. set-env.sh acme
. set-ca-msp.sh admin
chain.sh   upgrade-auto

If you do not have the chaincode installed then install/instantiate
. set-env.sh acme
. set-ca-msp.sh admin

set-chain-env.sh -n cid -v 1.0  -p token/cid -c '{"Args":["init"]}'

chain.sh install
chain.sh instantiate


In our test identity setup only mary & john are from dept accounting
so they are the only ones who can approve trade.
If trade is < 100K both john/mary can approve
If trade is >= 100K only mary can approve

set-chain-env.sh  -i '{"Args": ["ApproveTrade", "50000"]}'
. set-ca-msp.sh anil
chain.sh invoke   
. set-ca-msp.sh john
chain.sh invoke   

. set-ca-msp.sh mary
chain.sh invoke 

- Only Mary & John should be able to approve

set-chain-env.sh  -i '{"Args": ["ApproveTrade", "2000000"]}'
Do a chain.sh invoke   - Only mary should be able to approve

Anil should not be able to approve