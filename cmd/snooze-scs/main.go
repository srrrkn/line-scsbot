package main

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"os"
	"io/ioutil"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
	"github.com/joho/godotenv"
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
	// メッセージjsonの読み込み
	raw, err := ioutil.ReadFile("/go/cmd/snooze-scs/template.json")
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
		Addr:      "db:3306",
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
	// DBからリマインド対象取得
	rows, err := db.Query(`
			select 
			e.group_id
			from notif_event e
			join line_group g on e.group_id = g.group_id and g.invalid = 0
			join scheduling s on g.id = s.line_group_id and s.invalid = 0
			where e.invalid = 0 
			and DATE_ADD(e.last_notified_at, INTERVAL s.snooze_interval_minutes MINUTE) < now()
			and DATE_ADD(e.last_notified_at, INTERVAL s.snooze_limit_minutes MINUTE) > now()
			and replyed_at is null
			group by e.group_id;
		`)
	if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
	// 対象グループ全てに通知送信
	for rows.Next() {
        var group_id string
        rows.Scan(&group_id)
		fmt.Println(group_id)
		// グループへの安否確認送信
		contents, err := linebot.UnmarshalFlexMessageJSON([]byte(raw))
		resp, err := bot.PushMessage(
			group_id,
			linebot.NewFlexMessage("【再送】安否確認への回答をお願いします", contents),
		).Do()
		fmt.Println(resp)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		// 最終通知時刻を更新
		result, err := db.Exec(
			`update notif_event set last_notified_at = now() where group_id = ? and replyed_at is null and invalid = 0`,
			group_id,
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(rowsAffected)
	}
}