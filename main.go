package main

import (
	"fmt"

	"github.com/docktermj/go-proc-parse/proc/meminfo"
	"github.com/docktermj/go-proc-parse/proc/net/dev"
	"github.com/docktermj/go-proc-parse/proc/net/snmp"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

func displayBanner(title string) {
	fmt.Printf("\n---------- %s ------------------------------\n\n", title)
}

func demoProcMeminfo() {
	displayBanner("/proc/meminfo")
	contents, _ := meminfo.Get()
	fmt.Println(contents)
	fmt.Printf("\nMemory Total:  %d\n", contents["MemTotal"])
}

func demoProcNetDev() {
	displayBanner("/proc/net/dev")
	contents, _ := dev.Get()
	fmt.Println(contents)
	fmt.Printf("\nReceived bytes:  %d\n", contents["lo"]["receive-bytes"])
}

func demoProcNetSnmp() {
	displayBanner("/proc/net/snmp")
	contents, _ := snmp.Get()
	fmt.Println(contents)
	fmt.Printf("\nOutbound datagrams:  %d\n", contents["Udp"]["OutDatagrams"])
}

func main() {
	demoProcMeminfo()
	demoProcNetDev()
	demoProcNetSnmp()
}
