package main

import (
	"fmt"
	"strconv"
)

const TAB = "\t"

const (
	DATANULL = "DATANULL"
	AM = "AM"
	FM = "FM"
	DAB = "DAB"
	ANALOG = "ANALOG"
	HD = "HD"
	ALL = "ALL"
)

type Channel struct {
	Band            string
	Signal          string
	Frequency       string
	Latitude        float32
	Longitude       float32
	CallSign        string
	PiCode          string
	MulticastNumber int
	Eid             string
	Sid             string
	SCIdS           string
	BroadcastId     string
}

func NewChannel(band, signal, frequency string, latitude, longitude float32, multicastNumber int, eid, sid, sCIdS string) *Channel {
	return &Channel{
		Band:            band,
		Signal:          signal,
		Frequency:       frequency,
		Latitude:        latitude,
		Longitude:       longitude,
		CallSign:        DATANULL,
		PiCode:          DATANULL,
		MulticastNumber: multicastNumber,
		Eid:             eid,
		Sid:             sid,
		SCIdS:           sCIdS,
		BroadcastId:     DATANULL,
	}
}

func (c *Channel) String() string {
	return c.Band + TAB + c.Signal + TAB + c.Frequency +  TAB + fmt.Sprintf("%.3f", c.Latitude) +
		TAB + fmt.Sprintf("%.3f", c.Longitude) + TAB + c.CallSign + TAB + c.PiCode + TAB +
		strconv.Itoa(c.MulticastNumber) + TAB + c.Eid + TAB + c.Sid + TAB + c.SCIdS + TAB + c.BroadcastId
}
