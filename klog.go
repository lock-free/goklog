package klog

import (
    "fmt"
    "github.com/go-errors/errors"
    "os"
    "strconv"
    "sync"
    "time"
)

type ListenerFunc func(string)
func emptyListener(message string) {

}

type KLog struct {
    listener ListenerFunc
    inspectOpened bool
    formatter string
    printLevel int
}
var instance *KLog

var once sync.Once

const (
    NORMAL_LEVEL = iota
    ERROR_LEVEL
    VITAL
)

// copied from: https://stackoverflow.com/questions/40326540/how-to-assign-default-value-if-env-var-is-empty
func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

func getErrorCallStack(e error) string {
    if e == nil {
        return ""
    }

    e0, ok := e.(*errors.Error)
    if !ok {
        return e.Error()
    }
    callStack := string(e0.ErrorStack())
    return callStack
}

func GetInstance() *KLog {
    once.Do(func() {
    	level, err := strconv.ParseInt(getenv("KLOG_PRINT_LEVEL", "0"), 0, 8)
    	if err != nil {
    	    fmt.Println("System environment KLOG_PRINT_LEVEL is polluted")
    	    level = NORMAL_LEVEL
        }
        instance = &KLog{
            listener: emptyListener,
            inspectOpened: false,
            formatter: "2006-01-02-15.04.05.999",
            printLevel: int(level),
        }
    })
    return instance
}

func (klog *KLog) SetListener(listener ListenerFunc) {
    klog.listener = listener
}

func (klog *KLog) ToggleInspector() {
    klog.inspectOpened = !klog.inspectOpened
}

func (klog *KLog) LogIn(title string, text string) {
    if klog.inspectOpened {
        klog.Info(title, text, NORMAL_LEVEL)
    }
}

func (klog *KLog) Info(title string, text string, level int) {
    today := time.Now()
    logText := fmt.Sprintf("[klog:%s] (%s) %s", today.Format(klog.formatter), title, text)

    if level >= klog.printLevel {
        fmt.Println(logText)
    }

    klog.listener(logText)
}

func (klog *KLog) LogVital(title string, text string) {
    klog.Info(title, text, VITAL)
}

func (klog *KLog) LogError(title string, err error) {
    errorText := getErrorCallStack(err)
    klog.Info(fmt.Sprintf("error-response-%s", title), errorText, ERROR_LEVEL)
}

func (klog *KLog) LogNormal(title string, text string) {
    klog.Info(title, text, NORMAL_LEVEL)
}
