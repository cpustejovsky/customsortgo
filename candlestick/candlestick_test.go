package candlestick_test

import (
	"github.com/cpustejovsky/customsortgo/candlestick"
	"testing"
)

func TestAggregateTrades(t *testing.T) {
	trades := []float64{14.0, 14.0, 15.0, 15.0, 16.0, 15.0, 13.0, 12.0, 13.0}
	want := candlestick.TradeAggregate{Open: 14, Close: 13, High: 16, Low: 12}
	got := candlestick.NewTradeAggregate(trades)
	if want.Low != got.Low {
		t.Errorf("got:\n%+v\nwanted:\n%+v\n", got, want)
	}
	if want.High != got.High {
		t.Errorf("got:\n%+v\nwanted:\n%+v\n", got, want)
	}
	if want.Open != got.Open {
		t.Errorf("got:\n%+v\nwanted:\n%+v\n", got, want)
	}
	if want.Close != got.Close {
		t.Errorf("got:\n%+v\nwanted:\n%+v\n", got, want)
	}
	if trades[0] != 14.0 {
		t.Error("first trade should equal 14.0")
	}
}

func BenchmarkAggregateTrades(b *testing.B) {
	trades := []float64{14, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 13}
	for n := 0; n < b.N; n++ {
		candlestick.NewTradeAggregate(trades)
	}
}

func BenchmarkAggregateTradesSideEffects(b *testing.B) {
	trades := []float64{14, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 14, 15, 15, 16, 15, 13, 12, 13}
	for n := 0; n < b.N; n++ {
		candlestick.AggregateTrades(trades)
	}
}
