package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SimpleAsset struct {
}

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("success"))
}

func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    tByte, err := GetTimeByteArray()
    if err != nil {
        return shim.Error(err.Error())
    }
    err = stub.PutState("key", tByte)
    if err != nil {
        return shim.Error(err.Error())
    }
	return shim.Success([]byte("success"))
}

func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
