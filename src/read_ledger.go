package main

import (
  "github.com/hyperledger/fabric/core/chaincode/shim"
  pb "github.com/hyperledger/fabric/protos/peer"
)

func queryAll(APIstub shim.ChaincodeStubInterface) pb.Response {
  resultsIterator, err := APIstub.GetStateByRange("", "")
  if err != nil {
    return shim.Error(err.Error())
  }
  defer resultsIterator.Close()

  var queryValAsBytes []byte
  for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

    // queryKeyAsStr := aKeyValue.Key
    queryValAsBytes := queryResponse.Value
  }

  return shim.Success(queryValAsBytes)
}
