package web

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/kataras/iris"
	"sort"
	"strings"
)

func HomeController(ctx iris.Context) {
	ctx.View("index.html")
}

func auth(ctx iris.Context) {
	data := ctx.URLParams()
	if checkTelegramAuthorization(data) {
		ctx.Text("yes")

	} else {
		ctx.Text("error")
	}

}

func checkTelegramAuthorization(data map[string]string) bool {

	hash := data["hash"]
	delete(data, "hash")

	var dataCheckList []string
	for k, v := range data {
		dataCheckList = append(dataCheckList, k+"="+v)
	}
	sort.Strings(dataCheckList)
	dataCheckString := strings.Join(dataCheckList, "\n")

	sha256Hash := sha256.New()

	sha256Hash.Write([]byte(botToken))
	hmacHash := hmac.New(sha256.New, sha256Hash.Sum(nil))
	hmacHash.Write([]byte(dataCheckString))

	return hex.EncodeToString(hmacHash.Sum(nil)) == hash
}
