package larry

import (
	"strconv"

	"github.com/dhzjfhtm/ATH/realtime/api"
)

func GetTargetPrice(binanceFuture *api.BinanceFuture, symbol string) float64 {
	kline, _ := binanceFuture.GetBinanceFutureKlines(symbol, "d", 2)
	high, _ := strconv.ParseFloat(kline[1].High, 64)
	low, _ := strconv.ParseFloat(kline[1].Low, 64)
	TR := high - low
	open, _ := strconv.ParseFloat(kline[0].Open, 64)
	targetPrice := open + TR*0.5

	return targetPrice
}

func Run(binanceFuture *api.BinanceFuture) {

}

func Sell(binanceFuture *api.BinanceFuture) {

}
