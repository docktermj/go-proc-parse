package meminfo

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

// References:
// - http://man7.org/linux/man-pages/man5/proc.5.html  "/proc/meminfo"
type Meminfo struct {
	MemTotal          uint64 `json:"MemTotal"`
	MemFree           uint64 `json:"MemFree"`
	MemAvailable      uint64 `json:"MemAvailable"`
	Buffers           uint64 `json:"Buffers"`
	Cached            uint64 `json:"Cached"`
	SwapCached        uint64 `json:"SwapCached"`
	Active            uint64 `json:"Active"`
	Inactive          uint64 `json:"Inactive"`
	Active_anon       uint64 `json:"Active_anon"`
	Inactive_anon     uint64 `json:"Inactive_anon"`
	Active_file       uint64 `json:"Active_file"`
	Inactive_file     uint64 `json:"Inactive_file"`
	Unevictable       uint64 `json:"Unevictable"`
	Mlocked           uint64 `json:"Mlocked"`
	HighTotal         uint64 `json:"HighTotal"`
	HighFree          uint64 `json:"HighFree"`
	LowTotal          uint64 `json:"LowTotal"`
	LowFree           uint64 `json:"LowFree"`
	MmapCopy          uint64 `json:"MmapCopy"`
	SwapTotal         uint64 `json:"SwapTotal"`
	SwapFree          uint64 `json:"SwapFree"`
	Dirty             uint64 `json:"Dirty"`
	Writeback         uint64 `json:"Writeback"`
	AnonPages         uint64 `json:"AnonPages"`
	Mapped            uint64 `json:"Mapped"`
	Shmem             uint64 `json:"Shmem"`
	Slab              uint64 `json:"Slab"`
	SReclaimable      uint64 `json:"SReclaimable"`
	SUnreclaim        uint64 `json:"SUnreclaim"`
	KernelStack       uint64 `json:"KernelStack"`
	PageTables        uint64 `json:"PageTables"`
	Quicklists        uint64 `json:"Quicklists"`
	NFS_Unstable      uint64 `json:"NFS_Unstable"`
	Bounce            uint64 `json:"Bounce"`
	WritebackTmp      uint64 `json:"WritebackTmp"`
	CommitLimit       uint64 `json:"CommitLimit"`
	Committed_AS      uint64 `json:"Committed_AS"`
	VmallocTotal      uint64 `json:"VmallocTotal"`
	VmallocUsed       uint64 `json:"VmallocUsed"`
	VmallocChunk      uint64 `json:"VmallocChunk"`
	HardwareCorrupted uint64 `json:"HardwareCorrupted"`
	AnonHugePages     uint64 `json:"AnonHugePages"`
	ShmemHugePages    uint64 `json:"ShmemHugePages"`
	ShmemPmdMapped    uint64 `json:"ShmemPmdMapped"`
	CmaTotal          uint64 `json:"CmaTotal"`
	CmaFree           uint64 `json:"CmaFree"`
	HugePages_Total   uint64 `json:"HugePages_Total"`
	HugePages_Free    uint64 `json:"HugePages_Free"`
	HugePages_Rsvd    uint64 `json:"HugePages_Rsvd"`
	HugePages_Surp    uint64 `json:"HugePages_Surp"`
	Hugepagesize      uint64 `json:"Hugepagesize"`
	DirectMap4k       uint64 `json:"DirectMap4k"`
	DirectMap4M       uint64 `json:"DirectMap4M"`
	DirectMap2M       uint64 `json:"DirectMap2M"`
	DirectMap1G       uint64 `json:"DirectMap1G"`
}

// Allow filename to be specified by OS Environment variable: PROC_NET_SNMP
func GetFilename() string {
	result := os.Getenv("PROC_MEMINFO")
	if result == "" {
		result = "/proc/meminfo"
	}
	return result
}

func asUint64(value string) uint64 {
	result, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return uint64(0)
	}
	return result
}

