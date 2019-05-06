package cmcpro

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	ApiKey   string
	ApiPoint string

	transport *http.Transport
	timeout   time.Duration
}

func New(apiKey string, production bool, proxy string, timeout time.Duration) (*Client, error) {
	transport, err := createHttpTransport(proxy)
	if err != nil {
		return nil, err
	}
	apiPoint := ""
	if production {
		apiPoint = ApiPoint
	} else {
		apiPoint = ApiPointTest
	}
	return &Client{ApiKey: apiKey, ApiPoint: apiPoint, timeout: timeout, transport: transport}, nil
}

func (c *Client) exec(req *http.Request, obj interface{}) error {
	b, err := c.do(req)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}
	return nil
}

func (c *Client) createRequest(link string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.ApiPoint, link), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-CMC_PRO_API_KEY", c.ApiKey)
	return req, nil
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	client := &http.Client{
		Timeout:   c.timeout,
		Transport: c.transport,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer closer(resp.Body)
	return ioutil.ReadAll(resp.Body)
}

func createHttpTransport(proxy string) (*http.Transport, error) {
	transport := &http.Transport{}

	if proxy != "" {
		proxyUrl, err := url.Parse(proxy)
		if err != nil {
			log.Fatalf("Problem parse proxy %s", err)
		}
		transport.Proxy = http.ProxyURL(proxyUrl)
	}

	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	return transport, nil
}

func closer(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
