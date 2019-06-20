package klog

import (
    "sync"
)

type KLog struct {}
var instance *KLog
type ListenerFunc func(string)

var once sync.Once

func GetInstance() *KLog {
    once.Do(func() {
        instance = &KLog{}
    })
    return instance
}

func (klog *KLog) SetListener(listener ListenerFunc) {

}

func (klog *KLog) ToggleInspector() {

}

func (klog *KLog) LogIn(title string, text string) {

}

func (klog *KLog) Info(title string, text string, level int) {

}

func (klog *KLog) logVital(title string, text string) {

}

func (klog *KLog) logErr(title string, err error) {

}