package debug

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var prefix = ""

func SetPrefix(s string) {
	prefix = s
}

func SetOutput(io *os.File) {
	log.SetOutput(io)
}

func printf(skip int, fmt string, v ...interface{}) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2+skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	format := prefix + fmt
	format = strings.Replace(format, "{{file}}", frame.File, -1)
	format = strings.Replace(format, "{{line}}", strconv.Itoa(frame.Line), -1)
	format = strings.Replace(format, "{{func}}", frame.Function, -1)

	log.Printf(format, v...)
}

func Printf(fmt string, v ...interface{}) {
	printf(1, fmt, v...)
}

func Print(v ...interface{}) {
	printf(1, "%s", fmt.Sprint(v...))
}

func Println(v ...interface{}) {
	printf(1, "%s", fmt.Sprintln(v...))
}
