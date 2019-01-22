package httpstudy

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func Start() {
	//登陆就注册
	handlers := Handlers{
		"/login": HandleLogin,
	}
	for pattern, hander := range handlers {
		http.HandleFunc(pattern, hander)
	}
	//第二个参数是自定义路由方法
	//通过重写ServeHTTP方法来实现
	http.ListenAndServe(":8080", nil)
}
