package main // 每个可执行程序都需要这个

import (
	"fmt"      // 用于输入输出（打印和写入网页）
	"net/http" // 用于创建web服务器
)

func main() {
	// 定义一个处理函数，当有人访问网站时调用这个函数
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		// w 是用于向浏览器写内容的
		// r 是浏览器发来的请求信息
		fmt.Fprintf(w, "Hello, Cloud Native! 这是我的第一个Go程序。")
	}

	// 告诉服务器：当有人访问根路径"/"时，使用helloHandler函数来处理
	http.HandleFunc("/", helloHandler)

	// 打印提示信息
	fmt.Println("服务器启动在 http://localhost:8080")
	fmt.Println("请在浏览器中访问上面的地址")

	// 启动服务器，监听8080端口
	http.ListenAndServe(":8080", nil)
}
