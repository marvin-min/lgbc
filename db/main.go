package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/boxroot120new?charset=utf8")
	checkErr(err)
	//addGles(db)
	//queryAllRootFiles(db)
	//updateGles(db)
	deleteGles(db)
	db.Close()
}

func deleteGles(db *sql.DB)  {
	stmt, err := db.Prepare("delete from gles where id = ?")
	checkErr(err)
	res, err:= stmt.Exec(2)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}

func updateGles(db *sql.DB) {
	stmt,err := db.Prepare("update gles set version = ? where id = ?")
	checkErr(err)
	res,err := stmt.Exec("3.0",2)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

}

func addGles(db *sql.DB) {
	stmt, err := db.Prepare("INSERT gles set extensions = ?, version=?, renderer=?,device_id=?,created_at=?,updated_at=?")
	checkErr(err)
	res, err := stmt.Exec("扩展", "2.0", "手动渲染", 13, time.Now(), time.Now())
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func queryAllRootFiles(db *sql.DB) {
	rows, err := db.Query("select id,file,source_name,created_at from root_files")
	checkErr(err)
	for rows.Next() {
		var id int
		var file string
		var source_name string
		var created_at string

		err = rows.Scan(&id, &file, &source_name, &created_at)
		checkErr(err)
		fmt.Println(id, file, source_name, created_at)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}