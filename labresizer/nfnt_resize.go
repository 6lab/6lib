package labresizer

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/nfnt/resize"

	"github.com/6lab/6lib/labfile"
)

func nfnt_resize(path, name, dest string, size uint) error {
	img, err := openImage(path + name)
	if err != nil {
		return err
	}

	res := resize.Thumbnail(size, size, img, resize.Lanczos3)

	out, err := os.Create(path + dest + ".jpg")
	if err != nil {
		return err
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, res, nil)

	return nil
}

func openImage(fullFileName string) (image.Image, error) {
	f, err := os.Open(fullFileName)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return decodeImage(fullFileName, f)
}

func decodeImage(fullFileName string, f *os.File) (image.Image, error) {
	mime, err := xtfile.GetMime(fullFileName)
	if err != nil {
		return nil, err
	}

	switch mime {
	case "image/jpeg":
		// decode jpeg into image.Image
		return jpeg.Decode(f)

	case "image/png":
		// decode png into image.Image
		return png.Decode(f)
	}

	return nil, errors.New("this type of file has not been implemented")
}
