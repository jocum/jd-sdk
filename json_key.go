package sdk

import (
	"sync"
)

var (
	mapInited 	bool
	jsonKeyMap 	map[string]string
	mapMu 	 sync.Mutex
)
/*
	jd 联盟接口  不同接口请求参数的最上上层key 略有不同 用方法字符串 获取 这个key 空则没有
*/
func GetJsonKey(metch string) string {
	initedJsonMap()
	return jsonKeyMap[metch]
}


func initedJsonMap() {
	mapMu.Lock()
	defer mapMu.Unlock()
	if mapInited {
		return
	}
	jsonKeyMap = map[string]string {
		"jd.union.open.goods.query" : "goodsReqDTO",
		"jd.union.open.order.query" : "orderReq",
		"jd.union.open.order.bonus.query" : "orderReq",
		"jd.union.open.goods.jingfen.query": "goodsReq",
		"jd.union.open.goods.bigfield.query" : "goodsReq",
		"jd.union.open.goods.link.query" : "goodsReq",
		"jd.union.open.coupon.query" : "",
		"jd.union.open.category.goods.get" : "req",
		"jd.union.open.goods.stuprice.query" : "goodsReq",
		"jd.union.open.goods.seckill.query" : "goodsReq",
		"jd.union.open.goods.promotiongoodsinfo.query" : "",
		"jd.union.open.position.create" : "positionReq",
		"jd.union.open.promotion.applet.get" : "promotionCodeReq",
		"jd.union.open.promotion.bysubunionid.get" : "promotionCodeReq",
		"jd.union.open.promotion.byunionid.get" : "promotionCodeReq",
		"jd.union.open.promotion.common.get" : "promotionCodeReq",
		"jd.union.open.user.pid.get" : "pidReq",
		"jd.union.open.position.query" : "positionReq",
		"jd.union.open.coupon.importation" : "couponReq",
	}
	mapInited = true
	return 
}