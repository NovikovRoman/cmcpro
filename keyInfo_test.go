package cmcpro

import (
	"fmt"
	"github.com/NovikovRoman/cmcpro/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_KeyInfo(t *testing.T) {
	var (
		res    *types.KeyInfo
		status *types.Status
		err    error
	)
	res, status, err = cTest.KeyInfo()
	require.Nil(t, err)
	fmt.Println(res.Plan, res.Usage, status.CreditCount)
}
