package sdk

import (
	"testing"
	"fmt"
)

/*
	测试调用sdk 完成请求
*/
func TestJdSdk(t *testing.T) {
	sdk := NewJdSdk("appKey","secretKey","accessToken","version")
	r := GetRequest()
	r.SetBase("method","jd.union.open.category.goods.get").SetReq("parentId",0).SetReq("grade",0)
	res,err := sdk.Send(r)
	if err != nil {
		fmt.Println("get res err",err)		
	}
	fmt.Println("res json str",string(res))
}