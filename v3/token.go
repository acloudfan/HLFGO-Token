package main
/**
 * tokenv3
 * Shows the use of ChaincodeStub API
 *    A) Transaction information
 *    B) Transaction timestamp
 *    C) Channel ID
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
	fmt.Println("Invoke executed ")

	// V3
	fmt.Printf("GetTxID() => %s\n", stub.GetTxID())

	// V3
	TxTimestamp, _ := stub.GetTxTimestamp()
	fmt.Printf("GetTxTimestamp() => %s\n", time.Unix(TxTimestamp.GetSeconds(),0))

	// V4
	fmt.Println("GetChannelID() =>", stub.GetChannelID())

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