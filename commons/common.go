package commons

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/draw"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ResizeImg(x int, y int, sourcefile string, destinationfile string) {
	// Read input file.
	// f, err := os.Open("assets/photos/"+ k.Request.FormValue("userid") +"/photo." + k.Request.FormValue("filetype"))
	f, err := os.Open(sourcefile)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	src, _, err := image.Decode(f)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Scale down by a factor of 2.
	sb := src.Bounds()
	// dst := image.NewRGBA(image.Rect(0, 0, 64, 64))
	dst := image.NewRGBA(image.Rect(0, 0, x, y))
	draw.BiLinear.Scale(dst, dst.Bounds(), src, sb, draw.Over, nil)

	// Write output file.
	// if f, err = os.Create("assets/photos/"+ k.Request.FormValue("userid") +"/photo_64." + k.Request.FormValue("filetype")); err != nil {
	if f, err = os.Create(destinationfile); err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	var opt jpeg.Options
	opt.Quality = 80
	if err := jpeg.Encode(f, dst, &opt); err != nil {
		if err := png.Encode(f, dst); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func DecryptKey(sInput string) string {
	sKey := "Aa8B1bCcDd9Ee2F.fGg0HhIi3JjKkLlMm_Nn4OoPpQqRr5SsTtUuVv6WwXxYy7Zz -:"

	sRet := ""
	for i := 0; i < len(sInput); i++ {
		cText := sInput[i:(i + 1)]
		cPos := strings.Index(sKey, cText)
		nPos := 0
		if cPos-(len(sInput)+i) < 0 {
			nPos = len(sKey) - ((len(sInput) + i) - cPos)
		} else {
			nPos = cPos - (len(sInput) + i)
		}

		for {
			if nPos < 0 {
				nPos = (len(sKey)) + nPos
			} else {
				break
			}
		}

		rVal := sKey[nPos:(nPos + 1)]

		sRet = sRet + rVal
	}

	return sRet
}
