package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"server/common/cache"
	"server/common/etcd"
	"server/common/response"
	"server/constants"
	"server/utils"
	"strconv"
)

type Proxy struct {
	CacheBox *cache.CacheBox
}

type Config struct {
	Addr         string
	Etcd         string
	Log          logx.LogConf
	ReverseProxy bool
	Redis        struct {
		Addr     string
		Password string
	}
	WhiteList []string
	UseCache  bool
}

var c Config

// 鉴权服务
func auth(authAddr string, res http.ResponseWriter, req *http.Request) bool {
	authReq, err := http.NewRequest("POST", fmt.Sprintf("http://%s/api/auth/authentication", authAddr), nil)
	if err != nil {
		logx.Error("网关: 创建请求失败")
		return false
	}
	token := req.Header.Get("Authorization")
	if token != "" {
		authReq.Header.Set("Authorization", token)
	}
	authReq.Header.Set("ValidPath", req.URL.Path) // 设置请求路径 用于鉴权
	authResp, err := http.DefaultClient.Do(authReq)
	if err != nil {
		logx.Error("网关: 鉴权失败")
		return false
	}
	// 解析鉴权服务返回的数据
	type Response struct {
		Code uint32 `json:"code"`
		Msg  string `json:"msg"`
		Data *struct {
			UserId   uint   `json:"userId"`
			NickName string `json:"nickName"`
			Role     uint   `json:"role"`
		} `json:"data"`
	}
	var authResponse Response
	byteData, err := io.ReadAll(authResp.Body)
	err = json.Unmarshal(byteData, &authResponse)
	if err != nil {
		logx.Error("网关: 解析鉴权服务返回的数据失败")
		return false
	}
	if authResponse.Code != 0 {
		logx.Error("网关: 鉴权失败")
		return false
	}
	if authResponse.Code == 600 {
		return false
	}

	if authResponse.Data != nil { // 将用户信息传到 request 中
		req.Header.Set("userID", fmt.Sprintf("%d", authResponse.Data.UserId))
		req.Header.Set("nickName", authResponse.Data.NickName)
		req.Header.Set("role", strconv.Itoa(int(authResponse.Data.Role)))
	}
	return true
}

// CacheProxyHandler 缓存代理处理
func CacheProxyHandler(w http.ResponseWriter, r *http.Request, cacheBox *cache.CacheBox, proxyAddr string) {
	// 新建url,并复制一份r.URL到新建的url中
	newURI := new(url.URL)
	*newURI = *r.URL
	newURI.Host = proxyAddr
	newURI.Scheme = "http"
	//上面新建了一个 uri , 但是目标主机替换成了 proxyAddr, 也就是说请求的目标主机是 proxyAddr, 资源最终的目的地
	logx.Infof("准备读取缓存 Key: %s", newURI.String())
	// 从缓存中获取数据
	uriCache := cacheBox.Get(newURI.String())
	if uriCache != nil {
		if uriCache.Verify() {
			_, err := uriCache.WriteTo(w)
			if err != nil {
				logx.Error("网关: 从缓存中获取数据失败")
				return
			}
		}
	}
	// 没命中缓存, 从服务中获取数据
	// 处理请求
	reverseRequestHandler(r, proxyAddr)
	RemoveProxyHeader(r)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		logx.Error("网关(缓存处理): 从服务中获取数据失败")
		return
	}
	defer resp.Body.Close()
	// 将响应数据写入缓存
	cacheBox.CheckAndStore(newURI.String(), resp, proxyAddr)
	ClearHeaders(w.Header())
	CopyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		logx.Error("网关(缓存代理): 写入响应失败")
		return
	}
}

// HttpProxyHandler http代理处理
func HttpProxyHandler(w http.ResponseWriter, r *http.Request, proxyAddr string) {
	logx.Infof("网关: 请求路径: %s 服务主机: %s", r.URL.Path, proxyAddr)
	RemoveProxyHeader(r)
	reverseRequestHandler(r, proxyAddr)
	resp, err := http.DefaultTransport.RoundTrip(r) // 发送请求HttpProxyHandler
	if err != nil {
		logx.Error("网关: 从服务中获取数据失败")
		return
	}
	defer resp.Body.Close()
	ClearHeaders(w.Header())
	CopyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	nr, err := io.Copy(w, resp.Body)
	if err != nil {
		logx.Error("网关(HTTP代理): 写入响应失败")
		return
	} else {
		logx.Infof("网关(HTTP代理): 响应数据写入成功, 写入字节数: %d", nr)
	}

}

