/**
 * This implements the ERC20 standard token
 * https://theethereum.wiki/w/index.php/ERC20_Token_Standard
 **/
 package main

 import (
	"fmt"

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// peer.Response is in the peer package
	"github.com/hyperledger/fabric/protos/peer"

	// Conversion functions
	// "strconv"
)

// ERC20TokenChaincode Represents our chaincode object
type ERC20TokenChaincode struct {
}

// ERC20Token 
type ERC20Token struct {
	Symbol   		string   `json:"symbol"`
	TotalSupply     int      `json:"totalSupply"`
	Description		string   `json:"description"`
}


// Init Implements the Init method
// Receives 3 parameters = [0] Symbol [1] TotalSupply   [2] Description
func (token *ERC20TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed")
	args := stub.GetArgs()

	// Check if we received the right number of arguments
	if len(args) < 3 {
		return shim.Error("Failed - incorrect number of parameters!! ")
	}

	var erc20 ERC20Token = ERC20Token{Symbol: string(args[0])}

	stub.PutState("token", []byte(erc20))

	return shim.Success([]byte("Created the token!!"))
}

