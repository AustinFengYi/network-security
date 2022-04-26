// Golang 資料、資料型態與變數


package main // 可執行程式必須使用 main 封包
import "fmt" // 載入內建的 fmt 封包，用來做基本輸出輸入
func main(){ // 建立 main 程式，程式的進入點
		// 執行程式時，從這邊開始
		// 輸出訊息到終端: fmt.Println(資料,資料...)

		// 1. 資料型態介紹
		/*
		fmt.Println(3) //  int
		fmt.Println(4.3455) // float 64
		fmt.Println("測試中文") // String
		fmt.Println(true) //boolean
		fmt.Println('a') // rune 字符
		*/

		// 2. 變數的使用
		var x int // 宣告變數 x
		x = 4 // 指定資料
		fmt.Println(x) // 使用變數
		x = 10 // 指定新的資料
		fmt.Println(x)		

		var f float64 = 4.35355
		fmt.Println(f)

		var s string = "Hello, Golang"
		fmt.Println(s)

		var test bool=true
		fmt.Println(test)
}