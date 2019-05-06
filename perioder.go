package cmcpro

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

const MaxPeriodCount = 10000

type Perioder interface {
	AddQuery(query *url.Values, withoutTime bool)
	SetTimeEnd(*time.Time)
	SetTimeStart(*time.Time)
	SetCount(uint)
}

type Period struct {
	timeStart *time.Time
	timeEnd   *time.Time
	count     uint
}

func NewPeriod(timeStart *time.Time, timeEnd *time.Time, count uint) *Period {
	if count > MaxPeriodCount {
		count = MaxPeriodCount
	}
	return &Period{
		timeStart: timeStart,
		timeEnd:   timeEnd,
		count:     count,
	}
}

func (p *Period) AddQuery(query *url.Values, withoutTime bool) {
	if p.timeStart != nil {
		query.Add("time_start", p.timeStartToString(withoutTime))
	}

	if p.timeEnd != nil {
		query.Add("time_end", p.timeEndToString(withoutTime))
	}

	if p.count > 0 {
		query.Add("count", strconv.FormatUint(uint64(p.count), 10))
	}
}

func (p *Period) SetTimeStart(t *time.Time) {
	p.timeStart = t
}

func (p *Period) SetTimeEnd(t *time.Time) {
	p.timeEnd = t
}

func (p *Period) SetCount(count uint) {
	p.count = count
}

func (p *Period) timeStartToString(withoutTime bool) string {
	if p.timeStart == nil {
		return ""
	}

	dateTime := p.timeStart.Format(time.RFC3339)

	if withoutTime {
		dateTime = strings.Split(dateTime, "T")[0]
	}

	return dateTime
}

func (p *Period) timeEndToString(withoutTime bool) string {
	if p.timeEnd == nil {
		return ""
	}

	dateTime := p.timeEnd.Format(time.RFC3339)

	if withoutTime {
		dateTime = strings.Split(dateTime, "T")[0]
	}

	return dateTime
}
