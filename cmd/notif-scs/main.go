package main

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"os"
	"io/ioutil"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
	"flag"
	"github.com/joho/godotenv"
	"strings"
)

func main() {
	// .envの読み込み
	err := godotenv.Load("/.env")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// LINE BotのChannel SecretとChannel Access Tokenを設定
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// コマンドライン引数取得
	str_user_ids := flag.String("users", "", "User Id")
	group_id := flag.String("group", "", "Group Id")
	flag.Parse()
	user_ids := strings.Split(*str_user_ids, ",")
	// メッセージjsonの読み込み
	raw, err := ioutil.ReadFile("/go/cmd/notif-scs/template.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// グループへの安否確認送信
	contents, err := linebot.UnmarshalFlexMessageJSON([]byte(raw))
	resp, err := bot.PushMessage(
		*group_id,
		linebot.NewFlexMessage("安否確認への回答をお願いします", contents),
	).Do()
	fmt.Println(resp)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// ロケーション設定
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
	// DB接続設定
	c := mysql.Config{
		DBName:    "scsbot",
		User:      os.Getenv("MYSQL_USER"),
		Passwd:    os.Getenv("MYSQL_USER_PASSWORD"),
		Addr:      "line-scs-bot-db:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
	fmt.Println("db connected!!")
	// 通知イベントをDBに記録
	for _, user_id := range user_ids {
		result, err := db.Exec(`insert into notif_event(group_id, target_user) values(?, ?)`, *group_id, user_id)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(rowsAffected)
	}
}

