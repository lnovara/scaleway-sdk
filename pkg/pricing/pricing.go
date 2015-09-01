package pricing

import (
	"math/big"
	"time"
)

type PricingObject struct {
	Path             string
	Identifier       string
	Currency         string
	UsageUnit        string
	UnitPrice        *big.Rat
	UnitQuantity     *big.Rat
	UnitPriceCap     *big.Rat
	UsageGranularity time.Duration
}

type PricingList []PricingObject

// CurrentPricing tries to be up-to-date with the real pricing
// we cannot guarantee of these values since we hardcode values for now
// later, we should be able to call a dedicated pricing API
var CurrentPricing PricingList

func init() {
	CurrentPricing = PricingList{
		{
			Path:             "/compute/c1/run",
			Identifier:       "aaaaaaaa-aaaa-4aaa-8aaa-111111111111",
			Currency:         "EUR",
			UnitPrice:        big.NewRat(12, 1000),    // 0.012
			UnitQuantity:     big.NewRat(60000, 1000), // 60
			UnitPriceCap:     big.NewRat(6000, 1000),  // 6
			UsageGranularity: time.Minute,
		},
		{
			Path:             "/ip/dynamic",
			Identifier:       "467116bf-4631-49fb-905b-e07701c2db11",
			Currency:         "EUR",
			UnitPrice:        big.NewRat(4, 1000),     // 0.004
			UnitQuantity:     big.NewRat(60000, 1000), // 60
			UnitPriceCap:     big.NewRat(1999, 1000),  // 1.99
			UsageGranularity: time.Minute,
		},
		{
			Path:             "/ip/reserved",
			Identifier:       "467116bf-4631-49fb-905b-e07701c2db22",
			Currency:         "EUR",
			UnitPrice:        big.NewRat(4, 1000),     // 0.004
			UnitQuantity:     big.NewRat(60000, 1000), // 60
			UnitPriceCap:     big.NewRat(1990, 1000),  // 1.99
			UsageGranularity: time.Minute,
		},
		{
			Path:             "/storage/local/ssd/storage",
			Identifier:       "bbbbbbbb-bbbb-4bbb-8bbb-111111111113",
			Currency:         "EUR",
			UnitPrice:        big.NewRat(4, 1000),     // 0.004
			UnitQuantity:     big.NewRat(50000, 1000), // 50
			UnitPriceCap:     big.NewRat(2000, 1000),  // 2
			UsageGranularity: time.Hour,
		},
	}
}

// GetByPath returns an object matching a path
func (pl *PricingList) GetByPath(path string) *PricingObject {
	for _, object := range *pl {
		if object.Path == path {
			return &object
		}
	}
	return nil
}

// GetByIdentifier returns an object matching a identifier
func (pl *PricingList) GetByIdentifier(identifier string) *PricingObject {
	for _, object := range *pl {
		if object.Identifier == identifier {
			return &object
		}
	}
	return nil
}
