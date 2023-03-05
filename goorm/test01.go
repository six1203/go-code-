package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

//gorm.Model 定义
//type Model struct {
//	ID        uint `gorm:"primarykey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt DeletedAt `gorm:"index"`
//}

type Common struct {
	Id        int64     `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null;index:created_at"`
	UpdatedAt time.Time `gorm:"not null;index:updated_at"`
	IsDeleted bool      `gorm:"not null;default:false"`
}

type JSON json.RawMessage

func (j *JSON) Scan(val interface{}) error {
	//s := val.([]uint8)
	//ss := strings.Split(string(s), ",")
	//*r = ss
	//return nil

	bytes, ok := val.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", val))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

type Product struct {
	Common
	Code     string `gorm:"type:varchar(255);uniqueIndex"`
	CodeDesc string `gorm:"type:varchar(1000);default:''"`
	/* 在 GORM 中设置默认值为数字 0 时，它会将其转换为 SQL 中的数字 0，默认情况下会将其作为字符串插入到 DDL 语句中，因此你看到的是 '0'。
	要修改此行为，可以使用 GORM 模型标记 gorm:"default:0" 指定默认值。这将确保将数字 0 直接插入 DDL 语句中，而不是作为字符串。*/
	Price int  `gorm:"not null;default:0"`
	Tag   JSON `gorm:"type:json"`
}

// root用户名，没有密码，有密码的话就root:pwd,loc=Local 设置时区为当前时区
var mysqlUrl = "root@(127.0.0.1:3306)/dylan?charset=utf8&parseTime=True&loc=Local"

func (p *Product) MarshalJSON() ([]byte, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}
	type Alias Product
	return json.Marshal(&struct {
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		*Alias
	}{
		CreatedAt: p.CreatedAt.In(loc).Format("2006-01-02 15:04:05"),
		UpdatedAt: p.UpdatedAt.In(loc).Format("2006-01-02 15:04:05"),
		Alias:     (*Alias)(p),
	})
}

func main() {
	// 连接db的三种形式
	//db, err := gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{})
	/*等价于 gorm.Open("mysql", mysqlUrl)*/

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: mysqlUrl, // DSN data source name
		//DefaultStringSize:         256,   // string 类型字段的默认长度
		//DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		//SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 表名前缀，`Product` 的表名应该是 `t_products`， 默认不起用
			//TablePrefix: "",
			// 使用单数表名，启用该选项，默认启用，此时，`User` 的表名应该是 `t_user`
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed connect to database")
	}

	// 自动初始化表
	//db.AutoMigrate(&Product{})

	//product := Product{
	//	Common{
	//		CreatedAt: time.Now(),
	//		UpdatedAt: time.Now(),
	//		IsDeleted: false,
	//	},
	//	"123",
	//	"ceshi",
	//	123,
	//	JSON(`[]`),
	//}

	//tx := db.Create(&product)
	//fmt.Println(tx, product.Id, tx.Error, tx.Statement, tx.RowsAffected)

	var product Product
	db.First(&product, 3)
	fmt.Println(product.Tag, product.Price)
	//
	//db.First(&product, "code=?", "402")
	//db.First(&product, 1)
	//jsonData, err := json.Marshal(product)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(jsonData))
	//
	//fmt.Printf("%T, %v", product.Tag, product.Tag)
}
