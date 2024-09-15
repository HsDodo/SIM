package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"flag"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"regexp"
	log "server/common/logger"
)

var nacosConfPath = flag.String("nacos_conf_path", "config/nacos.yaml", "配置文件")

type Nacos struct {
	IpAddr string `json:"IpAddr"`
	Port   uint64 `json:"Port"`
}

func EncryptPassword(password string) string {
	encryptPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatalf("密码加密失败: %v", err)
	}
	return string(encryptPwd)
}

func ComparePassword(encryptPwd, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPwd), []byte(password))
	return err
}

func GetNacosConfig() Nacos {
	//var cf = "nacos.yaml"
	var nacos Nacos
	viper.SetConfigFile(*nacosConfPath) // 指定配置文件
	err := viper.ReadInConfig()         // 读取配置文件
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	err = viper.Unmarshal(&nacos)
	if err != nil {
		return Nacos{}
	}
	return nacos
}

func GetNamingClient() (naming_client.INamingClient, error) {
	flag.Parse()
	nacos := GetNacosConfig()
	// 客户端配置
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogLevel("debug"),
	)
	// nacos服务器配置
	sc := []constant.ServerConfig{
		{
			IpAddr: nacos.IpAddr,
			//ContextPath: "/nacos",Nacos的ContextPath，默认/nacos，在2.0中不需要设置
			Port: nacos.Port,
		},
	}
	// 创建服务发现客户端的另一种方式 (推荐)
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}

func GetServiceConfigYamlByte(group, dataId string) ([]byte, error) {
	nacos := GetNacosConfig()
	// 客户端配置
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogLevel("debug"),
	)
	// nacos服务器配置
	sc := []constant.ServerConfig{
		{
			IpAddr: nacos.IpAddr,
			//ContextPath: "/nacos",Nacos的ContextPath，默认/nacos，在2.0中不需要设置
			Port: nacos.Port,
		},
	}
	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 用动态配置客户端获取配置, 服务发现客户端获取服务
	_conf, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return []byte(_conf), nil
}

// ReflectStructToMap 反射结构体到 map
func ReflectStructToMap(obj interface{}, tag string) map[string]any {
	maps := map[string]any{}

	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		t, ok := field.Tag.Lookup(tag)
		if !ok {
			continue
		}
		val := objValue.Field(i)
		if val.IsZero() {
			continue
		}
		if field.Type.Kind() == reflect.Ptr {
			if field.Type.Elem().Kind() == reflect.Struct {
				newMaps := ReflectStructToMap(val.Elem().Interface(), tag)
				maps[t] = newMaps
				continue
			}
		}
		maps[t] = val.Interface()
	}
	return maps
}
func ReflectMapToStruct(m map[string]any, s any) error { // s 必须是指针, 因为下面解引用了,不是指针的话会报错
	sType := reflect.TypeOf(s).Elem()
	sValue := reflect.ValueOf(s).Elem()
	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		mapField, ok := m[tag]
		if !ok {
			continue
		}
		fieldValue := sValue.Field(i)
		if fieldValue.Kind() == reflect.Struct {
			// 如果是结构体,则递归
			err := ReflectMapToStruct(mapField.(map[string]any), fieldValue.Addr().Interface())
			if err != nil {
				return err
			}
		} else if fieldValue.CanSet() {
			// 判断 两者值类型是否一样
			v := reflect.ValueOf(mapField)
			if v.Type().ConvertibleTo(fieldValue.Type()) {
				fieldValue.Set(v.Convert(fieldValue.Type()))
			} else {
				return errors.New("反射：转换类型不匹配")
			}
		}
	}
	return nil
}

func InList(list []string, key string) (ok bool) {
	for _, s := range list {
		if s == key {
			return true
		}
	}
	return false
}

func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func InListByRegex(list []string, key string) (ok bool) {
	for _, s := range list {
		regex, err := regexp.Compile(s)
		if err != nil {
			log.Error(err)
			return
		}
		if regex.MatchString(key) {
			return true
		}
	}
	return false
}

func InIDsList[T uint | uint32](list []T, id T) (ok bool) {
	for _, s := range list {
		if s == id {
			return true
		}
	}
	return false
}
