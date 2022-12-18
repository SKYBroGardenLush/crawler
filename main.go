package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"net/http"
)

func main() {
	url := "https://www.thepaper.cn/"
	//url := "http://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	enc := mahonia.NewEncoder("gbk")
	fmt.Println(enc.ConvertString(string(body)))
}
