package main

/**
 * Demonstrates the use of Private Data Collections
 * Path:   token/priv
 * Requires the creation of the PDC definition using the pcollection.json 
 **/
import (
	"fmt"

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// peer.Response is in the peer package
	"github.com/hyperledger/fabric/protos/peer"
)

// CollectionName is the name of PDC
const CollectionName = "PrivateToken"

// PrivChaincode Represents our chaincode object
type PrivChaincode struct {
}

// Init Implements the Init method
func (privCode *PrivChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed")

	// Return success
	return shim.Success([]byte("true"))
}

// Invoke method
func (privCode *PrivChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the func name and parameters
	funcName, params := stub.GetFunctionAndParameters()

	fmt.Printf("funcName=%s  Params=%s \n", funcName, params)

	if funcName == "Set" {

		return privCode.Set(stub, params)

	} else if funcName == "Get" {

		return privCode.Get(stub)

	} 

	return shim.Error("Invalid Function Name: " + funcName)
}

// Set function
func (privCode *PrivChaincode) Set(stub shim.ChaincodeStubInterface, params []string) peer.Response {

	// Minimum of 2 args is needed - skipping the check for clarity
	// params[0]= Value for the token

	 
	tokenValue := params[0]

	err := stub.PutPrivateData(CollectionName, "token", []byte(tokenValue))
	if err != nil {
		return shim.Error("Error=" + err.Error())
	}

	return shim.Success([]byte("true"))
}

// Get gets the value of token from both collections
func (privCode *PrivChaincode) Get(stub shim.ChaincodeStubInterface) peer.Response {

	// Read the open data
	data, err := stub.GetPrivateData(CollectionName, "token")
	if err != nil {
		return shim.Error("Error=" + err.Error())
	}

	return shim.Success([]byte(data))
}


// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode. priv-token\n")
	err := shim.Start(new(PrivChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
