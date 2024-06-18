package userrepo

import (
	"fmt"

	"github.com/Ccmuyu/biz-demo/gomall/demo_thrift/biz/dal/model"
	"github.com/Ccmuyu/biz-demo/gomall/demo_thrift/biz/dal/mysql"
)

func TestData() {
	db := mysql.DB
	var err error

	err = db.Create(&model.User{Email: "xx@email.com", Name: "Jane"}).Error
	if err != nil {
		fmt.Printf("user 插入失败, err %v", err)
	}

	err = db.Model(&model.User{}).Where("email = ?", "xx@email.com").Update("name", "Jane_1").Error
	if err != nil {
		fmt.Printf("user update失败, err %v", err)
	}

	u := &model.User{}
	err = db.Model(&model.User{}).Where("email = ?", "xx@email.com").First(u).Error
	if err != nil {
		fmt.Printf("user 查询失败, err %v", err)
	}
	fmt.Printf("row: %+v", u)

	// 软删 deleteAt
	// db.Where("email = ?", "xx@email.com").Delete(&model.User{})

	// 物理删除
	db.Unscoped().Where("email = ?", "xx@email.com").Delete(&model.User{})
}
