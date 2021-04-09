package funcs

import (
	"bufio"
	"fmt"
	"os"

	luajson "github.com/layeh/gopher-json"
	lua "github.com/yuin/gopher-lua"
	luaparse "github.com/yuin/gopher-lua/parse"
	checker "gitlab.jiagouyun.com/cloudcare-tools/sec-checker"
)

type (
	luaLib struct {
		name            string
		fn              lua.LGFunction
		disabledMethods []string
	}

	ByteCode struct {
		Proto *lua.FunctionProto
	}

	ScriptRunTime struct {
		ls *lua.LState
	}
)

var (
	supportLuaLibs = []luaLib{
		{lua.LoadLibName, lua.OpenPackage, nil},
		{lua.BaseLibName, lua.OpenBase, nil},
		{lua.TabLibName, lua.OpenTable, nil},
		{lua.StringLibName, lua.OpenString, nil},
		{lua.MathLibName, lua.OpenMath, nil},
		{lua.DebugLibName, lua.OpenDebug, nil},
		{lua.OsLibName, lua.OpenOs, []string{"execute", "remove", "rename", "setenv", "setlocale"}},
	}

	//LuaExtendFuncs = []LuaFunc{
	// {`ls`, ls},
	// {`file_exist`, fileExist},
	// {`file_info`, fileInfo},
	// {`read_file`, readFile},
	// {`file_hash`, fileHash},
	// {`hostname`, hostname},
	// {`uptime`, uptime},
	// {`time_zone`, zone},
	// {`kernel_info`, kernelInfo},
	// {`kernel_modules`, kernelModules},
	// {`ulimit_info`, ulimitInfo},
	// {`mounts`, mounts},
	// {`processes`, processes},
	// {`interface_addresses`, interfaceAddresses},
	// {`interface_details`, interfaceDetails},
	// {`iptables`, ipTables},
	// {`process_open_sockets`, processOpenSockets},
	// {`listening_ports`, listeningPorts},
	// {`users`, users},
	// //{`logged_in_users`, loggedInUsers},
	// {`shadow`, shadow},
	// //{`last`, last},
	// {`shell_history`, shellHistory},
	// {`json_encode`, jsonEncode},
	// {`json_decode`, jsonDecode},
	// {`get_cache`, getCache},
	// {`set_cache`, setCache},
	//}
)

// // Register extended lua funcs to lua machine
// func (l *LuaExt) Register(lstate *lua.LState) error {
// 	//luajson.Preload(lstate) //for json parse
// 	if err := loadLuaLibs(lstate); err != nil {
// 		return err
// 	}
// 	for _, f := range luaExtendFuncs {
// 		lstate.Register(f.Name, f.Fn)
// 	}
// 	return nil
// }

func (r *ScriptRunTime) Close() {
	if r.ls != nil {
		if !r.ls.IsClosed() {
			r.ls.Close()
		}
	}
}

func (r *ScriptRunTime) PCall(bcode *ByteCode) error {
	lfunc := r.ls.NewFunctionFromProto(bcode.Proto)
	r.ls.Push(lfunc)
	return r.ls.PCall(0, lua.MultRet, nil)
}

func CompilesScript(filePath string) (*ByteCode, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	chunk, err := luaparse.Parse(reader, filePath)
	if err != nil {
		return nil, fmt.Errorf("fail to parse lua file '%s', err: %s", filePath, err)
	}
	proto, err := lua.Compile(chunk, filePath)
	if err != nil {
		return nil, fmt.Errorf("fail to compile lua file '%s', err: %s", filePath, err)
	}
	return &ByteCode{
		Proto: proto,
	}, nil
}

func unsupportFn(ls *lua.LState) int {
	lv := ls.Get(lua.UpvalueIndex(1))
	ls.RaiseError("'%s' diabled", lv.String())
	return 0
}

type ScriptGlobalCfg struct {
	RulePath string
}

func SetScriptGlobalConfig(l *lua.LState, cfg *ScriptGlobalCfg) {
	var t lua.LTable
	t.RawSetString("rulefile", lua.LString(cfg.RulePath))
	l.SetGlobal("__this_configuration", &t)
}

func GetScriptGlobalConfig(l *lua.LState) *ScriptGlobalCfg {
	lv := l.GetGlobal("__this_configuration")
	if lv.Type() == lua.LTTable {
		t := lv.(*lua.LTable)
		var cfg ScriptGlobalCfg
		v := t.RawGetString("rulefile")
		if v.Type() == lua.LTString {
			cfg.RulePath = string(v.(lua.LString))
		}
		return &cfg
	}
	return nil
}

func GetScriptRuntime(cfg *ScriptGlobalCfg) (*ScriptRunTime, error) {
	ls := lua.NewState(lua.Options{SkipOpenLibs: true})
	if err := LoadLuaLibs(ls); err != nil {
		ls.Close()
		return nil, err
	}
	SetScriptGlobalConfig(ls, cfg)
	luajson.Preload(ls) //for json parse
	for _, p := range checker.FuncProviders {
		for _, f := range p.Funcs() {
			ls.Register(f.Name, lua.LGFunction(f.Fn))
		}
	}
	return &ScriptRunTime{
		ls: ls,
	}, nil
}

func LoadLuaLibs(ls *lua.LState) error {

	for _, lib := range supportLuaLibs {

		err := ls.CallByParam(lua.P{
			Fn:      ls.NewFunction(lib.fn),
			NRet:    1,
			Protect: true,
		}, lua.LString(lib.name))
		if err != nil {
			return fmt.Errorf("load %s failed, %s", lib.name, err)
		}
		lvMod := ls.Get(-1)
		if lvMod.Type() == lua.LTTable {
			lt := lvMod.(*lua.LTable)
			for _, mth := range lib.disabledMethods {
				lt.RawSetString(mth, ls.NewClosure(unsupportFn, lua.LString(mth)))
			}
		}
		ls.Pop(1)
	}
	return nil
}
