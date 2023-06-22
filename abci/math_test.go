package abci_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/holiman/uint256"
	"github.com/skip-mev/slinky/abci"
	oracletypes "github.com/skip-mev/slinky/oracle/types"
)

type validator struct {
	stake   math.Int
	address sdk.ValAddress
}

var (
	validator1 = sdk.ValAddress("validator1")
	validator2 = sdk.ValAddress("validator2")
	validator3 = sdk.ValAddress("validator3")
)

func (suite *ABCITestSuite) TestStakeWeightedMedian() {
	cases := []struct {
		name              string
		providerPrices    oracletypes.AggregatedProviderPrices
		validators        []validator
		totalBondedTokens math.Int
		expectedPrices    map[oracletypes.CurrencyPair]*uint256.Int
	}{
		{
			name:              "no providers",
			providerPrices:    oracletypes.AggregatedProviderPrices{},
			validators:        []validator{},
			totalBondedTokens: math.NewInt(100),
			expectedPrices:    map[oracletypes.CurrencyPair]*uint256.Int{},
		},
		{
			name: "single provider entire stake + single price",
			providerPrices: oracletypes.AggregatedProviderPrices{
				validator1.String(): map[oracletypes.CurrencyPair]oracletypes.QuotePrice{
					{
						Base:          "BTC",
						Quote:         "USD",
						QuoteDecimals: 8,
					}: {
						Price: uint256.NewInt(100),
					},
				},
			},
			validators: []validator{
				{
					stake:   math.NewInt(100),
					address: validator1,
				},
			},
			totalBondedTokens: math.NewInt(100),
			expectedPrices: map[oracletypes.CurrencyPair]*uint256.Int{
				{
					Base:          "BTC",
					Quote:         "USD",
					QuoteDecimals: 8,
				}: uint256.NewInt(100),
			},
		},
		{
			name: "single provider with not enough stake + single price",
			providerPrices: oracletypes.AggregatedProviderPrices{
				validator1.String(): map[oracletypes.CurrencyPair]oracletypes.QuotePrice{
					{
						Base:          "BTC",
						Quote:         "USD",
						QuoteDecimals: 8,
					}: {
						Price: uint256.NewInt(100),
					},
				},
			},
			validators: []validator{
				{
					stake:   math.NewInt(50),
					address: validator1,
				},
			},
			totalBondedTokens: math.NewInt(100),
			expectedPrices:    map[oracletypes.CurrencyPair]*uint256.Int{},
		},
		{
			name: "single provider with just enough stake + multiple prices",
			providerPrices: oracletypes.AggregatedProviderPrices{
				validator1.String(): map[oracletypes.CurrencyPair]oracletypes.QuotePrice{
					{
						Base:          "BTC",
						Quote:         "USD",
						QuoteDecimals: 8,
					}: {
						Price: uint256.NewInt(100),
					},
					{
						Base:          "ETH",
						Quote:         "USD",
						QuoteDecimals: 8,
					}: {
						Price: uint256.NewInt(200),
					},
				},
			},
			validators: []validator{
				{
					stake:   math.NewInt(68),
					address: validator1,
				},
			},
			totalBondedTokens: math.NewInt(100),
			expectedPrices: map[oracletypes.CurrencyPair]*uint256.Int{
				{
					Base:          "BTC",
					Quote:         "USD",
					QuoteDecimals: 8,
				}: uint256.NewInt(100),
				{
					Base:          "ETH",
					Quote:         "USD",
					QuoteDecimals: 8,
				}: uint256.NewInt(200),
			},
		},
		{
			name: "2 providers with equal stake + single asset",
			providerPrices: oracletypes.AggregatedProviderPrices{
				validator1.String(): map[oracletypes.CurrencyPair]oracletypes.QuotePrice{
					{
						Base:          "BTC",
						Quote:         "USD",
						QuoteDecimals: 8,
					}: {
						Price: uint256.NewInt(100),
					},
				},
				validator2.String(): map[oracletypes.CurrencyPair]oracletypes.QuotePrice{
					{
						Base:          "BTC",
						Quote:         "USD",
						QuoteDecimals: 8,
					}: {
						Price: uint256.NewInt(200),
					},
				},
			},
			validators: []validator{
				{
					stake:   math.NewInt(50),
					address: validator1,
				},
				{
					stake:   math.NewInt(50),
					address: validator2,
				},
			},
			totalBondedTokens: math.NewInt(100),
			expectedPrices: map[oracletypes.CurrencyPair]*uint256.Int{
				{
					Base:          "BTC",
					Quote:         "USD",
					QuoteDecimals: 8,
				}: uint256.NewInt(100),
			},
		},
		{
			name: "3 providers with equal stake + single asset",
			providerPrices: oracletypes.AggregatedProviderPrices{
				validator1.String(): map[oracletypes.CurrencyPair]oracletypes.QuotePrice{
					{
						Base:          "BTC",
						Quote:         "USD",
						QuoteDecimals: 8,
					}: {
						Price: uint256.NewInt(100),
					},
				},
				validator2.String(): map[oracletypes.CurrencyPair]oracletypes.QuotePrice{
					{
						Base:          "BTC",
						Quote:         "USD",
						QuoteDecimals: 8,
					}: {
						Price: uint256.NewInt(200),
					},
				},
				validator3.String(): map[oracletypes.CurrencyPair]oracletypes.QuotePrice{
					{
						Base:          "BTC",
						Quote:         "USD",
						QuoteDecimals: 8,
					}: {
						Price: uint256.NewInt(300),
					},
				},
			},
			validators: []validator{
				{
					stake:   math.NewInt(33),
					address: validator1,
				},
				{
					stake:   math.NewInt(33),
					address: validator2,
				},
				{
					stake:   math.NewInt(33),
					address: validator3,
				},
			},
			totalBondedTokens: math.NewInt(100),
			expectedPrices: map[oracletypes.CurrencyPair]*uint256.Int{
				{
					Base:          "BTC",
					Quote:         "USD",
					QuoteDecimals: 8,
				}: uint256.NewInt(200),
			},
		},
		{
			name: "3 providers with equal stake + multiple assets",
			providerPrices: oracletypes.AggregatedProviderPrices{
				validator1.String(): map[oracletypes.CurrencyPair]oracletypes.QuotePrice{
					{
						Base:          "BTC",
						Quote:         "USD",
						QuoteDecimals: 8,
					}: {
						Price: uint256.NewInt(100),
					},
					{
						Base:          "ETH",
						Quote:         "USD",
						QuoteDecimals: 8,
					}: {
						Price: uint256.NewInt(200),
					},
				},
				validator2.String(): map[oracletypes.CurrencyPair]oracletypes.QuotePrice{
					{
						Base:          "BTC",
						Quote:         "USD",
						QuoteDecimals: 8,
					}: {
						Price: uint256.NewInt(300),
					},
					{
						Base:          "ETH",
						Quote:         "USD",
						QuoteDecimals: 8,
					}: {
						Price: uint256.NewInt(400),
					},
				},
				validator3.String(): map[oracletypes.CurrencyPair]oracletypes.QuotePrice{
					{
						Base:          "BTC",
						Quote:         "USD",
						QuoteDecimals: 8,
					}: {
						Price: uint256.NewInt(500),
					},
				},
			},
			validators: []validator{
				{
					stake:   math.NewInt(33),
					address: validator1,
				},
				{
					stake:   math.NewInt(33),
					address: validator2,
				},
				{
					stake:   math.NewInt(33),
					address: validator3,
				},
			},
			totalBondedTokens: math.NewInt(100),
			expectedPrices: map[oracletypes.CurrencyPair]*uint256.Int{ // only btc/usd should be included
				{
					Base:          "BTC",
					Quote:         "USD",
					QuoteDecimals: 8,
				}: uint256.NewInt(300),
			},
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			// Create a mock validator store.
			mockValidatorStore := suite.createMockValidatorStore(tc.validators, tc.totalBondedTokens)

			// Compute the stake weighted median.
			aggregateFn := abci.StakeWeightedMedian(suite.ctx, mockValidatorStore, abci.DefaultPowerThreshold)
			result := aggregateFn(tc.providerPrices)

			// Verify the result.
			suite.Require().Len(result, len(tc.expectedPrices))
			for currencyPair, expectedPrice := range tc.expectedPrices {
				suite.Require().Equal(expectedPrice, result[currencyPair])
			}
		})
	}
}

