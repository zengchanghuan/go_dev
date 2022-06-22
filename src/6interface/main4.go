package main

import (
	"fmt"
	"log"
	"net/http"
)

//创建 float32 类型的自定义类型，然后编写 String() 方法的自定义实现
type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

/*
写入 http.Handler 可使用的 ServeHTTP 方法的实现。
请注意，我们重新创建了自定义类型，但这次它是映射，而不是结构。
接下来，我们通过使用 database 类型作为接收方来写入 ServeHTTP 方法。
此方法的实现使用来自接收方的数据，然后对其进行循环访问，再输出每一项
*/
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

/*
我们将 database 类型实例化，并使用一些值对其进行初始化。
我们使用 http.ListenAndServe 函数启动了 HTTP 服务器，
在其中定义了服务器地址，包括要使用的端口和实现 ServerHTTP 方法自定义版本的 db 对象
*/
func main() {
	db := database{"Go T-Shirt": 25, "Go Jacket": 55}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
