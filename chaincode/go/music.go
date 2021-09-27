package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "strconv"

   "github.com/hyperledger/fabric/core/chaincode/shim"
   pb "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct{}

type Wallet struct {
   Name  string `json:"name"`
   ID    string `json:"id"`
   Token string `json:"token"`
}

type Goods struct {
   // Title    string `json:"title"`
   // Singer   string `json:"singer"`
   // Price    string `json:"price"`
   // WalletID string `json:"walletid"`
   State    string `json:"state"`
   Seller   string `json:"seller"`
   Category    string `json:"category"`
   Price    string `json:"price"`
   Content    string `json:"content"`
   Title    string `json:"title"`
   WalletID    string `json:"walletid"`
}

type GoodsKey struct {
   Key string
   Idx int
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
   return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
   function, args := APIstub.GetFunctionAndParameters()

   if function == "initWallet" {
      return s.initWallet(APIstub)
   } else if function == "getWallet" {
      return s.getWallet(APIstub, args)
   } else if function == "setGoods" {
      return s.setGoods(APIstub, args)
   } else if function == "getAllGoods" {
      return s.getAllGoods(APIstub)
   } else if function == "purchaseGoods" {
      return s.purchaseGoods(APIstub, args)
   }
   fmt.Println("Please check your function : " + function)
   return shim.Error("Unknown function")
}

func (s *SmartContract) initWallet(APIstub shim.ChaincodeStubInterface) pb.Response {

   //Declare wallets
   seller := Wallet{Name: "Hyper", ID: "1Q2W3E4R", Token: "100"}
   buyer := Wallet{Name: "Ledger", ID: "5T6Y7U8I", Token: "200"}

   // Convert seller to []byte
   SellerasJSONBytes, _ := json.Marshal(seller)
   err := APIstub.PutState(seller.ID, SellerasJSONBytes)
   if err != nil {
      return shim.Error("Failed to create asset " + seller.Name)
   }
   // Convert buyer to []byte
   BuyerasJSONBytes, _ := json.Marshal(buyer)
   err = APIstub.PutState(buyer.ID, BuyerasJSONBytes)
   if err != nil {
      return shim.Error("Failed to create asset " + buyer.Name)
   }

   return shim.Success(nil)
}

func (s *SmartContract) getWallet(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

   walletAsBytes, err := APIstub.GetState(args[0])
   if err != nil {
      fmt.Println(err.Error())
   }

   wallet := Wallet{}
   json.Unmarshal(walletAsBytes, &wallet)

   var buffer bytes.Buffer
   buffer.WriteString("[")
   bArrayMemberAlreadyWritten := false

   if bArrayMemberAlreadyWritten == true {
      buffer.WriteString(",")
   }
   buffer.WriteString("{\"Name\":")
   buffer.WriteString("\"")
   buffer.WriteString(wallet.Name)
   buffer.WriteString("\"")

   buffer.WriteString(", \"ID\":")
   buffer.WriteString("\"")
   buffer.WriteString(wallet.ID)
   buffer.WriteString("\"")

   buffer.WriteString(", \"Token\":")
   buffer.WriteString("\"")
   buffer.WriteString(wallet.Token)
   buffer.WriteString("\"")

   buffer.WriteString("}")
   bArrayMemberAlreadyWritten = true
   buffer.WriteString("]\n")

   return shim.Success(buffer.Bytes())

}
func generateGoodsKey(APIstub shim.ChaincodeStubInterface, key string) []byte {

   var isFirst bool = false

   goodskeyAsBytes, err := APIstub.GetState(key)
   if err != nil {
      fmt.Println(err.Error())
   }

   goodskey := GoodsKey{}
   json.Unmarshal(goodskeyAsBytes, &goodskey)
   var tempIdx string
   tempIdx = strconv.Itoa(goodskey.Idx)
   fmt.Println(goodskey)
   fmt.Println("Key is " + strconv.Itoa(len(goodskey.Key)))
   if len(goodskey.Key) == 0 || goodskey.Key == "" {
      isFirst = true
      goodskey.Key = "GD"
   }
   if !isFirst {
      goodskey.Idx = goodskey.Idx + 1
   }

   fmt.Println("Last GoodsKey is " + goodskey.Key + " : " + tempIdx)

   returnValueBytes, _ := json.Marshal(goodskey)

   return returnValueBytes
}

// func generateGoodsKey(APIstub shim.ChaincodeStubInterface, key string) []byte {

//    var isFirst bool = false

//    goodskeyAsBytes, err := APIstub.GetState(key)
//    if err != nil {
//       fmt.Println(err.Error())
//    }

//    goodskey := GoodsKey{}
//    json.Unmarshal(goodskeyAsBytes, &goodskey)
//    var tempIdx string
//    tempIdx = strconv.Itoa(goodskey.Idx)
//    fmt.Println(goodskey)
//    fmt.Println("Key is " + strconv.Itoa(len(goodskey.Key)))
//    if len(goodskey.Key) == 0 || goodskey.Key == "" {
//       isFirst = true
//       goodskey.Key = "GD"
//    }
//    if !isFirst {
//       goodskey.Idx = goodskey.Idx + 1
//    }

//    fmt.Println("Last GoodsKey is " + goodskey.Key + " : " + tempIdx)

