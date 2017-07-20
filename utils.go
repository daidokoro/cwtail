package main

// HandleError - exits on error
import (
	"bytes"
	"encoding/json"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

func handleError(err error) {
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func parseStatement(msg string) string {
	reg, _ := regexp.Compile(stateRegex)
	if !reg.MatchString(msg) {
		return msg
	}

	return reg.ReplaceAllStringFunc(msg, func(s string) string {
		return color.New(color.BgBlack).Add(color.Bold).SprintFunc()(s)
	})
}

func parseJSON(msg string) (output string) {
	reg, _ := regexp.Compile(JSONregex)

	if !reg.MatchString(msg) {
		output = msg
		return
	}

	output = reg.ReplaceAllStringFunc(msg, func(js string) string {

		js = strings.Replace(js, `'`, `"`, -1)
		var j bytes.Buffer

		_ = json.Indent(&j, []byte(js), "", "  ")

		reg, _ := regexp.Compile(outputJSONregex)

		return reg.ReplaceAllStringFunc(j.String(), func(s string) string {
			return log.ColorString(s, "yellow")
		})
	})
	return
}

func parse(msg *string) string {
	return parseStatement(parseJSON(*msg))
}
