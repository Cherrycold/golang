# golang

2019.1.26
---
### **git再学习**
remote的使用
git remote origin master
其中origin相当于远程仓库的一个别名
master是远程仓库的分支名

2019.1.25
---

### **反射的使用**
用在运行时操作对象
reflect.Value和reflect.type
用来获取任意对象的值和类型

使用 “值.Interface().(类型)”来还原

v.Elem()  
Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值。

v.Kind()
Kind返回v持有的值的分类，如果v是Value零值，返回值为Invalid



### **sync/atomic**  
原子操作，用来实现同步  
类型共有六种：int32, int64, uint32, uint64, uintptr, unsafe.Pinter  
数据的增减使用Add，读取使用Load
sync.WaitGroup
Add()，Done()，Wait()，加，减，阻塞直到所有goroutine执行结束


### **内嵌类型**  
可以在一个结构体或接口中嵌套一个接口或者结构体

### **Actor模型的学习**  
每个actor之间是独立的，他是封装的是状态和行为的对象
异步的，消息和发送者之间是解耦的

### **接口和断言**  
如果一个类型实现了这些方法，那么他就是这个接口类型的实例   
**只有接口存储的类型和对象都是nil，接口才是nil**  
类似于x.(T)，x是一个接口类型，T表示一个类型  
会返回一个结果，一个是目标类型的值，一个是bool  
对于x.(type)，可以搭配switch使用，进行类型的判断  


2019.1.24
---
### flag包学习

提供解析命令行参数的功能接口

flag.Parse()解析命令行到定义的flag

首先定义解析的方式，然后通过Parse解析

多余的参数使用flag.Args()进行接收

flag.String(),Bool(), Int()返回一个相对应的指针

flag.XxxVar()将flag绑定到一个变量，返回值类型

flag.var()绑定自定义类型，实现Value接口


接口的使用？？？？？？？？？？？？？？？？？？？？？？？？？？？？？？？？

看不懂，明天看



//2019.1.23

并发和并行的区别

并发相当于在同一个进程中，某一个线程执行一半，被暂停去执行另一个线程了

并行相当于同时有两个核，两个核分别处理两个线程，同时可以做到很多事情

可以说并行是某一时刻多个程序运行，并发是某一时段多个程序运行


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
```go
func(ResponseWriter, *Request)
```
第二个->
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```
