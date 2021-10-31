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
	Title    string `json:"title"`
	State    string `json:"state"`
	Seller   string `json:"seller"`
	Category string `json:"category"`
	Price    string `json:"price"`
	Content  string `json:"content"`
	WalletID string `json:"walletid"`
	Count    string `json:"count"`
}

type Event struct {
	Name     string `json:"name"`     // 행사 이름
	Seller   string `json:"seller"`   // 주최자
	Category string `json:"category"` // 종류
	Price    string `json:"price"`    // 가격
	Date     string `json:"date"`     // 행사 날짜 & 시간
	WalletID string `json:"walletid"`
	Count    string `json:"count"`
}

type GoodsKey struct {
	Key string
	Idx int
}

type EventKey struct {
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
		// } else if function == "getAllWallet" {
		// 	return s.getAllWallet(APIstub, args)
	} else if function == "setWallet" {
		return s.setWallet(APIstub, args)
	} else if function == "getGoods" {
		return s.getGoods(APIstub, args)
	} else if function == "setGoods" {
		return s.setGoods(APIstub, args)
	} else if function == "getEvent" {
		return s.getEvent(APIstub, args)
	} else if function == "setEvent" {
		return s.setEvent(APIstub, args)
	} else if function == "getAllGoods" {
		return s.getAllGoods(APIstub)
	} else if function == "getAllEvent" {
		return s.getAllEvent(APIstub)
	} else if function == "purchaseGoods" {
		return s.purchaseGoods(APIstub, args)
	} else if function == "purchaseEvent" {
		return s.purchaseEvent(APIstub, args)
	} else if function == "changeGoodsPrice" {
		return s.changeGoodsPrice(APIstub, args)
	} else if function == "changeEventPrice" {
		return s.changeEventPrice(APIstub, args)
	} else if function == "deleteGoods" {
		return s.deleteGoods(APIstub, args)
	} else if function == "deleteEvent" {
		return s.deleteEvent(APIstub, args)
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

// func (s *SmartContract) getAllWallet(APIstub shim.ChaincodeStubInterface) pb.Response {

// 	// Find latestKey

// 	walletAsBytes, err := APIstub.GetState("latestKey")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	wallet := Wallet{}
// 	json.Unmarshal(walletAsBytes, &wallet)
// 	idxStr := strconv.Itoa(wallet.Idx + 1)

// 	var buffer bytes.Buffer
// 	buffer.WriteString("[")
// 	bArrayMemberAlreadyWritten := false
// 	for resultsIter.HasNext() {

// 		if bArrayMemberAlreadyWritten == true {
// 			buffer.WriteString(",")
// 		}
// 		buffer.WriteString("{\"Name\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(wallet.Name)
// 		buffer.WriteString("\"")

// 		buffer.WriteString(", \"ID\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(wallet.ID)
// 		buffer.WriteString("\"")

// 		buffer.WriteString(", \"Token\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(wallet.Token)
// 		buffer.WriteString("\"")

// 		buffer.WriteString("}")
// 		bArrayMemberAlreadyWritten = true
// 		buffer.WriteString("]\n")

// 		return shim.Success(buffer.Bytes())
// 	}
// 	buffer.WriteString("]\n")
// 	return shim.Success(buffer.Bytes())
// }

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

func generateEventKey(APIstub shim.ChaincodeStubInterface, key string) []byte {

	var isFirst bool = false

	EventkeyAsBytes, err := APIstub.GetState(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	eventkey := EventKey{}
	json.Unmarshal(EventkeyAsBytes, &eventkey)
	var tempIdx string
	tempIdx = strconv.Itoa(eventkey.Idx)
	fmt.Println(eventkey)
	fmt.Println("Key is " + strconv.Itoa(len(eventkey.Key)))
	if len(eventkey.Key) == 0 || eventkey.Key == "" {
		isFirst = true
		eventkey.Key = "EV"
	}
	if !isFirst {
		eventkey.Idx = eventkey.Idx + 1
	}

	fmt.Println("Last EventKey is " + eventkey.Key + " : " + tempIdx)

	returnValueBytes, _ := json.Marshal(eventkey)

	return returnValueBytes
}

func (s *SmartContract) setWallet(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	var wallet = Wallet{Name: args[0], ID: args[1], Token: args[2]}

	WalletasJSONBytes, _ := json.Marshal(wallet)
	err := APIstub.PutState(wallet.ID, WalletasJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + wallet.Name)
	}

	return shim.Success(nil)
}

func (s *SmartContract) setGoods(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}
	var goodskey = GoodsKey{}
	json.Unmarshal(generateGoodsKey(APIstub, "latestKey"), &goodskey)
	keyidx := strconv.Itoa(goodskey.Idx)
	fmt.Println("Key : " + goodskey.Key + ", Idx : " + keyidx)

	var goods = Goods{Title: args[0], State: args[1], Seller: args[2], Category: args[3], Price: args[4], Content: args[5], WalletID: args[6], Count: "0"}
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

func (s *SmartContract) getGoods(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	goodsAsBytes, err := stub.GetState(args[0])
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

	buffer.WriteString(", \"State\":")
	buffer.WriteString("\"")
	buffer.WriteString(goods.State)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Seller\":")
	buffer.WriteString("\"")
	buffer.WriteString(goods.Seller)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Category\":")
	buffer.WriteString("\"")
	buffer.WriteString(goods.Category)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Price\":")
	buffer.WriteString("\"")
	buffer.WriteString(goods.Price)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Content\":")
	buffer.WriteString("\"")
	buffer.WriteString(goods.Content)
	buffer.WriteString("\"")

	buffer.WriteString(", \"WalletID\":")
	buffer.WriteString("\"")
	buffer.WriteString(goods.WalletID)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Count\":")
	buffer.WriteString("\"")
	buffer.WriteString(goods.Count)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	bArrayMemberAlreadyWritten = true
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) setEvent(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	var eventkey = EventKey{}
	json.Unmarshal(generateEventKey(APIstub, "latestKey"), &eventkey)
	keyidx := strconv.Itoa(eventkey.Idx)
	fmt.Println("Key : " + eventkey.Key + ", Idx : " + keyidx)

	var event = Event{Name: args[0], Seller: args[1], Category: args[2], Price: args[3], Date: args[4], WalletID: args[5], Count: "0"}
	eventAsJSONBytes, _ := json.Marshal(event)

	var keyString = eventkey.Key + keyidx
	fmt.Println("eventkey is " + keyString)

	err := APIstub.PutState(keyString, eventAsJSONBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record event catch: %s", eventkey))
	}

	eventkeyAsBytes, _ := json.Marshal(eventkey)
	APIstub.PutState("latestKey", eventkeyAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) getEvent(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	eventAsBytes, err := stub.GetState(args[0])
	if err != nil {
		fmt.Println(err.Error())
	}

	event := Event{}
	json.Unmarshal(eventAsBytes, &event)

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	if bArrayMemberAlreadyWritten == true {
		buffer.WriteString(",")
	}
	buffer.WriteString("{\"Name\":")
	buffer.WriteString("\"")
	buffer.WriteString(event.Name)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Seller\":")
	buffer.WriteString("\"")
	buffer.WriteString(event.Seller)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Category\":")
	buffer.WriteString("\"")
	buffer.WriteString(event.Category)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Price\":")
	buffer.WriteString("\"")
	buffer.WriteString(event.Price)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Date\":")
	buffer.WriteString("\"")
	buffer.WriteString(event.Date)
	buffer.WriteString("\"")

	buffer.WriteString(", \"WalletID\":")
	buffer.WriteString("\"")
	buffer.WriteString(event.WalletID)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Count\":")
	buffer.WriteString("\"")
	buffer.WriteString(event.Count)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	bArrayMemberAlreadyWritten = true
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
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

