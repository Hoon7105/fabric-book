
package main

import (
	"encoding/json"
	"fmt"
		  
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct{
}

type BOOK struct{
      ISBN string `json:"isbn"` //일련번호
      Name string `json:"name"` //book name
      Category string `json:"category"` //book category
}

type USER struct  {
      ObjectType string `json:"doctype"`
      ID string `json:"userid"`
      ReadBook []BOOK `json:"readbook"`
      RecommendBook []BOOK `json:"Recommendbok"` 
}  

func (s *SmartContract) Init(APIsub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
        

    function, args := APIstub.GetFunctionAndParameters()

    if function == "setUser" {
    	return s.setUser(APIstub, args)
    } else if function == "readbook" {
        return s.readBook(APIstub, args)
	} else if function == "getuserBookinfo" {
		return s.getuserBookinfo(APIstub, args)
	} else if function == "RecommendBook" {
		return s.RecommendBook(APIstub, args)
	} else if function == "getRecommendBookinfo" {
		return s.getRecommendBookinfo(APIstub, args)
	}
      
	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) setUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	docType := "user"
	id := docType + args[0]
      
	user := USER{ObjectType : docType, ID : id}

	userAsBytes, _ := json.Marshal(user)
	APIstub.PutState(id, userAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) readBook(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
      //args [userId, ISBN, Bookname , bookcategory]

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	user := USER{}

	docType := "user"
	id := docType + args[0]

	userAsBytes, _ := APIstub.GetState(id)

	err := json.Unmarshal(userAsBytes, &user)
	if err != nil {
		return shim.Error(err.Error())
	}

	book := BOOK{ISBN : args[1], Name : args[2], Category : args[3]}

	user.ReadBook = append(user.ReadBook,book)
	userAsBytes, _ = json.Marshal(user) 
	APIstub.PutState(id, userAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) getuserBookinfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

   if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	
	user := USER{}

	docType := "user"
	id := docType + args[0]

	userAsBytes, _ := APIstub.GetState(id)
	err := json.Unmarshal(userAsBytes, &user)
	if err != nil {
		return shim.Error(err.Error())
	}

	ReadBookAsBytes, _ := json.Marshal(user.ReadBook)     

	return shim.Success(ReadBookAsBytes)
      
}

func (s *SmartContract) RecommendBook(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
      
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	user := USER{}

	docType := "user"
	id := docType + args[0]

	userAsBytes, _ := APIstub.GetState(id)
	err := json.Unmarshal(userAsBytes, &user)
	if err != nil {
		return shim.Error(err.Error())
	}

	book := BOOK{ISBN : args[1], Name : args[2], Category : args[3]}

	user.RecommendBook = append(user.RecommendBook,book)
	userAsBytes, _ = json.Marshal(user) 
	APIstub.PutState(id, userAsBytes)

	return shim.Success(nil)
}  

func (s *SmartContract) getRecommendBookinfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
      
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
    
	user := USER{}

	docType := "user"
	id := docType + args[0]

	userAsBytes, _ := APIstub.GetState(id)
	err := json.Unmarshal(userAsBytes, &user)
    
	if err != nil {
		return shim.Error(err.Error())
	}

	RecommendBookAsBytes, _ := json.Marshal(user.RecommendBook)     

	return shim.Success(RecommendBookAsBytes)
}

func main() {
        // Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
} 
