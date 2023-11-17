package cmcpro

import (
	"testing"

	"github.com/NovikovRoman/cmcpro/types"
	"github.com/stretchr/testify/require"
)

func TestClient_KeyInfo(t *testing.T) {
	var (
		res    *types.KeyInfo
		status *types.Status
		err    error
	)
	res, status, err = cTest.KeyInfo(contextTest)
	require.Nil(t, err)

	require.NotNil(t, res.Usage)
	require.NotNil(t, res.Plan)
	require.EqualValues(t, status.CreditCount, 0)
}
