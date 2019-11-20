/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the card structure, with 4 properties.  Structure tags are used by encoding/json library
type Card struct {
	C1   string `json:"c1"`
	C2   string `json:"c2"`
	C3   string `json:"c3"`
	C4   string `json:"c4"`
	C5   string `json:"c5"`
	C6   string `json:"c6"`
	C7   string `json:"c7"`
	C8   string `json:"c8"`
	C9   string `json:"c9"`
	C10   string `json:"c10"`
	C11   string `json:"c11"`
	C12   string `json:"c12"`
	C13   string `json:"c13"`
	C14   string `json:"c14"`
	C15   string `json:"c15"`
	C16   string `json:"c16"`
	C17   string `json:"c17"`
	C18   string `json:"c18"`
	C19   string `json:"c19"`
	C20   string `json:"c20"`
	C21   string `json:"c21"`
	C22   string `json:"c22"`
	C23   string `json:"c23"`
	C24   string `json:"c24"`
	C25   string `json:"c25"`
	C26   string `json:"c26"`
	C27   string `json:"c27"`
	C28   string `json:"c28"`
	C29   string `json:"c29"`
	C30   string `json:"c30"`
	C31   string `json:"c31"`
	C32   string `json:"c32"`
	C33   string `json:"c33"`
	C34   string `json:"c34"`
	C35   string `json:"c35"`
	C36   string `json:"c36"`
	C37   string `json:"c37"`
	C38   string `json:"c38"`
	C39   string `json:"c39"`
	C40   string `json:"c40"`
	C41   string `json:"c41"`
	C42   string `json:"c42"`
	C43   string `json:"c43"`
	C44   string `json:"c44"`
	C45   string `json:"c45"`
	C46   string `json:"c46"`
	C47   string `json:"c47"`
	C48   string `json:"c48"`
	C49   string `json:"c49"`
	C50   string `json:"c50"`
	C51   string `json:"c51"`
	C52   string `json:"c52"`
}

/*
 * The Init method is called when the Smart Contract "fabcard" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcard"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryCard" {
		return s.queryCard(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createCard" {
		return s.createCard(APIstub, args)
	} else if function == "queryAllCards" {
		return s.queryAllCards(APIstub)
	} else if function == "changeCardOwner" {
		return s.changeCardOwner(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryCard(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	cardAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(cardAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	cards := []Card{
		Card{
				C1:"SPADE|1",
				C2: "SPADE|2",
				C3: "SPADE|3",
				C4: "SPADE|4",
				C5: "SPADE|5",
				C6: "SPADE|6",
				C7: "SPADE|7",
				C8: "SPADE|8",
				C9: "SPADE|9",
				C10: "SPADE|10",
				C11: "SPADE|11",
				C12: "SPADE|12",
				C13: "SPADE|13",
				C14: "HEART|1",
				C15: "HEART|2",
				C16: "HEART|3",
				C17: "HEART|4",
				C18: "HEART|5",
				C19: "HEART|6",
				C20: "HEART|7",
				C21: "HEART|8",
				C22: "HEART|9",
				C23: "HEART|10",
				C24: "HEART|11",
				C25: "HEART|12",
				C26: "HEART|13",
				C27: "CLUB|1",
				C28: "CLUB|2",
				C29: "CLUB|3",
				C30: "CLUB|4",
				C31: "CLUB|5",
				C32: "CLUB|6",
				C33: "CLUB|7",
				C34: "CLUB|8",
				C35: "CLUB|9",
				C36: "CLUB|10",
				C37: "CLUB|11",
				C38: "CLUB|12",
				C39: "CLUB|13",
				C40: "DIAMOND|1",
				C41: "DIAMOND|2",
				C42: "DIAMOND|3",
				C43: "DIAMOND|4",
				C44: "DIAMOND|5",
				C45: "DIAMOND|6",
				C46: "DIAMOND|7",
				C47: "DIAMOND|8",
				C48: "DIAMOND|9",
				C49: "DIAMOND|10",
				C50: "DIAMOND|11",
				C51: "DIAMOND|12",
				C52: "DIAMOND|13"
		}
	}

	i := 0
	for i < len(cards) {
		fmt.Println("i is ", i)
		cardAsBytes, _ := json.Marshal(cards[i])
		APIstub.PutState("CARD"+strconv.Itoa(i), cardAsBytes)
		fmt.Println("Added", cards[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createCard(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var card = Card{C1 : args[1], 
					C2 : args[2], 
					C3 : args[3], 
					C4 : args[4], 
					C5 : args[5], 
					C6 : args[6], 
					C7 : args[7], 
					C8 : args[8], 
					C9 : args[9], 
					C10: args[10],
					C11: args[11],
					C12: args[12],
					C13: args[13],
					C14: args[14],
					C15: args[15],
					C16: args[16],
					C17: args[17],
					C18: args[18],
					C19: args[19],
					C20: args[20],
					C21: args[21],
					C22: args[22],
					C23: args[23],
					C24: args[24],
					C25: args[25],
					C26: args[26],
					C27: args[27],
					C28: args[28],
					C29: args[29],
					C30: args[30],
					C31: args[31],
					C32: args[32],
					C33: args[33],
					C34: args[34],
					C35: args[35],
					C36: args[36],
					C37: args[37],
					C38: args[38],
					C39: args[39],
					C40: args[40],
					C41: args[41],
					C42: args[42],
					C43: args[43],
					C44: args[44],
					C45: args[45],
					C46: args[46],
					C47: args[47],
					C48: args[48],
					C49: args[49],
					C50: args[50],
					C51: args[51],
					C52: args[52]}

	cardAsBytes, _ := json.Marshal(card)
	APIstub.PutState(args[0], cardAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllCards(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "CARD0"
	endKey := "CARD999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCards:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeCardOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	cardAsBytes, _ := APIstub.GetState(args[0])
	card := Card{}

	json.Unmarshal(cardAsBytes, &card)
	card.Owner = args[1]

	cardAsBytes, _ = json.Marshal(card)
	APIstub.PutState(args[0], cardAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
