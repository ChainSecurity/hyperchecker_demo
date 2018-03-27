package main

import (
  "fmt"

  "github.com/hyperledger/fabric/core/chaincode/shim"
  "github.com/hyperledger/fabric/protos/peer"
)

type SimpleAsset struct {
}

type simpleStruct struct {
    key string
    data string
}

type simpleStruct2 struct {
    ss simpleStruct
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
  keyValuePair := simpleStruct {
      key: "testKey",
      data: "dataNew",
  }

  ss2 := simpleStruct2 {
      ss: keyValuePair,
  }

  err := stub.PutState(ss2.ss.key, []byte(keyValuePair.data))
  if err != nil {
     return shim.Error("could not write new data")
  }
  //vulnerability
  respone, err := stub.GetState(ss2.ss.key)
  if err != nil {
     return shim.Error("could not read data")
  }

  return shim.Success([]byte(respone))
}


func main() {
  if err := shim.Start(new(SimpleAsset)); err != nil {
     fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
  }
}
