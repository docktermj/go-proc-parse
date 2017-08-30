package dev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Allow filename to be specified by OS Environment variable: PROC_NET_SNMP
func GetFilename() string {
	result := os.Getenv("PROC_NET_DEV")
	if result == "" {
		result = "/proc/net/dev"
	}
	return result
}

// Get values of /proc/net/dev as a map of maps of uint64.
// Example:
//     myDev := dev.Get()
//     x := myDev["lo"]["receive-bytes"]
func Get() (map[string]map[string]uint64, error) {

	result := make(map[string]map[string]uint64)

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
		if strings.Contains(inputLine, "|") { // Avoid table headers.
			continue
		}

		splits := strings.Fields(inputLine)

		// Pull out the key.

		keySplits := strings.Split(splits[0], ":")
		key := keySplits[0]
		result[key] = make(map[string]uint64)

		// Convert strings to uint64.

		uint64splits := make([]uint64, len(splits))
		for index, split := range splits {
			value, err := strconv.ParseUint(split, 10, 64)
			if err != nil {
				continue
			}
			uint64splits[index] = value
		}

		// Pull out the values.

		result[key]["receive-bytes"] = uint64splits[1]
		result[key]["receive-packets"] = uint64splits[2]
		result[key]["receive-errs"] = uint64splits[3]
		result[key]["receive-drop"] = uint64splits[4]
		result[key]["receive-fifo"] = uint64splits[5]
		result[key]["receive-frame"] = uint64splits[6]
		result[key]["receive-compressed"] = uint64splits[7]
		result[key]["receive-multicast"] = uint64splits[8]
		result[key]["transmit-bytes"] = uint64splits[9]
		result[key]["transmit-packets"] = uint64splits[10]
		result[key]["transmit-errs"] = uint64splits[11]
		result[key]["transmit-drop"] = uint64splits[12]
		result[key]["transmit-fifo"] = uint64splits[13]
		result[key]["transmit-colls"] = uint64splits[14]
		result[key]["transmit-carrier"] = uint64splits[15]
		result[key]["transmit-compressed"] = uint64splits[16]
	}
	return result, nil
}
