package candlestick

import (
	csgo "github.com/cpustejovsky/customsortgo/sort"
	"sort"
)

type TradeAggregate struct {
	Open  float64
	Close float64
	High  float64
	Low   float64
}

func AggregateTrades(trades []float64) TradeAggregate {
	sortedTrades := csgo.NewSortedFloats(trades)
	return TradeAggregate{
		Open:  trades[0],
		Close: trades[len(trades)-1],
		Low:   sortedTrades[0],
		High:  sortedTrades[len(sortedTrades)-1],
	}
}

func AggregateTradesSideEffects(trades []float64) TradeAggregate {
	//Mark first and last items in list of trades
	ta := TradeAggregate{
		Open:  trades[0],
		Close: trades[len(trades)-1],
	}
	//Sort trades
	sort.Float64s(trades)
	//Mark high and low based on sorted trades
	ta.Low = trades[0]
	ta.High = trades[len(trades)-1]
	return ta
}
