package loger

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"time"
)

var PATH_LOGS = "./request-logs"

func IsResourceExist(fileName *string) bool {
	if _, err := os.Stat(*fileName); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
func InsertNewLog(fileName string, url string) {
	if !IsResourceExist(&fileName) {
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Create Error", err.Error())
		}
		file.Close()
	}
	data, errRead := ioutil.ReadFile(fileName)
	if errRead == nil {
		dataResult := string(data)
		if len(dataResult) > 0 {
			dataResult = dataResult + "\n"
		}
		dataResult = dataResult + "Date: " + time.Now().String() + " | Link: " + url

		ioutil.WriteFile(fileName, []byte(dataResult), 0644)
	}
}

func WriteLogFiles(url string, bodyText string) int {
	if !IsResourceExist(&PATH_LOGS) {
		os.MkdirAll(PATH_LOGS, os.ModePerm)
	}
	re := regexp.MustCompile("(http:|/)")
	fileName := PATH_LOGS + "/file_" + re.ReplaceAllString(url, "") + "_" + strconv.FormatInt(time.Now().UnixNano(), 10) + "_log"

	if !IsResourceExist(&fileName) {
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Create Error", err.Error())
			return -2
		}
		defer file.Close()
		len, err2 := file.WriteString(bodyText)
		if err2 != nil {
			fmt.Println("Write error", err2.Error())
			return -3
		}
		fmt.Println("Seccessful!!")
		return len
	}
	return -1
}

// func main() {
// 	fileName := "prueba.txt"
// 	if len(os.Args) >= 2 {
// 		fileName = os.Args[1]
// 	}
// 	bodyText := "el texto que se va a esctibir es este texto que se esta pasando por parametro"
// 	result := WriteFile(fileName, bodyText)
// 	fmt.Println(result)
// }
