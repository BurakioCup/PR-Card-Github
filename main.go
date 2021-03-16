package main

import (
	"PR-Card-Github/pkg/git"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func main(){
	Server = gin.Default()
	Server.GET("/github/read",git.ReadImage())
	Server.GET("/github/get",git.GetGit())
	Server.GET("/github/a",A())
	addr := GetServerPort()
	Server.Run(addr)
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
