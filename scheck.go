package securityChecker

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils"
	"gitlab.jiagouyun.com/cloudcare-tools/sec-checker/git"
)

const (
	OSWindows = `windows`
	OSLinux   = `linux`
	OSDarwin  = `darwin`

	OSArchWinAmd64    = "windows/amd64"
	OSArchWin386      = "windows/386"
	OSArchLinuxArm    = "linux/arm"
	OSArchLinuxArm64  = "linux/arm64"
	OSArchLinux386    = "linux/386"
	OSArchLinuxAmd64  = "linux/amd64"
	OSArchDarwinAmd64 = "darwin/amd64"

	CommonChanCap = 32

	// categories
	MetricDeprecated  = "/v1/write/metrics"
	Metric            = "/v1/write/metric"
	KeyEvent          = "/v1/write/keyevent"
	Object            = "/v1/write/object"
	CustomObject      = "/v1/write/custom_object"
	Logging           = "/v1/write/logging"
	LogFilter         = "/v1/logfilter/pull"
	Tracing           = "/v1/write/tracing"
	Rum               = "/v1/write/rum"
	Security          = "/v1/write/security"
	HeartBeat         = "/v1/write/heartbeat"
	Election          = "/v1/election"
	ElectionHeartbeat = "/v1/election/heartbeat"
	QueryRaw          = "/v1/query/raw"
)

var (
	Exit = cliutils.NewSem()
	WG   = sync.WaitGroup{}

	Docker     = false
	Version    = git.Version
	AutoUpdate = false

	InstallDir         = optionalInstallDir[runtime.GOOS+"/"+runtime.GOARCH]
	optionalInstallDir = map[string]string{
		OSArchWinAmd64: `C:\Program Files\scheck`,
		OSArchWin386:   `C:\Program Files (x86)\scheck`,

		OSArchLinuxArm:    `/usr/local/scheck`,
		OSArchLinuxArm64:  `/usr/local/scheck`,
		OSArchLinuxAmd64:  `/usr/local/scheck`,
		OSArchLinux386:    `/usr/local/scheck`,
		OSArchDarwinAmd64: `/usr/local/scheck`,
	}

	AllOS   = []string{OSWindows, OSLinux, OSDarwin}
	AllArch = []string{OSArchWinAmd64, OSArchWin386, OSArchLinuxArm, OSArchLinuxArm64, OSArchLinux386, OSArchLinuxAmd64, OSArchDarwinAmd64}

	UnknownOS   = []string{"unknown"}
	UnknownArch = []string{"unknown"}

	//DataDir  = filepath.Join(InstallDir, "data")
	//ConfdDir = filepath.Join(InstallDir, "conf.d")

	//MainConfPathDeprecated = filepath.Join(InstallDir, "datakit.conf")
	//MainConfPath           = filepath.Join(ConfdDir, "datakit.conf")

	//PipelineDir        = filepath.Join(InstallDir, "pipeline")
	//PipelinePatternDir = filepath.Join(PipelineDir, "pattern")
	GRPCDomainSock = filepath.Join(InstallDir, "datakit.sock")
	GRPCSock       = ""
)

const (
	ConfPerm = os.ModePerm
)

func Quit() {
	Exit.Close()
	WG.Wait()
	//service.Stop()
}
