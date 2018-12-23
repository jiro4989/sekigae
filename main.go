package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// Config は設定ファイル構造体
type Config struct {
	Sheets [][]bool `json:"sheets"`
	Ids    []string `json:"ids"`
}

var (
	configPaths = []string{
		".sekigae.json",
		HomeDir() + "/.config/sekigae/config.json",
	}
	sheets = [][]bool{
		[]bool{false, false, true, true, false, false},
		[]bool{true, true, true, true, true, true},
		[]bool{true, true, true, true, true, true},
		[]bool{true, true, true, true, true, true},
	}
	ids = []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22",
	}
)

func init() {
	// 設定ファイルが存在したら読み込み
	for _, p := range configPaths {
		if Exists(p) {
			b, err := ioutil.ReadFile(p)
			if err != nil {
				panic(err)
			}
			var config Config
			if err := json.Unmarshal(b, &config); err != nil {
				fmt.Fprintln(os.Stderr, "[err]設定ファイル("+p+")の書式が不正です。")
				panic(err)
			}
			sheets = config.Sheets
			ids = config.Ids
			break
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	Shuffle(ids)

	var (
		max    = MaxLen(ids)
		maxStr = fmt.Sprintf("%d", max)
		padFmt = "| %" + maxStr + "s "
	)

	newSheets := MakeSheets(sheets, ids)
	for _, sheet := range newSheets {
		var line string
		for _, s := range sheet {
			line += fmt.Sprintf(padFmt, s)
		}
		fmt.Println(line + "|")
	}
}

// MakeSheets は座席リストにIDを割り振って返す。
// sheetsのtrueの箇所が着席可能な箇所を指す。
// idsを着席可能な箇所に順に割り振る。
func MakeSheets(sheets [][]bool, ids []string) (newSheets [][]string) {
	var i int
	for _, sheet := range sheets {
		var lineSheet []string
		for _, isAvailable := range sheet {
			if isAvailable && i < len(ids) {
				lineSheet = append(lineSheet, ids[i])
				i++
				continue
			}
			lineSheet = append(lineSheet, "")
		}
		newSheets = append(newSheets, lineSheet)
	}
	return
}
