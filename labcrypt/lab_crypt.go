package labcrypt

import (
	"bytes"
	"crypto/md5"
	"crypto/sha512"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
)

// EncryptJSON encrypt any object into a base64 encoded JSON
func EncryptJSON(object interface{}) string {
	return EncodeBase64(EncodeJSON(object))
}

// EncodeJSON encode any object into JSON
func EncodeJSON(object interface{}) []byte {
	b, err := json.Marshal(object)
	if err != nil {
		log.Println(err)
	}

	return b
}

func EncryptSHA512Hash(s string) string {
	h512 := sha512.New()
	io.WriteString(h512, s)

	return convertBytes2Hexa(h512.Sum(nil))
}

func BcryptHash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func CompareHashAndPassword(hashedPassword, password string) bool {
	// nil means it is a match
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

func EncodeGOB(object interface{}) string {
	var buffer bytes.Buffer

	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(object)
	if err != nil {
		log.Println("encode:", err.Error())
	}

	return EncodeBase64(buffer.Bytes())
}

func GetEncryptedString(codesArray []int) (str string) {
	for _, code := range codesArray {
		if str != "" {
			str += "-"
		}

		str += castIntToString(code)
	}

	return
}

func EncodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func DecodeBase64(s string) []byte {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Println("Error", err)
	}

	return b
}

func Encrypt(str string, keyBytes []byte) (result []int) {
	keys := len(keyBytes)

	if str != "" && keys > 0 {
		textBytes := []byte(str)
		for index, textByte := range textBytes {
			offset := keyBytes[index%keys]
			result = append(result, int(textByte+offset))
		}

		return
	}

	return
}

func Decrypt(encryptedArray []int, keyBytes []byte) string {
	keys := len(keyBytes)

	if len(encryptedArray) > 0 && keys > 0 {
		// Convert text to bytes
		var textBytes []byte

		for index, encryptedInt := range encryptedArray {
			offset := keyBytes[index%keys]
			textBytes = append(textBytes, byte(encryptedInt-int(offset)))
		}

		return string(textBytes)
	}

	return ""
}

// Get the hash code of the given filename on disk
func GetFileHashCode(fileName string) (hashCode string) {
	content, err := ioutil.ReadFile(fileName)
	if err == nil {
		hashCode = GetHashCode(string(content))
	}

	return
}

// Get the hash code of the given string
func GetHashCode(s string) string {
	hashCode := md5.New()
	io.WriteString(hashCode, s)

	return convertBytes2Hexa(hashCode.Sum(nil))
}

/*
Private Methods
*/

// Convert an array of byte in an hexadecimal value
func convertBytes2Hexa(n []byte) string {
	return fmt.Sprintf("%X", n)
}

// CastStringToInt
func castIntToString(value int) string {
	return strconv.Itoa(value)
}
