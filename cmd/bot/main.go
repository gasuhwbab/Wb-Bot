package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gasuhwbab/wbbot/wbapi"
	"github.com/gasuhwbab/wbbot/wbapi/requests"
)

func main() {
	// pref := tgbotapi.Settings{
	// 	Token:  "6656832916:AAG3f2U6RAizDuPsDu6snP3W6qqR_k4JpGU",
	// 	Poller: &tgbotapi.LongPoller{Timeout: 10 * time.Second},
	// }
	// b, err := tgbotapi.NewBot(pref)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// tgbot := telegram.NewBot(b)
	// if err := tgbot.Start(); err != nil {
	// 	log.Fatal(err)
	// }
	wbClient := wbapi.NewWbClient("eyJhbGciOiJFUzI1NiIsImtpZCI6IjIwMjQwODE5djEiLCJ0eXAiOiJKV1QifQ.eyJlbnQiOjEsImV4cCI6MTczOTkxOTg2MiwiaWQiOiIzZjJhNmE3Ny0wZTdjLTQxYTItYWQxZi0xMDBmZGUxYzU1ZDgiLCJpaWQiOjU4MzkyNzM5LCJvaWQiOjEzOTI5ODYsInMiOjEwNzM3NDI4NDgsInNpZCI6Ijk4MzViODdjLThkMjMtNGM0NS1hNWNkLTdjOGEyZTI5OGQ3YSIsInQiOmZhbHNlLCJ1aWQiOjU4MzkyNzM5fQ.RD3ENKfvwPG_QHUz6qB09bfXJXXz-8nzCEFMB9DxdtV8E4oiUSf6swUbzBwspKjat9Cnq8K5adYQx6WClCrv5w")
	warehouses, err := requests.GetWarehouses(context.Background(), wbClient)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*warehouses)
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	coefficients, err := requests.GetCoefficients(context.Background(), wbClient)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*coefficients)
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	warehouseCoefficients, err := requests.GetWarehouseCoefficients(context.Background(), wbClient, 120762)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*warehouseCoefficients)
}
