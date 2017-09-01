package stat

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

// References:
// - http://man7.org/linux/man-pages/man5/proc.5.html  "/proc/[pid]/stat"
type Stat struct {
	Pid                   int    `json:"pid"`
	Comm                  string `json:"comm"`
	State                 string `json:"state"`
	Ppid                  int    `json:"ppid"`
	Pgrp                  int    `json:"pgrp"`
	Session               int    `json:"session"`
	Tty_nr                int    `json:"tty_nr"`
	Tpgid                 int    `json:"tpgid"`
	Flags                 uint   `json:"flags"`
	Minflt                uint64 `json:"minflt"`
	Cminflt               uint64 `json:"cminflt"`
	Majflt                uint64 `json:"majflt"`
	Cmajflt               uint64 `json:"cmajflt"`
	Utime                 uint64 `json:"utime"`
	Stime                 uint64 `json:"stime"`
	Cutime                int64  `json:"cutime"`
	Cstime                int64  `json:"cstime"`
	Priority              int64  `json:"priority"`
	Nice                  int64  `json:"nice"`
	Num_threads           int64  `json:"num_threads"`
	Itrealvalue           int64  `json:"itrealvalue"`
	Starttime             uint64 `json:"starttime"`
	Vsize                 uint64 `json:"vsize"`
	Rss                   int64  `json:"rss"`
	Rsslim                uint64 `json:"rsslim"`
	Startcode             uint64 `json:"startcode"`
	Endcode               uint64 `json:"endcode"`
	Startstack            uint64 `json:"startstack"`
	Kstkesp               uint64 `json:"kstkesp"`
	Kstkeip               uint64 `json:"kstkeip"`
	Signal                uint64 `json:"signal"`
	Blocked               uint64 `json:"blocked"`
	Sigignore             uint64 `json:"sigignore"`
	Sigcatch              uint64 `json:"sigcatch"`
	Wchan                 uint64 `json:"wchan"`
	Nswap                 uint64 `json:"nswap"`
	Cnswap                uint64 `json:"cnswap"`
	Exit_signal           int    `json:"exit_signal"`
	Processor             int    `json:"processor"`
	Rt_priority           uint   `json:"rt_priority"`
	Policy                uint   `json:"policy"`
	Delayacct_blkio_ticks uint64 `json:"delayacct_blkio_ticks"`
	Guest_time            uint64 `json:"guest_time"`
	Cguest_time           int64  `json:"cguest_time"`
	Start_data            uint64 `json:"start_data"`
	End_data              uint64 `json:"end_data"`
	Start_brk             uint64 `json:"start_brk"`
	Arg_start             uint64 `json:"arg_start"`
	Arg_end               uint64 `json:"arg_end"`
	Env_start             uint64 `json:"env_start"`
	Env_end               uint64 `json:"env_end"`
	Exit_code             int    `json:"exit_code"`
}

// Allow filename to be specified by OS Environment variable: PROC_PID_STAT
func getFilename(pid int64) string {
	result := os.Getenv("PROC_PID_STAT")
	if result == "" {
		pidString := strconv.FormatInt(pid, 10)
		result = "/proc/" + pidString + "/stat"
	}
	return result
}

func asInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return int(0)
	}
	return result
}

func asUint(value string) uint {
	result, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return uint(0)
	}
	return uint(result)
}

