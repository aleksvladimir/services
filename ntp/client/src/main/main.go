package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/beevik/ntp"
)

func isNil(err error, host string) bool {
	switch {
	case err == nil:
		return true
	case strings.Contains(err.Error(), "timeout"):
		fmt.Printf("[%s] Query timeout: %s\n", host, err)
		return false
	case strings.Contains(err.Error(), "kiss of death"):
		fmt.Printf("[%s] Query kiss of death: %s\n", host, err)
		return false
	default:
		// error
		fmt.Fprintf(os.Stderr, "[%s] Query failed: %s\n", host, err)
		return false
	}
}

func main() {
	var host string = "0.beevik-ntp.pool.ntp.org"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}
	fmt.Printf("[%s] ----------------------\n", host)
	fmt.Printf("[%s] NTP protocol version %d\n", host, 4)

	r, err := ntp.QueryWithOptions(host, ntp.QueryOptions{Version: 4})
	if !isNil(err, host) {
		os.Exit(1)
	}

	fmt.Printf("[%s]  LocalTime: %v\n", host, time.Now())
	fmt.Printf("[%s]   XmitTime: %v\n", host, r.Time)
	fmt.Printf("[%s]    RefTime: %v\n", host, r.ReferenceTime)
	fmt.Printf("[%s]        RTT: %v\n", host, r.RTT)
	fmt.Printf("[%s]     Offset: %v\n", host, r.ClockOffset)
	fmt.Printf("[%s]       Poll: %v\n", host, r.Poll)
	fmt.Printf("[%s]  Precision: %v\n", host, r.Precision)
	fmt.Printf("[%s]    Stratum: %v\n", host, r.Stratum)
	fmt.Printf("[%s]      RefID: 0x%08x\n", host, r.ReferenceID)
	fmt.Printf("[%s]  RootDelay: %v\n", host, r.RootDelay)
	fmt.Printf("[%s]   RootDisp: %v\n", host, r.RootDispersion)
	fmt.Printf("[%s]   RootDist: %v\n", host, r.RootDistance)
	fmt.Printf("[%s]   MinError: %v\n", host, r.MinError)
	fmt.Printf("[%s]       Leap: %v\n", host, r.Leap)
	fmt.Printf("[%s]   KissCode: %v\n", host, stringOrEmpty(r.KissCode))
}

func stringOrEmpty(s string) string {
	if s == "" {
		return "<empty>"
	}
	return s
}
