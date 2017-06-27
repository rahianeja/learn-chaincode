/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"

)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct{
}
type Truck struct{
	Address string `json:"address"`
	Lattitude string `json:"lat"`
	Longitude string `json:"long"`
	Name string `json:"name"`
	Status string `json:"status"`
	Time string `json:"time"`
	Type string `json:"type"`
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var TruckA, TruckB string
	var truckALoc, truckBLoc int
	var err error

	//shim.SetLoggingLevel(shim.LogLevel("DEBUG"))

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	//initialize the chaincode
	TruckA = args[0];
	truckALoc, err = strconv.Atoi(args[1]);
	if err != nil {
		return nil, errors.New("Expecting int value for truck A location");
	}

	TruckB = args[2];
	truckBLoc, err = strconv.Atoi(args[3]);
	if err != nil {
		return nil, errors.New("Expecting int value for truck B location");
	}

	err = stub.PutState(TruckA, []byte(strconv.Itoa(truckALoc)))
	if err != nil {
		return nil, err
	}
	err = stub.PutState(TruckB, []byte(strconv.Itoa(truckBLoc)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)
	var err error
	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		dataFromEnd := args[0]

	//	var truckData Truck
	//	json.Unmarshal([]byte(dataFromEnd), &truckData)

	//	jsonAsBytes, err :=	json.Marshal(truckData)
	//	if err != nil {
	//	fmt.Println("error:", err)
	//	}

		err = stub.PutState("data", []byte(dataFromEnd))
		//err = stub.PutState("data", jsonAsBytes)
		if err != nil {
		    return nil, err
		}
		return t.Init(stub, "init", args)
	}
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
	var jsonAsBytes []byte
	var err error
	// Handle different functions
	if function == "query" {

	/*	truck := Truck{
			Address: "1,Delhi",
			Lattitude: "1.2",
			Longitude:"1.3",
			Name:"Vinayak",
			Status:"Enroute",
			Time:"33:88",
			Type:"16 Wheeler",
		} */
		jsonAsBytes, err = stub.GetState("a")
		if err != nil{
			return nil, err
		}


		//var out Truck

		//err :=	json.Unmarshal(jsonAsBytes, &out)
		//if err != nil {
		//fmt.Println("error:", err)
		//}

		var TruckA string
	//	var err error
		if len(args) != 1 {
				return nil, errors.New("Incorrect number of arguments, expecting truck name")
		}
		TruckA = args[0] //read a variable
		//Get state from ledger
		//Avalbytes, err := stub.GetState(TruckA)
		if err != nil {
			jsonResp :="{\"Error\":\"Failed to get state for" + TruckA + "\"}"
			return nil, errors.New(jsonResp)
		}
		return jsonAsBytes, nil
	}
	fmt.Println("query did not find func: " + function)						//error
	return nil, errors.New("Received unknown function query: " + function)
}
