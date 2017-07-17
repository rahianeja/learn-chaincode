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
	"github.com/hyperledger/fabric/core/chaincode/shim"
  "encoding/json"

)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct{
}
type TruckData struct {
	Truck1 struct {
		Address string `json:"address"`
		Lat float64 `json:"lat"`
		Long float64 `json:"long"`
		Name string `json:"name"`
		Shock float64 `json:"shock"`
		Status string `json:"status"`
		Time float64 `json:"time"`
		Type string `json:"type"`
	} `json:"Truck1"`
	Truck2 struct {
		Address string `json:"address"`
		Lat float64 `json:"lat"`
		Long float64 `json:"long"`
		Name string `json:"name"`
		Shock float64 `json:"shock"`
		Status string `json:"status"`
		Time float64 `json:"time"`
		Type string `json:"type"`
	} `json:"Truck2"`
	Truck3 struct {
		Address string `json:"address"`
		Lat float64 `json:"lat"`
		Long float64 `json:"long"`
		Name string `json:"name"`
		Shock float64 `json:"shock"`
		Status string `json:"status"`
		Time float64 `json:"time"`
		Type string `json:"type"`
	} `json:"Truck3"`
	Truck4 struct {
		Address string `json:"address"`
		Lat float64 `json:"lat"`
		Long float64 `json:"long"`
		Name string `json:"name"`
		Shock float64 `json:"shock"`
		Status string `json:"status"`
		Time float64 `json:"time"`
		Type string `json:"type"`
	} `json:"Truck4"`
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
	var jsonData string
	var err error

	//shim.SetLoggingLevel(shim.LogLevel("DEBUG"))

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	//initialize the chaincode
	jsonData = args[0];
	err = stub.PutState("truckData", []byte(jsonData))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)
	var err error
	var updatedJsonData string
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	if function == "updateLoc" {
		updatedJsonData = args[0];
													//initialize the chaincode state, used as reset
		err = stub.PutState("truckData", []byte(updatedJsonData))
		if err != nil {
			return nil, err
		}
		var stateArg TruckData
		err = json.Unmarshal([]byte(updatedJsonData), &stateArg)
		if err != nil {
		return nil, errors.New("Truckdata argument unmarshal failed: " + fmt.Sprint(err))
		}

		return nil, nil
	}
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
	//var err error
	var latestTruckData []byte
	// Handle different functions
	if function == "query" {
		//latestTruckData, err = stub.GetState("truckData")
		latestTruckData = byte[]("There are these two young fish swimming along and they happen to meet an older fish swimming the other way")
		
		return latestTruckData, nil
	}
	fmt.Println("query did not find func: " + function)						//error
	return nil, errors.New("Received unknown function query: " + function)
}
