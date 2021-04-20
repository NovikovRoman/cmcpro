package cmcpro

import (
	"context"
	"github.com/NovikovRoman/cmcpro/types"
	"net/http"
)

func (c *Client) KeyInfo(ctx context.Context) (keyInfo *types.KeyInfo, status *types.Status, err error) {
	var (
		req *http.Request
	)

	if req, err = c.createRequest(ctx, "/key/info"); err != nil {
		return
	}

	respInfo := struct {
		Data   *types.KeyInfo `json:"data"`
		Status *types.Status  `json:"status"`
	}{}

	if err = c.exec(req, &respInfo); err != nil {
		return
	}

	keyInfo = respInfo.Data
	status = respInfo.Status
	return
}
