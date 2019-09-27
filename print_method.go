package log

import (
	"fmt"
	"os"
)

func (l mylog) Trace(v ...interface{}) {
	trace := l.loggers[Leveltrace]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message =  l.joint(lTrace, s, WinColorGreen)
		defer winKernelColse()
	} else {
		message =  l.joint(lTrace, s, colorGreen)
	}
	trace.Println(message)
}

func (l mylog) Tracef(format string, v ...interface{}) {
	trace := l.loggers[Leveltrace]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message =  l.joint(lTrace, s, WinColorGreen)
		defer winKernelColse()
	} else {
		message =  l.joint(lTrace, s, colorGreen)
	}
	trace.Println(message)
}

func (l mylog) Info(v ...interface{}) {
	info := l.loggers[Levelinfo]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = l.joint(lInfo, s, WinColorBlue)
		defer winKernelColse()
	} else {
		message = l.joint(lInfo, s, colorBlue)
	}
	info.Println(message)
}

func (l mylog) Infof(format string, v ...interface{}) {
	info := l.loggers[Levelinfo]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = l.joint(lInfo, s, WinColorBlue)
		defer winKernelColse()
	} else {
		message = l.joint(lInfo, s, colorBlue)
	}
	info.Println(message)
}

func (l mylog) Debug(v ...interface{}) {
	debug := l.loggers[Leveldebug]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = l.joint(lDebug, s, WinColorDarkblue)
		defer winKernelColse()
	} else {
		message = l.joint(lDebug, s, colorDarkblue)
	}
	debug.Println(message)
}

func (l mylog) Debugf(format string, v ...interface{}) {
	debug := l.loggers[Leveldebug]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = l.joint(lDebug, s, WinColorDarkblue)
		defer winKernelColse()
	} else {
		message = l.joint(lDebug, s, colorDarkblue)
	}
	debug.Println(message)
}

func (l mylog) Warn(v ...interface{}) {
	warn := l.loggers[Levelwarn]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = l.joint(lWarn, s, WinColorYellow)
		defer winKernelColse()
	} else {
		message = l.joint(lWarn, s, colorYellow)
	}
	warn.Println(message)
}

func (l mylog) Warnf(format string, v ...interface{}) {
	warn := l.loggers[Levelwarn]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = l.joint(lWarn, s, WinColorYellow)
		defer winKernelColse()
	} else {
		message = l.joint(lWarn, s, colorYellow)
	}
	warn.Println(message)
}

func (l mylog) Error(v ...interface{}) {
	e := l.loggers[Levelerror]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = l.joint(lError, s, WinColorRed)
		defer winKernelColse()
	} else {
		message = l.joint(lError, s, colorRed)
	}
	e.Println(message)
}

func (l mylog) Errorf(format string, v ...interface{}) {
	e := l.loggers[Levelerror]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = l.joint(lError, s, WinColorRed)
		defer winKernelColse()
	} else {
		message = l.joint(lError, s, colorRed)
	}
	e.Println(message)
}

func (l mylog) Panic(v ...interface{}) {
	p := l.loggers[Levelpanic]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = l.joint(lPanic, s, WinColorPurple)
		defer winKernelColse()
	} else {
		message = l.joint(lPanic, s, colorPurple)
	}
	p.Panicln(message)
}

func (l mylog) Panicf(format string, v ...interface{}) {
	p := l.loggers[Levelpanic]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = l.joint(lPanic, s, WinColorPurple)
		defer winKernelColse()
	} else {
		message = l.joint(lPanic, s, colorPurple)
	}
	p.Panicln(message)
}

func (l mylog) Fatal(v ...interface{}) {
	falat := l.loggers[Levelfatal]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = l.joint(lFatal, s, WinColorPurple)
	} else {
		message = l.joint(lFatal, s, colorPurple)
	}
	falat.Println(message)
	winKernelColse()
	os.Exit(1)
}

func (l mylog) Fatalf(format string, v ...interface{}) {
	falat := l.loggers[Levelfatal]
	s := fmt.Sprint(v...)
	var message string
	if isColor && systemType == "windows" {
		message = l.joint(lFatal, s, WinColorPurple)
	} else {
		message = l.joint(lFatal, s, colorPurple)
	}
	falat.Println(message)
	winKernelColse()
	os.Exit(1)
}
