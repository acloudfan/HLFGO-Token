Requires CouchDB to be enabled

Chaincode Setup
===============
# Launch the environment with CouchDB enabled
dev-init.sh -d -s

# Deploy the chaincode
. set-env.sh  acme
reset-chain-env.sh
set-chain-env.sh -n CryptocurrencyTxn -v 1.0 -p token/qry/v3 -c '{"Args":[]}'
cc-run.sh

chain.sh install
chain.sh instantiate

# Upload the data (It may take 15+ minutes)
. set-env.sh  acme
- Change directory to qry/v3/data
./setup-data.sh

ExecuteRichQuery
================
set-chain-env.sh -q '{"Args":["ExecuteRichQuery","{\"selector\":{\"txnDate\": \"2009-12-12T00:00:00Z\"}}"]}'
chain.sh query

# Utility for executing queries:
PS: Remember query attributes limit & skip are ignored by Fabric

qry/v3/samples
./run-query.sh    #NUMBER#      Executes the query against chaincode
./run-query.sh    1.1           Executes the query in file:  query-1.1.json

GetDatesByPrice
===============
- Get the dates on which the price of crypto was $19200 and above
set-chain-env.sh -q '{"Args":["GetDatesByPrice","19200"]}'
chain.sh query

- Get the dates on which the price of crypto was $15000 and above
set-chain-env.sh -q '{"Args":["GetDatesByPrice","15000"]}'
chain.sh query

GetAveragesBetweenDates (Exercise)
==================================
set-chain-env.sh -q '{"Args":["GetAveragesBetweenDates","2009-01-01T00:00:00Z","2019-02-15T00:00:00Z"]}'
chain.sh query

set-chain-env.sh -q '{"Args":["GetAveragesBetweenDates","2017-12-01T00:00:00Z","2018-06-31T00:00:00Z"]}'
chain.sh query

set-chain-env.sh -q '{"Args":["GetAveragesBetweenDates","2017-12-01T00:00:00Z","2018-01-31T00:00:00Z"]}'
chain.sh query



Queries/Selectors
=================
http://docs.couchdb.org/en/latest/api/database/find.html#find-selectors
Query => A JSON object with standard "attribute" names for specifying the query criteria
"selector": Object specifying the document selection criteria
            Operators:
            Selectors may use the operators for defining the criteria
            $eq     Equal
            $gt     Greater than
            $lt     Less than

"fields":   An array of field names that need to be returned
"sort":     An array of field names for sorting result set


Data on Bitcoin
===============
#1 Downloaded the data from this site in csv format
https://coinmetrics.io/data-downloads/

#2 Remove the (USD) from headers - Replaced (USD)  with ""
Open in VSC editor and just replace

#3 Converted the data from csv to json format
http://www.convertcsv.com/csv-to-json.htm



CouchDB Indexes
===============
http://docs.couchdb.org/en/stable/api/database/find.html#db-index
