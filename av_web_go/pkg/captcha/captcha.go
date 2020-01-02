package captcha

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/golang/freetype"
)

const (
	fontFile = "./pkg/captcha/font/Arial.ttf"
	fontDPI  = 72
	fontSize = 40
)

// CreateCaptcha 返回一张图片
func CreateCaptcha() ([]byte, string, error) {
	// base64Captcha.ConfigCharacter
	// imgfile, _ := os.Create("tmp.png")
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		// fmt.Println(err)
		return nil, "", err
	}
	backFile, _ := os.Open("./pkg/captcha/font/background.jpg")
	defer backFile.Close()
	backImg, err := png.Decode(backFile)
	if err != nil {
		// fmt.Println(err)
		return nil, "", err
	}
	img := image.NewNRGBA(backImg.Bounds())
	// img := resize.Resize(100, 40, backImg, resize.Lanczos3)
	draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Src)
	draw.Draw(img, img.Bounds(), backImg, image.Point{X: 0, Y: 0}, draw.Src)

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		// fmt.Println(err)
		return nil, "", err
	}

	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.Black)
	c.SetHinting(2)

	x, y := randomPlace()
	// fmt.Println(x, y)
	pt := freetype.Pt(x, y)

	rs := getRandomString(5)
	_, err = c.DrawString(rs, pt)
	if err != nil {
		// fmt.Println(err)
		return nil, "", err
	}
	var outImg []byte
	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		// fmt.Println(err)
		return nil, "", err
	}
	outImg = buf.Bytes()
	return outImg, rs, nil
}

func randomPlace() (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(60), rand.Intn(40) + 30
}

func getRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
