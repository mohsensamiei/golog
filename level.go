package log

type Level int

const (
	FatalLevel Level = iota
	ErrorLevel
	WarningLevel
	InfoLevel
	DebugLevel
)

func (lvl Level) String() string {
	switch lvl {
	case FatalLevel:
		return "Fatal"
	case ErrorLevel:
		return "Error"
	case WarningLevel:
		return "Warning"
	case InfoLevel:
		return "Info"
	case DebugLevel:
		return "Debug"
	default:
		return "Unknown"
	}
}