//    returnValueBytes, _ := json.Marshal(goodskey)

//    return returnValueBytes
// }

func (s *SmartContract) setGoods(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

   if len(args) != 5 {
      return shim.Error("Incorrect number of arguments. Expecting 5")
   }
   var goodskey = GoodsKey{}
   json.Unmarshal(generateGoodsKey(APIstub, "latestKey"), &goodskey)
   keyidx := strconv.Itoa(goodskey.Idx)
   fmt.Println("Key : " + goodskey.Key + ", Idx : " + keyidx)

   var goods = Goods{Title: args[0], Content: args[1], Price: args[2], Category: args[3], WalletID: args[4]}
   goodsAsJSONBytes, _ := json.Marshal(goods)

   var keyString = goodskey.Key + keyidx
   fmt.Println("goodskey is " + keyString)

   err := APIstub.PutState(keyString, goodsAsJSONBytes)
   if err != nil {
      return shim.Error(fmt.Sprintf("Failed to record goods catch: %s", goodskey))
   }

   goodskeyAsBytes, _ := json.Marshal(goodskey)
   APIstub.PutState("latestKey", goodskeyAsBytes)

   return shim.Success(nil)
}

func (s *SmartContract) getAllGoods(APIstub shim.ChaincodeStubInterface) pb.Response {

   // Find latestKey
   goodskeyAsBytes, _ := APIstub.GetState("latestKey")
   goodskey := GoodsKey{}
   json.Unmarshal(goodskeyAsBytes, &goodskey)
   idxStr := strconv.Itoa(goodskey.Idx + 1)

   var startKey = "GD0"
   var endKey = goodskey.Key + idxStr
   fmt.Println(startKey)
   fmt.Println(endKey)

   resultsIter, err := APIstub.GetStateByRange(startKey, endKey)
   if err != nil {
      return shim.Error(err.Error())
   }
   defer resultsIter.Close()

   var buffer bytes.Buffer
   buffer.WriteString("[")
   bArrayMemberAlreadyWritten := false
   for resultsIter.HasNext() {
      queryResponse, err := resultsIter.Next()
      if err != nil {
         return shim.Error(err.Error())
      }

      if bArrayMemberAlreadyWritten == true {
         buffer.WriteString(",")
      }
      buffer.WriteString("{\"Key\":")
      buffer.WriteString("\"")
      buffer.WriteString(queryResponse.Key)
      buffer.WriteString("\"")

      buffer.WriteString(", \"Record\":")

      buffer.WriteString(string(queryResponse.Value))
      buffer.WriteString("}")
      bArrayMemberAlreadyWritten = true
   }
   buffer.WriteString("]\n")
   return shim.Success(buffer.Bytes())
}

func (s *SmartContract) purchaseGoods(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
   var tokenFromKey, tokenToKey int // Asset holdings
   var goodsprice int               // Transaction value
   var err error

   if len(args) != 2 {
      return shim.Error("Incorrect number of arguments. Expecting 2")
   }

   goodsAsBytes, err := APIstub.GetState(args[1])
   if err != nil {
      return shim.Error(err.Error())
   }

   goods := Goods{}
   json.Unmarshal(goodsAsBytes, &goods)
   goodsprice, _ = strconv.Atoi(goods.Price)

   SellerAsBytes, err := APIstub.GetState(goods.WalletID)
   if err != nil {
      return shim.Error("Failed to get state")
   }
   if SellerAsBytes == nil {
      return shim.Error("Entity not found")
   }
   seller := Wallet{}
   json.Unmarshal(SellerAsBytes, &seller)
   tokenToKey, _ = strconv.Atoi(seller.Token)

   BuyerAsBytes, err := APIstub.GetState(args[0])
   if err != nil {
      return shim.Error("Failed to get state")
   }
   if BuyerAsBytes == nil {
      return shim.Error("Entity not found")
   }

   buyer := Wallet{}
   json.Unmarshal(BuyerAsBytes, &buyer)
   tokenFromKey, _ = strconv.Atoi(string(buyer.Token))

   buyer.Token = strconv.Itoa(tokenFromKey - goodsprice)
   seller.Token = strconv.Itoa(tokenToKey + goodsprice)
   updatedBuyerAsBytes, _ := json.Marshal(buyer)
   updatedSellerAsBytes, _ := json.Marshal(seller)
   APIstub.PutState(args[0], updatedBuyerAsBytes)
   APIstub.PutState(goods.WalletID, updatedSellerAsBytes)

   // buffer is a JSON array containing QueryResults
   var buffer bytes.Buffer
   buffer.WriteString("[")

   buffer.WriteString("{\"Buyer Token\":")
   buffer.WriteString("\"")
   buffer.WriteString(buyer.Token)
   buffer.WriteString("\"")

   buffer.WriteString(", \"Seller Token\":")
   buffer.WriteString("\"")
   buffer.WriteString(seller.Token)
   buffer.WriteString("\"")

   buffer.WriteString("}")
   buffer.WriteString("]\n")

   return shim.Success(buffer.Bytes())
}

func main() {

   err := shim.Start(new(SmartContract))
   if err != nil {
      fmt.Printf("Error starting Simple chaincode: %s", err)
   }
}