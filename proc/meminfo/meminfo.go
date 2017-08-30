package meminfo

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Allow filename to be specified by OS Environment variable: PROC_NET_SNMP
func GetFilename() string {
	result := os.Getenv("PROC_MEMINFO")
	if result == "" {
		result = "/proc/meminfo"
	}
	return result
}

// Get values of /proc/meminfo as a map of uint64.
// Example:
//     myMeminfo := meminfo.Get()
//     x := myMeminfo["MemTotal"]
func Get() (map[string]uint64, error) {

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
