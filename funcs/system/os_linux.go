package system

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"

	lua "github.com/yuin/gopher-lua"
	"gitlab.jiagouyun.com/cloudcare-tools/sec-checker/funcs/impl"
)

var log = logger.DefaultSLogger("system")

func loadOtherOS() map[string]lua.LGFunction {
	return linuxAPI
}

var linuxAPI = map[string]lua.LGFunction{
	"kernel_info":        KernelInfo,
	"kernel_modules":     KernelModules,
	"users":              Users,
	"logged_in_users":    LoggedInUsers,
	"last":               Last,
	"lastb":              Lastb,
	"shadow":             Shadow,
	"ulimit_info":        UlimitInfo,
	"processes":          Processes,
	"process_open_files": ProcessOpendFiles,
	"shell_history":      ShellHistory,
	"crontab":            Crontab,
	"sysctl":             Sysctl,
	"rpm_list":           RpmList,
	"rpm_query":          RpmQuery,
	"process_command":    getProcessCommand,
}

// nolint
func KernelInfo(l *lua.LState) int {
	var kernel lua.LTable
	if data, err := ioutil.ReadFile(`/proc/cmdline`); err == nil {
		args := strings.Split(string(data), " ")
		additionalArguments := ""
		boot := 11
		argL := 5
		for _, arg := range args {
			if len(arg) > boot && arg[0:boot] == "BOOT_IMAGE=" {
				kernel.RawSetString("path", lua.LString(arg[boot:]))
			} else if len(arg) > argL && arg[0:argL] == "root=" {
				kernel.RawSetString("device", lua.LString(arg[argL:]))
			} else {
				if additionalArguments != "" {
					additionalArguments += " "
				}
				additionalArguments += arg
			}
		}
		if additionalArguments != "" {
			kernel.RawSetString("arguments", lua.LString(strings.TrimSpace(additionalArguments)))
		}
	}

	if data, err := ioutil.ReadFile(`/proc/version`); err == nil {
		details := strings.Split(string(data), " ")
		if len(details) > 2 && details[1] == "version" {
			kernel.RawSetString("version", lua.LString(details[2]))
		}
	}

	l.Push(&kernel)
	return 1
}

func KernelModules(l *lua.LState) int {
	var result lua.LTable
	partsMax := 5
	if data, err := ioutil.ReadFile(`/proc/modules`); err == nil {
		mods := strings.Split(string(data), "\n")
		for _, mod := range mods {
			mod = strings.TrimSpace(mod)
			parts := strings.FieldsFunc(mod, func(r rune) bool {
				return r == ' '
			})
			if len(parts) < partsMax {
				continue
			}
			for idx, p := range parts {
				if len(p) > 0 && p[len(p)-1] == ',' {
					parts[idx] = p[0 : len(p)-1]
				}
			}
			var mt lua.LTable
			mt.RawSetString("name", lua.LString(parts[0]))
			mt.RawSetString("size", lua.LString(parts[1]))
			mt.RawSetString("used_by", lua.LString(parts[2]))
			mt.RawSetString("status", lua.LString(parts[3]))
			mt.RawSetString("address", lua.LString(parts[4]))
			result.Append(&mt)
		}
	}

	l.Push(&result)
	return 1
}

func Users(l *lua.LState) int {
	us, err := impl.GetUserDetail("")
	if err != nil {
		l.RaiseError("%s", err)
		return lua.MultRet
	}

	var result lua.LTable

	for _, u := range us {
		var ut lua.LTable
		ut.RawSetString("username", lua.LString(u.Name))
		uid := -1
		gid := -1
		if n, err := strconv.Atoi(u.UID); err == nil {
			uid = n
		}
		if n, err := strconv.Atoi(u.GID); err == nil {
			uid = n
		}
		ut.RawSetString("uid", lua.LNumber(uid))
		ut.RawSetString("gid", lua.LNumber(gid))
		ut.RawSetString("directory", lua.LString(u.Home))
		ut.RawSetString("shell", lua.LString(u.Shell))
		result.Append(&ut)
	}

	l.Push(&result)
	return 1
}

