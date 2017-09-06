package labimage

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// Get the image size of the given file on disk
func GetSize(fileName string) (width, height int, err error) {
	reader, err := os.Open(fileName)
	if err != nil {
		return 0, 0, err
	}

	defer reader.Close()

	im, _, err := image.DecodeConfig(reader)
	if err == nil {
		width = im.Width
		height = im.Height
	}

	return
}

func GetResolution(biggest bool) (width, height int) {
	out, err := exec.Command("system_profiler", "SPDisplaysDataType").Output()
	if err != nil {
		log.Println(err)
	}

	results := split(string(out), "\n")

	for _, property := range results {
		if strings.Index(property, "Resolution:") > 0 {
			res := replace(property, "Resolution:", "")
			res = strings.Trim(res, " ")

			values := split(res, " ")

			if len(values) >= 2 {
				w := castStringToInt(values[0])
				h := castStringToInt(values[2])

				if biggest {
					if w > width {
						width = w
						height = h
					}
				} else {
					if width == 0 || w < width {
						width = w
						height = h
					}
				}
			}
		}
	}

	return
}

func GetFormat(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		return ""
	}

	bytes := make([]byte, 4)
	n, _ := file.ReadAt(bytes, 0)

	if n < 4 {
		return ""
	}

	if bytes[0] == 0x89 && bytes[1] == 0x50 && bytes[2] == 0x4E && bytes[3] == 0x47 {
		return ".png"
	}

	if bytes[0] == 0xFF && bytes[1] == 0xD8 {
		return ".jpg"
	}

	if bytes[0] == 0x47 && bytes[1] == 0x49 && bytes[2] == 0x46 && bytes[3] == 0x38 {
		return ".gif"
	}

	if bytes[0] == 0x42 && bytes[1] == 0x4D {
		return ".bmp"
	}

	return ""
}

func GetExtension(fileName string) string {
	ext := GetFormat(fileName)

	if ext != "" {
		return ext
	}

	return getFileExtension(fileName)
}

func GetShortExtension(ext string) string {
	if ext == ".jpeg" {
		return ".jpg"
	}

	if ext == ".tiff" {
		return ".tif"
	}

	return ext
}

/*
Private Methods
*/

// CastStringToInt
func castIntToString(value int) string {
	return strconv.Itoa(value)
}

// Convert string to int
func castStringToInt(value string) int {
	res, _ := strconv.Atoi(value)

	return res
}

func split(s, sep string) []string {
	return strings.Split(s, sep)
}

func replace(s, sOld, sNew string) string {
	// -1 is for all instances of the string
	return strings.Replace(s, sOld, sNew, -1)
}

// Get the extension from a full filename
func getFileExtension(fileName string) string {
	return filepath.Ext(strings.ToLower(fileName))
}
