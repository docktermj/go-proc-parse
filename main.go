package main

import (
	"fmt"
	"strconv"

	"github.com/docktermj/go-proc-parse/proc/_pid_/stat"
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

func demoProcPidStat() {
	pid := int64(5)
	displayBanner("/proc/" + strconv.FormatInt(pid, 10) + "/stat")
	contents, _ := stat.Get(pid)
	contentsAsJson, _ := stat.GetAsJson(pid)
	contentsAsMap, _ := stat.GetAsMap(pid)
	fmt.Printf("%+v\n\n", contents)
	fmt.Printf("%s\n\n", contentsAsJson)
	fmt.Printf("%+v\n\n", contentsAsMap)
	fmt.Printf("Stat:  %d\n", contents.Blocked)
}

func demoProcMeminfo() {
	displayBanner("/proc/meminfo")
	contents, _ := meminfo.Get()
	contentsAsJson, _ := meminfo.GetAsJson()
	contentsAsMap, _ := meminfo.GetAsMap()
	fmt.Printf("%+v\n\n", contents)
	fmt.Printf("%s\n\n", contentsAsJson)
	fmt.Printf("%+v\n\n", contentsAsMap)
	fmt.Printf("\nMemory Available:  %d\n", contents.MemAvailable)
}

func demoProcNetDev() {
	displayBanner("/proc/net/dev")
	contents, _ := dev.Get()
	contentsAsJson, _ := dev.GetAsJson()
	contentsAsMap, _ := dev.GetAsMap()
	fmt.Printf("%+v\n\n", contents)
	fmt.Printf("%s\n\n", contentsAsJson)
	fmt.Printf("%+v\n\n", contentsAsMap)
	fmt.Printf("\nReceived bytes:  %d\n", contentsAsMap["lo"]["receive-bytes"])
}

func demoProcNetSnmp() {
	displayBanner("/proc/net/snmp")
	contentsAsJson, _ := snmp.GetAsJson()
	contentsAsMap, _ := snmp.GetAsMap()
	fmt.Printf("%s\n\n", contentsAsJson)
	fmt.Printf("%+v\n\n", contentsAsMap)
	fmt.Printf("\nOutbound datagrams:  %d\n", contentsAsMap["Udp"]["OutDatagrams"])
}

func main() {
	demoProcPidStat()
	demoProcMeminfo()
	demoProcNetDev()
	demoProcNetSnmp()
}