// nolint
func Shadow(l *lua.LState) int {
	partsL := 9
	file, err := os.Open("/etc/shadow")
	if err != nil {
		l.RaiseError("%s", err)
		return lua.MultRet
	}
	defer func() {
		_ = file.Close()
	}()

	var result lua.LTable
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				l.RaiseError("%s", err)
				return lua.MultRet
			}
		}

		parts := strings.Split(line, ":")
		if len(parts) < partsL {
			continue
		}

		var status string
		lastChange := -1
		expire := -1

		psw := parts[1]
		if psw == "" {
			status = "empty"
		} else if psw == "!!" {
			status = "not-set"
		} else if psw[:1] == "*" || psw[:1] == "!" || psw[:1] == "x" {
			status = "locked"
		} else {
			status = "active"
		}

		if d, err := strconv.Atoi(parts[2]); err == nil && d > 0 {
			lastChange = d
		}
		if d, err := strconv.Atoi(parts[8]); err == nil && d > 0 {
			expire = d
		}

		var ut lua.LTable
		ut.RawSetString("username", lua.LString(parts[0]))
		ut.RawSetString("password_status", lua.LString(status))
		ut.RawSetString("last_change", lua.LNumber(lastChange))
		ut.RawSetString("expire", lua.LNumber(expire))
		result.Append(&ut)
	}

	l.Push(&result)
	return 1
}

func UlimitInfo(l *lua.LState) int {
	var result lua.LTable
	l.Push(&result)
	return 1
}

func Processes(l *lua.LState) int {
	pslist, err := impl.GetProcesses()
	if err != nil {
		l.RaiseError("%s", err)
		return lua.MultRet
	}
	var result lua.LTable
	for _, p := range pslist {
		var pt lua.LTable
		pt.RawSetString("pid", lua.LNumber(p.Pid))
		pt.RawSetString("parent", lua.LNumber(p.Parent))
		pt.RawSetString("path", lua.LString(p.Path))
		pt.RawSetString("name", lua.LString(p.Name))
		pt.RawSetString("pgroup", lua.LString(p.PGroup))
		pt.RawSetString("state", lua.LString(p.State))
		pt.RawSetString("nice", lua.LString(p.Nice))
		pt.RawSetString("threads", lua.LNumber(p.Threads))
		pt.RawSetString("cmdline", lua.LString(p.Cmdline))
		pt.RawSetString("cwd", lua.LString(p.Cwd))
		pt.RawSetString("root", lua.LString(p.Root))
		pt.RawSetString("uid", lua.LString(p.UID))
		pt.RawSetString("euid", lua.LString(p.EUID))
		pt.RawSetString("suid", lua.LString(p.SUID))
		pt.RawSetString("gid", lua.LString(p.GID))
		pt.RawSetString("egid", lua.LString(p.EGID))
		pt.RawSetString("sgid", lua.LString(p.SGID))
		pt.RawSetString("on_disk", lua.LNumber(p.OnDisk))
		pt.RawSetString("resident_size", lua.LNumber(p.ResidentSize))
		pt.RawSetString("total_size", lua.LNumber(p.TotalSize))
		pt.RawSetString("user_time", lua.LNumber(p.UserTime))
		pt.RawSetString("system_time", lua.LNumber(p.SystemTime))
		pt.RawSetString("start_time", lua.LNumber(p.StartTime))
		pt.RawSetString("disk_bytes_read", lua.LNumber(p.DiskBytesRead))
		pt.RawSetString("disk_bytes_written", lua.LNumber(p.DiskBytesWritten))

		result.Append(&pt)
	}

	l.Push(&result)
	return 1
}

func ProcessOpendFiles(l *lua.LState) int {
	var pids []int
	lv := l.Get(1)
	if lv != lua.LNil {
		if lv.Type() != lua.LTNumber {
			l.TypeError(1, lua.LTNumber)
			return lua.MultRet
		}
		pids = append(pids, int(lv.(lua.LNumber)))
	}

	fds, err := impl.EnumProcessesFds(pids)
	if err != nil {
		l.RaiseError("%s", err)
		return lua.MultRet
	}

	var result lua.LTable

	for _, fd := range fds {
		var t lua.LTable
		t.RawSetString("pid", lua.LNumber(fd.PID))
		t.RawSetString("fd", lua.LString(fd.Fd))
		t.RawSetString("path", lua.LString(fd.LinkPath))
		result.Append(&t)
	}

	l.Push(&result)
	return 1
}

