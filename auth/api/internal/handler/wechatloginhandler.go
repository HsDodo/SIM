package handler

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"server/auth/api/internal/svc"
	"sort"
	"strings"
)

func wechatLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 微信公众平台测试，Token验证,微信服务器携带4个参数过来，需要验证签名
		signature := r.URL.Query().Get("signature")
		timestamp := r.URL.Query().Get("timestamp")
		nonce := r.URL.Query().Get("nonce")
		echostr := r.URL.Query().Get("echostr")
		// 验证签名
		token := "hsen1015"
		//开发者通过检验signature对请求进行校验（下面有校验方式）。若确认此次GET请求来自微信服务器，请原样返回echostr参数内容，则接入生效，成为开发者成功，否则接入失败。加密/校验流程如下：
		//1）将token、timestamp、nonce三个参数进行字典序排序
		//2）将三个参数字符串拼接成一个字符串进行sha1加密
		//3）开发者获得加密后的字符串可与signature对比，标识该请求来源于微信
		tmpArr := []string{token, timestamp, nonce}
		sort.Strings(tmpArr)
		tmpStr := strings.Join(tmpArr, "")
		h := sha1.New()
		io.WriteString(h, tmpStr)
		if fmt.Sprintf("%x", h.Sum(nil)) == signature { //验证成功
			_, err := w.Write([]byte(echostr))
			if err != nil {
				fmt.Println("write err:", err)
			}
		}
	}
}
