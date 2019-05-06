package cmcpro

import (
	"net/url"
	"strconv"
	"strings"
)

type Converter interface {
	IsByID() bool
	IsByCode() bool
	ID() []int
	Code() []string
	AddQuery(query *url.Values)
}

type Convert struct {
	ids   []int
	codes []string
}

func NewConvertByID(id ...int) *Convert {
	return &Convert{
		ids: id,
	}
}

func NewConvertByCodes(code ...string) *Convert {
	return &Convert{
		codes: code,
	}
}

func (c *Convert) ID() []int {
	return c.ids
}

func (c *Convert) Code() []string {
	return c.codes
}

func (c *Convert) IsByID() bool {
	return len(c.ids) > 0
}

func (c *Convert) IsByCode() bool {
	return len(c.codes) > 0
}

func (c *Convert) AddQuery(query *url.Values) {
	if len(c.codes) > 0 {
		query.Add("convert", strings.Join(c.codes, ","))
		return
	}

	ids := make([]string, len(c.ids))
	for k, v := range c.ids {
		ids[k] = strconv.Itoa(v)
	}
	query.Add("convert_id", strings.Join(ids, ","))
}