func CopyHeaders(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func reverseRequestHandler(req *http.Request, proxyAddr string) {
	// 反向代理处理, 处理请求, 将请求转发到对应的服务
	req.Host = proxyAddr
	req.URL.Host = proxyAddr
	req.URL.Scheme = "http"
	fmt.Printf("%v", req.RequestURI)
}

func RemoveProxyHeader(req *http.Request) {
	req.RequestURI = ""
	req.Header.Del("Proxy-Connection")
	req.Header.Del("Proxy-Authenticate")
	req.Header.Del("Proxy-Authorization")
	req.Header.Del("Connection")
	req.Header.Del("Upgrade")
	req.Header.Del("Keep-Alive")
	req.Header.Del("Transfer-Encoding")
	req.Header.Del("TE")
	req.Header.Del("Trailer")
}

func ClearHeaders(headers http.Header) {
	for key := range headers {
		headers.Del(key)
	}
}

func InWhiteList(serName string, whiteList []string) bool {
	for _, v := range whiteList {
		if serName == v {
			return true
		}
	}
	return false
}

func (p Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 1. 解析请求
	// 2. 鉴权
	// 3. 反向代理

	//获取请求路径  网关传入路径为 http://{host}/api/{服务名}/{接口名}  需要捕获服务名,然后使用反向代理将请求转发到对应的服务
	// 服务名 类似 auth.api、user.api 等
	reg, err := regexp.Compile(`/api/(.*?)/`)
	if err != nil {
		logx.Error("正则表达式错误")
		return
	}
	fmt.Printf("网关: 请求路径: %s\n", r.URL.Path)

	// 获取服务名和接口名
	matches := reg.FindStringSubmatch(r.URL.Path)
	if len(matches) != 2 {
		logx.Error("请求路径错误")
		return
	}
	serName := fmt.Sprintf("%s.api", matches[1]) // rpc 对内, api 对外, 也就是网关这里只用暴露api的服务就行, rpc对内使用
	var authAddr string
	// 判断是否在服务白名单里

	if !utils.InListByRegex(c.WhiteList, r.URL.Path) {
		// 如果不在白名单里，则需要先鉴权
		authAddr = etcd.GetServiceAddr(c.Etcd, constants.AUTH_API_ETCD_KEY)
		isAuth := auth(authAddr, w, r) // 鉴权, 也就是 request 中要携带 Authorization 值, Authorization 值是用户的 token, 在用户登录成功后颁发token
		if !isAuth {
			logx.Error("网关：鉴权失败")
			response.Response(w, nil, errors.New("鉴权失败"))
			return
		}
	}

	// 从 etcd 中获取微服务地址
	servAddr := etcd.GetServiceAddr(c.Etcd, serName) // 从 etcd 中获取服务地址 ip:port
	if servAddr == "" {
		logx.Error("网关: 获取服务地址失败")
		return
	}
	// 反向代理, 可以处理http, WebSocket,请求
	serverUrl, err := url.Parse(fmt.Sprintf("http://%s", servAddr))
	reverseProxy := httputil.NewSingleHostReverseProxy(serverUrl)
	reverseProxy.ServeHTTP(w, r)
	//if c.ReverseProxy {
	//	if c.UseCache {
	//		CacheProxyHandler(w, r, p.CacheBox, servAddr)
	//	} else {
	//		HttpProxyHandler(w, r, servAddr)
	//	}
	//}
}

//var configFile = flag.String("config", "config.yaml", "config file")

// 启动网关
func main() {
	//flag.Parse()
	// 加载配置文件
	cfYamlByte, err := utils.GetServiceConfigYamlByte("server", "gateway")
	if err != nil {
		//logx.Errorf("获取Nacos配置失败: %v", err)
		fmt.Printf("获取Nacos配置失败: %v\n", err)
		return
	}
	err = conf.LoadFromYamlBytes(cfYamlByte, &c)
	if err != nil {
		//logx.Errorf("load config error: %v", err)
		fmt.Printf("load config error: %v\n", err)
		return

	}
	err = logx.SetUp(c.Log)
	if err != nil {
		logx.Error("网关: 设置日志失败")
		return
	}
	//启动服务
	logx.Infof("网关启动成功, 监听地址: %s", c.Addr)
	err = http.ListenAndServe(c.Addr, Proxy{CacheBox: cache.NewCacheBox(c.Redis.Addr, c.Redis.Password)})
	if err != nil {
		logx.Error("网关启动失败!")
		panic(err)
	}
}
