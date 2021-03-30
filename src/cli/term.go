package cli

import (
	"fmt"
	"runtime"
	"strings"
)

var Reset  = "\033[0m"
var Red    = "\033[31m"
var LRed = "\033[91m"
var Green  = "\033[32m"
var Yellow = "\033[33m"
var Blue   = "\033[34m"
var Purple = "\033[35m"
var Cyan   = "\033[36m"
var Gray   = "\033[37m"
var White  = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		Reset  = ""
		Red    = ""
		LRed   = ""
		Green  = ""
		Yellow = ""
		Blue   = ""
		Purple = ""
		Cyan   = ""
		Gray   = ""
		White  = ""
	}
}

func FString (text string) string {
	text = strings.Replace(text, "{reset}", Reset, -1)
	text = strings.Replace(text, "{red}", Red, -1)
	text = strings.Replace(text, "{l_red}", LRed, -1)
	text = strings.Replace(text, "{green}", Green, -1)
	text = strings.Replace(text, "{yellow}", Yellow, -1)
	text = strings.Replace(text, "{blue}", Blue, -1)
	text = strings.Replace(text, "{purple}", Purple, -1)
	text = strings.Replace(text, "{cyan}", Cyan, -1)
	text = strings.Replace(text, "{gray}", Gray, -1)
	text = strings.Replace(text, "{white}", White, -1)
	
	return text
}

func Println (text string) {
	fmt.Println(FString(text), Reset)
}