package converter

import (
	"context"
	"fmt"
	"time"

	cex "github.com/binance/binance-connector-go"
	"github.com/f-taxes/binance_conversion/global"
	"github.com/shopspring/decimal"
	"go.uber.org/ratelimit"
)

var limiter = ratelimit.New(60, ratelimit.Per(time.Minute))
var srv = cex.NewClient("", "").NewKlinesService()

func PriceAtTime(asset, targetCurrency string, ts time.Time) (decimal.Decimal, error) {
	limiter.Take()
	ts = global.StartOfMinute(ts)

	srv.Symbol(fmt.Sprintf("%s%s", asset, targetCurrency))
	srv.StartTime(uint64(ts.Add(-5 * time.Minute).UnixMilli()))
	srv.Interval("1m")
	srv.Limit(10)
	resp, err := srv.Do(context.Background())

	if err != nil {
		return decimal.Zero, err
	}

	for _, rec := range resp {
		t := time.UnixMilli(int64(rec.OpenTime))
		if ts.Equal(t) {
			return global.StrToDecimal(rec.Open), nil
		}
	}

	return decimal.Zero, nil
}
