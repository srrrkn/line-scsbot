package main

import (
        "fmt"
        "log"
        "net/http"
		"encoding/json"
		"database/sql"
		"github.com/go-sql-driver/mysql"
		"time"
		"os"
		"github.com/joho/godotenv"
)

type Reply struct {
	UserId  string `json:"events[0].source.userId"`
	GroupId string `json:"events[0].source.groupId"`
}

func reflectReply(w http.ResponseWriter, r *http.Request){
	// // request bodyをそのままresponseとして返す
	// len := r.ContentLength
	// body := make([]byte, len)
	// r.Body.Read(body)
	// fmt.Println(string(body))
	// fmt.Fprintln(w, string(body))
	// グループID、ユーザーID取得
	var reply Reply
	err := json.NewDecoder(r.Body).Decode(&reply)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(reply)
	// // グループID、ユーザーID両方存在している場合のみ実行
	// if reply.GroupId == "" || reply.UserId == "" {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// // ロケーション設定
	// jst, err := time.LoadLocation("Asia/Tokyo")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }
	// // DB接続設定
	// c := mysql.Config{
	// 	DBName:    "scsbot",
	// 	User:      os.Getenv("MYSQL_USER"),
	// 	Passwd:    os.Getenv("MYSQL_USER_PASSWORD"),
	// 	Addr:      "db:3306",
	// 	Net:       "tcp",
	// 	ParseTime: true,
	// 	Collation: "utf8mb4_unicode_ci",
	// 	Loc:       jst,
	// }
	// db, err := sql.Open("mysql", c.FormatDSN())
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }
	// fmt.Println("db connected!!")
	// // 通知イベントをDBに記録
	// result, err := db.Exec(
	// 	`update notif_event set replyed_at = now() where group_id = ? and ( target_user = ? or target_user is null or target_user = "" ) and invalid = 0;`,
	// 	reply.GroupId,
	// 	reply.UserId,
	// )
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(rowsAffected)
}

func handleRequests() {
	http.HandleFunc("/webhook", reflectReply)
	log.Fatal(http.ListenAndServeTLS(":443", "/ssl/letsencrypt-all.crt", "/ssl/letsencrypt.key", nil))
}

func main() {
	// .envの読み込み
	err := godotenv.Load("/.env")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	handleRequests()
}
