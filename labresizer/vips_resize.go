package labresizer

import (
	"errors"
	"os/exec"

	"github.com/h2non/bimg" // 10x faster than nfnt  but needs libvips (brew install vips)
)

func vips_resize(path, name, dest string, width, height int) error {
	if width == 0 || height == 0 {
		return errors.New("image size can't be null")
	}

	buffer, err := bimg.Read(path + name)
	if err != nil {
		return err
	}

	newImage, err := bimg.NewImage(buffer).Resize(width, height)
	if err != nil {
		return err
	}

	//size, err := bimg.NewImage(newImage).Size()
	//if size.Width == 400 && size.Height == 300 {
	//	fmt.Println("The image size is valid")
	//}

	if _, err := bimg.NewImage(newImage).Size(); err != nil {
		return err
	}

	bimg.Write(path+dest+".jpg", newImage)

	return nil
}

func checkVipsInstalled() bool {
	out, err := exec.Command("vips", "--version").Output()
	if err != nil {
		return false
	}

	return string(out)[:4] == "vips"
}

func getNewSize(width, height, size int) (newWidth int, newHeight int) {
	if width == 0 || height == 0 {
		return
	}

	if width > height {
		newWidth = size
		newHeight = size * height / width
	} else {
		newHeight = size
		newWidth = size * width / height
	}

	return
}
