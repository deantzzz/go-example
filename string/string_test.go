package string_test

import "testing"

func TestMarketMaker_SortString(t *testing.T) {
	usdt := "usdt"
	try := "try"
	eth := "eth"
	btc := "btc"
	if usdt <= try {
		t.Fatalf("expect %s > %s but not", usdt, try)
	}
	if usdt <= eth {
		t.Fatalf("expect %s > %s but not", usdt, eth)
	}
	if usdt <= btc {
		t.Fatalf("expect %s > %s but not", usdt, btc)
	}
	if try <= btc {
		t.Fatalf("expect %s > %s but not", try, btc)
	}
	if try <= eth {
		t.Fatalf("expect %s > %s but not", try, eth)
	}
	if eth <= btc {
		t.Fatalf("expect %s > %s but not", eth, btc)
	}
}
