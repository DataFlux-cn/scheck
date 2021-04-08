
## Security Checker
一般在运维过程中有非常重要的工作就是对系统，软件，包括日志等一系列的状态进行巡检，传统方案往往是通过工程师编写shell（bash）脚本进行类似的工作，通过一些远程的脚本管理工具实现集群的管理。但这种方法实际上非常危险，由于系统巡检操作存在权限过高的问题，往往使用root方式运行，一旦恶意脚本执行，后果不堪设想。  
实际情况中存在两种恶意脚本，一种是恶意命令，如 `rm -rf`，另外一种是进行数据偷窃，如将数据通过网络 IO 泄露给外部。 
因此 Security Checker 希望提供一种新型的安全的脚本方式（限制命令执行，限制本地IO，限制网络IO）来保证所有的行为安全可控，并且 Security Checker 将以日志方式通过统一的网络模型进行巡检事件的收集。同时 Security Checker 将提供海量的可更新的规则库脚本，包括系统，容器，网络，安全等一系列的巡检。


## 安装/更新

*安装*：  
```Shell
bash -c "$(curl https://zhuyun-static-files-testing.oss-cn-hangzhou.aliyuncs.com/security-checker/install.sh)"
```

*更新*：  
```Shell
bash -c "$(curl https://zhuyun-static-files-testing.oss-cn-hangzhou.aliyuncs.com/security-checker/install.sh) --upgrade"
```

安装完成后即以服务的方式运行，服务名为`security-checker`，使用服务管理工具来控制程序的启动/停止：  
`systemctl start/stop/restart security-checker`  
或 `service security-checker start/stop/restart`  


>Security Checker 目前仅支持 Linux

## 配置

