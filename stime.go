package stime

import (
	"github.com/beevik/ntp"
	"time"
)

var response = ntpResponse()

func ntpResponse() ntp.Response {
	response, err := ntp.Query("time.google.com")
	if err != nil {
		return ntp.Response{}
	}
	return *response
}

func Now() time.Time {
	return time.Now().Add(response.ClockOffset).UTC()
}