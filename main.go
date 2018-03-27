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
// vulnerability!
var globalValue = ""

func (t SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
  fn, args := stub.GetFunctionAndParameters()

  if fn == "setValue" {
     globalValue = args[0]
     myValue := globalValue
     return shim.Success([]byte(myValue))
  } else if fn == "getValue" { // assume 'get' even if fn is nil
      myValue := globalValue
     return shim.Success([]byte(myValue))
  }
  return shim.Error("not a valid function")
}

func main() {
  if err := shim.Start(new(SimpleAsset)); err != nil {
     fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
  }
}
