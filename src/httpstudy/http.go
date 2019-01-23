package httpstudy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//构造http响应，   http接收到的请求
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	/* if r.Method != "POST" {
		return
	} */
	//Method是请求方法
	if r.Method == "POST" {
		data, err := ioutil.ReadAll(r.Body)
		fmt.Println(data)
		if err != nil {
			w.WriteHeader(400)
			fmt.Println("400")
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(200)
		//将返回的内容保存在w中
		fmt.Fprintln(w, "200OK")
		fmt.Println("200")
		w.Write(data)
		return
	}
	w.Write([]byte("hello,world"))
}

type handelers func(http.ResponseWriter, *http.Request)
type Handlers map[string]handelers

func createNewMux() {
	//这个函数可以用来新建一个路由的规则
	var mux = http.NewServeMux()
	//创建一个http请求路由器
	//Handler是处理器

	//用307将请求的任意网页都重定向到百度
	rh := http.RedirectHandler("http://www.baidu.com", 307)
	mux.Handle("/", rh)

	//http.ListenAndServe("8080",mux)
	//ServeMux也有ServeHTTP因此可以作为第二个参数进行使用
}

type Test struct {
	timing string
}

func (t Test) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(t.timing)
	w.Write([]byte("time is :" + tm))
}

func Start() {
	//1.
	//普通使用，登陆就注册
	/* handlers := Handlers{
		"/login": HandleLogin,
	}
	for pattern, hander := range handlers {
		http.HandleFunc(pattern, hander)

	http.ListenAndServe(":8080", nil)
	} */

	//2.
	//第二个参数是自定义路由方法
	//通过重写ServeHTTP方法来实现
	test1 := Test{time.RFC1123}
	http.ListenAndServe(":8080", test1)
}
