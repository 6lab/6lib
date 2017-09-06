package labfile

import (
	"fmt"
	"testing"
)

/* GetMime */

/* EXAMPLE */

func ExampleGetMime() {
	mime, err := GetMime("test/test.jpg")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(mime)
	}
	// Output: image/jpeg
}

/* TEST */

// jpg file
func TestGetMime_jpg(t *testing.T) {
	mime, err := GetMime("test/test.jpg")
	if err != nil {
		t.Error(err.Error())
	}

	if mime != "image/jpeg" {
		t.Error("expected 'image/jpeg', got '" + mime + "'")
	}
}

// png file
func TestGetMime_png(t *testing.T) {
	mime, err := GetMime("test/test.png")
	if err != nil {
		t.Error(err.Error())
	}

	if mime != "image/png" {
		t.Error("expected 'image/png', got '" + mime + "'")
	}
}

// gif file
func TestGetMime_gif(t *testing.T) {
	mime, err := GetMime("test/test.gif")
	if err != nil {
		t.Error(err.Error())
	}

	if mime != "image/gif" {
		t.Error("expected 'image/gif', got '" + mime + "'")
	}
}

// docx file
func TestGetMime_docx(t *testing.T) {
	mime, err := GetMime("test/test.docx")
	if err != nil {
		t.Error(err.Error())
	}

	if mime != "application/zip" {
		t.Error("expected 'application/zip', got '" + mime + "'")
	}
}

// xlsx file
func TestGetMime_xlsx(t *testing.T) {
	mime, err := GetMime("test/test.xlsx")
	if err != nil {
		t.Error(err.Error())
	}

	if mime != "application/zip" {
		t.Error("expected 'application/zip', got '" + mime + "'")
	}
}

// rtf file
func TestGetMime_rtf(t *testing.T) {
	mime, err := GetMime("test/test.rtf")
	if err != nil {
		t.Error(err.Error())
	}

	if mime != "text/plain; charset=utf-8" {
		t.Error("expected 'application/zip', got '" + mime + "'")
	}
}

// txt file
func TestGetMime_txt1(t *testing.T) {
	mime, err := GetMime("test/test1.txt")
	if err != nil {
		t.Error(err.Error())
	}

	if mime != "text/plain; charset=utf-16le" {
		t.Error("expected 'text/plain; charset=utf-16le', got '" + mime + "'")
	}
}

// txt file
func TestGetMime_txt2(t *testing.T) {
	mime, err := GetMime("test/test2.txt")
	if err != nil {
		t.Error(err.Error())
	}

	if mime != "text/plain; charset=utf-8" {
		t.Error("expected 'text/plain; charset=utf-8', got '" + mime + "'")
	}
}
