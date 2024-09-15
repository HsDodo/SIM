### 认证服务

> 认证服务：
>
> - 用户登录  `login`
> - 用户注册 `待写`
> - 第三方用户登录 `OpenLogin`
> - 鉴权 `Authentication`

#### 1. 编写认证服务配置文件

```yaml
# 认证服务配置
Host: "0.0.0.0"
Port: 9700
Name: "auth"
# etcd 服务器配置
Etcd:
  Hosts:
    - "localhost:2379"
# 用户微服务配置
UserRpc:
  Etcd:
    Hosts:
      - "localhost:2379"
    Key: user.rpc
# Mysql 配置
Mysql:
  DataSource: "root:root@tcp(localhost:3306)/sim_server_db?charset=utf8&parseTime=True&loc=Local"
# Jwt 鉴权配置
Auth:
  AccessSecret: "Hsen1015"
  AccessExpire: 24 # 小时 
# Logx 配置
Log:
  ServiceName: auth
  Encoding: plain
  Stat: false
  TimeFormat: 2006-01-02 15:04:05
  Colors:
    debug: "\x1b[36m"
    info: "\x1b[32m" 
    warn: "\x1b[33m"
    error: "\x1b[31m"
 # Redis 配置
Redis:
  Addr: "127.0.0.1:6379"
  Pwd: "5566"
  DB: 0
# 第三方登录配置
OpenLoginList:
  Wechat:
    Appid: "wx93daaac2d32872ce"
    AppSecret: "b3cd8687029de78b9952c44f9ddb029c"  # 微信开放平台的appSecret
    Icon: "https://open.weixin.qq.com/connect/qrconnect?appid=wx39c379788eb1286a&scope=snsapi_login&redirect_uri=http%3A%2F%2Fmp.weixin.qq.com%2Fdebug%2Fcgi-bin%2Fsandbox%3Ft%3Dsandbox%2Flogin"
    RedirectURI: "http://hsenn.nat300.top/api/auth/open_login"  # 微信回调地址,本地服务器的地址,（这里用来内网穿透，这个域名会变的）
    Endpoint:
      TokenURL: "https://api.weixin.qq.com/sns/oauth2/access_token" # 微开放平台的tokenURL,获取access_token ,首先需要得到 code
      AuthURL: "https://open.weixin.qq.com/connect/oauth2/authorize" # 微信开放平台的authURL,获取code
      UserInfoURL: "https://api.weixin.qq.com/sns/userinfo" # 微信开放平台的userInfoURL,获取用户信息
```

#### 2. 编写 API 接口 (待完善)

> 按照 go-zero API 文件编写规范来写

```go
syntax = "v1"

info (
	title:   "认证"
	desc:    "鉴权认证"
	author:  "sen"
	version: "v0.01"
)

type LoginRequest {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResponse {
	Token string `json:"token"`
}

type OpenLoginInfo {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Href string `json:"href"` // 跳转地址
}

type OpenLoginInfoResponse {
	RedirectURI string `json:"redirect_uri"`
}

type AuthenticationReponse {
	UserId uint `json:"userId"`
	NickName string `json:"nickName"`
}

@server (
	prefix: /v1
)
service auth {
	@handler login
	post /api/auth/login (LoginRequest) returns (LoginResponse)

	@handler authentication
	post /api/auth/authentication (string) returns (AuthenticationReponse) //认证接口

	@handler logout
	post /api/auth/logout (string) returns (string) //退出登录

	@handler openLoginInfo
	get /api/auth/open_login_info returns (OpenLoginInfoResponse) //第三方登录信息

	@handler openLogin
	get /api/auth/open_login returns (LoginResponse) //第三方登录
}

```

> 使用 `goctl api go --api auth.api --dir .` 来生成 API 服务文件

 **目录树：**

```shell
└─api
    ├─etc                # 配置文件
    ├─internal           # 内部服务
    │  ├─config		  # config.go 配置文件, 用于加载etc里面的配置
    │  ├─handler		 # Handler 处理器
    │  ├─logic		   # 服务逻辑
    │  ├─svc 			# 用于加载服务环境, 配置上下文 Context
    │  └─types		   # 类型结构体放置处
    └─log
    └─auth.go			# 服务启动入口
    └─auth.api 		  # 服务API文件
```

