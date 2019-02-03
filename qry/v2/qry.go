package main

/**
 * v2 shows the use of Composite Key functions
 **/

import (
	// For printing messages on console
	"fmt"

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// peer.Response is in the peer package
	"github.com/hyperledger/fabric/protos/peer"

	// KV Interface
	"github.com/hyperledger/fabric/protos/ledger/queryresult"

	// JSON Encoding
	"encoding/json"

	// Conversion functions
	"strconv"

	// Used for converting []string to string
	"strings"
)

// QueryChaincode Represents our chaincode object
type QueryChaincode struct {
}

// TokenBalance represent balances of different type of tokens
type TokenBalance struct {
	Symbol      string `json:"symbol"`		// symbol of the token
	Owner       string `json:"owner"`		// unique identity of the owner
	Country		string `json:"country"`		// country of residence
	Balance     uint   `json:"balance"`		// balance of the token
}

const indexBalance = "owner~country~symbol"

// Init Implements the Init method
func (token *QueryChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

		// Simply print a message
	fmt.Println("Init executed in qry")


	// Add the data
	token.SetupSampleData(stub)

	// Return success
	return shim.Success(nil)
}

// Invoke method
func (token *QueryChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the function name and parameters
	funcName, args := stub.GetFunctionAndParameters()

	if funcName == "getStateOnKey" {
		return token.GetTokenByCompositeKey(stub, args)
	} else if funcName == "getStateRangeOnKey" {
		return token.GetTokensByPartialCompositeKey(stub, args)
	}
	
	// This is not good
	return shim.Error(("Bad Function Name = !!!"))
}

func (token *QueryChaincode) GetTokensByPartialCompositeKey(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	
	if len(args) == 0 {
		return shim.Error("Must provide args for the key!!")
	}
	

	fmt.Printf("Exec qry with ")

	QryIterator, err := stub.GetStateByPartialCompositeKey(indexBalance,args)
	if err != nil {
		fmt.Printf("Error in getting by range=" + err.Error())
		return shim.Error(err.Error())
	}
	var resultJSON = "["
	counter := 0
	for QryIterator.HasNext() {
		// Hold pointer to the query result
		var resultKV *queryresult.KV
		var err error

		// Get the next element
		resultKV, err = QryIterator.Next()
		if err != nil {
			fmt.Println("Err=" + err.Error())
		} else {
			
			key,arr, _ := stub.SplitCompositeKey(resultKV.GetKey())
			fmt.Println(key)
			fmt.Println(arr)
			resultJSON += " ["+strings.Join(arr,"~")+"] "
			counter++
		}
	}
	resultJSON += "]"
	resultJSON = "Counter="+strconv.Itoa(counter)+"  "+resultJSON
	fmt.Println("Done.")
	return shim.Success([]byte(resultJSON))
}


// GetTokenByCompositeKey executes the Range query on state data
func (token *QueryChaincode) GetTokenByCompositeKey(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("MUST provide start & end Key!!")
	}
	
	startKey, errkey := stub.CreateCompositeKey(indexBalance, args)
	if errkey != nil {
		fmt.Printf("Error in creating key =" + errkey.Error())
		return shim.Error(errkey.Error())
	}

	fmt.Println("-->")
	fmt.Println([]byte(startKey))

	// var counter = 0
	var resultJSON = "["
	
	dat, _ := stub.GetState(startKey)

	fmt.Println("dat="+string(dat))

	return shim.Success([]byte(resultJSON))
}

// SetupSampleData creates multiple instances of the ERC20Token
func (token *QueryChaincode) SetupSampleData(stub shim.ChaincodeStubInterface)  {
	
	value := []byte{0x00}
	stub.PutState(indexBalance, value)

	addData(stub, "BTC", "john", "USA", 26)
	addData(stub, "ETH", "john", "USA", 4)
	addData(stub, "ETH", "john", "UK", 21)
	
	addData(stub, "BTC", "sam", "USA", 42)
	addData(stub, "BTC", "kim", "UK", 12)
	addData(stub, "ETH", "sam", "USA", 1)
	

	fmt.Println("Added the data!!")
}

func addData(stub shim.ChaincodeStubInterface, symbol, owner, country string, balance uint ) {
	tokbal := TokenBalance{Symbol: symbol, Owner: owner,  Country: country, Balance: balance}
	jsonTokbal, _ := json.Marshal(tokbal)
	balanceIndexKey, _ := stub.CreateCompositeKey(indexBalance, []string{tokbal.Owner, tokbal.Country, tokbal.Symbol})
	stub.PutState(balanceIndexKey, jsonTokbal)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode. token/qry/v1\n")
	err := shim.Start(new(QueryChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}