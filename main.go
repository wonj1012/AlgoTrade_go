package main

import (
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2/futures"
	"github.com/dhzjfhtm/ATH/config"
	"github.com/dhzjfhtm/ATH/realtime/api"
	"github.com/dhzjfhtm/ATH/record"
	"github.com/dhzjfhtm/ATH/strategy/larry"
	"github.com/dhzjfhtm/ATH/trade"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	logger := record.NewLogger()
	binanceFuture := api.NewBinanceFuture(logger)
	//
	config.SetFutureConfig(binanceFuture, logger)

	flowStage := 0

	for {
		now := time.Now()
		if now.Hour() == 8 && now.Minute() == 55 && (flowStage == 0 || flowStage == 2) {

			// position 조회 -> 0보다 크면 전량 매도 (시장가로)

			flowStage = 1

		} else if now.Hour() == 9 && now.Minute() == 0 && flowStage == 1 {
			// 잔고조회 -> usdt 수량 확인
			// usdtTrade := usdt / len(config.Symbols)
			usdtTrade := 100.

			for i := 0; i < len(config.Symbols); i++ {
				symbol := config.Symbols[i]

				// set leverage and isolated
				config.SetFutureConfig(binanceFuture, logger)

				targetPrice := larry.GetTargetPrice(binanceFuture, symbol)
				quantity := usdtTrade / targetPrice
				trade.NewStopLimitOrder(binanceFuture, symbol, string(futures.SideTypeBuy), quantity, targetPrice)
			}
			flowStage = 2

		}
		if flowStage == 2 {
			larry.Run(binanceFuture)
			time.Sleep(1 * time.Second)
		}

	}

}
