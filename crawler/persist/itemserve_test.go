package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"ccmouse/crawler/model"
	"testing"
)

func TestItemSaver(t *testing.T) {
	user := model.User{
		Name: "cocoyo",  // 姓名
		Gender: "男",      // 性别
		Age: 18,            // 年龄
		Height: 180,         // 身高
		City: "深圳",        // 城市
		Income: "20000",      // 收入
		Marriage: "未婚",    // 婚姻状况
		Education: "本科",   // 学历
		Occupation: "程序员",  //职业
		Hokou: "广东",        // 户口 -> 籍贯
		Xinzuo: "天蝎座",      // 星座
	}

	id, err := save(user)
	if err != nil {
		t.Error(err)
	}

	client, _ := elastic.NewClient(elastic.SetSniff(false))

	response, _ := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())

	//response , err := client.Index().Index("dating_profile").Type("cocoyo").BodyJson(user).Do(context.Background())

	var responseUser model.User

	json.Unmarshal(*response.Source, &responseUser)

	t.Log(responseUser)
}
