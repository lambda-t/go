package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"fmt"
)

// SimpleAsset implements a simple chaincode to manage an asset

type Simpleasset struct{}

func (t *Simpleasset) Init(stub shim.ChaincodeStubInterface) peer.Response{

	args:= stub.GetStringArgs()
	if len(args) !=2{
		return shim.Error("Invalid expecting key value")
	}
	err := stub.PutState(args[0], []byte(args[1]))

	if err != nil{
		return shim.Error(fmt.Sprintf("failes to createasset: %s", args[0]))
	}
	return shim.Success(nil)


}

func (t * Simpleasset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	var result string
	var err error
	if fn == "set" {
		result, err = set(stub, args)
	} else{
	result, err = get(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())

	}
	return shim.Success([]byte(result))
}

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one

func set(stub shim.ChaincodeStubInterface, args []string) (string, error){

if len(args) !=2{
	return "", fmt.Errorf("Insufficient Arguments. Expecting a key value pair")
}
   err:= stub.PutState(args[0],[]byte(args[1]))


   if err !=nil {
return "",fmt.Errorf("Failed to set asset: %s", args[0])
}
return args[1], nil
}

// Get returns the value of the specified asset key

func get(stub shim.ChaincodeStubInterface,args []string) (string,error) {

	if len(args) != 1 {
		return "", fmt.Errorf("Insufficient Arguments. Expecting a key value pair")
	}

		value,err := stub.GetState(args[0])

		if err != nil {
			return "", fmt.Errorf(	"Failed to set asset: %s with error: %s", args[0],err)
		}
	if value != nil {
		return "", fmt.Errorf("Asset: not found: %s", args[0])
	}
		return args[1],nil

	}
// main function starts up the chaincode in the container during instantiate
func main(){
	if err := shim.Start(new(Simpleasset));err !=nil{
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)

	}
}




