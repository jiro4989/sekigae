package main

import (
	"math/rand"
	"os"
)

// HomeDir はホームディレクトリを取得する。
// HOMEがあればHOMEを返し、USERPROFILEがあればUSERPROFILEを返す。
// 両方なければエラーでアプリが死ぬ
func HomeDir() string {
	home := os.Getenv("HOME")
	if home == "" {
		home := os.Getenv("USERPROFILE")
		if home == "" {
			panic("環境変数が定義されてません。環境変数を確認してください。")
		}
		return home
	}
	return home
}

// Shuffle は配列をシャッフルする。
// シャッフルする際の欄数値は関数内ではセットしないので
// 起動する都度異なる並びにシャッフルしたい場合は、以下のコードを呼ぶ。
//
//     rand.Seed(time.Now().UnixNano())
//     Shuffle(array)
func Shuffle(data []string) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

// MaxLen は文字列スライスのもっとも長い値を返す。
func MaxLen(ns []string) (max int) {
	for _, n := range ns {
		l := len(n)
		if max < l {
			max = l
		}
	}
	return
}

// Exists はファイルの有無を返す。
func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
