# gin-web

#### 项目介绍
基于gin-gonic开发的api接口框架，通过特征实现版本迭代。

#### 软件架构

|-- app
	|-- common
		|-- helpers #工具助手
		|-- libraries #本地库
		|-- models  #数据模型
		|-- pools   #连接池
		|-- repositories #数据仓库
		|-- services #数据服务
	|-- config		#配置类
	|-- modules #模块文件
		|-- api
			|-- config  #模块配置文件
			|-- controllers #模块控制器
			|-- middlewares #模块中间件
			|-- router.go  #模块路由
|-- gin-web.go

#### 使用说明

$ go build gin-web.go