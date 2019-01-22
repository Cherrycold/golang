# golang
学习http服务器的创建
需要头文件 net/http

首先，最简单的使用就是http.HandleFunc和http.ListenAndServe
HandleFunc用来处理客户端发来的http请求
第一个参数用来指定路由的路径
ListenAndserve用来监听请求，然后创建一个goroutine去进行服务
第二个参数可以自定义路由的内容

（才知道go的方法是有重写这一说的，学习了）

其中Request负责接收请求的内容
Method是HTTP方法	GET POST等等
ResponseWriter负责写入返回的内容
返回头WriterHeader
返回内容Write