func asInt64(value string) int64 {
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return int64(0)
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

func Get(pid int64) (Stat, error) {

	result := Stat{}

	// Open the file.

	fileName := getFilename(pid)
	file, err := os.Open(fileName)
	if err != nil {
		return result, err
	}

	// Read the file.

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine := scanner.Text()
		splits := strings.Fields(inputLine)

		result.Pid = asInt(splits[0])
		result.Comm = splits[1]
		result.State = splits[2]
		result.Ppid = asInt(splits[3])
		result.Pgrp = asInt(splits[4])
		result.Session = asInt(splits[5])
		result.Tty_nr = asInt(splits[6])
		result.Tpgid = asInt(splits[7])
		result.Flags = asUint(splits[8])
		result.Minflt = asUint64(splits[9])
		result.Cminflt = asUint64(splits[10])
		result.Majflt = asUint64(splits[11])
		result.Cmajflt = asUint64(splits[12])
		result.Utime = asUint64(splits[13])
		result.Stime = asUint64(splits[14])
		result.Cutime = asInt64(splits[15])
		result.Cstime = asInt64(splits[16])
		result.Priority = asInt64(splits[17])
		result.Nice = asInt64(splits[18])
		result.Num_threads = asInt64(splits[19])
		result.Itrealvalue = asInt64(splits[20])
		result.Starttime = asUint64(splits[21])
		result.Vsize = asUint64(splits[22])
		result.Rss = asInt64(splits[23])
		result.Rsslim = asUint64(splits[24])
		result.Startcode = asUint64(splits[25])
		result.Endcode = asUint64(splits[26])
		result.Startstack = asUint64(splits[27])
		result.Kstkesp = asUint64(splits[28])
		result.Kstkeip = asUint64(splits[29])
		result.Signal = asUint64(splits[30])
		result.Blocked = asUint64(splits[31])
		result.Sigignore = asUint64(splits[32])
		result.Sigcatch = asUint64(splits[33])
		result.Wchan = asUint64(splits[34])
		result.Nswap = asUint64(splits[35])
		result.Cnswap = asUint64(splits[36])
		result.Exit_signal = asInt(splits[37])
		result.Processor = asInt(splits[38])
		result.Rt_priority = asUint(splits[39])
		result.Policy = asUint(splits[40])
		result.Delayacct_blkio_ticks = asUint64(splits[41])
		result.Guest_time = asUint64(splits[42])
		result.Cguest_time = asInt64(splits[43])
		result.Start_data = asUint64(splits[44])
		result.End_data = asUint64(splits[45])
		result.Start_brk = asUint64(splits[46])
		result.Arg_start = asUint64(splits[47])
		result.Arg_end = asUint64(splits[48])
		result.Env_start = asUint64(splits[49])
		result.Env_end = asUint64(splits[50])
		result.Exit_code = asInt(splits[51])
	}
	return result, nil
}

func GetAsJson(pid int64) ([]byte, error) {
	content, err := Get(pid)
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
func GetAsMap(pid int64) (map[string]interface{}, error) {

	result := make(map[string]interface{})
	stat, err := Get(pid)
	if err != nil {
		return result, err
	}

	result["pid"] = stat.Pid
	result["comm"] = stat.Comm
	result["state"] = stat.State
	result["ppid"] = stat.Ppid
	result["pgrp"] = stat.Pgrp
	result["session"] = stat.Session
	result["tty_nr"] = stat.Tty_nr
	result["tpgid"] = stat.Tpgid
	result["flags"] = stat.Flags
	result["minflt"] = stat.Minflt
	result["cminflt"] = stat.Cminflt
	result["majflt"] = stat.Majflt
	result["cmajflt"] = stat.Cmajflt
	result["utime"] = stat.Utime
	result["stime"] = stat.Stime
	result["cutime"] = stat.Cutime
	result["cstime"] = stat.Cstime
	result["priority"] = stat.Priority
	result["nice"] = stat.Nice
	result["num_threads"] = stat.Num_threads
	result["itrealvalue"] = stat.Itrealvalue
	result["starttime"] = stat.Starttime
	result["vsize"] = stat.Vsize
	result["rss"] = stat.Rss
	result["rsslim"] = stat.Rsslim
	result["startcode"] = stat.Startcode
	result["endcode"] = stat.Endcode
	result["startstack"] = stat.Startstack
	result["kstkesp"] = stat.Kstkesp
	result["kstkeip"] = stat.Kstkeip
	result["signal"] = stat.Signal
	result["blocked"] = stat.Blocked
	result["sigignore"] = stat.Sigignore
	result["sigcatch"] = stat.Sigcatch
	result["wchan"] = stat.Wchan
	result["nswap"] = stat.Nswap
	result["cnswap"] = stat.Cnswap
	result["exit_signal"] = stat.Exit_signal
	result["processor"] = stat.Processor
	result["rt_priority"] = stat.Rt_priority
	result["policy"] = stat.Policy
	result["delayacct_blkio_ticks"] = stat.Delayacct_blkio_ticks
	result["guest_time"] = stat.Guest_time
	result["cguest_time"] = stat.Cguest_time
	result["start_data"] = stat.Start_data
	result["end_data"] = stat.End_data
	result["start_brk"] = stat.Start_brk
	result["arg_start"] = stat.Arg_start
	result["arg_end"] = stat.Arg_end
	result["env_start"] = stat.Env_start
	result["env_end"] = stat.Env_end
	result["exit_code"] = stat.Exit_code

	return result, nil
}