func (suite *ABCITestSuite) TestComputeStakeWeightedMedian() {
	cases := []struct {
		name      string
		priceInfo abci.StakeWeightPriceInfo
		expected  *uint256.Int
	}{
		{
			name: "single price",
			priceInfo: abci.StakeWeightPriceInfo{
				Prices: []abci.StakeWeightPrice{
					{
						StakeWeight: sdk.NewInt(1),
						Price:       uint256.NewInt(100),
					},
				},
				TotalWeight: sdk.NewInt(1),
			},
			expected: uint256.NewInt(100),
		},
		{
			name: "two prices that are equal",
			priceInfo: abci.StakeWeightPriceInfo{
				Prices: []abci.StakeWeightPrice{
					{
						StakeWeight: sdk.NewInt(1),
						Price:       uint256.NewInt(100),
					},
					{
						StakeWeight: sdk.NewInt(1),
						Price:       uint256.NewInt(100),
					},
				},
				TotalWeight: sdk.NewInt(2),
			},
			expected: uint256.NewInt(100),
		},
		{
			name: "two prices that are not equal",
			priceInfo: abci.StakeWeightPriceInfo{
				Prices: []abci.StakeWeightPrice{
					{
						StakeWeight: sdk.NewInt(1),
						Price:       uint256.NewInt(100),
					},
					{
						StakeWeight: sdk.NewInt(1),
						Price:       uint256.NewInt(200),
					},
				},
				TotalWeight: sdk.NewInt(2),
			},
			expected: uint256.NewInt(100),
		},
		{
			name: "two prices that are not equal with different weights",
			priceInfo: abci.StakeWeightPriceInfo{
				Prices: []abci.StakeWeightPrice{
					{
						StakeWeight: sdk.NewInt(10),
						Price:       uint256.NewInt(100),
					},
					{
						StakeWeight: sdk.NewInt(20),
						Price:       uint256.NewInt(200),
					},
				},
				TotalWeight: sdk.NewInt(30),
			},
			expected: uint256.NewInt(200),
		},
		{
			name: "three prices that are not equal with different weights",
			priceInfo: abci.StakeWeightPriceInfo{
				Prices: []abci.StakeWeightPrice{
					{
						StakeWeight: sdk.NewInt(10),
						Price:       uint256.NewInt(100),
					},
					{
						StakeWeight: sdk.NewInt(20),
						Price:       uint256.NewInt(200),
					},
					{
						StakeWeight: sdk.NewInt(30),
						Price:       uint256.NewInt(300),
					},
				},
				TotalWeight: sdk.NewInt(60),
			},
			expected: uint256.NewInt(200),
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			result := abci.ComputeStakeWeightedMedian(tc.priceInfo)
			suite.Require().Equal(tc.expected, result)
		})
	}
}