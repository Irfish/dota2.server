### dota2.server  一些dota2地图的服务器
- 每个服务器均可独立运行
- 每个服务器启动之后都会注册到etcd中
- 每当有服务器注册之后，网关服务器都会与之保持连接（web服务器除外）
#### service-db DB服务器
- 其他服务器与db(如mysql等)的中间人
#### service-gw 网管服务器
- 与客户端保持长链接
- 消息转发
- 隐藏内部服务器
#### serivce-log 日志服务器
- 收集服务器日志 
#### service-login 登录服务器
- 用户登录
#### service-web web服务器
- 后台管理系统
#### sql 相关sql文件
#### config 服务器相关配置
- login.json 登录服务器相关配置 
- gw.json  网关服务器相关配置 
- xorm.cfg mysql 配置
- token.cfg 全局token
- redis.cfg redis 配置


