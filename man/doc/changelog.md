# Scheck 版本历史

## 1.0.2（2021/09/23）

### 发布说明

本次发布对 Scheck 做了较大的调整，主要涉及性能、稳定性方面。

- 内存 性能优化
- 调整文件监听检测方式，替换原来文件缓存的形式
- 规则脚本除了间歇运行方式，增加一种常驻运行分类，后者常用于监听类场景,如:[文件变更](https://www.yuque.com/dataflux/sec_checker/funcs#sc_path_watch) 等。
- 增加 Lua 脚本执行超时控制
- 增加 Lua 运行的[统计信息](https://www.yuque.com/dataflux/sec_checker/best-practices#c5609495)
- 增加命令行 `-check` [功能](https://www.yuque.com/dataflux/sec_checker/best-practices#c5609495)

## v1.0.1-67-gd445240（2021/6/18）
### 发布说明

脚本相关：

- 添加mysql 弱密码扫描

新功能相关：

- 添加3个func


## v2.0.0-67-gd445240（2021/8/27）
### 发布说明

脚本相关：

- 添加多个容器检测脚本

新功能相关：

- 修改了scheck 配置
- 添加阿里云日志对接
- 添加自定义rule路径和用户自定义lua.libs路径
- 添加cgroup功能
- 用户自定义rule脚本自动更新功能
- windows的powershell安装和linux环境下的shell安装
- 服务第一次启动时发送info信息
- output支持多端发送信息

优化相关：
- cpu 性能优化
- 语雀文档重构


## v1.0.1-67-gd445240（2021/6/18）
### 发布说明

脚本相关：

- 添加mysql 弱密码扫描

新功能相关：

- 添加3个func




## v1.0.1-62-g7715dc6
### 发布说明

本次发布，对Security Checker 操作层面名称统一确认为scheck。修改了相关bug

脚本相关：

- 修改了desc内容间隔
- 调整脚本的执行频率

新功能相关：

- 添加scheck 自身md5选项

### Bug 修复

- 优化脚本运行性能
