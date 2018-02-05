# IceCream

IceCream.Server

## Installation

```sh
$ git clone https://github.com/XMatrixStudio/IceCream.git
$ cd IceCream
$ go run main.go
```

## Documents

```
# APIs
Base: /ice-cream/v1

/User/Login :POST   登录

/Comment    :GET    获取评论 加载博客时展示
/Comment    :POST   创建评论
/Comment    :DELETE 删除评论

/Article    :GET    获取博客 修改博客前获取
/Article    :POST   新建博客
/Article    :PATCH  修改博客
/Article    :DELETE 删除博客
```
