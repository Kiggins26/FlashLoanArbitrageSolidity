package main 

import (
    "fmt"
    "log"
    "os"
    "io/ioutil"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

type ContractDetails struct {
    ContractAddress string `json:"address"`
    EthRPCURL string `json:"url"`
}

func callContractUniSwap( contractDetails *ContractDetails) {
    
    client, err := ethclient.Dial(contractDetails.EthRPCURL)
    if err != nil {
        log.Fatal(err)
    }


    address := common.HexToAddress(contractDetails.ContractAddress)
    // replace with contract call instance, err := store.NewStore(address, client)
    if err != nil {
        log.Fatal(err)
    }
}

func callContractSushiSwap( contractDetails *ContractDetails) {
    
    client, err := ethclient.Dial(contractDetails.EthRPCURL)
    if err != nil {
        log.Fatal(err)
    }


    address := common.HexToAddress(contractDetails.ContractAddress)
    // replace with contract call instance, err := store.NewStore(address, client)
    if err != nil {
        log.Fatal(err)
    }
}
func main(){
    jsonPath := "./arbitageDetails.json"
    jsonFile, err := os.Open(jsonPath)
    
    if err != nil {
        log.Fatalf("CAN NOT OPEN FILE: %s \n", jsonPath)
    }

    defer jsonFile.Close()
    
    byteVal, err := ioutil.RealAll(jsonFile)
    
    if err != nil {
        log.Fatalf("UNABLE TO READ THE CONTENT OF %s \n", jsonPath)
    }
    
    var contractDetails ContractDetails

    json.Unmarshal(byteVal, &contractDetails)


    // Get prices and look for price diff, then call the code below
    priceFavour := "equal"
    while priceFavour == "equal" {
        //get uniswap and sushiswap diff
        priceFavour = "uniswap"
        switch priceFavour {
            case "uniswap":
                callContractUniSwap(&contractDetails)
            case "sushiswap":
                callContractSushiSwap(&contractDetails) 
            default:
               log.print("Invalid string") 
        }
    }
}
