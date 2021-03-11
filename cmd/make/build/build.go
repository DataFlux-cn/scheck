package build

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/sec-checker/git"
)

var (
	/* Use:
		go tool dist list
	to get current os/arch list */

	OSArches = []string{ // supported os/arch list
		`linux/386`,
		`linux/amd64`,
		`linux/arm`,
		`linux/arm64`,

		`darwin/amd64`,

		`windows/amd64`,
		`windows/386`,
	}
)

func runEnv(args, env []string) ([]byte, error) {
	cmd := exec.Command(args[0], args[1:]...)
	if env != nil {
		cmd.Env = append(os.Environ(), env...)
	}

	return cmd.CombinedOutput()
}

var (
	l = logger.DefaultSLogger("build")

	BuildDir     = "build"
	PubDir       = "pub"
	AppName      = "sec-checker"
	AppBin       = "sec-checker"
	OSSPath      = "sec-checker"
	Archs        string
	Release      string
	MainEntry    string
	DownloadAddr string
	ReleaseType  string
)

func prepare() {

	os.RemoveAll(BuildDir)
	_ = os.MkdirAll(BuildDir, os.ModePerm)
	_ = os.MkdirAll(filepath.Join(PubDir, Release), os.ModePerm)

	// create version info
	vd := &versionDesc{
		Version:  strings.TrimSpace(git.Version),
		Date:     git.BuildAt,
		Uploader: git.Uploader,
		Branch:   git.Branch,
		Commit:   git.Commit,
		Go:       git.Golang,
	}

	versionInfo, err := json.MarshalIndent(vd, "", "    ")
	if err != nil {
		l.Fatal(err)
	}

	if err := ioutil.WriteFile(filepath.Join(PubDir, Release, "version"), versionInfo, 0666); err != nil {
		l.Fatal(err)
	}
}

func Compile() {
	start := time.Now()

	prepare()

	var archs []string

	switch Archs {
	case "all":
		archs = OSArches

		// read cmd-line env
		if x := os.Getenv("ALL_ARCHS"); x != "" {
			archs = strings.Split(x, "|")
		}
	case "local":
		archs = []string{runtime.GOOS + "/" + runtime.GOARCH}
	default:
		archs = strings.Split(Archs, "|")
	}

	for _, arch := range archs {

		parts := strings.Split(arch, "/")
		if len(parts) != 2 {
			l.Fatalf("invalid arch %q", parts)
		}

		goos, goarch := parts[0], parts[1]

		dir := fmt.Sprintf("%s/%s-%s-%s", BuildDir, AppName, goos, goarch)

		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			l.Fatalf("failed to mkdir: %v", err)
		}

		dir, err = filepath.Abs(dir)
		if err != nil {
			l.Fatal(err)
		}

		compileArch(AppBin, goos, goarch, dir)

		// if goos == "windows" {
		// 	installerExe = fmt.Sprintf("installer-%s-%s.exe", goos, goarch)
		// } else {
		// 	installerExe = fmt.Sprintf("installer-%s-%s", goos, goarch)
		// }

		// buildInstaller(filepath.Join(PubDir, Release), goos, goarch)
	}

	l.Infof("Done!(elapsed %v)", time.Since(start))
}

func compileArch(bin, goos, goarch, dir string) {

	output := filepath.Join(dir, bin)
	if goos == "windows" {
		output += ".exe"
	}

	args := []string{
		"go", "build",
		"-o", output,
		"-ldflags",
		fmt.Sprintf("-w -s -X main.ReleaseType=%s", ReleaseType),
		MainEntry,
	}

	env := []string{
		"GOOS=" + goos,
		"GOARCH=" + goarch,
		`GO111MODULE=off`,
		"CGO_ENABLED=0",
	}

	l.Debugf("building %s", fmt.Sprintf("%s-%s/%s", goos, goarch, bin))
	msg, err := runEnv(args, env)
	if err != nil {
		l.Fatalf("failed to run %v, envs: %v: %v, msg: %s", args, env, err, string(msg))
	}
}

// type installInfo struct {
// 	Name         string
// 	DownloadAddr string
// 	Version      string
// }

// func buildInstaller(outdir, goos, goarch string) {

// 	l.Debugf("building %s-%s/installer...", goos, goarch)

// 	args := []string{
// 		"go", "build",
// 		"-o", filepath.Join(outdir, installerExe),
// 		"-ldflags",
// 		fmt.Sprintf("-w -s -X main.DataKitBaseURL=%s -X main.DataKitVersion=%s", DownloadAddr, git.Version),
// 		"cmd/installer/installer.go",
// 	}

// 	env := []string{
// 		"GOOS=" + goos,
// 		"GOARCH=" + goarch,
// 	}

// 	msg, err := runEnv(args, env)
// 	if err != nil {
// 		l.Fatalf("failed to run %v, envs: %v: %v, msg: %s", args, env, err, string(msg))
// 	}
// }
