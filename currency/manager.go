package currency

import (
	"errors"

	"github.com/thrasher-/gocryptotrader/exchanges/assets"
)

// Get gets the currency pair config based on the asset type
func (p *PairsManager) Get(a assets.AssetType) *PairStore {
	p.m.Lock()
	defer p.m.Unlock()
	c, ok := p.Pairs[a]
	if !ok {
		return nil
	}
	return c
}

// Store stores a new currency pair config based on its asset type
func (p *PairsManager) Store(a assets.AssetType, ps PairStore) {
	p.m.Lock()

	if p.Pairs == nil {
		p.Pairs = make(map[assets.AssetType]*PairStore)
	}

	p.Pairs[a] = &ps
	p.m.Unlock()
}

// GetPairs gets a list of stored pairs based on the asset type and whether
// they're enabled or not
func (p *PairsManager) GetPairs(a assets.AssetType, enabled bool) (Pairs, error) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.Pairs == nil {
		return nil, errors.New("pairs manager not initialised")
	}

	c, ok := p.Pairs[a]
	if !ok {
		return nil, errors.New("pair manager specified asset type doesn't exist")
	}

	var pairs Pairs
	if enabled {
		pairs = c.Enabled
	} else {
		pairs = c.Available
	}

	return pairs, nil
}

// StorePairs stores a list of pairs based on the asset type and whether
// they're enabled or not
func (p *PairsManager) StorePairs(a assets.AssetType, pairs Pairs, enabled bool) error {
	p.m.Lock()
	defer p.m.Unlock()

	if p.Pairs == nil {
		return errors.New("pairs manager not initialised")
	}

	c, ok := p.Pairs[a]
	if !ok {
		c = new(PairStore)
	}

	if enabled {
		c.Enabled = pairs
	} else {
		c.Available = pairs
	}

	p.Pairs[a] = c
	return nil
}
