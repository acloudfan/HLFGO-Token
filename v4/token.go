package main
/**
 * tokenv4
 * Shows the use of ChaincodeStub API for getting the arguments sent from client
 * set-chain-env.sh -q '{"Args":["FunctionName","Arg-1", "Arg-2"]}'
 **/
import (
	"fmt"

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"
	// peer.Response is in the peer package
	"github.com/hyperledger/fabric/protos/peer"
	// Used for formatting the timestamp
	"time"
)

// TokenChaincode Represents our chaincode object
type TokenChaincode struct {
}


// Init Implements the Init method
func (token *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed")

	// Return success
	return shim.Success(nil)
}

// Invoke method
func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	TxTimestamp, _ := stub.GetTxTimestamp()
	fmt.Printf("Invoke executed TxID=%s   TxTimestamp=%s\n", stub.GetTxID(), time.Unix(TxTimestamp.GetSeconds(),0))

	// V4
	// Get the args in a 2D array of byte
	argsArray := stub.GetArgs()
	for ndx, arg := range argsArray {
		fmt.Printf("[%d]=%s  ", ndx, string(arg))
	}
	fmt.Println()

	// V4
	// Get the Args[] sent by the client
	fmt.Println(stub.GetStringArgs())
	

	// V4
	argsSlice,_ := stub.GetArgsSlice()
	fmt.Println(string(argsSlice))


	// V4
	// Get the function & parameters
	funcName, args := stub.GetFunctionAndParameters()
	fmt.Printf("Function=%s  Args=%s\n", funcName, args)

	return shim.Success(nil)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode.")
	err := shim.Start(new(TokenChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}