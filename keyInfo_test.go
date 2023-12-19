package cmcpro

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_KeyInfo(t *testing.T) {
	res, status, err := cTest.KeyInfo(contextTest)
	require.Nil(t, err)

	require.NotNil(t, res.Usage)
	require.NotNil(t, res.Plan)
	require.EqualValues(t, status.CreditCount, 0)
}
