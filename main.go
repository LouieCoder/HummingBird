package main

/*
1 读取目录下所有excel文件
2 遍历这些文件并将它们保存到数据库
*/
import (
	_ "HummingBird/dao"
	"HummingBird/service"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go http.ListenAndServe(":8080",nil)

	service.SaveExcelFormDir("exl/")
	//阻塞
	select {}
}
