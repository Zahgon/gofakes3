package gofakes3

import "log"

type LogLevel string

const (
	LogErr   LogLevel = "ERR"
	LogWarn  LogLevel = "WARN"
	LogInfo  LogLevel = "INFO"
	LogDebug LogLevel = "DEBUG"
)

// Logger provides a very minimal target for logging implementations to hit to
// allow arbitrary logging dependencies to be used with GoFakeS3.
//
// Only an interface to the standard library's log package is provided with
// GoFakeS3, other libraries will require an adapter. Adapters are trivial to
// write.
//
// For zap:
//
//	type LogrusLog struct {
//		log *zap.Logger
//	}
//
//	func (l LogrusLog) Print(level LogLevel, v ...interface{}) {
//		switch level {
//		case gofakes3.LogErr:
//			l.log.Error(fmt.Sprint(v...))
//		case gofakes3.LogWarn:
//			l.log.Warn(fmt.Sprint(v...))
//		case gofakes3.LogInfo:
//			l.log.Info(fmt.Sprint(v...))
//		default:
//			panic("unknown level")
//		}
//	}
//
// For logrus:
//
//	type LogrusLog struct {
//		log *logrus.Logger
//	}
//
//	func (l LogrusLog) Print(level LogLevel, v ...interface{}) {
//		switch level {
//		case gofakes3.LogErr:
//			l.log.Errorln(v...)
//		case gofakes3.LogWarn:
//			l.log.Warnln(v...)
//		case gofakes3.LogInfo:
//			l.log.Infoln(v...)
//		default:
//			panic("unknown level")
//		}
//	}
type Logger interface {
	Print(level LogLevel, v ...interface{})
}

// GlobalLog creates a Logger that uses the global log.Println() function.
//
// All levels are reported by default. If you pass levels to this function,
// it will act as a level whitelist.
func GlobalLog(levels ...LogLevel) Logger { _ = "STUB: not implemented"; return *new(Logger) }

// StdLog creates a Logger that uses the stdlib's log.Logger type.
//
// All levels are reported by default. If you pass levels to this function,
// it will act as a level whitelist.
func StdLog(log *log.Logger, levels ...LogLevel) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

// DiscardLog creates a Logger that discards all messages.
func DiscardLog() Logger { _ = "STUB: not implemented"; return *new(Logger) }

type stdLog struct {
	log    func(v ...interface{})
	levels map[LogLevel]bool
}

func newStdLog(log func(v ...interface{}), levels ...LogLevel) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

func (s *stdLog) Print(level LogLevel, v ...interface{}) { _ = "STUB: not implemented"; return }

type discardLog struct{}

func (d discardLog) Print(level LogLevel, v ...interface{}) { _ = "STUB: not implemented"; return }

func MultiLog(loggers ...Logger) Logger { _ = "STUB: not implemented"; return *new(Logger) }

type multiLog struct {
	loggers []Logger
}

func (m multiLog) Print(level LogLevel, v ...interface{}) { _ = "STUB: not implemented"; return }
