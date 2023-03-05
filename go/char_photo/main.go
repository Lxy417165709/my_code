package main

import (
	"fmt"
	"github.com/Lxy417165709/gcm/char_photo"
)

func main() {
	path := `C:\Users\李学悦\Desktop\all\github\Lxy417165709\my_code\go\char_photo\photo\xm.jpg`
	pixelMatrix, err := char_photo.BuildPixelMatrixByImgPath(path, 100, 50, char_photo.CharSet2)
	if err != nil {
		panic(err)
	}
	fmt.Println(pixelMatrix.ToCharPhotoColorful())
}
