package log

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"strings"
)

type Logger interface {
	Log(Entry) error
}

func NewConsoleLogger() Logger {
	return new(consoleLogger)
}

type consoleLogger struct {
}

func (consoleLogger) Log(entry Entry) error {
	raise := entry.Raise.Format("2006-01-02 15:04:05.999-07:00")

	level := strings.ToUpper(entry.Level.String())
	switch entry.Level {
	case FatalLevel:
		level = fmt.Sprintf("\033[1;31m%s\033[0m", level)
	case ErrorLevel:
		level = fmt.Sprintf("\033[0;31m%s\033[0m", level)
	case WarningLevel:
		level = fmt.Sprintf("\033[0;33m%s\033[0m", level)
	case InfoLevel:
		level = fmt.Sprintf("\033[0;36m%s\033[0m", level)
	case DebugLevel:
		level = fmt.Sprintf("\033[0;37m%s\033[0m", level)
	}

	row := fmt.Sprintf("%v | %v | %v \n\tsource: %v", raise, level, strings.TrimSpace(entry.Message), entry.Src)
	if entry.Fields != nil && len(entry.Fields) > 0 {
		bytes, _ := yaml.Marshal(entry.Fields)
		row = fmt.Sprintf("%s\n\t%s", row, strings.TrimSpace(strings.ReplaceAll(string(bytes), "\n", "\n\t")))
	}

	fmt.Println(row)
	return nil
}
