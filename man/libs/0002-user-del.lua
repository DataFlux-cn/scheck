local system = require("system")local cache = require("cache")local common = require("common")local sc_file = require("file")local function user_del_check(...) local cache_key="current_users" local currents=system.users() local old=cache.get_cache(cache_key) if old == nil then  old = system.users() end local content='' local count = 0 for i,v in ipairs(old) do  local exist=false  for ii,vv in ipairs(currents) do   if vv["username"] == v["username"] and vv["uid"] == vv["uid"] then    exist = true    break   end  end  if not exist then   if content ~= '' then content=content.."; " end   count = count + 1   content = content..string.format("%s", v["username"])  end end if content ~= '' then  trigger({Content=content, Num=tostring(count)}) end  cache.set_cache(cache_key, currents)endlocal function check() local cache_key="current_users" local currents=system.users() --local old=cache.get_cache(cache_key) cache.set_cache(cache_key, currents) local file = "/etc/passwd" while not exit do  if sc_file.file_exist(file)  then   common.watcher(file, {1, 2, 4, 8, 16}, user_del_check)  end  system.sc_sleep(5) endendcheck()