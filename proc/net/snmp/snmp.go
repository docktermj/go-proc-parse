package snmp

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Allow filename to be specified by OS Environment variable: PROC_NET_SNMP
func GetFilename() string {
	result := os.Getenv("PROC_NET_SNMP")
	if result == "" {
		result = "/proc/net/snmp"
	}
	return result
}

// Get values of /proc/net/snmp as a map of maps of uint64.
// Example:
//     mySnmp := snmp.Get()
//     x := mySnmp["Ip"]["InHdrErrors"]
func Get() (map[string]map[string]uint64, error) {

	result := make(map[string]map[string]uint64)

	// Open the file.

	fileName := GetFilename()
	file, err := os.Open(fileName)
	if err != nil {
		return result, err
	}

	// Oscillate between header and non-header lines in file.

	header := true
	headerSplits := []string{}

	// Read the file.

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine := scanner.Text()
		splits := strings.Fields(inputLine)

		// Pull out the first-level key for the map-map structure.

		keySplits := strings.Split(splits[0], ":")
		key := keySplits[0]

		// Handle header and data lines in the file.

		if header {
			headerSplits = splits
			result[key] = make(map[string]uint64)
		} else {

			// Transform string data to uint64 and put in result.

			for index, split := range splits {
				value, err := strconv.ParseUint(split, 10, 64)
				if err != nil {
					continue
				}
				result[key][headerSplits[index]] = value
			}
		}
		header = !header // Oscillate between header and non-header lines in file.
	}
	return result, nil
}
