package labpath

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/6lab/6lib/labfile"
	"github.com/6lab/6lib/labrand"
	"github.com/6lab/6lib/labstring"
)

const (
	PathDB         = "database/"
	PathWork       = "ExpoZ/"
	PathLog        = "log/"
	PathTemp       = "temp/"
	PathMedia      = "media/"
	PathClipart    = "cliparts/"
	PathAutoImport = "autoimport/"
	PathUpload     = "upload/"
	PathDownload   = "download/"
	PathCert       = "cert/"
)

type PATH struct {
	Path string `json:"path"`
}

func GetPath(iniPath, pathType string) (path string) {
	// Assign Path
	switch pathType {
	case "db":
		path = GetDatabasePath(iniPath)
	case "media":
		path = GetMediaPath()
	case "cliparts":
		path = GetClipartPath()
	}

	return
}

/* Paths Configurable in INI */

func GetWorkingPath(iniPath string) string {
	path := xtstring.GetValue(GetDefaultExpoZPath(), iniPath)
	return xtfile.CreatePathIfNotExists(path)
}

func GetDatabasePath(iniPath string) string {
	path := xtstring.GetValue(GetWorkingPath(iniPath)+PathDB, iniPath)
	return xtfile.CreatePathIfNotExists(path)
}

func GetLogPath(iniPath string) string {
	path := xtstring.GetValue(GetWorkingPath(iniPath)+PathLog, iniPath)
	return xtfile.CreatePathIfNotExists(path)
}

func GetCertPath(iniPath string) string {
	path := xtstring.GetValue(CleanPath(GetEXEPath())+PathCert, iniPath)
	return xtfile.CreatePathIfNotExists(path)
}

/*  Fix Paths according to the previous Paths */

func GetMediaPath() string {
	return xtfile.CreatePathIfNotExists(GetWorkingPath("") + PathMedia)
}

func GetClipartPath() string {
	return xtfile.CreatePathIfNotExists(GetMediaPath() + PathClipart)
}

func GetAutoImportPath() string {
	return xtfile.CreatePathIfNotExists(GetMediaPath() + PathAutoImport)
}

/* OS depending Path */

func GetTempPath(unique bool) string {
	path := GetMediaPath() + PathTemp

	if unique {
		path = path + xtrand.GetUUID() + "/"
	}

	return xtfile.CreatePathIfNotExists(path)
}

func GetTempFile(ext string) string {
	return GetMediaPath() + xtrand.GetUUID() + "." + ext
}

func GetUserPath() string {
	user, err := user.Current()
	if err != nil {
		log.Println(err)
	}

	return user.HomeDir + "/"
}

// Private
func GetEXEPath() string {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
	}

	return path
}

/* Other Paths */

func GetDownloadPath() string {
	return xtfile.CreatePathIfNotExists(GetMediaPath() + PathDownload)
}

func GetUploadPath(unique bool) string {
	path := GetMediaPath() + PathUpload

	if unique {
		path = path + xtrand.GetUUID() + "/"
	}

	return xtfile.CreatePathIfNotExists(path)
}

func GetDownloadTempPath(unique bool) string {
	path := GetDownloadPath() + PathTemp

	if unique {
		path = path + xtrand.GetUUID() + "/"
	}

	return xtfile.CreatePathIfNotExists(path)
}

func GetDefaultExpoZPath() string {
	// By Default, the ExpoZ Path is the Parent Directory of the EXE Path
	return CleanPath(filepath.Dir(GetEXEPath()))
}

func CleanPath(path string) string {
	return xtstring.Replace(xtfile.CompleteDir(path), "$HOME/", GetUserPath())
}
