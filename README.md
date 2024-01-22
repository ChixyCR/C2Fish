
**C2FISH**
---
半成品项目，失业在家无聊和朋友写的，写一半不想写了，毕竟半途而废才是日站的主旋律。
代码基本上是从https://github.com/XZB-1248/Spark和https://github.com/MorouU/nightingale抄过来的，能用就行



## 安装教程

Mysql数据库自己安装好,并新建库导入c2fish.sql文件

在sever/config/config.go中更改数据库配置

然后按照Spark的说明

```
# 开始编译前端页面。
$ cd ./web
# 安装所有的依赖，然后编译。
$ npm install
$ npm run build-prod


# 通过statik，将前端文件嵌入到服务端里。
$ cd ..
$ go install github.com/rakyll/statik
$ statik -m -src="./web/dist" -f -dest="./server/embed" -p web -ns web


# 开始编译客户端。
# 在使用类Unix系统时，运行以下命令。
$ mkdir ./built
$ go mod tidy
$ go mod download
$ ./scripts/build.client.sh


# 最终开始编译服务端。
$ mkdir ./releases
$ ./scripts/build.server.sh

# 复制config.json到releases
$ cd ..
$ cp config.json ./releases

# 使用适合你系统的服务端运行
$ cd releases
$ ./server_
```

## To do

没有Todo，能用就行，想起来就写，想不起来算辣



## 开源协议

本项目基于 [BSD-2 协议](./LICENSE) 。