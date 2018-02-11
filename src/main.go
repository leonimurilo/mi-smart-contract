package main

import (
  "./models"
  "bytes"
  "encoding/json"
  "fmt"
	"strconv"
  "github.com/hyperledger/fabric/core/chaincode/shim"
  pb "github.com/hyperledger/fabric/protos/peer"
)

//docs: https://godoc.org/github.com/hyperledger/fabric/core/chaincode/shim#MockStub.GetStateByRange

type SmartContract struct {
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

func (sc *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (sc *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "recordLotCall" {
		// return sc.recordLotCall(APIstub, args)
	} else if function == "recordLotDelivery" {
		// return sc.recordLotDelivery(APIstub)
	} else if function == "recordLotReceipt" {
		// return sc.recordLotReceipt(APIstub, args)
	} else if function == "queryAll" {
		return sc.queryAll(APIstub)
	} else if function == "initLedger" {
		return sc.initLedger(APIstub)
	}

	return shim.Error("Invalid Smart Contract function name.")
}
