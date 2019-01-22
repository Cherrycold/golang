package httpstudy

import "net/http"

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	/* if r.Method != "POST" {
		return
	} */
	w.Write([]byte("hello,world"))
}
func Start() {
	http.HandleFunc("/", HandlerIndex)
	http.ListenAndServe(":8080", nil)
}
