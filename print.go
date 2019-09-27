package log

import (
	"fmt"
	"os"
)

func Trace(v ...interface{}) {
	trace := loggers[Leveltrace]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = joint(lTrace, s, WinColorGreen)
		defer winKernelColse()
	} else {
		message = joint(lTrace, s, colorGreen)
	}
	trace.Println(message)
}

func Tracef(format string, v ...interface{}) {
	trace := loggers[Leveltrace]
	s := fmt.Sprintf(format, v...)
	var message string
	if isColor && systemType == "windows" {
		message = joint(lTrace, s, WinColorGreen)
		defer winKernelColse()
	} else {
		message = joint(lTrace, s, colorGreen)
	}
	trace.Println(message)
}

func Info(v ...interface{}) {
	info := loggers[Levelinfo]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = joint(lInfo, s, WinColorBlue)
		defer winKernelColse()
	} else {
		message = joint(lInfo, s, colorBlue)
	}
	info.Println(message)
}

func Infof(format string, v ...interface{}) {
	info := loggers[Levelinfo]
	s := fmt.Sprintf(format, v...)
	var message string
	if isColor && systemType == "windows" {
		message = joint(lInfo, s, WinColorBlue)
		defer winKernelColse()
	} else {
		message = joint(lInfo, s, colorBlue)
	}
	info.Println(message)
}

func Debug(v ...interface{}) {
	debug := loggers[Leveldebug]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = joint(lDebug, s, WinColorDarkblue)
		defer winKernelColse()
	} else {
		message = joint(lDebug, s, colorDarkblue)
	}
	debug.Println(message)
}

func Debugf(format string, v ...interface{}) {
	debug := loggers[Leveldebug]
	s := fmt.Sprintf(format, v...)
	var message string
	if isColor && systemType == "windows" {
		message = joint(lDebug, s, WinColorDarkblue)
		defer winKernelColse()
	} else {
		message = joint(lDebug, s, colorDarkblue)
	}
	debug.Println(message)
}

func Warn(v ...interface{}) {
	warn := loggers[Levelwarn]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = joint(lWarn, s, WinColorYellow)
		defer winKernelColse()
	} else {
		message = joint(lWarn, s, colorYellow)
	}
	warn.Println(message)
}

func Warnf(format string, v ...interface{}) {
	warn := loggers[Levelwarn]
	s := fmt.Sprintf(format, v...)
	var message string
	if isColor && systemType == "windows" {
		message = joint(lWarn, s, WinColorYellow)
		defer winKernelColse()
	} else {
		message = joint(lWarn, s, colorYellow)
	}
	warn.Println(message)
}

func Error(v ...interface{}) {
	e := loggers[Levelerror]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = joint(lError, s, WinColorRed)
		defer winKernelColse()
	} else {
		message = joint(lError, s, colorRed)
	}
	e.Println(message)
}

func Errorf(format string, v ...interface{}) {
	e := loggers[Levelerror]
	s := fmt.Sprintf(format, v...)
	var message string
	if isColor && systemType == "windows" {
		message = joint(lError, s, WinColorRed)
		defer winKernelColse()
	} else {
		message = joint(lError, s, colorRed)
	}
	e.Println(message)
}

func Panic(v ...interface{}) {
	p := loggers[Levelpanic]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = joint(lPanic, s, WinColorPurple)
		defer winKernelColse()
	} else {
		message = joint(lPanic, s, colorPurple)
	}
	p.Panicln(message)
}

func Panicf(format string, v ...interface{}) {
	p := loggers[Levelpanic]
	s := fmt.Sprintf(format, v...)
	var message string
	if isColor && systemType == "windows" {
		message = joint(lPanic, s, WinColorPurple)
		defer winKernelColse()
	} else {
		message = joint(lPanic, s, colorPurple)
	}
	p.Panicln(message)
}

func Fatal(v ...interface{}) {
	falat := loggers[Levelfatal]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = joint(lFatal, s, WinColorPurple)
	} else {
		message = joint(lFatal, s, colorPurple)
	}
	falat.Println(message)
	winKernelColse()
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	falat := loggers[Levelfatal]
	s := fmt.Sprintf(format, v...)
	var message string
	if isColor && systemType == "windows" {
		message = joint(lFatal, s,  WinColorPurple)
	} else {
		message = joint(lFatal, s, colorPurple)
	}
	falat.Println(message)
	winKernelColse()
	os.Exit(1)
}
