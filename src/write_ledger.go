package main

import (
  "encoding/json"
  "fmt"
  "github.com/hyperledger/fabric/core/chaincode/shim"
  pb "github.com/hyperledger/fabric/protos/peer"
)

func initLedger(APIstub shim.ChaincodeStubInterface) pb.Response {
	lots := []Lot{
		Lot{
      ID: "LOT00001",
      PurchasedProductsQuantity: 300,
      Supplier: Supplier{
        ID: "SUP00001",
        Name: "Frutas Dist. LTDA",
        UF: "PR",
      },
      Recipient: Recipient{
        ID: "SCH00001",
        Name: "Escola municipal central de campinas",
        UF: "SP",
      },
      Product: Product{
        ID: "PDT00001",
        Name: "Maçãs frescas tabajara",
        Type: "Organic",
        Weight: 150,
      },
      Status: "called",
      ExpectedDeliveryDate: "2018-01-30T00:00:01.00Z",
      CallDate: "2018-01-03T00:00:01.00Z",
    },
    Lot{
      ID: "LOT00002",
      PurchasedProductsQuantity: 440,
      Supplier: Supplier{
        ID: "SUP00001",
        Name: "Frutas Dist. LTDA",
        UF: "PR",
      },
      Recipient: Recipient{
        ID: "SCH00002",
        Name: "Escola de barão geraldo",
        UF: "SP",
      },
      Product: Product{
        ID: "PDT00009",
        Name: "Banana da terra",
        Type: "Organic",
        ExpirationDate: "2018-03-15T00:00:01.00Z",
        ProductionDate: "2018-01-30T00:00:01.00Z",
        Producer: "Colheitas da tia Nastácia",
        Weight: 140,
      },
      Status: "departed",
      ExpectedDeliveryDate: "2018-01-03T00:00:01.00Z",
      CallDate: "2018-01-03T00:00:01.00Z",
    },
	}

	i := 0
	for i < len(lots) {
		lotAsBytes, _ := json.Marshal(lots[i])
		APIstub.PutState(lots[i].ID, lotAsBytes)
		fmt.Println("Added", lots[i])
		i = i + 1
	}

	return shim.Success(nil)
}

// receives array of lots
// func (sc *SmartContract) recordLotCall(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
//   if len(args) != 5 {
// 		return shim.Error("Incorrect number of arguments. Expecting 5")
// 	}
//
// 	var car = Car{Make: args[1], Model: args[2], Colour: args[3], Owner: args[4]}
//
// 	carAsBytes, _ := json.Marshal(car)
// 	APIstub.PutState(args[0], carAsBytes)
//
// 	return shim.Success(nil)
// }
