# golang
学习http服务器的创建
需要头文件 net/http

首先，最简单的使用就是http.HandleFunc和http.ListenAndServe
HandleFunc用来处理客户端发来的http请求