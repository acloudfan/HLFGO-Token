package main

/**
 * Demonstrates the use of Private Data Collections
 **/
import (
	"fmt"

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// peer.Response is in the peer package
	"github.com/hyperledger/fabric/protos/peer"
)

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

	// Minimum of 2 args is needed - skipping the check
	// params[0]=Collection name
	// params[1]=Value for the token

	err := stub.PutPrivateData(params[0], "token", []byte(params[1]))
	if err != nil {
		return shim.Error("Error=" + err.Error())
	}

	return shim.Success([]byte("true"))
}

// Gets the value of "token" from both the collections
func (privCode *PrivChaincode) Get(stub shim.ChaincodeStubInterface) peer.Response {

	// This is returned
	resultString := "{}"

	// Read the open data
	dataOpen, err1 := stub.GetPrivateData("AcmeBudgetOpen", "token")
	if err1 != nil {
		return shim.Error("Error1=" + err1.Error())
	}

	// Read the acme private data
	dataSecret, err2 := stub.GetPrivateData("AcmePrivate", "token")

	accessError := "N.A."
	if err2 != nil {
		//return shim.Error("Error="+err1.Error())
		fmt.Println("Error2=" + err2.Error())
		accessError = err2.Error()
		dataSecret = []byte("**** Not Allowed ***")
	}

	// resultString = "{open:\""+string(dataOpen)+"\", secret:\""+string(dataSecret)+"\"}"
	// resultString = "{\"open\":\""+string(dataOpen)+"\", \"secret\":\""+string(dataSecret)+"\",\"error\":\""+err2.Error()+  "\"}"
	resultString = "{open:\"" + string(dataOpen) + "\", secret:\"" + string(dataSecret) + "\" , error:\"" + accessError + "\"}"

	return shim.Success([]byte(resultString))
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode. priv\n")
	err := shim.Start(new(PrivChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
