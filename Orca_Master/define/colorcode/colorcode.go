package colorcode

import (
	"fmt"
	"time"
)

const (
	END          = "\033[0m"
	COLOR_GREY   = "\033[1;30m"
	COLOR_RED    = "\033[1;31m"
	COLOR_GREEN  = "\033[1;32m"
	COLOR_YELLOW = "\033[1;33m"
	COLOR_BLUE   = "\033[1;34m"
	COLOR_PURPLE = "\033[1;35m"
	COLOR_CYAN   = "\033[1;36m"
	COLOR_SHINY  = "\033[6;20m"
)

var (
	SIGN_WARNING = COLOR_YELLOW + "[!]" + END
	SIGN_ERROR   = COLOR_RED + "[!]" + END
	SIGN_FAIL    = COLOR_RED + "[-]" + END
	SIGN_SUCCESS = COLOR_GREEN + "[+]" + END
	SIGN_NOTICE  = COLOR_BLUE + "[*]" + END
	SIGN_QUEST   = COLOR_PURPLE + "[?]" + END
)

func PrintMessage(signStr string, message string) {
	currentTime := time.Now().Format("2006/01/02 15:04:05")
	timeSign := COLOR_GREY + currentTime + END
	fmt.Println(timeSign, signStr, message)
}

func Colorf(colorStr string, message string) string {
	return fmt.Sprintf("%s%s%s", colorStr, message, END)
}

func OutputMessage(signStr string, message string) string {
	var currentTime = time.Now().Format("2006/01/02 15:04:05")
	timeSign := COLOR_GREY + currentTime + END
	return fmt.Sprintf("%s %s %s\n", timeSign, signStr, message)
}
