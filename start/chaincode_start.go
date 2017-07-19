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
	"strconv"

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

	fmt.Println("Inside Init block ! Awesome")

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


func FloatToString(input_num float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'f', 6, 64)
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("metal invoke is running " + function)
	var err error
	var updatedJsonData string
	var violatedTruckData []byte
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
		fmt.Println("metal copied json to struct")
		if stateArg.Truck1.Shock > 2.8 {

			fmt.Println("Truck 1 Violated shock")
			//var result []byte = float64ToByte(stateArg.Truck1.Shock)
			//err = stub.PutState("truck1Violations", result)
			violatedTruckData, err = stub.GetState("truck1ViolationsCount")
			if err != nil {
			 fmt.Println("Could get Truck1 Violation count in bytes")
			 return nil, err
			}

			strVoilatedData, err := string(violatedTruckData)
			if err != nil {
			 fmt.Println("Could not convert Truck1 Violation count byte to string")
			 return nil, err
			}

			count, err := strconv.Atoi(strVoilatedData)
			if err != nil {
			 fmt.Println("Could not convert Truck1 Violation string to int")
			 return nil, err
			}

			count = count + 1

			strNewViolatedData, err := Itoa(count)
			if err != nil {
			 fmt.Println("Could not convert Truck1 Violation int to string")
			 return nil, err
			}

			err = stub.PutState("truck1ViolationsCount",[]byte(strNewViolatedData))
			if err != nil {
			 fmt.Println("Could not save Truck1 Violation count")
			 return nil, err
			}

			var shockStr string
			shockStr = FloatToString(stateArg.Truck1.Shock)
			fmt.Println("Truck 1 Violated shock" + shockStr)
			err = stub.PutState("truck1Violations", []byte(shockStr))
			 if err != nil {
			 	fmt.Println("Could not save Truck1 Violation")
			 	return nil, err
			 }
			fmt.Println("Truck 1 Violation saved")
		}

		if stateArg.Truck2.Shock > 2.8 {
			err = stub.GetState("truck2Violations")
			if err != nil {
			 fmt.Println("Could not save Truck2 Violation")
			 return nil, err
			}
			fmt.Println("Truck 2 Violated shock")
			//var result []byte = float64ToByte(stateArg.Truck1.Shock)
			//err = stub.PutState("truck1Violations", result)
			var shockStr string
			shockStr = FloatToString(stateArg.Truck2.Shock)
			fmt.Println("Truck 2 Violated shock" + shockStr)
			err = stub.PutState("truck2Violations", []byte(shockStr))
			 if err != nil {
			 	fmt.Println("Could not save Truck2 Violation")
			 	return nil, err
			 }
			fmt.Println("Truck 2 Violation saved")
		}

		if stateArg.Truck3.Shock > 2.8 {
			fmt.Println("Truck 3 Violated shock")
			//var result []byte = float64ToByte(stateArg.Truck1.Shock)
			//err = stub.PutState("truck1Violations", result)
			var shockStr string
			shockStr = FloatToString(stateArg.Truck3.Shock)
			fmt.Println("Truck 3 Violated shock" + shockStr)
			err = stub.PutState("truck3Violations", []byte(shockStr))
			 if err != nil {
			 	fmt.Println("Could not save Truck3 Violation")
			 	return nil, err
			 }
			fmt.Println("Truck 3 Violation saved")
		}

		if stateArg.Truck4.Shock > 2.8 {
			fmt.Println("Truck 4 Violated shock")
			//var result []byte = float64ToByte(stateArg.Truck1.Shock)
			//err = stub.PutState("truck1Violations", result)
			var shockStr string
			shockStr = FloatToString(stateArg.Truck4.Shock)
			fmt.Println("Truck 4 Violated shock" + shockStr)
			err = stub.PutState("truck1Violations", []byte(shockStr))
			 if err != nil {
			 	fmt.Println("Could not save Truck1 Violation")
			 	return nil, err
			 }
			fmt.Println("Truck 4 Violation saved")
		}
		return nil, nil
	}
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
	var err error
	var latestTruckData []byte
	//var keyHistory HistoryQueryIteratorInterface
	// Handle different functions
	if function == "query" {
	latestTruckData, err = stub.GetState("truckData")
	//latestTruckData, err = stub.GetState("truck1Violations")

		if err != nil{
  			return nil, err
  		}
		return latestTruckData, nil
	}

	if function == "violation" {
	//latestTruckData, err = stub.GetState("truckData")
	latestTruckData, err = stub.GetState("truck1Violations")

		if err != nil{
  			return nil, err
  		}
		return latestTruckData, nil
	}

	if function == "violationCount" {
	//latestTruckData, err = stub.GetState("truckData")
	latestTruckData, err = stub.GetState("truck1ViolationsCount")

		if err != nil{
  			return nil, err
  		}
		return latestTruckData, nil
	}


	//if function == "keyHistory" {
	//latestTruckData, err = stub.GetState("truckData")
	// latestTruckData, err = stub.GetHistoryForKey("truck3Violations")
	//
	// 	if err != nil{
  // 			return nil, err
  // 		}
	// 	return latestTruckData, nil
	// }
	fmt.Println("query did not find func: " + function)						//error
	return nil, errors.New("Received unknown function query: " + function)
}
