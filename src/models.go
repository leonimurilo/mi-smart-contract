package main

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
