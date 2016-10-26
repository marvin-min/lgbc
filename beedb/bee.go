package main

import (
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Gles struct {
	Id int `orm:"auto"`
	Extensions string
	Version string
	Renderer string
	Vendor string
	DeviceId int
	CreatedAt time.Time
	UpdatedAt time.Time
}
func init() {
	// register model
	orm.RegisterModel(new(Gles))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:@tcp(localhost:3306)/boxroot120new?charset=utf8", 30)
}
func main() {
	orm := orm.NewOrm()
	gle :=new(Gles)
	gle.DeviceId = 1
	gle.Extensions="adf"
	gle.Renderer="self"
	gle.Vendor="iob"
	gle.CreatedAt = time.Now()
	gle.UpdatedAt = time.Now()
	rs,err := orm.Insert(gle)
	checkErr(err)
	fmt.Println(rs)

	orm.Delete(gle)
	//orm.Commit()

}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}