进入默认安装目录 `/usr/local/security-checker`，打开配置文件 `checker.conf`，配置文件采用 [TOML](https://toml.io/en/) 格式，说明如下：

```toml
# ##(必选) 存放检测脚本的目录
rule_dir='/usr/local/security-checker/rules.d'

# ##(必选) 将检测结果采集到哪里，支持本地文件或http(s)链接
# ##本地文件时需要使用前缀 file://， 例：file:///your/file/path
# ##远程server，例：http(s)://your.url
output=''

# ##(可选) 程序本身的日志配置
log='/usr/local/security-checker/log'
log_level='info'
```

>配置文件的修改需要重启服务才生效


## 检测规则
检测规则放在规则目录中：由配置文件中 `rule_dir` 指定。每项规则对应两个文件：  
1. 脚本文件：使用 [Lua](http://www.lua.org/) 语言来编写，必须以 `.lua` 为后缀名。    
2. 清单文件：使用 [TOML](https://toml.io/en/) 格式，必须以 `.manifest` 为后缀名，[详情](#清单文件)。  
脚本文件和清单文件必须同名。   

Security Checker 会定时周期性(由清单文件的 `cron` 指定)地执行检测脚本，lua脚本代码每次执行时检测相关安全事件(比如文件被改动、有新用户登录等)是否触发，如果触发了，则使用 [trig]() 函数将事件(以行协议格式)发送到由配置文件中 `output` 指定的目标中。   

Security Checker 定义了若干 lua 扩展函数，并且为确保安全，禁用了一些lua包或函数，仅支持以下 lua 内置包/函数：  
```
内置基础函数，如print
table
math
string
debug
os - 其中以下函数被禁用："execute", "remove", "rename", "setenv", "setlocale"
```

>添加/修改规则不需要重启服务，Security Checker 会每隔 10 秒扫描规则目录。

### 清单文件 
清单文件是对当前规则所检测内容的一个描述，比如检测文件变化，端口启停等。最终的行协议数据中只会包含清单文件中的字段。详细内容如下：    
```toml
# ##必选字段

# ##事件的规则编号，如 k8s-pod-001，将作为行协议的指标名
id=''

# ##事件的分类，根据业务自定义，如 security，network
category=''

# ##当前事件的危险等级，根据业务自定义，比如warn，error
level=''

# ##当前事件的标题，描述检测内容，如 "敏感文件改动"
title=''

# ##当前事件的内容（支持模板，详情见下）
desc=''

# ##配置事件的执行周期（使用 linux crontab 的语法规则）
cron=''

# ##可选字段

# ##禁用当前规则
#disabled=false

# ##默认在tag中添加hostname
#omit_hostname=false

# ##显式设置hostname
#hostname=''


# ##自定义字段

# ##支持添加自定义key-value，且value必须为字符串
#instanceID=''

```

#### 模板支持
清单文件中 `desc` 的字符串中支持设置模板变量，语法为 `{{.<Variable>}}`，比如 `文件{{.FileName}}被修改，改动后的内容为: {{.Content}}` 表示 FileName 和 Diff 是模板变量，将会被替换(包括前面的点号`.`)。变量的替换发生在调用 `trig` 函数时，该函数可传入一个 `table` 变量，其中包含了模板变量的替换值，假设传入如下参数：  
```lua
tmpl_vals={
    FileName="/etc/passwd",
    Content="delete user demo"
}
trig(tmpl_vals)
```
则最终的 `desc` 值为：`文件/etc/passwd被修改，改动内容为: delete user demo`


## 测试规则
在编写规则代码的时候，可以使用 Security Checker 来测试代码是否正确：  
```
checker --test  /ruledir/demo
```


## lua 函数
见 [函数](./funcs.md)


---

## 示例

### 一. 检查敏感文件的变动，将事件记录到文件 `/var/log/security-checker/event.log` 中。    

1. 进入安装目录，编辑配置文件 `checker.conf` 的 `output` 字段：  
```toml
rule_dir='/usr/local/security-checker/rules.d'

output='/var/log/security-checker/event.log'

log='/usr/local/security-checker/log'
log_level='info'
```

2. 在目录`/usr/local/security-checker/rules.d`(即以上配置文件中的`rule_dir`)下新建清单文件 `files.manifest`，编辑如下：  
```toml
id='check-file'
category='security'
level='warn'
title='监视文件变动'
desc='文件 {{.File}} 发生了变化'
cron='*/10 * * * *' #表示每10秒执行该lua脚本
```

3. 在清单文件同级目录下新建脚本文件 `files.lua`，编辑如下：
```lua
local files={
	'/etc/passwd',
	'/etc/group'
}
local function check(file)
	local cache_key=file
	local hashval = file_hash(file)

	local old = get_cache(cache_key)
	if not old then
		set_cache(cache_key, hashval)
		return
	end

	if old ~= hashval then
		err = trig({File=file})
		if err == "" then
			set_cache(cache_key, hashval)
		else
			print("error: "..err)
		end
	end
end

for i,v in ipairs(files) do
	check(v)
end
```

4. 当有敏感文件被改动后，下一个10秒会检测到并触发 trig 函数，从而将事件发送到文件 `/var/log/security-checker/event.log` 中，添加一行数据，例：  
```
check-file-01,category=security,level=warn,title=监视文件变动 message="文件 /etc/passwd 发生了变化" 1617262230001916515
```

---

### 二. 监控系统用户的变化。    
省略其它步骤，只显示规则代码和清单文件。  

`users.manifest`：  

```toml
id='users-checker'
category='security'
level='warn'
title='监视系统用户变动'
desc='{{.Content}}'
cron='*/10 * * * *'

instanceId='id-xxx'
```

`users.lua`：  

```lua
local function check()
    local cache_key="current_users"
    local currents=users()

    local old=get_cache(cache_key)
    if not old then
        set_cache(cache_key, currents)
        return
    end

    local adds={}
    for i,v in ipairs(currents) do
        local exist=false
        for ii,vv in ipairs(old) do
            if vv["username"] == v["username"] then
                exist = true
                break
            end
        end
        if not exist then
            table.insert(adds, v["username"])
        end
    end

    local dels={}
    for i,v in ipairs(old) do
        local exist=false
        for ii,vv in ipairs(currents) do
            if vv["username"] == v["username"] then
                exist = true
                break
            end
        end
        if not exist then
            table.insert(dels, v["username"])
        end
    end

    local content=''
    if #adds > 0 then
        content=content..'新用户: '..table.concat(adds, ',')
    end
    if #dels > 0 then
        if content ~= '' then content=content..'; ' end
        content=content..'删除的用户: '..table.concat(dels, ',')
    end
    if content ~= '' then
        local err=trig({Content=content})
        if err == ''  then
            set_cache(cache_key, currents)
        end
    end
end

check()
```

---

### 三. 监控是否有新的端口开启，将事件发送到 DataKit 中。    

1. 将配置文件 `checker.conf` 的 `output` 设置为 DataKit 的server地址：  
```toml
rule_dir='/usr/local/security-checker/rules.d'

output='http://datakit-ip:9529/v1/write/log'

log='/usr/local/security-checker/log'
log_level='info'
```

2. 在目录`/usr/local/security-checker/rules.d`(即以上配置文件中的`rule_dir`)下新建清单文件 `ports.manifest`，编辑如下：  
```toml
id='check-listening-ports'
category='security'
level='warn'
title='监视新端口'
desc='{{.Content}}'
cron='*/10 * * * *'
```

3. 在清单文件同级目录下新建脚本文件 `ports.lua`，编辑如下：
```lua
ocal function check()
    local cache_key='check_ports'
    local ports = listening_ports()
    local old=get_cache(cache_key)
    if not old then
        set_cache(cache_key, ports)
        return
    end

    local content=''
    for i,v in ipairs(ports) do
        if v["family"] == 'AF_INET' then
            local exist=false
            for ii,vv in ipairs(old) do
                if vv["port"] == v["port"] and vv["address"] == v["address"]  then
                    exist = true
                    break
                end
            end
            if not exist then
                if content ~= '' then content=content.."; " end
                content = content..string.format("%d(%s)", v["port"], v["protocol"])
            end
        end
    end

    if content ~= '' then
        local err=trig({Content="发现新端口: "..content})
        if err == ''  then
            set_cache(cache_key, ports)
        end
    end
end

check()
```
