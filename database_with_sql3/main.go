package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DBConnection *sql.DB

// Game ゲームソフトクラス
type Game struct {
	Title string
	Maker string
	Price uint
}

func main() {
	// データベースへのアクセス開始
	DBConnection, _ := sql.Open("sqlite3", "./example.sql")

	// Openを呼び出す場合は必ず実行する
	defer DBConnection.Close()

	// テーブル作成コマンド
	cmd := `CREATE TABLE IF NOT EXISTS game(
				title STRING,
				maker STRING,
				price INT)`

	// コマンドを実行しつつ、エラーハンドリング
	_, err := DBConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// データを入れてみる - SQLインジェクションを防止するために、valuesはここで直打ちしないこと
	// cmd = `INSERT INTO game (title, maker, price) VALUES (?, ?, ?)`
	// _, err = DBConnection.Exec(cmd, "Pawapuro Go", "Gonami", 5800)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// DBのデータを読み出す
	cmd = "SELECT * from game"
	// Queryで得られる結果は必ず使用後にクローズすること
	rows, _ := DBConnection.Query(cmd)
	defer rows.Close()
	var gamesoft []Game

	// データのスキャン
	for rows.Next() {
		var g Game
		err := rows.Scan(&g.Title, &g.Maker, &g.Price)
		if err != nil {
			log.Println(err)
		}
		gamesoft = append(gamesoft, g)
	}

	for _, g := range gamesoft {
		fmt.Println(g.Title, g.Maker, g.Price)
	}
}
