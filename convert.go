package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tonyHuinker/ehop"
)

func askForInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	response, _ := reader.ReadString('\n')
	fmt.Println("\nThank You")
	return strings.TrimSpace(response)
}

func typeConverter(ehopType string) (string, error) {
	if ehopType == "app" {
		return "ApiObject", nil

	} else if ehopType == "s" {
		return "string", nil

	} else if ehopType == "dev" {
		return "ApiObject", nil

	} else if ehopType == "addr4" {
		return "ApiObject", nil

	} else if ehopType == "addr6" {
		return "ApiObject", nil

	} else if ehopType == "n" {
		return "int", nil

	} else if ehopType == "b" {
		return "bool", nil
	}
	return "", errors.New("not a valud Record Format")
}

func main() {
	answer := askForInput("Enter location of Record Format")
	recordName := askForInput("Enter the name you'd like to use for this struct")
	f, _ := ioutil.ReadFile(answer)
	var recordFormat []ehop.Metric
	json.NewDecoder(bytes.NewReader(f)).Decode(&recordFormat)
	fmt.Println(recordFormat[2].DisplayName)

	structString := "type " + recordName + " struct{\n    "

	for _, metric := range recordFormat {
		s, _ := typeConverter(metric.DataType)
		structString = structString + strings.ToUpper(string(metric.Name[0])) + string(metric.Name[1:]) + "    " + s + "    " + "`json:\"" + metric.Name + "\"`" + "\n    "
	}

	structString = structString + "\n}"

	ioutil.WriteFile(recordName+".output", []byte(structString), 0644)
	fmt.Println(structString)

}
