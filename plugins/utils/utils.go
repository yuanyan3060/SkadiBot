package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
)

func Image2Base64(image image.Image) []byte {
	buffer := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, buffer)
	var opt jpeg.Options
	opt.Quality = 70
	_ = jpeg.Encode(encoder, image, &opt)
	err := encoder.Close()
	if err != nil {
		return nil
	}
	return buffer.Bytes()
}

func Bytes2Base64(body []byte) string {
	return base64.StdEncoding.EncodeToString(body)
}

func init() {

}
