package main

import (
  "fmt"
  "github.com/hyperledger/fabric/core/chaincode/shim"
  "github.com/hyperledger/fabric/protos/peer"
)

type SimpleAsset struct {
    //Vulnerability
    globalValue string
}

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
  return shim.Success(nil)
}

func (t SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
  fn, args := stub.GetFunctionAndParameters()

  if fn == "setValue" {
     t.globalValue = args[0]
     return shim.Success([]byte(t.globalValue))
  } else if fn == "getValue" { // assume 'get' even if fn is nil
      myValue := t.globalValue
     return shim.Success([]byte(myValue))
  }
  return shim.Error("not a valid function")
}

func main() {
  if err := shim.Start(new(SimpleAsset)); err != nil {
     fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
  }
}
