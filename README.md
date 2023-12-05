# go-quick-start
> Go语言工程项目开发学习

## GO开发脚手架

<div align=center>
<img src="https://img.shields.io/badge/golang-1.16-blue"/>
<img src="https://img.shields.io/badge/gin-v1.7.7-lightBlue"/>
<img src="https://img.shields.io/badge/gorm-v1.25-red"/>
</div>

* 后端：用 [Gin](https://gin-gonic.com/) 快速搭建基础restful风格API。
* 数据库：支持`MySQL`, `PostgreSQL`, `SQLite`, `Oracle`, 使用 [gorm](http://gorm.cn) 实现对数据库的基本操作。
* 缓存：使用Redis实现记录当前活跃用户的jwt令牌并实现多点登录限制。
* 配置文件：使用 [fsnotify](https://github.com/fsnotify/fsnotify) 和 [viper](https://github.com/spf13/viper) 实现yaml格式的配置文件。
* 日志：使用 [zap](https://github.com/uber-go/zap) 实现日志记录。
* 参考 [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin) 功能精简出本项目开发脚手架
* 参考 [project-layout](https://github.com/golang-standards/project-layout)结构化目录结构规范，对`gin-vue-admin`目录进行改造

## 目录结构
```
.
├── api             (API 接口定义文件)
│   └── swagger
├── cmd             (组件 main 函数)
│   ├── gin_app
├── configs         (配置文件)
├── docs            (存放文档)
├── internal        (私有应用和库代码)
│   ├── app         (目录中存放真实的应用代码)
│   └── pkg         (存放项目内可共享，项目外不共享的包)
├── pkg             (可以被外部应用使用的代码库)
│   ├── log
│   ├── recovery
│   └── utils
├── scripts         (存放脚本文件)
└── web             (前端代码存放目录)
```

## 相关技术Demo
* gRPC
* Kafka
* Redis

