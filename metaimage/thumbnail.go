package metaimage

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
)

func Thumbnail(path string, maxWidth, minWidth uint) image.Image {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	// decode file into image.Image
	image, _, err := image.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
	defer file.Close()

	t1 := resize.Thumbnail(maxWidth, minWidth, image, resize.Bilinear)

	return t1

}
