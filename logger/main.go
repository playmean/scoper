package logger

import (
	"fmt"
	"log"
	"strings"
)

func buildLine(tag, format string, v ...interface{}) string {
	var line string

	if len(tag) > 0 {
		line += "[" + tag + "] "
	}

	if strings.ContainsAny(format, "%") {
		line += fmt.Sprintf(format, v...)
	} else {
		line += format + " "

		for _, chunk := range v {
			line += fmt.Sprintf("%v ", chunk)
		}
	}

	return line
}

// Log data to stdout
func Log(tag, format string, v ...interface{}) {
	line := buildLine(tag, format, v...)

	log.Println(line)
}

// Fatal after logging
func Fatal(tag, format string, v ...interface{}) {
	line := buildLine(tag, format, v...)

	log.Fatalln(line)
}
