package main

import (
	"log"
	"strings"
	"time"

	tgbotapi "gopkg.in/telebot.v3"
)

func main() {

	pref := tgbotapi.Settings{
		Token:  "6656832916:AAG3f2U6RAizDuPsDu6snP3W6qqR_k4JpGU",
		Poller: &tgbotapi.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tgbotapi.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	vehiclesKb := tgbotapi.ReplyMarkup{
		OneTimeKeyboard: true,
		ResizeKeyboard:  true,
		Placeholder:     "Выберите категорию",
	}
	vehiclesKb.ReplyKeyboard = [][]tgbotapi.ReplyButton{
		{
			*vehiclesKb.Text("Мотоцикл").Reply(),
			*vehiclesKb.Text("Автомобиль").Reply(),
		},
	}

	b.Handle("/start", func(ctx tgbotapi.Context) error {
		return ctx.Send("Какой вид транспорта вы хотите приобрести?", &vehiclesKb)
	})
	b.Handle(tgbotapi.OnText, func(ctx tgbotapi.Context) error {
		motorcyclesKb := tgbotapi.ReplyMarkup{
			OneTimeKeyboard: true,
			ResizeKeyboard:  true,
			Placeholder:     "Выберите мотоцикл",
		}
		motorcyclesKb.ReplyKeyboard = [][]tgbotapi.ReplyButton{
			{
				*motorcyclesKb.Text("Harley Davidson").Reply(),
				*motorcyclesKb.Text("Ducati").Reply(),
			},
		}

		autoKb := tgbotapi.ReplyMarkup{
			OneTimeKeyboard: true,
			ResizeKeyboard:  true,
			Placeholder:     "Выберите автомобиль",
		}
		autoKb.ReplyKeyboard = [][]tgbotapi.ReplyButton{
			{
				*autoKb.Text("BMW").Reply(),
				*autoKb.Text("Audi").Reply(),
				*autoKb.Text("Mercedes").Reply(),
			},
		}
		switch {
		case strings.EqualFold(ctx.Text(), "мотоцикл"):
			return ctx.Send("Выберите желаемый вариант", &motorcyclesKb)
		case strings.EqualFold(ctx.Text(), "автомобиль"):
			return ctx.Send("Выберите желаемый вариант", &autoKb)
			// default:
			// 	return ctx.Send(fmt.Sprintf("Такой категории не существует: %s\nВыберете из предложенных вариантов",
			// 		ctx.Text()), &vehiclesKb)
		}
		return nil
	})

	b.Start()

	// tgbot := telegram.NewBot(b)
	// if err := tgbot.Start(); err != nil {
	// 	log.Fatal(err)
	// }

	// wbClient := wbapi.NewWbClient("eyJhbGciOiJFUzI1NiIsImtpZCI6IjIwMjQwODE5djEiLCJ0eXAiOiJKV1QifQ.eyJlbnQiOjEsImV4cCI6MTczOTkxOTg2MiwiaWQiOiIzZjJhNmE3Ny0wZTdjLTQxYTItYWQxZi0xMDBmZGUxYzU1ZDgiLCJpaWQiOjU4MzkyNzM5LCJvaWQiOjEzOTI5ODYsInMiOjEwNzM3NDI4NDgsInNpZCI6Ijk4MzViODdjLThkMjMtNGM0NS1hNWNkLTdjOGEyZTI5OGQ3YSIsInQiOmZhbHNlLCJ1aWQiOjU4MzkyNzM5fQ.RD3ENKfvwPG_QHUz6qB09bfXJXXz-8nzCEFMB9DxdtV8E4oiUSf6swUbzBwspKjat9Cnq8K5adYQx6WClCrv5w")

	// warehouses, err := requests.GetWarehouses(context.Background(), wbClient)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(*warehouses)
	// fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")

	// coefficients, err := requests.GetCoefficients(context.Background(), wbClient)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(*coefficients)
	// fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")

	// warehouseCoefficients, err := requests.GetWarehouseCoefficients(context.Background(), wbClient, 120762)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(*warehouseCoefficients)
}
