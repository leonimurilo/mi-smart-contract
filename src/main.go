package main

import (
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

type Supplier struct {
  ID string `json:"id"`
  Name string `json:"name"`
  UF string `json:"uf"`
}

type Recipient struct {
  ID string `json:"id"`
  Name string `json:"name"`
  UF string `json:"uf"`
}

type Product struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Type string `json:"type"`
  ExpirationDate string `json:"expirationDate"`
  ProductionDate string `json:"productionDate"`
  Producer string `json:"producer"`
  Weight float32 `json:"weight"`
}

// steps: purchased -> called -> departed -> arrived | -> invalidated -> returned |
type Lot struct {
  ObjectType string `json:"docType"` //field for couchdb
  ID string `json:"id"`
  PurchasedProductsQuantity int `json:"purchasedProductsQuantity"`
  ReceivedProductsQuantity int `json:"receivedProductsQuantity"`
  Supplier Supplier `json:"supplier"`
  Recipient Recipient `json:"recipient"`
  Product Product `json:"product"`
  Status string `json:"status"`
  ExpectedDeliveryDate string `json:"expectedDeliveryDate"`
  DeliveryDate string `json:"deliveryDate"`
  CallDate string `json:"callDate"`
}

// func (lot *Lot) getID() string {
//   return lot.ID
// }

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