func ShellHistory(l *lua.LState) int {
	targetUser := ""
	lv := l.Get(1)
	if lv.Type() != lua.LTNil {
		if lv.Type() != lua.LTString {
			l.TypeError(1, lua.LTString)
			return lua.MultRet
		}
		targetUser = lv.String()
	}

	users, err := impl.GetUserDetail(targetUser)
	if err != nil {
		l.RaiseError("%s", err)
		return lua.MultRet
	} else if len(users) == 0 {
		l.RaiseError("user '%s' not exists", targetUser)
		return lua.MultRet
	}

	shellHistoryFiles := []string{
		".bash_history",
		".zsh_history",
		".zhistory",
		".history",
		".sh_history",
	}

	var result lua.LTable
	for _, u := range users {
		if u.Home == "" {
			continue
		}
		if u.Shell == "/bin/false" || strings.HasSuffix(u.Shell, "/nologin") {
			continue
		}

		for _, f := range shellHistoryFiles {
			cmds, err := impl.GenShellHistoryFromFile(filepath.Join(u.Home, f))
			if err != nil {
				l.RaiseError("%s", err)
				return lua.MultRet
			}

			for _, cmd := range cmds {
				var item lua.LTable
				uid := -1
				if n, err := strconv.Atoi(u.UID); err == nil {
					uid = n
				}
				item.RawSetString("uid", lua.LNumber(uid))
				item.RawSetString("history_file", lua.LString(filepath.Join(u.Home, f)))
				item.RawSetString("command", lua.LString(cmd.Command))
				item.RawSetString("time", lua.LNumber(cmd.Time))
				result.Append(&item)
			}
		}
	}

	l.Push(&result)
	return 1
}

func parseUtmpFile(file string) (*lua.LTable, error) {
	f, err := os.Open(filepath.Clean(file))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	var result lua.LTable
	utmps, err := impl.ParseUtmp(f)
	if err != nil {
		return nil, err
	}
	for _, utmp := range utmps {
		var item lua.LTable
		item.RawSetString("pid", lua.LNumber(utmp.Pid))
		item.RawSetString("username", lua.LString(utmp.User))
		item.RawSetString("tty", lua.LString(utmp.Device))
		item.RawSetString("host", lua.LString(utmp.Host))
		item.RawSetString("type", lua.LNumber(utmp.Type))
		item.RawSetString("time", lua.LNumber(utmp.Time))
		result.Append(&item)
	}
	return &result, nil
}

func Last(l *lua.LState) int {
	utmps, err := parseUtmpFile("/var/log/wtmp")
	if err != nil {
		l.RaiseError("%s", err)
		return lua.MultRet
	}

	l.Push(utmps)
	return 1
}

func Lastb(l *lua.LState) int {
	utmps, err := parseUtmpFile("/var/log/btmp")
	if err != nil {
		l.RaiseError("%s", err)
		return lua.MultRet
	}

	l.Push(utmps)
	return 1
}

func LoggedInUsers(l *lua.LState) int {
	utmps, err := parseUtmpFile("/var/run/utmp")
	if err != nil {
		l.RaiseError("%s", err)
		return lua.MultRet
	}

	l.Push(utmps)
	return 1
}

// nolint
func Crontab(l *lua.LState) int {
	targetUser := ""
	lv := l.Get(1)
	if lv.Type() != lua.LTNil {
		if lv.Type() != lua.LTString {
			l.TypeError(1, lua.LTString)
			return lua.MultRet
		}
		targetUser = lv.String()

	}

	crondir := `/var/spool/cron`

	files, err := ioutil.ReadDir(crondir)
	if err != nil {
		l.RaiseError("%s", err)
		return lua.MultRet
	}

	parseLine := func(path, line string, num int) *lua.LTable {
		cronL := 5
		if line == "" {
			return nil
		}
		var cron lua.LTable
		columns := strings.Split(line, "\t")
		if len(columns) == 1 {
			columns = strings.Split(columns[0], " ")
		}
		if len(columns) < cronL {
			log.Warnf("unknown line: %s[%d]", path, num)
			return nil
		}
		var command string
		keys := []string{"minute", "hour", "day_of_month", "month", "day_of_week"}

		for i := 0; i < len(columns); i++ {
			if i < cronL {
				cron.RawSetString(keys[i], lua.LString(columns[i]))
			} else if i == cronL {
				command = columns[i]
			} else {
				command += " " + columns[i]
			}
		}
		if command == "" {
			log.Warnf("no cron command found: %s[%d]", path, num)
			return nil
		}
		cron.RawSetString("command", lua.LString(command))
		cron.RawSetString("path", lua.LString(path))
		return &cron
	}

	var result lua.LTable
	for _, file := range files {
		if targetUser != "" && file.Name() != targetUser {
			continue
		}
		fpath := filepath.Join(crondir, file.Name())
		content, err := ioutil.ReadFile(fpath)
		if err != nil {
			log.Errorf("%s", err)
			continue
		}
		lines := strings.Split(string(content), "\n")
		for num, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || line[0] == '#' {
				continue
			}
			cron := parseLine(fpath, line, num)
			if cron != nil {
				result.Append(cron)
			}
		}
	}

	l.Push(&result)
	return 1
}

