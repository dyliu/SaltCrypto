package SaltCrypto

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test_Crypto(t *testing.T) {
	//测试次数
	countCase := 20
	pwds := make(map[string]string, countCase)
	password := "abc123456"

	Convey("测试", t, func() {
		for index := 0; index < countCase; index++ {
			//取新盐值
			salt := strconv.Itoa(int(rand.Int31n(10)))
			//加密
			cryptoedPwd := Crypto(password, salt)
			fmt.Printf("Salt:%s,PWD:%s\n", salt, cryptoedPwd)
			//生成32位加密密码
			ShouldEqual(32, len(cryptoedPwd))

			if _, has := pwds[salt]; !has {
				pwds[salt] = cryptoedPwd
			} else {
				//盐值相等，结果相等，反之
				ShouldEqual(pwds[salt], cryptoedPwd)
			}

		}
	})
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
