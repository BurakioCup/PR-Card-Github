package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kbinani/screenshot"
	"image/png"
	"io"
	"net/http"
	"os"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func main(){
	Server = gin.Default()
	Server.GET("/github/read",ReadImage())
	Server.GET("/github/a",A())
	addr := GetServerPort()
	Server.Run(addr)
}

func ReadImage()gin.HandlerFunc {
	return func(c *gin.Context) {
		user :=c.GetHeader("user")
		var url string
		url = "https://github-readme-stats.vercel.app/api?username="+user+"&count_private=true&show_icons=true"
		fmt.Println(url)
		response, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		file, err := os.Create("save.png")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		io.Copy(file, response.Body)
	}
}

func A()gin.HandlerFunc {
	return func(c *gin.Context) {
		bounds := screenshot.GetDisplayBounds(0)
		img, _ := screenshot.CaptureRect(bounds)

		file, _ := os.Create("out.png")
		defer file.Close()
		png.Encode(file, img)
	}
}


func GetServerPort() string{
	var addr string
	port := os.Getenv("PORT")
	if (port == "") {
		port = "8080"
	}
	// 接続情報は以下のように指定する.
	// user:password@tcp(host:port)/database
	flag.StringVar(&addr, "addr", ":"+port, "tcp host:port to connect")
	flag.Parse()
	return addr
}
