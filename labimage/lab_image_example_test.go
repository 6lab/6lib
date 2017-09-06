package labimage

import (
	"testing"
)

// import (
// 	"fmt"
// )

// func ExampleStringBuild() {
// 	fmt.Println(StringBuild("SELECT * FROM tProject WHERE name LIKE '%%1%' AND idVersion = '%2'", "Technologies", "IDCH4"))
// 	// Output: SELECT * FROM tProject WHERE name LIKE '%Technologies%' AND idVersion = 'IDCH4'
// }

// jpg file
func TestGetSize_jpg(t *testing.T) {
	width, height, err := GetSize("test/test.jpg")
	if err != nil {
		t.Error(err.Error())
	}

	if width == 0 || height == 0 {
		t.Error(err.Error())
	}
}

// png file
func TestGetSize_png(t *testing.T) {
	width, height, err := GetSize("test/test.png")
	if err != nil {
		t.Error(err.Error())
	}

	if width == 0 || height == 0 {
		t.Error(err.Error())
	}
}

// gif file
func TestGetSize_gif(t *testing.T) {
	width, height, err := GetSize("test/test.gif")
	if err != nil {
		t.Error(err.Error())
	}

	if width == 0 || height == 0 {
		t.Error(err.Error())
	}
}

func BenchmarkGetSize_jp(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		GetSize("test/image.jpg")
	}
}
