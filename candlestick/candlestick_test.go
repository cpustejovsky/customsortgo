package candlestick_test

import (
	"github.com/cpustejovsky/customsortgo/candlestick"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAggregateTrades(t *testing.T) {
	trades := []int{14, 14, 15, 15, 16, 15, 13, 12, 13}
	want := candlestick.TradeAggregate{Open: 14, Close: 13, High: 16, Low: 12}
	got := candlestick.AggregateTrades(trades)
	assert.Equal(t, want.Open, got.Open, "Open")
	assert.Equal(t, want.Close, got.Close, "Close")
	assert.Equal(t, want.High, got.High, "High")
	assert.Equal(t, want.Low, got.Low, "Low")
}

func BenchmarkAggregateTrades(b *testing.B) {
	trades := []int{14, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 13}
	for n := 0; n < b.N; n++ {
		candlestick.AggregateTrades(trades)
	}
}

func BenchmarkAggregateTradesSideEffects(b *testing.B) {
	trades := []int{14, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 13}
	for n := 0; n < b.N; n++ {
		candlestick.AggregateTradesSideEffects(trades)
	}
}
