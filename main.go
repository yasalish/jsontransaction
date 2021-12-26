package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Transaction struct {
	Txhash          string  `json:"Txhash"`
	Blockno         int     `json:"Blockno"`
	UnixTimestamp   int     `json:"UnixTimestamp"`
	DateTime        string  `json:"DateTime"`
	From            string  `json:"From"`
	To              string  `json:"To"`
	ContractAddress string  `json:"ContractAddress"`
	Value_IN        float64 `json:"Value_IN(BNB)"`
	Value_OUT       float64 `json:"Value_OUT(BNB)"`
	CurrentValue    float64 `json:"CurrentValue @ $542.86/BNB"`
	TxnFeeB         float64 `json:"TxnFee(BNB)"`
	TxnFeeU         float64 `json:"TxnFee(USD)"`
	Historical      float64 `json:"Historical $Price/BNB"`
	Status          string  `json:"Status"`
	ErrCode         string  `json:"ErrCode"`
	Method          string  `json:"Method"`
}

func main() {
	// read data from CSV file

	csvFile, err := os.Open("./transactions.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneTransaction Transaction
	var allTransactions []Transaction

	for _, each := range csvData {
		oneTransaction.Txhash = each[0]
		oneTransaction.Blockno, _ = strconv.Atoi(each[1])
		oneTransaction.UnixTimestamp, _ = strconv.Atoi(each[2])
		oneTransaction.DateTime = each[3]
		oneTransaction.From = each[4]
		oneTransaction.To = each[5]
		oneTransaction.ContractAddress = each[6]
		oneTransaction.Value_IN, _ = strconv.ParseFloat(each[7], 64)
		oneTransaction.Value_OUT, _ = strconv.ParseFloat(each[8], 64)
		oneTransaction.CurrentValue, _ = strconv.ParseFloat(each[9], 64)
		oneTransaction.TxnFeeB, _ = strconv.ParseFloat(each[10], 64)
		oneTransaction.TxnFeeU, _ = strconv.ParseFloat(each[11], 64)
		oneTransaction.Historical, _ = strconv.ParseFloat(each[12], 64)
		oneTransaction.Status = each[13]
		oneTransaction.ErrCode = each[14]
		oneTransaction.Method = each[15]

		allTransactions = append(allTransactions, oneTransaction)
	}

	jsondata, err := json.MarshalIndent(allTransactions, "", "  ")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsondata))

	jsonFile, err := os.Create("./tranactions.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsondata)
	jsonFile.Close()
}
