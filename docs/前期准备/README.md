

### 前期准备

#### 各服务地址

```yaml

网关: 8888 
API:
	- 认证服务: 9000   # auth.api 服务端口
    - 用户服务: 9001   # user.api 服务端口
	- 文件服务: 9002   # file.api 服务端口
	- 会话服务: 9003   # chat.api 服务端口
	- 群聊服务: 9004   # group.api 服务端口
	- 日志服务: 9005   # logs.api 服务端口
RPC:
	- user.rpc: 9100  # user.rpc 服务端口
	- chat.rpc: 9101  # chat.rpc 服务端口
	- file.rpc: 9102  # file.rpc 服务端口
	- group.rpc: 9103 # group.rpc 服务端口
	- logs.rpc: 9104  # logs.rpc 服务端口

```

