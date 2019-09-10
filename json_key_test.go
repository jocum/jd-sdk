package sdk

import (
	"testing"
	"fmt"
)

func TestGetJsonKey(t *testing.T) {
	jsonkey := GetJsonKey("jd.union.open.goods.query")
	fmt.Println(jsonkey)
}