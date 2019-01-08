#!/bin/bash

# Extend the test cases from V6

source ../v5/test.sh

#################################### Test 4 ###############################
# Last  TestCase in V5 has deleted the token, we will check if add succeeds

set_test_case   'Add Token should recerate the token with value=2000'
export CC_INVOKE_ARGS='{"Args":["add"]}'
chain_invoke

# A call to get should return -1
export CC_QUERY_ARGS='{"Args":["get"]}'
chain_query 
MY_TOKEN_VALUE="$QUERY_RESULT" 
assert_number  $MY_TOKEN_VALUE "-eq"  "2000"