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

func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
  _, args := stub.GetFunctionAndParameters()
  if len(args) == 0 {
     return shim.Error("No argument specified.")
  }
  //second return value stands for potential error
  result, _ := stub.GetState(args[0])

  return shim.Success(result)
}

func main() {
  if err := shim.Start(new(SimpleAsset)); err != nil {
     fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
  }
}
