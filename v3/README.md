# Shows the use of shim functions
- token.go      V3 of token
- proposal.go   Shows extraction of info from SignedProposal

# Note: 
- In order to follow the code in proposal.go you MUST understan ProtoBuffers
- Refer to the documentation/links on the details of the various buffer structures
- Peer Proto Definition 
https://godoc.org/github.com/hyperledger/fabric/protos/peer
- Common Proto Definitions
https://godoc.org/github.com/hyperledger/fabric/protos/common
- Proposal proto buffer definitions
https://github.com/hyperledger/fabric/blob/release-1.4/protos/peer/proposal.proto


go get -u github.com/golang/protobuf/protoc-gen-go