// Package checker start to scheck
package checker

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	lua "github.com/yuin/gopher-lua"
	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/sec-checker/config"
	"gitlab.jiagouyun.com/cloudcare-tools/sec-checker/internal/cgroup"
	"gitlab.jiagouyun.com/cloudcare-tools/sec-checker/internal/global"
	lua2 "gitlab.jiagouyun.com/cloudcare-tools/sec-checker/internal/lua"
	"gitlab.jiagouyun.com/cloudcare-tools/sec-checker/internal/luafuncs"
	"gitlab.jiagouyun.com/cloudcare-tools/sec-checker/output"
)

const (
	PanicBuf = 2048
)

var (
	Chk *Checker
	l   = logger.DefaultSLogger("check")
)

type Checker struct {
	rulesDir      string
	customRuleDir string
	taskScheduler *TaskScheduler
}

// Start : start check.
func Start(ctx context.Context, confSys *config.System, outputpath *config.ScOutput) {
	l = logger.SLogger("checker")
	luafuncs.Start()
	Chk = newChecker(confSys)
	output.Start(outputpath)
	Chk.start(ctx)
}

func newChecker(confsys *config.System) *Checker {
	lua.LuaPathDefault = filepath.Join(confsys.RuleDir, global.LuaLocalLibPath, global.PublicLuaLib)
	_, err := os.Stat(confsys.CustomRuleLibDir)
	if err == nil {
		dir := filepath.Join(confsys.CustomRuleLibDir, global.PublicLuaLib)
		lua.LuaPathDefault += ";" + dir
	}

	c := &Checker{
		rulesDir:      confsys.RuleDir,
		customRuleDir: confsys.CustomRuleDir,
		taskScheduler: NewTaskScheduler(confsys.RuleDir, confsys.CustomRuleDir, confsys.LuaHotUpdate),
	}
	if confsys.LuaCap <= 0 || confsys.LuaInitCap <= 0 || confsys.LuaInitCap > confsys.LuaCap {
		confsys.LuaInitCap = global.DefLuaPoolCap
		confsys.LuaCap = global.DefLuaPoolMaxCap
	}
	InitStatePool(confsys.LuaInitCap, confsys.LuaCap)
	return c
}

func GetManifestByName(fileName string) (*RuleManifest, error) {
	if Chk != nil && Chk.taskScheduler != nil {
		rule := Chk.taskScheduler.GetRuleByName(fileName)
		if rule != nil && rule.manifest != nil {
			l.Debugf("find by name=%s from scheduler...", fileName)
			return rule.manifest, nil
		}
	}
	// test测试时 传递的是绝对路径
	return GetManifest(fileName)
}

func GetAllManifest() map[string]*RuleManifest {
	if Chk == nil || Chk.taskScheduler == nil {
		return nil
	}
	ms := make(map[string]*RuleManifest)
	for k, rule := range Chk.taskScheduler.tasks {
		if rule != nil && rule.manifest != nil {
			ms[k] = rule.manifest
		}
	}
	for k, rule := range Chk.taskScheduler.onceTasks {
		if rule != nil && rule.manifest != nil {
			ms[k] = rule.manifest
		}
	}
	return ms
}

func GetManifest(filename string) (*RuleManifest, error) {
	filename = strings.TrimSuffix(filename, global.LuaExt)
	if !strings.HasSuffix(filename, global.LuaManifestExt) {
		filename += global.LuaManifestExt
	}
	m := &RuleManifest{path: filename}
	if err := m.parse(); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *Checker) start(ctx context.Context) {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, PanicBuf)
			n := runtime.Stack(buf, false)
			l.Errorf("panic %s", e)
			l.Errorf("%s", string(buf[:n]))
		}
		output.Close()
		l.Info("checker exit")
	}()

	if pool != nil {
		l.Infof("scheduler start")
		go c.taskScheduler.run()     //nolint:contextcheck
		go c.taskScheduler.runOnce() //nolint:contextcheck
	} else {
		l.Fatalf("pool is nil ,exit!")
	}

	select {
	case <-ctx.Done():
		return
	default:
	}

	go cgroup.Run(config.Cfg.Cgroup)

	firstTrigger()
	<-ctx.Done()
	c.taskScheduler.Stop()
}

func InitLuaGlobalFunc() {
	lua2.HandleFunc("trigger", Trigger)
}

func ShowFuncs() {
	lua2.InitModules()
	InitLuaGlobalFunc()
	lua2.ShowModule()
}
