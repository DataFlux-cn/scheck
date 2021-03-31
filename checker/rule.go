package checker

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	luaparse "github.com/yuin/gopher-lua/parse"

	"github.com/influxdata/toml"
)

// Rule corresponding to a lua script file
type Rule struct {
	File  string
	Proto *lua.FunctionProto

	Hahs       string
	LastModify time.Duration

	ruleCfg *RuleCfg
}

type RuleCfg struct {
	RuleID   string `toml:"id"`
	Category string `toml:"category"`
	Level    string `toml:"level"`
	Title    string `toml:"title"`
	Desc     string `toml:"desc"`
	Cron     string `toml:"cron"`
	Author   string `toml:"author,omitempty"`
}

func (c *RuleCfg) toLuaTable() *lua.LTable {
	var t lua.LTable
	t.RawSetString("id", lua.LString(c.RuleID))
	t.RawSetString("category", lua.LString(c.Category))
	t.RawSetString("level", lua.LString(c.Level))
	t.RawSetString("title", lua.LString(c.Title))
	t.RawSetString("desc", lua.LString(c.Desc))
	t.RawSetString("author", lua.LString(c.Author))
	return &t
}

func (r *Rule) run(ls *lua.LState) error {
	defer func() {
		if e := recover(); e != nil {
			log.Errorf("panic, %v", e)
		}
	}()
	var err error
	lfunc := ls.NewFunctionFromProto(r.Proto)
	ls.Push(lfunc)
	err = ls.PCall(0, lua.MultRet, nil)
	ls.Pop(ls.GetTop())
	return err
}

func (c *Checker) newRuleFromFile(path string) (*Rule, error) {

	manifestFile := strings.TrimRight(path, filepath.Ext(path)) + ".conf"

	manifest, err := ioutil.ReadFile(manifestFile)
	if err != nil {
		log.Errorf("%s", err)
		return nil, err
	}

	var rc RuleCfg
	if err := toml.Unmarshal(manifest, &rc); err != nil {
		log.Errorf("%s", err)
		return nil, err
	}

	if _, err := specParser.Parse(rc.Cron); err != nil {
		log.Errorf("invalid cron in: %s, %s", manifestFile, err)
		return nil, err
	}

	proto, err := compileLua(path)
	if err != nil {
		return nil, err
	}

	r := &Rule{
		File:    path,
		Proto:   proto,
		ruleCfg: &rc,
	}

	return r, nil
}

func compileLua(filePath string) (*lua.FunctionProto, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	chunk, err := luaparse.Parse(reader, filePath)
	if err != nil {
		log.Errorf("fail to parse lua file '%s', err: %s", filePath, err)
		return nil, err
	}
	proto, err := lua.Compile(chunk, filePath)
	if err != nil {
		log.Errorf("fail to compile lua file '%s', err: %s", filePath, err)
		return nil, err
	}
	return proto, nil
}
