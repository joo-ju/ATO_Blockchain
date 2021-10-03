package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {}
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
	function, args := APIstub.GetFunctionAndParameters()
	
	if function == "initWallet" {
			return s.initWallet(APIstub)
	} else if function == "getWallet" {
			return s.getWallet(APIstub, args)
	} else if function == "setWallet" {
			return s.setWallet(APIstub, args)
	} else if function == "getGoods" {
			return s.getGoods(APIstub, args)
	} else if function == "setGoods" {
			return s.setGoods(APIstub, args)
	} else if function == "getAllGoods" {
			return s.getAllGoods(APIstub)
	} else if function == "purchaseGoods" {
			return s.purchaseGoods(APIstub, args)
	} else if function == "changeGoodsPrice" {
			return s.changeGoodsPrice(APIstub, args)
	} else if function == "deleteGoods" {
			return s.deleteGoods(APIstub, args)
	}
	fmt.Println("Please check your function : " + function)
	return shim.Error("Unknown function")
}


func main() {

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
type Wallet struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Token string `json:"token"`
}
func (s *SmartContract) initWallet(APIstub shim.ChaincodeStubInterface) pb.Response {

	//Declare wallets
	seller := Wallet{Name: "Hyper", ID: "1Q2W3E4R", Token: "100"}
	customer := Wallet{Name: "Ledger", ID: "5T6Y7U8I", Token: "200"}

	// Convert seller to []byte
	SellerasJSONBytes, _ := json.Marshal(seller)
	err := APIstub.PutState(seller.ID, SellerasJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + seller.Name)
	}
	// Convert customer to []byte
	CustomerasJSONBytes, _ := json.Marshal(customer)
	err = APIstub.PutState(customer.ID, CustomerasJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + customer.Name)
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
type Music struct {
	Title    string `json:"title"`
	Singer   string `json:"singer"`
	Price    string `json:"price"`
	WalletID    string `json:"walletid"`
	Count        string `json:"count"`
}

type MusicKey struct {
	Key string
	Idx int
}
func generateKey(APIstub shim.ChaincodeStubInterface, key string) []byte {

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
		goodskey.Key = "MS"
	}
	if !isFirst {
		goodskey.Idx = goodskey.Idx + 1
	}

	fmt.Println("Last GoodsKey is " + goodskey.Key + " : " + tempIdx)

	returnValueBytes, _ := json.Marshal(goodskey)

	return returnValueBytes
}
func (s *SmartContract) setWallet(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
			return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	var wallet = Wallet{Name: args[0], ID: args[1], Token:  args[2]}
	
	WalletasJSONBytes, _ := json.Marshal(wallet)
	err := APIstub.PutState(wallet.ID, WalletasJSONBytes)
	if err != nil {
			return shim.Error("Failed to create asset " + wallet.Name)
	}
	return shim.Success(nil)
}

func (s *SmartContract) setGoods(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	
	var goodskey = GoodsKey{}
	json.Unmarshal(generateKey(APIstub, "latestKey"), &goodskey)
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
// func (s *SmartContract) purchaseGoods(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	var tokenFromKey, tokenToKey int // Asset holdings
// 	var musicprice int // Transaction value
// 	var musiccount int
// 	var err error

// 	if len(args) != 2 {
// 		return shim.Error("Incorrect number of arguments. Expecting 2")
// 	}

// 	musicAsBytes, err := APIstub.GetState(args[1])
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	music := Music{}
// 	json.Unmarshal(musicAsBytes, &music)
// 	musicprice, _ = strconv.Atoi(music.Price)
// 	musiccount, _ = strconv.Atoi(music.Count)

// 	SellerAsBytes, err := APIstub.GetState(music.WalletID)
// 	if err != nil {
// 		return shim.Error("Failed to get state")
// 	}
// 	if SellerAsBytes == nil {
// 		return shim.Error("Entity not found")
// 	}
// 	seller := Wallet{}
// 	json.Unmarshal(SellerAsBytes, &seller)
// 	tokenToKey, _ = strconv.Atoi(seller.Token)

// 	CustomerAsBytes, err := APIstub.GetState(args[0])
// 	if err != nil {
// 		return shim.Error("Failed to get state")
// 	}
// 	if CustomerAsBytes == nil {
// 		return shim.Error("Entity not found")
// 	}

// 	customer := Wallet{}
// 	json.Unmarshal(CustomerAsBytes, &customer)
// 	tokenFromKey, _ = strconv.Atoi(string(customer.Token))

// 	customer.Token = strconv.Itoa(tokenFromKey - musicprice)
// 	seller.Token = strconv.Itoa(tokenToKey + musicprice)
// 	music.Count = strconv.Itoa(musiccount + 1)
// 	updatedCustomerAsBytes, _ := json.Marshal(customer)
// 	updatedSellerAsBytes, _ := json.Marshal(seller)
// 	updatedMusicAsBytes, _ := json.Marshal(music)
// 	APIstub.PutState(args[0], updatedCustomerAsBytes)
// 	APIstub.PutState(music.WalletID, updatedSellerAsBytes)
// 	APIstub.PutState(args[1], updatedMusicAsBytes)

// 	// buffer is a JSON array containing QueryResults
// 	var buffer bytes.Buffer
// 	buffer.WriteString("[")

// 	buffer.WriteString("{\"Customer Token\":")
// 	buffer.WriteString("\"")
// 	buffer.WriteString(customer.Token)
// 	buffer.WriteString("\"")

// 	buffer.WriteString(", \"Seller Token\":")
// 	buffer.WriteString("\"")
// 	buffer.WriteString(seller.Token)
// 	buffer.WriteString("\"")

// 	buffer.WriteString("}")
// 	buffer.WriteString("]\n")

// 	return shim.Success(buffer.Bytes())
// }

func (s *SmartContract) getGoods(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	goodsAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
			fmt.Println(err.Error())
	}

	goods := Goods{}
	json.Unmarshal(goodsAsBytes, &goods)

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
	}
	buffer.WriteString("{\"Title\":")
	buffer.WriteString("\"")
	buffer.WriteString(goods.Title)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Seller\":")
	buffer.WriteString("\"")
	buffer.WriteString(goods.Seller)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Price\":")
	buffer.WriteString("\"")
	buffer.WriteString(goods.Price)
	buffer.WriteString("\"")

	buffer.WriteString(", \"WalletID\":")
	buffer.WriteString("\"")
	buffer.WriteString(goods.WalletID)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Category\":")
	buffer.WriteString("\"")
	buffer.WriteString(goods.Category)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	bArrayMemberAlreadyWritten = true
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}
func (s *SmartContract) changeGoodsPrice(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
    if len(args) != 2 {
        return shim.Error("Incorrect number of arguments. Expecting 2")
    }
    goodsbytes, err := APIstub.GetState(args[0])
    if err != nil {
        return shim.Error("Could not locate goods")
    }
    goods := Goods{}
    json.Unmarshal(goodsbytes, &goods)
    
    goods.Price = args[1]
    goodsbytes, _ = json.Marshal(goods)
    err2 := APIstub.PutState(args[0], goodsbytes)
    if err2 != nil {
        return shim.Error(fmt.Sprintf("Failed to change goods price: %s", args[0]))
    }
    return shim.Success(nil)
}
func (s *SmartContract) deleteGoods(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1")
    }

    A := args[0]

    // Delete the key from the state in ledger
    err := APIstub.DelState(A)
    if err != nil {
        return shim.Error("Failed to delete state")
    }

    return shim.Success(nil)
}
