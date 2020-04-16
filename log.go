package log

import "os"

var (
	level     = InfoLevel
	logger    = NewConsoleLogger()
	constants = make(map[string]interface{})
)

func SetConstant(key string, value interface{}) {
	constants[key] = value
}
func SetLevel(lvl Level) {
	level = lvl
}
func SetLogger(log Logger) {
	logger = log
}
func Info(message string) {
	NewEntry().Info(message)
}
func Error(message string) {
	NewEntry().Error(message)
}
func Debug(message string) {
	NewEntry().Debug(message)
}
func Warning(message string) {
	NewEntry().Warning(message)
}
func Fatal(message string) {
	NewEntry().Fatal(message)
	os.Exit(1)
}
func With(data interface{}) *Entry {
	return NewEntry().With(data)
}
func Data(key string, value interface{}) *Entry {
	return NewEntry().Data(key, value)
}
func Source(src string) *Entry {
	return NewEntry().Source(src)
}
