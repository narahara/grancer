package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// PrettyJson 整形されたJSON文字列を返す
func PrettyJson(v interface{}) string {
	in, err := json.Marshal(v)
	if err == nil {
		var out bytes.Buffer
		err = json.Indent(&out, in, "", "  ")
		if err == nil {
			return out.String()
		}
	}
	return ""
}

// PrintJson 整形されたJSON文字列を標準出力に表示する
func PrintJson(v interface{}) {
	in, err := json.Marshal(v)
	if err == nil {
		var out bytes.Buffer
		err = json.Indent(&out, in, "", "  ")
		if err == nil {
			fmt.Println(out.String())
		}
	} else {
		fmt.Println(err)
	}
}
