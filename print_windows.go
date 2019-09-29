package log

import (
	"fmt"
	"os"
)

func Trace(v ...interface{}) {
	trace := loggers[Leveltrace]
	s := fmt.Sprint(v...)
	var message string
	message = joint(lTrace, s, WinColorGreen)
	defer winKernelColse()
	trace.Println(message)
}

func Tracef(format string, v ...interface{}) {
	trace := loggers[Leveltrace]
	s := fmt.Sprintf(format, v...)
	var message string
	message = joint(lTrace, s, WinColorGreen)
	trace.Println(message)
}

func Info(v ...interface{}) {
	info := loggers[Levelinfo]
	s := fmt.Sprint(v...)
	var message string
	message = joint(lInfo, s, WinColorBlue)
	defer winKernelColse()
	info.Println(message)
}

func Infof(format string, v ...interface{}) {
	info := loggers[Levelinfo]
	s := fmt.Sprintf(format, v...)
	var message string
	message = joint(lInfo, s, WinColorBlue)
	defer winKernelColse()
	info.Println(message)
}

func Debug(v ...interface{}) {
	debug := loggers[Leveldebug]
	s := fmt.Sprint(v...)
	var message string
	message = joint(lDebug, s, WinColorDarkblue)
	defer winKernelColse()
	debug.Println(message)
}

func Debugf(format string, v ...interface{}) {
	debug := loggers[Leveldebug]
	s := fmt.Sprintf(format, v...)
	var message string
	message = joint(lDebug, s, WinColorDarkblue)
	defer winKernelColse()
	debug.Println(message)
}

func Warn(v ...interface{}) {
	warn := loggers[Levelwarn]
	s := fmt.Sprint(v...)
	var message string
	message = joint(lWarn, s, WinColorYellow)
	defer winKernelColse()
	warn.Println(message)
}

func Warnf(format string, v ...interface{}) {
	warn := loggers[Levelwarn]
	s := fmt.Sprintf(format, v...)
	var message string
	message = joint(lWarn, s, WinColorYellow)
	defer winKernelColse()
	warn.Println(message)
}

func Error(v ...interface{}) {
	e := loggers[Levelerror]
	s := fmt.Sprint(v...)
	var message string
	message = joint(lError, s, WinColorRed)
	defer winKernelColse()
	e.Println(message)
}

func Errorf(format string, v ...interface{}) {
	e := loggers[Levelerror]
	s := fmt.Sprintf(format, v...)
	var message string
	message = joint(lError, s, WinColorRed)
	defer winKernelColse()
	e.Println(message)
}

func Panic(v ...interface{}) {
	p := loggers[Levelpanic]
	s := fmt.Sprint(v...)
	var message string
	message = joint(lPanic, s, WinColorPurple)
	defer winKernelColse()
	p.Panicln(message)
}

func Panicf(format string, v ...interface{}) {
	p := loggers[Levelpanic]
	s := fmt.Sprintf(format, v...)
	var message string
	message = joint(lPanic, s, WinColorPurple)
	defer winKernelColse()
	p.Panicln(message)
}

func Fatal(v ...interface{}) {
	falat := loggers[Levelfatal]
	s := fmt.Sprint(v...)
	var message string
	message = joint(lFatal, s, WinColorPurple)
	falat.Println(message)
	winKernelColse()
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	falat := loggers[Levelfatal]
	s := fmt.Sprintf(format, v...)
	var message string
	message = joint(lFatal, s, WinColorPurple)
	falat.Println(message)
	winKernelColse()
	os.Exit(1)
}

func (l mylog) Trace(v ...interface{}) {
	trace := l.loggers[Leveltrace]
	s := fmt.Sprint(v...)
	var message string
	message = l.joint(lTrace, s, WinColorGreen)
	defer winKernelColse()
	trace.Println(message)
}

func (l mylog) Tracef(format string, v ...interface{}) {
	trace := l.loggers[Leveltrace]
	s := fmt.Sprint(v...)
	var message string
	message = l.joint(lTrace, s, WinColorGreen)
	defer winKernelColse()
	trace.Println(message)
}

func (l mylog) Info(v ...interface{}) {
	info := l.loggers[Levelinfo]
	s := fmt.Sprint(v...)
	var message string
	message = l.joint(lInfo, s, WinColorBlue)
	defer winKernelColse()
	info.Println(message)
}

func (l mylog) Infof(format string, v ...interface{}) {
	info := l.loggers[Levelinfo]
	s := fmt.Sprint(v...)
	var message string
	message = l.joint(lInfo, s, WinColorBlue)
	defer winKernelColse()
	info.Println(message)
}

func (l mylog) Debug(v ...interface{}) {
	debug := l.loggers[Leveldebug]
	s := fmt.Sprint(v...)
	var message string
	message = l.joint(lDebug, s, WinColorDarkblue)
	defer winKernelColse()
	debug.Println(message)
}

func (l mylog) Debugf(format string, v ...interface{}) {
	debug := l.loggers[Leveldebug]
	s := fmt.Sprint(v...)
	var message string
	message = l.joint(lDebug, s, WinColorDarkblue)
	defer winKernelColse()
	debug.Println(message)
}

func (l mylog) Warn(v ...interface{}) {
	warn := l.loggers[Levelwarn]
	s := fmt.Sprint(v...)
	var message string
	message = l.joint(lWarn, s, WinColorYellow)
	defer winKernelColse()
	warn.Println(message)
}

func (l mylog) Warnf(format string, v ...interface{}) {
	warn := l.loggers[Levelwarn]
	s := fmt.Sprint(v...)
	var message string
	message = l.joint(lWarn, s, WinColorYellow)
	defer winKernelColse()
	warn.Println(message)
}

func (l mylog) Error(v ...interface{}) {
	e := l.loggers[Levelerror]
	s := fmt.Sprint(v...)
	var message string
	message = l.joint(lError, s, WinColorRed)
	defer winKernelColse()
	e.Println(message)
}

func (l mylog) Errorf(format string, v ...interface{}) {
	e := l.loggers[Levelerror]
	s := fmt.Sprint(v...)
	var message string
	message = l.joint(lError, s, WinColorRed)
	defer winKernelColse()
	e.Println(message)
}

func (l mylog) Panic(v ...interface{}) {
	p := l.loggers[Levelpanic]
	s := fmt.Sprint(v...)
	var message string
	message = l.joint(lPanic, s, WinColorPurple)
	defer winKernelColse()
	p.Panicln(message)
}

func (l mylog) Panicf(format string, v ...interface{}) {
	p := l.loggers[Levelpanic]
	s := fmt.Sprint(v...)
	var message string
	message = l.joint(lPanic, s, WinColorPurple)
	defer winKernelColse()
	p.Panicln(message)
}

func (l mylog) Fatal(v ...interface{}) {
	falat := l.loggers[Levelfatal]
	s := fmt.Sprint(v...)
	var message string
	message = l.joint(lFatal, s, WinColorPurple)
	falat.Println(message)
	winKernelColse()
	os.Exit(1)
}

func (l mylog) Fatalf(format string, v ...interface{}) {
	falat := l.loggers[Levelfatal]
	s := fmt.Sprint(v...)
	var message string
	message = l.joint(lFatal, s, WinColorPurple)
	falat.Println(message)
	winKernelColse()
	os.Exit(1)
}
