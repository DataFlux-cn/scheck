local sc_file = require("file")
local cache = require("cache")
local function grep_sudofile(file)
    if sc_file.file_exist(file) then
        local data = sc_file.grep("-Ei",'^\\s*Defaults\\s+([^#;]+,\\s*)?logfile\\s*=\\s*(")?[^#;]+(")?', file)
        if data ~= '' then
            return true
        end
    end
    return false
end

function log_file()
   local flag = false
   if sc_file.file_exist("/etc/sudoers") then
       if grep_sudofile("/etc/sudoers") then
           flag = true
       end
   end
   if sc_file.file_exist("/etc/sudoers.d/") then
        local filelist = sc_file.ls("/etc/sudoers.d/")
        for ii,value in ipairs(filelist) do
            if grep_sudofile(value) then
                flag = true
                break
            end
        end
   end
   return flag
end

local function check()
    local cache_key = "/var/log/sudo.log"
    local current = log_file()
   -- print(current)
    local old = cache.get_cache(cache_key)

    if old == nil then
        if not current
        then
            trigger()
        end
        cache.set_cache(cache_key, current)
        return
    end

    if old ~= current then
        if not current
        then
            trigger()
        end
        cache.set_cache(cache_key, current)
    end

end
check()
