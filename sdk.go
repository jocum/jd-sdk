package sdk

import (
	"time"
	"sync"
)

var (
	inited 	bool
	sdk 	*JdSdk
	jdmu 	sync.Mutex	
)

/*
	jd 联盟 sdk 
*/
type JdSdk struct {
	appKey 	string
	secretKey 	string 		// 加密sign 所需要的字符串
	accessToken string
	version 	string
}


/*
	@Description 初始化sdk 结构体  
	@Params 		
	@Return 	*JdSdk
	@Author  cwy
*/
func NewJdSdk(appKey , secretKey , accessToken , version string) *JdSdk{
	//	读写锁定
	jdmu.Lock()
	defer jdmu.Unlock()
	if inited {
		return sdk
	}
	sdk = &JdSdk{
		appKey : appKey,
		secretKey : secretKey,
		accessToken : accessToken,
		version : version,
	}
	inited = true
	return sdk
}



/*
	@Description 	设置access_token  采用OAuth授权方式为必填参数
*/
func (j *JdSdk) SetAccessToken(accessToken string) *JdSdk {
	j.accessToken = accessToken
	return j
}

/*
	@Ddescription 	发送请求
	@Params 
	@Return 	res 	[]byte
 							err		error
	@Author cwy
*/
func (j *JdSdk) Send(r Request) ([]byte,error) {
	//	设置基础base 参数
	r.SetBase("app_key",j.appKey).
		SetBase("access_token",j.accessToken).
		SetBase("format","json").
		SetBase("v",j.version).
		SetBase("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	//	json 化 req 参数  =>  签名sign  =>  生成请求url
	url := r.reqMapToJson().sign(j.secretKey).generateUrl()
	//	请求jd 接口 返回数据
	return Get(url)
}