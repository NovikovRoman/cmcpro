package cmcpro

import (
	"testing"

	"github.com/NovikovRoman/cmcpro/types"
	"github.com/stretchr/testify/assert"
)

func TestClient_ToolsPriceConversion(t *testing.T) {
	res, status, _ := cTest.ToolsPriceConversionByID(
		contextTest, 10, 1, NewConvertByCodes("RUB"), nil,
	)

	assert.Equal(t, res.Name, "Bitcoin")
	assert.Equal(t, res.Amount, 10.0)
	assert.True(t, res.Quote["RUB"].Price > 0)
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 1)
	}

	var res2 []types.PriceConversion
	res2, status, _ = cTest.ToolsPriceConversionBySymbol(
		contextTest, 10, "BTC", NewConvertByCodes("RUB"), nil,
	)

	assert.Equal(t, res2[0].Name, "Bitcoin")
	assert.Equal(t, res2[0].Amount, 10.0)
	assert.True(t, res2[0].Quote["RUB"].Price > 0)
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 1)
	}
}
