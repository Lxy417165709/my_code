package main

import (
	"fmt"
	"github.com/aybabtme/rgbterm"
	"github.com/nfnt/resize"
	"image"
	_ "image/jpeg" // 一定要导入，导入后才能解析相应格式的图片，否则会报错: image: unknown format。
	_ "image/png"  // 一定要导入，导入后才能解析相应格式的图片，否则会报错: image: unknown format。
	"log"
	"os"
	"strings"
)

// 字符集。
var (
	charSet1 = []byte(`@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\|()1{}[]?-_+~<>i!lI;:,\"^'. `)
	charSet2 = []byte(`@&W0*+. `)
	charSet3 = []byte(`@ `)
)

// main 主函数。
func main() {
	// 1. 打开图片。
	imgPath := `C:\Users\李学悦\Desktop\all\github\Lxy417165709\my_code\go\char_photo\photo\xm.jpg`
	img, err := OpenImage(imgPath)
	if err != nil {
		log.Fatalf("OpenImage fail. err=[%+v].", err)
		return
	}

	// 2. 缩小图片。
	width := uint(img.Bounds().Dx())
	height := uint(img.Bounds().Dy())
	scalingFactor := uint(5) // 缩小倍数。
	ratio := uint(2)         // 终端中单个字符的高度/宽度的比值。
	img = resize.Resize(width/scalingFactor, height/scalingFactor/ratio, img, resize.Lanczos3)

	// 3. 形成像素矩阵。
	pixelMatrix := buildPixelMatrix(img, charSet2)

	// 4. 在终端输出图像。
	fmt.Println(pixelMatrix.toCharPhotoGray())
	fmt.Println(pixelMatrix.toCharPhotoColorful())
}

// OpenImage 打开图片，返回图片对象。
func OpenImage(path string) (image.Image, error) {
	// 1. 打开文件。
	imgFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("Open fail. err=[%+v].", err)
		return nil, err
	}
	defer imgFile.Close()

	// 2. 解析为图片对象。
	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Fatalf("Decode fail. err=[%+v].", err)
		return nil, err
	}

	// 3. 返回。
	return img, nil
}

// -------------------------------------------------------------------------------------------------------------------

// PixelMatrix 像素矩阵。
type PixelMatrix struct {
	Matrix [][]*Pixel
}

// toCharPhotoGray 转为字符图片。 (灰白)
func (m *PixelMatrix) toCharPhotoGray() string {
	charLines := make([]string, 0)
	for y := 0; y < len(m.Matrix); y++ {
		var charLine string
		for x := 0; x < len(m.Matrix[y]); x++ {
			charLine += string(m.Matrix[y][x].Char)
		}
		charLines = append(charLines, charLine)
	}
	return strings.Join(charLines, "\n")
}

func (m *PixelMatrix) toCharPhotoColorful() string {
	charLines := make([]string, 0)
	for y := 0; y < len(m.Matrix); y++ {
		var charLine string
		for x := 0; x < len(m.Matrix[y]); x++ {
			pixel := m.Matrix[y][x]
			charLine += decorateWithColor(pixel.Char, pixel.Color)
		}
		charLines = append(charLines, charLine)
	}
	return strings.Join(charLines, "\n")
}

// ----------------------------------------------

// Pixel 像素。
type Pixel struct {
	Color *Color // 颜色。
	Char  byte   // 字符。
}

// Color 颜色。
type Color struct {
	R, G, B uint8
}

// -------------------------------------------------------------------------------------------------------------------

// buildPixelMatrix 构建像素矩阵。
func buildPixelMatrix(img image.Image, charSet []byte) *PixelMatrix {
	matrix := make([][]*Pixel, 0)
	for y := 0; y < img.Bounds().Dy(); y++ {
		pixels := make([]*Pixel, 0)
		for x := 0; x < img.Bounds().Dx(); x++ {
			r, g, b := get256RGB(img, x, y)
			gray := 0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)
			charIndex := int(gray / 256.0 * float64(len(charSet)))
			pixels = append(pixels, &Pixel{
				Color: &Color{R: r, G: g, B: b},
				Char:  charSet[charIndex],
			})
		}
		matrix = append(matrix, pixels)
	}
	return &PixelMatrix{Matrix: matrix}
}

// get256RGB 获取数值在区间[0, 256)的RGB。
func get256RGB(img image.Image, x int, y int) (uint8, uint8, uint8) {
	r, g, b, _ := img.At(x, y).RGBA() // 返回的数值区间为[0, 65536)。
	return uint8(r / 256), uint8(g / 256), uint8(b / 256)
}

// decorateWithColor 用颜色装饰字符，返回ANSI序列。
// 序列字符串如: "\x1b[38;5;245m9\x1b[0m"，在使用 fmt.Print 输出该字符串，终端会输出一个灰色的 9。
func decorateWithColor(char byte, color *Color) string {
	return rgbterm.FgString(string([]byte{char}), color.R, color.G, color.B)
}
