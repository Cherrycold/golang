# golang
//2019.1.23
并发和并行的区别
并发相当于在同一个进程中，某一个线程执行一半，被暂停去执行另一个线程了
并行相当于同时有两个核，两个核分别处理两个线程，同时可以做到很多事情
学习gorilla/mux库

//2019.1.22
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

HandleFunc和Handle两个函数的区别是：
参数使用的不同
第一个->
func(ResponseWriter, *Request)
第二个->
type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
}