# ncov-report-wx-server-Golang

> 这里是NCOV 2020疫情防控-人员健康管理平台开源项目的小程序后端--Golang版本。  详细项目记录及学习文档见[doc文件夹](https://github.com/huagua/MiniProgram-server-Golang/tree/master/doc)

主项目入口 >> https://github.com/2020NCOV/ncov-report  

![build](https://github.com/2020NCOV/MiniProgram-server-Golang/workflows/build/badge.svg)

## 项目导航
- [ncov-report-wx-server-Golang](#ncov-report-wx-server-golang)
    - [项目导航](#项目导航)
    - [目录结构](#目录结构)
    - [与小程序交互流程](#与小程序交互流程)
    - [项目本地配置](#项目本地配置)

## 目录结构
```
|-- MiniProgram-server-Golang
    |-- .env.example —— 这个是对环境变量如何设置的示例，在自己的项目中改名为.env并加入.gitignore
    |-- .gitignore—— 上传代码时一些不需要上传的文件就加入.gitignore
    |-- go.mod ——  Golang包管理工具
    |-- main.go  —— 项目入口
    |-- **api**——  定义接口
    |   |-- main.go							通用接口
    |   |-- user.go				  			用户相关的接口
    |-- conf ——  配置文件
    |   |-- conf.go
    |-- middleware—— 中间件
    |   |-- cors.go  			    			跨域相关设置
    |-- **model**——主要数据库表的设计，这里是重点需要修改的地方，需要配合统一的数据库表结构
    |   |-- code.go			
    |   |-- init.go
    |   |-- migration.go
    |   |-- report.go
    |   |-- user.go
    |-- **serializer**——序列化器，返回请求时用来序列化数据
    |   |-- main.go
    |   |-- report.go
    |   |-- user.go
    |-- **server**——主要定义了路由
    |   |-- router.go  
    |-- **service** —— 针对每个请求的具体服务逻辑
        |-- check_is_registered.go
        |-- check_user.go
        |-- get_corpname_service.go
        |-- get_info_service.go
        |-- save_daily_info_service.go
        |-- user_bind_service.go
        |-- user_openid_service.go
        |-- wexin_user_register.go
```
**核心部分就是以上加星部分的文件**

## 与小程序交互流程 
   以getcode接口为例

![流程图](http://q6uspeueh.bkt.clouddn.com/requestRoute.png)

## 项目本地配置
### 1.导入Goland
    可能会下载一会儿包，稍微等待一下。将项目创建在%GOPATH%/src目录下可以加载之前已经安装在%GOPATH%/pkg中的包。

### 2.在项目根目录新建文件.env, 内容如下：
```
MYSQL_DSN="db_user:db_passwd@tcp(127.0.0.1:3306)/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接配置
GIN_MODE="debug"            # 设置gin的运行模式，有 debug 和 release
APP_ID=""                   #appid
APP_SECRET=""              #appsecret
```
注：
- 其实redis没有用到，环境变量中不写也可
- 将db_user和db_passwd修改为自己本地的mysql连接的用户名和密码，并创建相应的数据库，将db_name替换成对应数据库名
- 补充自己注册的小程序的app_id和app_secret

### 3. 执行命令go run main.go, 如果RUN窗口出现如下字样，则代表后端程序启动成功
```
[GIN-debug] POST   /index/login/getcode      --> Miniprogram-server-Golang/api.UserLogin (3 handlers)
[GIN-debug] POST   /index/login/check_is_registered --> Miniprogram-server-Golang/api.UserIsReg (3 handlers)
[GIN-debug] POST   /index/login/check_user   --> Miniprogram-server-Golang/api.CheckUser (3 handlers)
[GIN-debug] POST   /index/login/register     --> Miniprogram-server-Golang/api.WeixinUsrRegister (3 handlers)
[GIN-debug] POST   /index/login/getcorpname  --> Miniprogram-server-Golang/api.GetCorp (3 handlers)
[GIN-debug] POST   /index/login/bind         --> Miniprogram-server-Golang/api.UserBind (3 handlers)
[GIN-debug] POST   /index/login/unbind       --> Miniprogram-server-Golang/api.UserUnBind (3 handlers)
[GIN-debug] POST   /index/report/save        --> Miniprogram-server-Golang/api.SaveInfo (3 handlers)
[GIN-debug] POST   /index/report/getlastdata --> Miniprogram-server-Golang/api.GetInfo (3 handlers)
[GIN-debug] POST   /index/info/getmyinfo           --> Miniprogram-server-Golang/api.GetUserInfo (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```
 ### 4.修改小程序端的baseURL,在/ncov-report-mini-program/util/config.js文件中
 ```
 const baseURL = 'http://127.0.0.1:8080/index'; //这表示小程序访问的是本机的8080端口，正是后端程序监听的端口
 ```
 ### 5.测试接口
 - 编译运行小程序
 - 打开调试器，点击network
 - 查看小程序发出的请求getcode，如果返回status code是200OK则表示前后端通信成功
