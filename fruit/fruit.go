package fruit

import (
	"log"
	"io/ioutil"
	"encoding/json"
)

type Product struct {
	Name	string
}

type Data struct {
	Products	[]*Product
}

var dataFile = "fruit/data.json" // go runを実行した場所から対象のjsonファイルへのパス

func init() {
	log.Println("include fruit.")
}

func GetList() Data {
	// ioutilパッケージのReadFile関数にファイルパスを渡してファイルを読む
	// 戻り値が「ファイルの中身」と「エラー」の２つ。カンマ区切りで変数を並べて受け取れる
	bytes, err := ioutil.ReadFile(dataFile)
	// err変数が空でなかった場合
	if err != nil {
		// ログを出力する
		log.Fatalf("%s", err)
	}
	// 空のlist変数を用意する
	var list []*Product
	// bytesに格納されたJSON文字列を、構造体へ変換しlistに格納
	// err変数が空でなかった場合
	if err := json.Unmarshal(bytes, &list); err != nil {
		// ログを出力する
		log.Fatal(err)
	}
	// Dataの構造体のProductsにjsonファイルから取得したデータセットして返す
	return Data{Products: list}
}