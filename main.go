package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SimpleAsset struct {
}

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	returnValue := 0
	myMap := make(map[int]int)
	for i, ii := range myMap {
		returnValue = returnValue * i - ii
	}
	return shim.Success([]byte("value: " + string(returnValue)))
}

func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
