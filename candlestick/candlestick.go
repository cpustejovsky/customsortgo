package candlestick

import (
	csgo "github.com/cpustejovsky/customsortgo/sort"
	"sort"
)

type TradeAggregate struct {
	Open  int
	Close int
	High  int
	Low   int
}

func AggregateTrades(trades []int) TradeAggregate {
	sortedTrades := csgo.NewSortedIntegers(trades)
	return TradeAggregate{
		Open:  trades[0],
		Close: trades[len(trades)-1],
		Low:   sortedTrades[0],
		High:  sortedTrades[len(sortedTrades)-1],
	}
}

func AggregateTradesSideEffects(trades []int) TradeAggregate {
	//Mark first and last items in list of trades
	ta := TradeAggregate{
		Open:  trades[0],
		Close: trades[len(trades)-1],
	}
	//Sort trades
	sort.Ints(trades)
	//Mark high and low based on sorted trades
	ta.Low = trades[0]
	ta.High = trades[len(trades)-1]
	return ta
}
