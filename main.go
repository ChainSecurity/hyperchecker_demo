package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SimpleAsset struct {
}

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	key := "testKey"
	data := "dataInit"
	err := stub.PutState(key, []byte(data))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	key := "testKey"

	iterator, _ := stub.GetHistoryForKey(key)
	data, _ := iterator.Next()

	err := stub.PutState(key, data.Value)
	if err != nil {
		return shim.Error("could not write new data")
	}

	return shim.Success([]byte("stored"))
}

func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