func (s *SmartContract) getAllEvent(APIstub shim.ChaincodeStubInterface) pb.Response {

	// Find latestKey
	eventkeyAsBytes, _ := APIstub.GetState("latestKey")
	eventkey := EventKey{}
	json.Unmarshal(eventkeyAsBytes, &eventkey)
	idxStr := strconv.Itoa(eventkey.Idx + 1)

	var startKey = "EV0"
	var endKey = eventkey.Key + idxStr
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
	var goodscount int
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
	goodscount, _ = strconv.Atoi(goods.Count)

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
	goods.Count = strconv.Itoa(goodscount + 1)
	updatedBuyerAsBytes, _ := json.Marshal(buyer)
	updatedSellerAsBytes, _ := json.Marshal(seller)
	updatedGoodsAsBytes, _ := json.Marshal(goods)
	APIstub.PutState(args[0], updatedBuyerAsBytes)
	APIstub.PutState(goods.WalletID, updatedSellerAsBytes)
	APIstub.PutState(args[1], updatedGoodsAsBytes)

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

func (s *SmartContract) purchaseEvent(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	var tokenFromKey, tokenToKey int // Asset holdings
	var eventprice int               // Transaction value
	var eventcount int
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	eventAsBytes, err := APIstub.GetState(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}

	event := Event{}
	json.Unmarshal(eventAsBytes, &event)
	eventprice, _ = strconv.Atoi(event.Price)
	eventcount, _ = strconv.Atoi(event.Count)

	SellerAsBytes, err := APIstub.GetState(event.WalletID)
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

	buyer.Token = strconv.Itoa(tokenFromKey - eventprice)
	seller.Token = strconv.Itoa(tokenToKey + eventprice)
	event.Count = strconv.Itoa(eventcount + 1)
	updatedBuyerAsBytes, _ := json.Marshal(buyer)
	updatedSellerAsBytes, _ := json.Marshal(seller)
	updatedEventAsBytes, _ := json.Marshal(event)
	APIstub.PutState(args[0], updatedBuyerAsBytes)
	APIstub.PutState(event.WalletID, updatedSellerAsBytes)
	APIstub.PutState(args[1], updatedEventAsBytes)

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

func (s *SmartContract) changeGoodsPrice(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	goodsbytes, err := APIstub.GetState(args[0])
	if err != nil {
		fmt.Println(err.Error())
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

func (s *SmartContract) changeEventPrice(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	eventbytes, _ := APIstub.GetState(args[0])
	if eventbytes != nil {
		return shim.Error("Could not locate event")
	}

	event := Event{}
	json.Unmarshal(eventbytes, &event)

	event.Price = args[1]
	eventbytes, _ = json.Marshal(event)
	err := APIstub.PutState(args[0], eventbytes)

	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change event holder: %s", args[0]))
	}
	return shim.Success(nil)
}

func (s *SmartContract) deleteGoods(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	err := APIstub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

func (s *SmartContract) deleteEvent(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	err := APIstub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

func main() {

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
