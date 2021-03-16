package git

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func GetGit()gin.HandlerFunc {
	return func(c *gin.Context) {
		user :=c.GetHeader("user")
		var url string
		url = "https://github-readme-stats.vercel.app/api?username="+user+"&count_private=true&show_icons=true"
		resp, _ := http.Get(url)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(byteArray)) // htmlをstringで取得
	}
}

func ReadImage()gin.HandlerFunc {
	return func(c *gin.Context) {
		user :=c.GetHeader("user")
		var url string
		url = "https://github-readme-stats.vercel.app/api/top-langs/?username="+user
		response, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		file, err := os.Create("save.svg")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		io.Copy(file, response.Body)

		a:=decode()
		fmt.Println(a)
		//a:=decode()
		//fmt.Println(a)
		//url = "https://initial-practice.s3-ap-northeast-1.amazonaws.com/second-sprintReview/save.svg"
		//
		//response, _ = http.Get(url)
		//fmt.Println(response)
	}
}

func decode()string{
	w, h := 512, 512
	in, err := os.Open("save.svg")
	if err != nil {
		panic(err)
	}
	defer in.Close()
	icon, _ := oksvg.ReadIconStream(in)
	icon.SetTarget(0, 0, float64(w), float64(h))
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)

	out, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = png.Encode(out, rgba)
	if err != nil {
		panic(err)
	}
	return ""
}
//file, _ := os.Open("save.svg")
//defer file.Close()
//
//fi, _ := file.Stat() //FileInfo interface
//size := fi.Size()    //ファイルサイズ
//
//data := make([]byte, size)
//file.Read(data)
//
//
//return base64.StdEncoding.EncodeToString(data)