package initializers

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/labstack/echo"
	"sort"
	"strconv"
	"time"
)

//签名校验
func checkSign(context echo.Context, secretKey string, params *map[string]string) bool {
	sign := (*params)["signature"]
	targetStr := context.Request().Method + "|" + context.Path() + "|"
	keys := make([]string, len(*params)-1)
	i := 0
	for k, _ := range *params {
		if k == "signature" {
			continue
		}
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for i, key := range keys {
		if i > 0 {
			targetStr += "&"
		}
		targetStr += key + "=" + (*params)[key]
	}
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(targetStr))
	return sign == fmt.Sprintf("%x", mac.Sum(nil))
}

//邮箱限流
func LimitTrafficWithEmail(context echo.Context) bool {

	return false
}

//IP限流
func LimitTrafficWithIp(context echo.Context) bool {
	return false
}

//时间戳校验
func checkTimestamp(context echo.Context, params *map[string]string) bool {
	timestamp, _ := strconv.Atoi((*params)["timestamp"])
	now := time.Now()
	if int(now.Add(-time.Second*60*5).Unix()) <= timestamp && timestamp <= int(now.Add(time.Second*60*5).Unix()) {
		return true
	}
	return false
}
