package cmcpro

import (
	"context"
	"log"
	"os"
)

var (
	cTest       *Client
	contextTest context.Context
	prodTest    bool
)

func init() {
	var (
		err error
	)
	prodTest = os.Getenv("PRODUCTION") != ""
	if cTest, err = New(os.Getenv("API-KEY"), prodTest, "", Timeout); err != nil {
		log.Fatalln(err)
	}

	contextTest = context.Background()
}
