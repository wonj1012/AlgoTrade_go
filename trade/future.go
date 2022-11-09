package trade

import (
	"fmt"

	"github.com/adshao/go-binance/v2/futures"
	"github.com/dhzjfhtm/ATH/realtime/api"
)

func NewStopLimitOrder(bf *api.BinanceFuture, symbol, side string, quantity, price float64) (*futures.CreateOrderResponse, error) {
	q := fmt.Sprintf("%f", quantity)
	p := fmt.Sprintf("%f", price)
	// stop limit order
	order, err := bf.NewBinanceFutureOrder(symbol, side, "STOPLIMIT", q, p)
	if err != nil {
		// bf.logger.Error("NewStopLimitOrder", err)
		return nil, err
	}
	return order, nil
}

func FutureTrade(binanceFuture *api.BinanceFuture) {
	price, err := binanceFuture.GetBinanceFuturePrice("XRPUSDT")
	if err != nil {
		fmt.Println(err)
		return
	}

	account := binanceFuture.GetBinanceFutureAccount()
	assets := account.Assets
	for _, asset := range assets {
		if asset.WalletBalance != "0.00000000" {
			fmt.Println(asset.Asset, ":", asset.WalletBalance)
		}
	}

	order, err := binanceFuture.NewBinanceFutureOrder("XRPUSDT", "BUY", "LIMIT", "40", price)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)
}
