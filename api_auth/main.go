// 基礎的なリクエストユーザー判別&データ受信を行うためのAPI認証
package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

// DB : user database
var DB = map[string]string{
	"Taro":   "Apex46Uy",
	"Hanako": "Dong97Cc",
}

// CheckCorrect :
func CheckCorrect(apiKey string, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	if sign == expectedHMAC {
		fmt.Println(string(data))
		return
	}

	fmt.Println("Invalid User.")
}

func main() {
	// apiKeyへ標準入力から取り込み
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	apiKey := scanner.Text()
	apiSecret := "Apex46Uy"

	data := []byte("Hello World.")
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	sign := hex.EncodeToString(h.Sum(nil))

	CheckCorrect(apiKey, sign, data)
}
