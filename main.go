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

	csvTXFile, err := os.Open("./transactions.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer csvTXFile.Close()

	TXreader := csv.NewReader(csvTXFile)

	TXreader.FieldsPerRecord = -1

	csvTXs, err := TXreader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneTransaction Transaction
	var allTransactions []Transaction

	for _, item := range csvTXs {
		oneTransaction.Txhash = item[0]
		oneTransaction.Blockno, _ = strconv.Atoi(item[1])
		oneTransaction.UnixTimestamp, _ = strconv.Atoi(item[2])
		oneTransaction.DateTime = item[3]
		oneTransaction.From = item[4]
		oneTransaction.To = item[5]
		oneTransaction.ContractAddress = item[6]
		oneTransaction.Value_IN, _ = strconv.ParseFloat(item[7], 64)
		oneTransaction.Value_OUT, _ = strconv.ParseFloat(item[8], 64)
		oneTransaction.CurrentValue, _ = strconv.ParseFloat(item[9], 64)
		oneTransaction.TxnFeeB, _ = strconv.ParseFloat(item[10], 64)
		oneTransaction.TxnFeeU, _ = strconv.ParseFloat(item[11], 64)
		oneTransaction.Historical, _ = strconv.ParseFloat(item[12], 64)
		oneTransaction.Status = item[13]
		oneTransaction.ErrCode = item[14]
		oneTransaction.Method = item[15]

		allTransactions = append(allTransactions, oneTransaction)
	}

	jsonTXdata, err := json.MarshalIndent(allTransactions, "", "  ")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsonTXdata))

	jsonTXFile, err := os.Create("./tranactions.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonTXFile.Close()

	jsonTXFile.Write(jsonTXdata)
	jsonTXFile.Close()
}
