package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dataSourceName = "root:root@/school?charset=utf8"
)

func query() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/school")
	check(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM student")

	for rows.Next() {
		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		//将数据保存到 record 字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
	rows.Close()
}

func query2() {
	fmt.Println("Query2")
	db, err := sql.Open("mysql", dataSourceName)
	check(err)
	defer db.Close()

	rows, err := db.Query("SELECT name,age,school,sex FROM student")
	check(err)

	for rows.Next() {
		var name string
		var age string
		var school string
		var sex string
		//注意这里的Scan括号中的参数顺序，和 SELECT 的字段顺序要保持一致。
		if err := rows.Scan(&name, &age, &school, &sex); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %s %s %s\n", name, age, school, sex)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
}

func insert() {
	db, err := sql.Open("mysql", dataSourceName)
	check(err)
	defer db.Close()

	stmt, err := db.Prepare(`INSERT student (name, age, school, sex) VALUES (?, ?, ?, ?)`)
	check(err)

	res, err := stmt.Exec("乐惠2", 4, nil, "保密")
	check(err)

	id, err := res.LastInsertId()
	check(err)

	fmt.Println(id)
	stmt.Close()
}

func update() {
	db, err := sql.Open("mysql", dataSourceName)
	check(err)
	defer db.Close()

	stmt, err := db.Prepare("UPDATE student set age=?, school=?, sex=? WHERE name=?")
	check(err)

	res, err := stmt.Exec(5, "未知", "男", "乐惠1")
	check(err)

	num, err := res.RowsAffected()
	check(err)

	fmt.Println(num)
	stmt.Close()
}

func remove() {
	db, err := sql.Open("mysql", dataSourceName)
	check(err)
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM student WHERE name=?")
	check(err)

	res, err := stmt.Exec("乐惠2")
	check(err)

	num, err := res.RowsAffected()
	check(err)

	fmt.Println(num)
	stmt.Close()
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
