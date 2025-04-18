package logger

type Level int

type Logger struct {
	level Level
}

var log *Logger

const (
	Debug_Level Level = iota
	Info_Level
	Warn_Level
	Error_Level
)

func InitializeLogger(level Level) {
	log = &Logger{level: level}
}

func Debug(msg string) {
	if log.level <= Debug_Level {
		println("DEBUG:", msg)
	}
}

func Info(msg string) {
	if log.level <= Info_Level {
		println("INFO:", msg)
	}
}

func Warn(msg string) {
	if log.level <= Warn_Level {
		println("WARN:", msg)
	}
}

func Error(msg string, err error) {
	if log.level <= Error_Level {
		println("ERROR:", msg, "Error:", err.Error())
	}
}
