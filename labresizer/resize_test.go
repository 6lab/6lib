package labresizer

import (
	"fmt"
	"testing"
	"time"

	"github.com/6lab/6lib/labfile"
)

func TestResize1(t *testing.T) {
	path := "test/file1/"

	xtfile.RemoveFile(path + "cover.jpg")
	xtfile.RemoveFile(path + "thumbnail.jpg")

	if err := Resize(path + "image1.jpg"); err != nil {
		t.Error(err.Error())
	}

	if xtfile.CheckIfFileExist(path + "cover.jpg") {
		t.Error("cover.jpg must not be generated")
	}

	if !xtfile.CheckIfFileExist(path + "thumbnail.jpg") {
		t.Error("thumbnail.jpg has not been generated")
	}
}

func TestResize2(t *testing.T) {
	path := "test/file2/"

	xtfile.RemoveFile(path + "cover.jpg")
	xtfile.RemoveFile(path + "thumbnail.jpg")

	if err := Resize(path + "image2.jpg"); err != nil {
		t.Error(err.Error())
	}

	if !xtfile.CheckIfFileExist(path + "cover.jpg") {
		t.Error("cover.jpg has not been generated")
	}

	if !xtfile.CheckIfFileExist(path + "thumbnail.jpg") {
		t.Error("thumbnail.jpg has not been generated")
	}
}

func TestResize3(t *testing.T) {
	path := "test/file3/"

	xtfile.RemoveFile(path + "cover.jpg")
	xtfile.RemoveFile(path + "thumbnail.jpg")

	t0 := time.Now()

	if err := Resize(path + "image3.jpg"); err != nil {
		t.Error(err.Error())
	}

	if !xtfile.CheckIfFileExist(path + "cover.jpg") {
		t.Error("cover.jpg has not been generated")
	}

	if !xtfile.CheckIfFileExist(path + "thumbnail.jpg") {
		t.Error("thumbnail.jpg has not been generated")
	}

	fmt.Println("Converted in", time.Since(t0))
}
