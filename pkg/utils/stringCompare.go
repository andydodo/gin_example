package utils

import (
	"fmt"
	"strings"
)

var str = `{"status":"success","code":200,"data":[{"title":"AI\u533b\u7597\u62a5\u544a\u4fc3\u9500","banner":"https:\/\/cdn.itjuzi.com\/images\/61f03d7573429ca5ab2968c99c1fa85d.jpg?imageView2\/0\/q\/100","url":"https:\/\/detail.youzan.com\/show\/goods?alias=2frllica94ywq&activity_alias=undefined","location":1},{"title":"6\u5e74\u62a5\u544a\u5408\u8f91","banner":"https:\/\/cdn.itjuzi.com\/images\/486f607ff37d1c2a70180c2bdd89216d.jpg?imageView2\/0\/q\/100","url":"https:\/\/mp.weixin.qq.com\/s\/-mrNTia9XmpyIH9ivloThA","location":2},{"title":"\u6392\u961fIPO","banner":"https:\/\/cdn.itjuzi.com\/images\/dc0267bc42769a7fe83975d70fa0ceaf.png?imageView2\/0\/q\/100","url":"https:\/\/www.itjuzi.com\/pre-ipo","location":2},{"title":"\u521b\u65b0\u51b2\u523a","banner":"https:\/\/cdn.itjuzi.com\/images\/92845a288bd30735be8ca9f976229c4d.jpg?imageView2\/0\/q\/100","url":"https:\/\/mp.weixin.qq.com\/s\/l8oL1JYG24fdByedvJ-mng","location":4}]} <nil>`

func JsonToMap(j string, expect string, flag int) bool {
	/**
	flag 1表示模糊匹配
	j 是输出数据
	expect是预期结果
	*/

	fmt.Println(str)

	return false

}

func JsonToMapStr(j string, expect string, flag int) bool {

	if flag == 1 {
		return strings.Contains(str, `{"status":"success","code":200`)

	}

	return false
}
