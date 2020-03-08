# ncov-report-wx-server-Golang

> 这里是NOCO 2020疫情防控-人员健康管理平台开源项目的小程序后端--Golang版本。

主项目入口 >> https://github.com/2020NCOV/ncov-report

## 目录结构

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

**核心部分就是以上加粗部分的文件**

## 与小程序交互流程 —— 以getcode接口为例

![image-20200308094548037](/Users/hua/Library/Application Support/typora-user-images/image-20200308094548037.png)