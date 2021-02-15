package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个结构体, 需要大写开头哦, 字段名也需要大写开头哦, 否则json模块会识别不了
// 结构体成员仅大写开头外界才能访问
type Student struct {
	Name   string `json:"name"`
	Age    string `json:"age"`
	School string `json:"school"`
	Sex    string `json:"sex"`
}

// 一如既往的main方法
func QueryStudent() {
	// 格式有点怪, @tcp 是指网络协议(难道支持udp?), 然后是域名和端口
	db, e := sql.Open("mysql", "root:root@tcp(localhost:3306)/school?charset=utf8")
	if e != nil { //如果连接出错,e将不是nil的
		fmt.Println("ERROR?")
		return
	}
	// 提醒一句, 运行到这里, 并不代表数据库连接是完全OK的, 因为发送第一条SQL才会校验密码 汗~!
	_, e2 := db.Query("select 1") //生产环境去掉这句，不然会有内存泄漏
	if e2 == nil {
		fmt.Println("DB OK")
		rows, e := db.Query("select name,age,school,sex from student")
		if e != nil {
			fmt.Print("query error!!%v\n", e)
			return
		}
		if rows == nil {
			fmt.Println("Rows is nil")
			return
		}
		for rows.Next() { //跟java的ResultSet一样,需要先next读取
			student := new(Student)
			// rows貌似只支持Scan方法 继续汗~! 当然,可以通过GetColumns()来得到字段顺序
			row_err := rows.Scan(&student.Name, &student.Age, &student.School, &student.Sex)
			if row_err != nil {
				fmt.Println("Row error!!")
				return
			}
			b, _ := json.Marshal(student)
			fmt.Println(string(b)) // 这里没有判断错误, 呵呵, 一般都不会有错吧
		}
		fmt.Println("Done")
	} else {
		fmt.Println(e)
	}
}
