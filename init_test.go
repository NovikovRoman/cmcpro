package cmcpro

import (
	"log"
	"os"
)

var (
	cTest *Client
)

func init() {
	var (
		err error
	)

	if cTest, err = New(os.Getenv("API-KEY"), false, "", Timeout); err != nil {
		log.Fatalln(err)
	}
}
