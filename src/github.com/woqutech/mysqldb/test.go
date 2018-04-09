package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/golang/glog"
)

const querySql = "select * from demo order by time desc limit 1"
const updateSql = "update demo set num=?,time=? where num=?"
const connSleepTime = 3 //重连延迟
const operateSleepTime = 3 // 操作延迟
const insertSql = "insert into demo(num, time) value(?, ?)"
func main() {
	//定时任务
	//for {
	//	db, err := initDB()
	//	for err!=nil {
	//		time.Sleep(3 * time.Second) // 连接数据库出错则延迟3s继续请求连接
	//		db, err = initDB()
	//	}
	//	oldN, oldD, err := query(db, querySql)
	//	update(db, updateSql, oldN, oldD, oldN+1)
	//	time.Sleep(3 * time.Second)
	//}
//begin:
//	db, err := initDB()
//	for err!=nil{
//		time.Sleep(connSleepTime * time.Second) // 连接数据库出错则延迟3s继续请求连接
//		db, err = initDB()
//	}
//	for {
//		oldN, oldD, err := query(db, querySql)
//		if err!=nil && strings.Contains(err.Error(), " No connection"){
//			db.Close()
//			goto begin
//		}
//		err = update(db, updateSql, oldN, oldD, oldN+1)
//		if err!=nil && strings.Contains(err.Error(), " No connection"){
//			db.Close()
//			goto begin
//		}
//		time.Sleep(operateSleepTime * time.Second)
//	}
	db, _ := initDB()
	for {
		oldN, _, _ := query(db, querySql)

		insert(db, insertSql, oldN+1)

		time.Sleep(operateSleepTime * time.Second)

	}
}

/**
连接数据库
*/
func initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Liuxin@950326@tcp(10.10.20.4:3306)/test")
	if err != nil {
		glog.Info(err)
		return db, err
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		glog.Info(err)
		return db, err
	}
	return db, err
}

/**
查询
*/
func query(db *sql.DB, querySql string) (int, string, error) {
	var oldN int
	var oldD string
	err := db.QueryRow(querySql).Scan(&oldN, &oldD)
	if err != nil {
		glog.Info(err)
		return 0, "", err
	}
	return oldN, oldD, nil
}

/**
更新
*/
func update(db *sql.DB, updateSql string, oldN int, oldD string, newN int) error {
	stmt, err := db.Prepare(updateSql)
	if err != nil {
		glog.Info(err)
		return err
	}
	newD := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec(newN, newD, oldN)
	if err != nil {
		glog.Info(err)
		return err
	}
	affectRowCnt, err := res.RowsAffected()
	if err != nil {
		glog.Info(err)
		return err
	}
	glog.Info("----->UPDATE", " affectRowCnt:", affectRowCnt, " oldN:", oldN, " oldD", oldD, " newN", newN, " newD:", newD)
	return nil
}


func insert(db *sql.DB, insertSql string, newN int) error{
	stmt, err := db.Prepare(insertSql)
	if err != nil {
		glog.Info(err)
		return err
	}
	newD := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec(newN, newD)
	if err != nil {
		glog.Info(err)
		return err
	}
	affectRowCnt, err := res.RowsAffected()
	if err != nil {
		glog.Info(err)
		return err
	}
	glog.Info("----->INSERT", " affectRowCnt:", affectRowCnt, " newN", newN, " newD:", newD)
	return nil
}