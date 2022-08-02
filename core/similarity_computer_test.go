package core

import (
	"fmt"
	"github.com/xiiv13/address_match_recommend/models"
	"testing"
)

func TestFindsimilarAddress(t *testing.T) {
	// querys := FindsimilarAddress("北京海淀区丹棱街18号创富大厦1106", 5, true)
	// querys := FindsimilarAddress("江苏连云港赣榆区江苏省赣榆县青口镇黄海路56号电影公司宿舍楼北楼1-901", 5, true)
	// querys := FindsimilarAddress("江津连云港赣榆区江苏省赣榆县青口镇黄海路56号科技有限公司宿舍楼北楼", 5, true)
	//querys := FindsimilarAddress("山东省济南市章丘区山东省章丘区明水开发区环路海尔公司", 5, true)
	querys := FindsimilarAddress("湖北武汉汉阳区汉阳经济技术开发区车城东路901号", 5, true)
	for _, v := range querys.SimiDocs {
		fmt.Println(models.NewAddressPersister().LoadAddr(v.Doc.Id).AddressText)
		//fmt.Println(v.Doc.Id)
		//fmt.Println(v.MatchedTerms)
	}
}
