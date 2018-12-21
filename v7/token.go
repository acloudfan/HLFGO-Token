package main
/**
 * tokenv6
 * Adds the token
 **/
import (
	"fmt"

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"
	// peer.Response is in the peer package
	"github.com/hyperledger/fabric/protos/peer"

	// Conversion functions
	"strconv"
)

// TokenChaincode Represents our chaincode object
type TokenChaincode struct {
}


// Init Implements the Init method
func (token *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// V6
	return AddToken(stub)
}

// Invoke method
func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	funcName, _ := stub.GetFunctionAndParameters()
	fmt.Println("Function=", funcName)


	// V5
	if(funcName == "set"){
		// Sets the value
		return SetToken(stub)

	} else if(funcName == "get"){

		// Gets the value
		return GetToken(stub)

	} else if(funcName == "delete"){

		// Delete the token
		return DeleteToken(stub)
	}  else if(funcName == "add"){	//V6

		// Add the token
		return AddToken(stub)
	}
	
	// This is not good
	return shim.Error(("Bad Function Name = "+funcName+"!!!"))
}


// SetToken sets the value of the token
// V5
// Returns true if successful
func SetToken(stub shim.ChaincodeStubInterface) peer.Response{
	value, err := stub.GetState("MyToken")
	if(err != nil){
		fmt.Println("MyToken Not found!!!")
		return  shim.Error("Token not found in State Database!!!")
	}
	// Otherwise increment the value by 10
	intValue, err :=  strconv.Atoi(string(value))

	if err != nil {
		// May also return sh.Error 
		return shim.Success([]byte("false"))
	}
	intValue += 10
	stub.PutState("MyToken", []byte(strconv.Itoa(intValue)))

	return shim.Success([]byte("true"))
}

// GetToken reads the value of the token from the database
// Reurns the value or -1 in case MyToken doesn't exist
func  GetToken(stub shim.ChaincodeStubInterface) peer.Response {
	var myToken string

	if value, err := stub.GetState("MyToken"); err != nil {
		fmt.Println("Get Failed!!! ", err.Error())
		return shim.Error(("Get Failed!! "+err.Error()+"!!!"))
	} else {
		// nil indicates non existent key
		if (value == nil) {
			myToken = "-1"
		} else {
			myToken = string(value)
		}
	}
	
	return shim.Success([]byte(myToken))
}

// DeleteToken deletes the token from the database
// Returns 0 if successful
func  DeleteToken(stub shim.ChaincodeStubInterface) peer.Response {

	// Check if the key exists - if not then return false
	value, _ := stub.GetState("MyToken")
	if value == nil {
		return shim.Success([]byte("false"))
	}
	// Delete the key
	if err := stub.DelState("MyToken"); err != nil {
		fmt.Println("Delete Failed!!! ", err.Error())
		return shim.Error(("Delete Failed!! "+err.Error()+"!!!"))
	}

	return shim.Success([]byte("true"))
}

// AddTooken adds the token back to  the State
// Moved the init logic to this function
// V6
func AddToken(stub shim.ChaincodeStubInterface) peer.Response  {
	// If token exist then return false
	value, _ := stub.GetState("MyToken")
	if value != nil {
		return shim.Success([]byte("false"))
	}

	// Simply print a message
	fmt.Println("Init executed")

	// Lets create a Key-Value pair
	stub.PutState("MyToken", []byte("2000"))

	// Return success
	return shim.Success([]byte("true"))
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode.\n")
	err := shim.Start(new(TokenChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}