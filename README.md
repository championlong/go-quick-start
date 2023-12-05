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

## 相关技术Demo
* gRPC
* Kafka
* Redis

