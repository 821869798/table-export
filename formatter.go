package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type Formatter struct{}

func (m *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string

	fieldsStr := ""
	if len(entry.Data) > 0 {
		fields := make([]string, 0, len(entry.Data))
		for k, v := range entry.Data {
			fields = append(fields, fmt.Sprintf("%s=%v", k, v))
		}
		fieldsStr = " " + strings.Join(fields, ", ")
	}

	//HasCaller()为true才会有调用信息
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("\x1b[%dm[%s] [%s] [%s:%d %s]\x1b[0m %s%s\n",
			levelColor, timestamp, entry.Level, fName, entry.Caller.Line, entry.Caller.Function, entry.Message, fieldsStr)
	} else {
		newLog = fmt.Sprintf("\x1b[%dm[%s] [%s]\x1b[0m %s%s\n", levelColor, timestamp, entry.Level, entry.Message, fieldsStr)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}
