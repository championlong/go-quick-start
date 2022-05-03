package utils

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strings"
)

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")


//去掉uuid的连接符，得到32位的uuid
func GetPureUUID() string {
	uid, _ := uuid.NewUUID()
	return strings.Replace(uid.String(), "-", "", -1)
}

// 4位，100并发10000次请求，重复率千分之一
func GetRandNChars(n int) (string, error) {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b), nil
}

func GetRandNCharsOld(numOfChar int) (string, error) {
	if numOfChar > 32 {
		return "", errors.New("md5 max length exceed")
	}
	u := GetPureUUID()
	md5Inst := md5.New()
	md5Inst.Write([]byte(u))
	result := md5Inst.Sum([]byte(""))
	resultStr := fmt.Sprintf("%x", result)
	return resultStr[:numOfChar], nil
}
