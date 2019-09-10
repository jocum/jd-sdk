package sdk

import (
	"fmt"
	"encoding/json"
	"net/url"
	"strings"
	"sort"
	"crypto/md5"
	"encoding/hex"
)


/*
	@Description jdsdk 参数请求对象
*/
type Request struct {
	Base 	map[string]string	// 公共参数
	Req 	map[string]interface{} 		//	接口参数
	Host 	string 		//调用域名
}

/*
	@Description 初始化 Request 对象
*/
func GetRequest() Request{
	return Request{
		Base : make(map[string]string),
		Req : make(map[string]interface{}),
		Host : JD_API_URL,
	}
}

/*
	@Description 	设置请求域名  如果不需要默认的域名的话
	@Params 	api 	string
	@Return 	Request
	@Author 	cwy
*/
func (r Request) SetHost(host string)  Request{
	r.Host = host
	return r
}

/*
	@Description 	设置 公用参数方法 
	@Params 	key	string
							  val 	interface{}
	@Retrurn 	*JdSdk
	@Author 	cwy
*/
func (r Request) SetBase(key string , val string) Request {
	r.Base[key] = val
	return r
}

/*
	@Description 	设置 请求参数
	@Params 	key 	string
							val 	interface{}
	@Return		*JdSdk
	@Author 	cwy
*/
func (r Request) SetReq(key string ,val interface{}) Request {
	r.Req[key] = val
	return r
}


/*
	@Description 	请求参数 生成 请求json  
	@Params 		jsonKey 	string 		jd的请求参数 不同的接口 最上层的key 值都不一样 传入一个字符串来完成参数json化
在jd 联盟接口 https://union.jd.com/openplatform/api/628  业务参数 最上层 名称的 字符串  
例:   jd.union.open.goods.query  接口的 jsonKey  就是   goodsReqDTO
	@Return 	*JdSdk
	@Author 	cwy
*/
func (r Request) reqMapToJson() Request {
	var 	jsonMap 	map[string]interface{}
	//	判断是否有最上级 key
	topKey := GetJsonKey(r.Base["method"])
	if 	topKey != "" {
		topKeyMap := make(map[string]interface{})	
		topKeyMap[topKey]	=	r.Req
		jsonMap = topKeyMap
	}else {
		jsonMap = r.Req
	}
	// 将请求参数转化成json 
	jsonByte , err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println("json encoding err ",err)
	}
	//	jd 接口非常怪异  不懂这个点 为什么是  360buy 但是测试工具上面是这一个 也只能这么来  json 不字符做url 转化

	r.Base["360buy_param_json"] = string(jsonByte)
	return r
}


/*
	@Description 	生成请求url 
	@Params 	req 	map[string]string
							 *Client.Url
	@Return 	url 	string
	@Author 	cwy
*/
func (r Request) generateUrl() string{
	var data string
	//	循环拼接参数
	for k,v := range r.Base {
		data += fmt.Sprintf("%s=%s&",k,url.QueryEscape(v))
	}
	//	 尾部出去 &
	trimData := strings.TrimRight(data,"&")
	//	拼接url
	urlStr  :=   fmt.Sprintf("%s?%s",r.Host,trimData)
	//	url 编码
	return urlStr
}


/*
	@Description 	生成签名
	@Params 	
	@Return 	
	@Author cwy
*/
func (r Request) sign(secretKey string) Request{
	// 按 key排序
	var keys []string
	for k2 := range r.Base {
		keys = append(keys,k2)
	}
	sort.Strings(keys)
	//拼接
	var str string
	for _,v1 := range keys {
		str += fmt.Sprintf("%s%s",v1, r.Base[v1])
	}
	str = secretKey+str+secretKey
	//	写入sign
	r.Base["sign"] = strings.ToUpper(MD5(str))
	return r 
}

// 生成32位MD5
func MD5(text string) string{
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
 }