package util

import (
	"errors"
	"github.com/disintegration/imaging"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Trimming 切割图片
func Trimming(beforeFilename string, afterFilename string, x, y, w, h int) (string, error) {
	src, err := LoadImage(beforeFilename)
	if err != nil {
		log.Println("load image fail..")
		return "", errors.New("load image fail..")
	}
	img, err := ImageCopy(src, x, y, w, h)
	if err != nil {
		log.Println("image copy fail..")
		return "", errors.New("image copy fail..")
	}
	saveErr := SaveImage(afterFilename, img)
	if saveErr != nil {
		log.Println("save image fail..")
		return "", errors.New("image copy fail..")
	}
	return afterFilename, nil
}

// LoadImage 加载图片
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

// SaveImage 保存图片
func SaveImage(p string, src image.Image) error {
	f, err := os.OpenFile(p, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	ext := filepath.Ext(p)
	if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {
		err = jpeg.Encode(f, src, &jpeg.Options{Quality: 80})
	} else if strings.EqualFold(ext, ".png") {
		err = png.Encode(f, src)
	} else if strings.EqualFold(ext, ".gif") {
		err = gif.Encode(f, src, &gif.Options{NumColors: 256})
	}
	return err
}

// ImageCopy 裁剪图片
func ImageCopy(src image.Image, x, y, w, h int) (image.Image, error) {
	var subImg image.Image
	if rgbImg, ok := src.(*image.YCbCr); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.RGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.RGBA) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.NRGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.NRGBA) //图片裁剪x0 y0 x1 y1
	} else {
		return subImg, errors.New("图片解码失败")
	}
	return subImg, nil
}

// Thumbnail 生成缩略图
func Thumbnail(beforeFilename string, afterFilename string, width int, height int) (string, error) {
	src, err := imaging.Open(beforeFilename)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
		return "", err
	}
	imgWidth := 0
	imgHeight := 0
	if width > height {
		imgWidth = 180
		imgHeight = 120
	} else {
		imgWidth = 120
		imgHeight = 180
	}
	//生成缩略图，从中间截取
	//imaging.Center, imaging.Lanczos
	dsc := imaging.Fill(src, imgWidth, imgHeight, imaging.Center, imaging.Lanczos)
	err = imaging.Save(dsc, afterFilename)
	return afterFilename, nil
}
