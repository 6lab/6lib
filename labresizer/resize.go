package labresizer

import (
	"github.com/6lab/6lib/labfile"
	"github.com/6lab/6lib/labimage"
)

const (
	SizeCover         = 360
	SizeMin4Cover     = 500
	SizeThumbnail     = 120
	SizeMin4Thumbnail = 200
)

var (
	canUseVips bool
)

func init() {
	canUseVips = checkVipsInstalled()
}

// Resize generates 2 new files (cover.jpg and thumbnail.jpg )
func Resize(fullFileName string) error {
	// Path + filename
	path := xtfile.GetFilePath(fullFileName) + "/"
	name := xtfile.GetFileName(fullFileName)

	width, height := xtimage.GetSize(fullFileName)

	// Resize as Cover and preserve aspect ratio
	if width > SizeMin4Cover || height > SizeMin4Cover {
		if canUseVips {
			newWidth, newHeight := getNewSize(width, height, SizeCover)
			vips_resize(path, name, "cover", newWidth, newHeight)
		} else {
			nfnt_resize(path, name, "cover", SizeCover)
		}
	}

	// Resize as Thumbnail and preserve aspect ratio
	if width > SizeMin4Thumbnail || height > SizeMin4Thumbnail {
		if canUseVips {
			newWidth, newHeight := getNewSize(width, height, SizeCover)
			vips_resize(path, name, "cover", newWidth, newHeight)
		} else {
			nfnt_resize(path, name, "thumbnail", SizeThumbnail)
		}
	}

	return nil
}

func CheckUseVips() bool {
	return canUseVips
}