func Sysctl(l *lua.LState) int {
	key := ""
	lv := l.Get(1)
	if lv.Type() != lua.LTNil {
		if lv.Type() != lua.LTString {
			l.TypeError(1, lua.LTString)
			return lua.MultRet
		}
		key = lv.String()
	}

	cmd := exec.Command("sysctl", "-a")
	buf := bytes.NewBuffer([]byte{})
	errbuf := bytes.NewBuffer([]byte{})
	cmd.Stdout = buf
	cmd.Stderr = errbuf
	if err := cmd.Run(); err != nil {
		l.RaiseError("%s, %s", err, errbuf.String())
		return lua.MultRet
	}

	lines := strings.Split(buf.String(), "\n")
	var result lua.LTable
	lineL := 2
	for _, line := range lines {
		line = strings.TrimSpace(line)
		fields := strings.Split(line, "=")
		if len(fields) < lineL {
			continue
		}
		k := strings.TrimSpace(fields[0])
		if key != "" && k != key {
			continue
		}
		result.RawSetString(k, lua.LString(strings.TrimSpace(fields[1])))
		if key != "" {
			break
		}
	}

	l.Push(&result)
	return 1
}

func RpmList(l *lua.LState) int {
	cmd := exec.Command("rpm", "-qa")
	buf := bytes.NewBuffer([]byte{})
	errbuf := bytes.NewBuffer([]byte{})
	cmd.Stdout = buf
	cmd.Stderr = errbuf
	if err := cmd.Run(); err != nil {
		l.RaiseError("%s, %s", err, errbuf.String())
		return lua.MultRet
	}

	l.Push(lua.LString(buf.String()))
	return 1
}

func RpmQuery(l *lua.LState) int {
	pkg := ""
	lv := l.Get(1)
	if lv.Type() != lua.LTNil {
		if lv.Type() != lua.LTString {
			l.TypeError(1, lua.LTString)
			return lua.MultRet
		}
		pkg = lv.String()
	}

	cmd := exec.Command("rpm", "-q", pkg)
	buf := bytes.NewBuffer([]byte{})
	errbuf := bytes.NewBuffer([]byte{})
	cmd.Stdout = buf
	cmd.Stderr = errbuf
	if err := cmd.Run(); err != nil {
		l.Push(lua.LString(""))
		return 1
	}

	l.Push(lua.LString(buf.String()))
	return 1
}

// 获取一个进程的启动的参数：config，crt信息，等
// 比如：获取scheck的config参数  getProcessCommand（“scheck”，“config”）
func getProcessCommand(l *lua.LState) int {
	lv := l.Get(1)
	if lv.Type() != lua.LTString {
		l.TypeError(1, lua.LTString)
		return lua.MultRet
	}
	processName := string(lv.(lua.LString))

	lv2 := l.Get(2)
	if lv.Type() != lua.LTString {
		l.TypeError(1, lua.LTString)
		return lua.MultRet
	}
	command := string(lv2.(lua.LString))
	cmdline := impl.GetCommandLine(processName)
	if cmdline == "" {
		l.RaiseError("can not find this process:%s", processName)
		return lua.MultRet
	}
	arguments := impl.GetCmdline(cmdline)
	for key, val := range arguments {
		if key == command {
			l.Push(lua.LString(val))
			return 1
		}
	}
	// 当没有此参数时或参数为空， 应当返回空 而不是错误。
	l.Push(lua.LString(""))
	return 1
}
