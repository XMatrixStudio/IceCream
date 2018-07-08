# IceCream


[![version](https://img.shields.io/badge/alpha-v0.1-blue.svg)](https://github.com/XMatrixStudio/IceCream)

## 什么是冰淇淋

**冰淇淋博客框架(IceCream Blog Framework)**是一个简单的、支持静态化的博客框架。

不喜欢类似WordPress的复杂博客框架？冰淇淋旨在实现一个从界面的操作尽可能简洁化的框架。

不喜欢类似Hexo的完全静态博客框架？冰淇淋在支持大量页面静态化的前提下，同时支持动态交互。

冰淇淋后端基于Golang、MongoDB，得益于两者的性能，使得冰淇淋在服务层能更快速地静态化网站。

冰淇淋前端基于Golang Text/HTML Template，模块化的前端设计使得前端主题的分离实现更简单。

## 如何使用冰淇淋

```sh
$ git clone https://github.com/XMatrixStudio/IceCream.git
$ cd IceCream
$ go run main.go
```

使用冰淇淋需要提前创建好MongoDB数据库，并且将配置信息写入`config/config.yaml`中。

## 版本迭代

* v0.2 (等待开发~)
  * 权限管理：最高管理员、管理员、作者、普通用户、小黑屋用户
  * 评论用户的评论
  * 个人用户黑名单
  * 文件、图片上传
  * 自定义页面静态化生成
  * 网站前端主题切换
  * 网页通知系统、邮箱通知系统
  * 优化静态化过程
  * 完全静态化
* v0.1
  * 基于Violet中央授权系统的用户模块
  * 支持文章的评论
  * 支持About关于界面的Markdown解析
  * 支持修改网站显示名字和显示URL

## 关于我们

冰淇淋由XMatrix工作室开发，如果有什么建议，不妨在[Issues](https://github.com/XMatrixStudio/IceCream/issues)里面联系我们，谢谢！

