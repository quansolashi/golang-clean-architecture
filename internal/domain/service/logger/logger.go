package logger

type LogType uint8

const (
	LogTypeString LogType = iota + 1
	LogTypeInt
	LogTypeError
	LogTypeAny
)

type LogField struct {
	Key       string
	Integer   int64
	String    string
	Err       error
	Interface interface{}
	Type      LogType
}

type Logger interface {
	Debug(msg string, fields ...LogField)
	Info(msg string, fields ...LogField)
	Warn(msg string, fields ...LogField)
	Error(msg string, fields ...LogField)
	Fatal(msg string, fields ...LogField)
}

func String(key, value string) LogField {
	return LogField{Key: key, String: value, Type: LogTypeString}
}

func Int(key string, value int64) LogField {
	return LogField{Key: key, Integer: value, Type: LogTypeInt}
}

func Error(err error) LogField {
	return LogField{Err: err, Type: LogTypeError}
}

func Any(key string, value interface{}) LogField {
	return LogField{Key: key, Interface: value, Type: LogTypeAny}
}
