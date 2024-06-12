package candlestick

import (
	csgo "github.com/cpustejovsky/customsortgo/pure"
	"sort"
)

type TradeAggregate struct {
	Open  float64
	Close float64
	High  float64
	Low   float64
}

// NewTradeAggregate returns a TradeAggregate based on a copy of the income trades slice
// This is nearly 4x slower than AggregateTrades
// Use if you do not want to mutate your data
func NewTradeAggregate(trades []float64) TradeAggregate {
	sortedTrades := csgo.NewSortedFloats(trades)
	return TradeAggregate{
		Open:  trades[0],
		Close: trades[len(trades)-1],
		Low:   sortedTrades[0],
		High:  sortedTrades[len(sortedTrades)-1],
	}
}

// AggregateTrades returns a TradeAggregate based on the incoming trades slice, mutating the slice
// This is significantly faster than NewTradeAggregate
// Use if the original order of incoming slice does not need to be retained
func AggregateTrades(trades []float64) TradeAggregate {
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
