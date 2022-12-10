package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	mathRand "math/rand"
	"strings"
	"time"

	"github.com/nsrvel/go-fiber-boilerplate/config"
	uuid "github.com/nu7hatch/gouuid"
	"golang.org/x/exp/utf8string"
)

func GenerateUUID() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}

func GenerateStringWithCharset(length int, charset string) string {
	var seededRand *mathRand.Rand = mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateTimeNowJakarta() time.Time {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println(err)
	}
	now := time.Now().In(loc)
	return now
}

func GenerateRandomNumber(length int) string {
	var letterRunes = []rune("1234567890")

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateMaskPhoneNumber(phoneNumber string) string {
	length := len(phoneNumber)
	str := phoneNumber[0:3]
	var str1 string
	if strings.Contains(str, "0") {
		str1 = phoneNumber[0:2]
	} else if strings.Contains(str, "+") {
		str1 = phoneNumber[0:4]
	} else {
		str1 = phoneNumber[0:3]
	}
	str2 := phoneNumber[length-2 : length]
	var str3 string
	for i := 0; i < length; i++ {
		if i == 0 || i == 1 {
			str3 = str3 + "a"
		} else if i == length-2 || i == length-1 {
			str3 = str3 + "b"
		} else {
			str3 = str3 + "*"
		}
	}
	str4 := strings.Replace(str3, "aa", str1, 1)
	str4 = strings.Replace(str4, "bb", str2, 1)
	return str4
}

func GenerateMaskEmail(email string) string {
	array := strings.Split(email, "@")
	str1 := array[0]
	str2 := array[1]
	length := len(str1)
	strs := str1[0:1]
	stre := str1[length-1 : length]
	var mask string
	for i := 0; i < length-2; i++ {
		mask = mask + "*"
	}
	return fmt.Sprintf("%s%s%s@%s", strs, mask, stre, str2)
}

func GenerateBasicToken(conf *config.Config, timeStamp string) string {

	apiKey := conf.Authorization.Basic.ApiKey
	apiSecret := conf.Authorization.Basic.ApiSecret
	message := fmt.Sprintf("%s:%s:%v", apiKey, apiSecret, timeStamp)

	hash := sha256.Sum256([]byte(message))

	wordArray := utf8string.NewString(fmt.Sprintf("%s:%x", apiKey, hash))

	secret_buffer := base64.StdEncoding.EncodeToString([]byte(wordArray.String()))

	return secret_buffer
}
