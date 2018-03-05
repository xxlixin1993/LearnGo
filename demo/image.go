package main

import (
	"code.google.com/p/graphics-go/graphics"
	"image"
	_ "image/jpeg"   //必须import，否则会出现：unknown format，其余类似
	"image/png"
	"log"
	"os"
)

//读取文件
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("open file fail:", err)
		return
	}

	defer file.Close()
	img, _, err = image.Decode(file)   //解码图片
	return
}

//保存文件
func SaveImage(path string, img image.Image) (err error) {
	imgfile, err := os.Create(path)

	defer imgfile.Close()
	err = png.Encode(imgfile, img)   //编码图片
	if err != nil {
		log.Fatal("Save fail:", err)
	}
	return
}

func main() {
	src, err := LoadImage("./demo/1.jpg")
	if err != nil {
		log.Fatal("open fail:", err)
	}
	dst := image.NewRGBA(image.Rect(0, 0, 100, 100))
	err = graphics.Scale(dst, src)   //缩小图片
	if err != nil {
		log.Fatal(err)
	}
	SaveImage("./demo/thumbnailimg.png", dst)
}