package currency

import (
	"testing"

	"github.com/thrasher-/gocryptotrader/exchanges/assets"
)

var p PairsManager

func initTest(t *testing.T) {
	p.Store(assets.AssetTypeSpot,
		PairStore{
			Available: NewPairsFromStrings([]string{"BTC-USD", "LTC-USD"}),
			Enabled:   NewPairsFromStrings([]string{"BTC-USD"}),
			RequestFormat: &PairFormat{
				Uppercase: true,
			},
			ConfigFormat: &PairFormat{
				Uppercase: true,
				Delimiter: "-",
			},
		},
	)
}

func TestGet(t *testing.T) {
	initTest(t)

	if p.Get(assets.AssetTypeSpot) == nil {
		t.Error("Test failed. Spot assets shouldn't be nil")
	}

	if p.Get(assets.AssetTypeFutures) != nil {
		t.Error("Test Failed. Futures should be nil")
	}
}

func TestStore(t *testing.T) {
	p.Store(assets.AssetTypeFutures,
		PairStore{
			Available: NewPairsFromStrings([]string{"BTC-USD", "LTC-USD"}),
			Enabled:   NewPairsFromStrings([]string{"BTC-USD"}),
			RequestFormat: &PairFormat{
				Uppercase: true,
			},
			ConfigFormat: &PairFormat{
				Uppercase: true,
				Delimiter: "-",
			},
		},
	)

	if p.Get(assets.AssetTypeFutures) == nil {
		t.Error("Test failed. Futures assets shouldn't be nil")
	}
}

func TestGetPairs(t *testing.T) {
	p.Pairs = nil
	_, err := p.GetPairs(assets.AssetTypeSpot, true)
	if err == nil {
		t.Errorf("GetPairs should throw an error but didn't")
	}

	initTest(t)
	_, err = p.GetPairs(assets.AssetTypeSpot, true)
	if err != nil {
		t.Errorf("TestGetPairs failed, error %v", err)
	}

	_, err = p.GetPairs("blah", true)
	if err == nil {
		t.Errorf("TestGetPairs return a result when it should be nil")
	}
}

func TestStorePairs(t *testing.T) {
	p.Pairs = nil
	err := p.StorePairs(assets.AssetTypeSpot, nil, false)
	if err == nil {
		t.Errorf("GetPairs should throw an error but didn't")
	}

	initTest(t)
	err = p.StorePairs(assets.AssetTypeSpot, NewPairsFromStrings([]string{"ETH-USD"}), false)
	if err != nil {
		t.Errorf("TestStorePairs failed, error %v", err)
	}

	pairs, err := p.GetPairs(assets.AssetTypeSpot, false)
	if err != nil {
		t.Errorf("TestStorePairs failed, error %v", err)
	}

	if !pairs.Contains(NewPairFromString("ETH-USD"), true) {
		t.Errorf("TestStorePairs failed, unexpected result")
	}

	err = p.StorePairs(assets.AssetTypeFutures, NewPairsFromStrings([]string{"ETH-KRW"}), true)
	if err != nil {
		t.Errorf("TestStorePairs failed, error %v", err)
	}

	pairs, err = p.GetPairs(assets.AssetTypeFutures, true)
	if err != nil {
		t.Errorf("TestStorePairs failed, error %v", err)
	}

	if !pairs.Contains(NewPairFromString("ETH-KRW"), true) {
		t.Errorf("TestStorePairs failed, unexpected result")
	}
}
