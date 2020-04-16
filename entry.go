package log

import (
	"fmt"
	"github.com/mohsensamiei/golog/pkg/errorsext"
	"github.com/mohsensamiei/golog/pkg/stringsext"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type Entry struct {
	Raise   time.Time
	Level   Level
	Src     string
	Message string
	Fields  map[string]interface{}
}

func NewEntry() *Entry {
	var src string
	if _, file, line, ok := runtime.Caller(2); ok {
		//slash := strings.LastIndex(file, "/")
		//if slash >= 0 {
		//	file = file[slash+1:]
		//}
		src = fmt.Sprintf("%s:%d", file, line)
	}
	entry := &Entry{
		Src:    src,
		Raise:  time.Now(),
		Fields: make(map[string]interface{}),
	}
	for key, value := range constants {
		entry.Fields[key] = value
	}
	return entry
}

func (entry *Entry) Info(message string) {
	entry.Level = InfoLevel
	entry.log(message)
}
func (entry *Entry) Error(message string) {
	entry.Level = ErrorLevel
	entry.log(message)
}
func (entry *Entry) Debug(message string) {
	entry.Level = DebugLevel
	entry.log(message)
}
func (entry *Entry) Warning(message string) {
	entry.Level = WarningLevel
	entry.log(message)
}
func (entry *Entry) Fatal(message string) {
	entry.Level = FatalLevel
	entry.log(message)
	os.Exit(1)
}
func (entry *Entry) With(data interface{}) *Entry {
	switch value := data.(type) {
	case nil:
		return entry
	case errorsext.StackTracer:
		var traces []string
		for _, f := range value.StackTrace() {
			traces = append(traces, fmt.Sprintf("%s:%d", f, f))
		}
		entry.Fields["error"] = errorsext.StackError{
			Message:    value.Error(),
			StackTrace: strings.Join(traces, " < "),
		}
	case error:
		entry.Fields["error"] = errorsext.StackError{
			Message: value.Error(),
		}
	default:
		if reflect.TypeOf(value).Kind() == reflect.Ptr {
			value = reflect.ValueOf(value).Elem().Interface()
		}
		refType := reflect.TypeOf(value)
		switch refType.Kind() {
		case reflect.Map:
			dic := reflect.ValueOf(value)
			for _, key := range dic.MapKeys() {
				entry.Fields[key.String()] = dic.MapIndex(key).Interface()
			}
		case reflect.Struct:
			entry.Fields[stringsext.ToSnakeCase(refType.Name())] = value
		default:
			var data []interface{}
			if entry.Fields["data"] != nil {
				data = entry.Fields["data"].([]interface{})
			}
			entry.Fields["data"] = append(data, value)
		}
	}

	return entry
}
func (entry *Entry) Data(key string, value interface{}) *Entry {
	entry.Fields[key] = value
	return entry
}
func (entry *Entry) log(message string) {
	entry.Message = message
	if entry.Level <= level {
		if err := logger.Log(*entry); err != nil {
			log.Printf("can not log with %t: %v", logger, err)
		}
	}
}
func (entry *Entry) Source(src string) *Entry {
	entry.Src = src
	return entry
}
