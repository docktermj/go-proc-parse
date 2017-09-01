package dev

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

// References:
// - http://man7.org/linux/man-pages/man5/proc.5.html  "/proc/net/dev"
type Dev struct {
	ReceiveBytes       uint64 `json:"ReceiveBytes"`
	ReceivePackets     uint64 `json:"ReceivePackets"`
	ReceiveErrs        uint64 `json:"ReceiveErrs"`
	ReceiveDrop        uint64 `json:"ReceiveErrs"`
	ReceiveFifo        uint64 `json:"ReceiveFifo"`
	ReceiveFrame       uint64 `json:"ReceiveFrame"`
	ReceiveCompressed  uint64 `json:"ReceiveCompressed"`
	ReceiveMulticast   uint64 `json:"ReceiveCompressed"`
	TransmitBytes      uint64 `json:"TransmitBytes"`
	TransmitPackets    uint64 `json:"TransmitPackets"`
	TransmitErrs       uint64 `json:"TransmitErrs"`
	TransmitDrop       uint64 `json:"TransmitDrop"`
	TransmitFifo       uint64 `json:"TransmitFifo"`
	TransmitColls      uint64 `json:"TransmitColls"`
	TransmitCarrier    uint64 `json:"TransmitCarrier"`
	TransmitCompressed uint64 `json:"TransmitCompressed"`
}

type Devs map[string]Dev

// Allow filename to be specified by OS Environment variable: PROC_NET_SNMP
func GetFilename() string {
	result := os.Getenv("PROC_NET_DEV")
	if result == "" {
		result = "/proc/net/dev"
	}
	return result
}

func Get() (Devs, error) {

	result := Devs{}

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

		aDev := Dev{
			ReceiveBytes:       uint64splits[1],
			ReceivePackets:     uint64splits[2],
			ReceiveErrs:        uint64splits[3],
			ReceiveDrop:        uint64splits[4],
			ReceiveFifo:        uint64splits[5],
			ReceiveFrame:       uint64splits[6],
			ReceiveCompressed:  uint64splits[7],
			ReceiveMulticast:   uint64splits[8],
			TransmitBytes:      uint64splits[9],
			TransmitPackets:    uint64splits[10],
			TransmitErrs:       uint64splits[11],
			TransmitDrop:       uint64splits[12],
			TransmitFifo:       uint64splits[13],
			TransmitColls:      uint64splits[14],
			TransmitCarrier:    uint64splits[15],
			TransmitCompressed: uint64splits[16],
		}
		result[key] = aDev
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

// Get values of /proc/net/dev as a map of maps of uint64.
// Example:
//     myDev := dev.Get()
//     x := myDev["lo"]["receive-bytes"]
func GetAsMap() (map[string]map[string]uint64, error) {

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

		result[key]["ReceiveBytes"] = uint64splits[1]
		result[key]["ReceivePackets"] = uint64splits[2]
		result[key]["ReceiveErrs"] = uint64splits[3]
		result[key]["ReceiveDrop"] = uint64splits[4]
		result[key]["ReceiveFifo"] = uint64splits[5]
		result[key]["ReceiveFrame"] = uint64splits[6]
		result[key]["ReceiveCompressed"] = uint64splits[7]
		result[key]["ReceiveMulticast"] = uint64splits[8]
		result[key]["TransmitBytes"] = uint64splits[9]
		result[key]["TransmitPackets"] = uint64splits[10]
		result[key]["TransmitErrs"] = uint64splits[11]
		result[key]["TransmitDrop"] = uint64splits[12]
		result[key]["TransmitFifo"] = uint64splits[13]
		result[key]["TransmitColls"] = uint64splits[14]
		result[key]["TransmitCarrier"] = uint64splits[15]
		result[key]["TransmitCompressed"] = uint64splits[16]
	}
	return result, nil
}