// Get values of /proc/meminfo as a map of uint64.
// Example:
//     myMeminfo := meminfo.Get()
//     x := myMeminfo["MemTotal"]
func Get() (Meminfo, error) {

	result := Meminfo{}

	// Open the file.

	fileName := GetFilename()
	file, err := os.Open(fileName)
	if err != nil {
		return result, err
	}

	// Read the file.

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine := scanner.Text()
		splits := strings.Fields(inputLine)

		// Pull out the key.

		keySplits := strings.Split(splits[0], ":")
		key := keySplits[0]

		// Pull out the value.

		value, err := strconv.ParseUint(splits[1], 10, 64)
		if err != nil {
			continue
		}

		switch key {
		case "MemTotal":
			result.MemTotal = value
		case "MemFree":
			result.MemFree = value
		case "MemAvailable":
			result.MemAvailable = value
		case "Buffers":
			result.Buffers = value
		case "Cached":
			result.Cached = value
		case "SwapCached":
			result.SwapCached = value
		case "Active":
			result.Active = value
		case "Inactive":
			result.Inactive = value
		case "Active(anon)":
			result.Active_anon = value
		case "Inactive(anon)":
			result.Inactive_anon = value
		case "Active(file)":
			result.Active_file = value
		case "Inactive(file)":
			result.Inactive_file = value
		case "Unevictable":
			result.Unevictable = value
		case "Mlocked":
			result.Mlocked = value
		case "HighTotal":
			result.HighTotal = value
		case "HighFree":
			result.HighFree = value
		case "LowTotal":
			result.LowTotal = value
		case "LowFree":
			result.LowFree = value
		case "MmapCopy":
			result.MmapCopy = value
		case "SwapTotal":
			result.SwapTotal = value
		case "SwapFree":
			result.SwapFree = value
		case "Dirty":
			result.Dirty = value
		case "Writeback":
			result.Writeback = value
		case "AnonPages":
			result.AnonPages = value
		case "Mapped":
			result.Mapped = value
		case "Shmem":
			result.Shmem = value
		case "Slab":
			result.Slab = value
		case "SReclaimable":
			result.SReclaimable = value
		case "SUnreclaim":
			result.SUnreclaim = value
		case "KernelStack":
			result.KernelStack = value
		case "PageTables":
			result.PageTables = value
		case "Quicklists":
			result.Quicklists = value
		case "NFS_Unstable":
			result.NFS_Unstable = value
		case "Bounce":
			result.Bounce = value
		case "WritebackTmp":
			result.WritebackTmp = value
		case "CommitLimit":
			result.CommitLimit = value
		case "Committed_AS":
			result.Committed_AS = value
		case "VmallocTotal":
			result.VmallocTotal = value
		case "VmallocUsed":
			result.VmallocUsed = value
		case "VmallocChunk":
			result.VmallocChunk = value
		case "HardwareCorrupted":
			result.HardwareCorrupted = value
		case "AnonHugePages":
			result.AnonHugePages = value
		case "ShmemHugePages":
			result.ShmemHugePages = value
		case "ShmemPmdMapped":
			result.ShmemPmdMapped = value
		case "CmaTotal":
			result.CmaTotal = value
		case "CmaFree":
			result.CmaFree = value
		case "HugePages_Total":
			result.HugePages_Total = value
		case "HugePages_Free":
			result.HugePages_Free = value
		case "HugePages_Rsvd":
			result.HugePages_Rsvd = value
		case "HugePages_Surp":
			result.HugePages_Surp = value
		case "Hugepagesize":
			result.Hugepagesize = value
		case "DirectMap4k":
			result.DirectMap4k = value
		case "DirectMap4M":
			result.DirectMap4M = value
		case "DirectMap2M":
			result.DirectMap2M = value
		case "DirectMap1G":
			result.DirectMap1G = value
		}

	}
	return result, nil
}

func GetAsJson() ([]byte, error) {
	content, err := Get()
	if err != nil {
		return []byte{}, err
	}
	result, _ := json.Marshal(content)
	return result, nil
}

// Get values of /proc/meminfo as a map of uint64.
// Example:
//     myMeminfo := meminfo.Get()
//     x := myMeminfo["MemTotal"]
func GetAsMap() (map[string]uint64, error) {

	result := make(map[string]uint64)

	// Open the file.

	fileName := GetFilename()
	file, err := os.Open(fileName)
	if err != nil {
		return result, err
	}

	// Read the file.

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine := scanner.Text()
		splits := strings.Fields(inputLine)

		// Pull out the key.

		keySplits := strings.Split(splits[0], ":")
		key := keySplits[0]

		// Pull out the value.

		value, err := strconv.ParseUint(splits[1], 10, 64)
		if err != nil {
			continue
		}

		result[key] = value
	}
	return result, nil
}
