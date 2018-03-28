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
	key1 := "testKey1"

    err := writeToLedger(stub, &key1)
    if err != nil {
        return shim.Error("could not write data")
    }

	//vulnerability
	respone, err := stub.GetState(key1)
	if err != nil {
		return shim.Error("could not read data")
	}

	return shim.Success([]byte(respone))
}

func writeToLedger(stub shim.ChaincodeStubInterface, key *string) error {
	return stub.PutState(*key, []byte("data"))
}

func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
