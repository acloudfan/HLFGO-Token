package main
/**
 * tokenv3
 * Shows the use of ChaincodeStub API
 *    A) Transaction information
 *    B) Transaction timestamp
 *    C) Channel ID
 * https://github.com/hyperledger/fabric/blob/release-1.4/protos/peer/proposal.pb.go
 https://github.com/hyperledger/fabric/blob/release-1.4/protos/common/common.pb.go
 https://gowalker.org/github.com/hyperledger/fabric/protos/common#ChannelHeader
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

	// V3
	fmt.Println("GetChannelID() =>", stub.GetChannelID())

	// V3
	// Will receieve empty map unless client set the transient data in Tx Proposal
	transientData, _ := stub.GetTransient()
	fmt.Println("GetTransient() =>", transientData)

	// V3
	
	// fmt.Println("signedProposal.String() =>", signedProposal.String(),err)
	
	// fmt.Println("proposal.GetHeader()=>",string(proposal.GetHeader()))

	// Read the header
	// header:= &common.Header{}
	// proto.Unmarshal(proposal.GetHeader(), header)
	// fmt.Println("header=>", header)

	// channelHeader:= &common.ChannelHeader{}
	// fmt.Println("channelHeader=>", string(header.GetChannelHeader()))
	// proto.Unmarshal(header.GetChannelHeader(), channelHeader)

	// fmt.Println("proposal.GetHeader().channelHeader.GetType()=>", channelHeader.GetType())
	// fmt.Println("proposal.GetHeader().channelHeader.GetChannelId=>", channelHeader.GetTxId())

	// signatureHeader:= &common.SignatureHeader{}
	// proto.Unmarshal(proposal.GetHeader(), signatureHeader)
	// fmt.Println("signatureHeader=>", signatureHeader.Creator)
	// fmt.Println("proposal.GetHeader().signatureHeader.GetChannelId=>", signatureHeader.GetTxId())

	PrintSignedProposalInfo(stub)

	// ext := proposal.GetExtension()
	// fmt.Println("proposal.GetExtension()=>", ext)

	// V3
	// Requires use of Handler for decorators on the peer
	// Refer: https://hyperledger-fabric.readthedocs.io/en/release-1.3/pluggable_endorsement_and_validation.html?highlight=handlers
	// decorations := stub. GetDecorations()
	// fmt.Println("GetDecorations() =>", decorations, err)

	// V3
	// binding, err := stub. GetBinding()
	// fmt.Println("GetBinding() =>", string(binding), err)

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