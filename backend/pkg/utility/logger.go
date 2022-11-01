package utility

import (
	"fmt"
)

func LogInfo(module string, msg string) {
	fmt.Println(fmt.Sprintf(prefix()+cyan("[%s] %s"), module, msg))
}

func LogError(module string, msg string) {
	fmt.Println(fmt.Sprintf(prefix()+red("[%s] %s"), module, msg))
}

func LogSuccess(module string, msg string) {
	fmt.Println(fmt.Sprintf(prefix()+green("[%s] %s"), module, msg))
}

func LogMonitor(module string, msg string) {
	fmt.Println(fmt.Sprintf(prefix()+yellow("[%s] %s"), module, msg))
}
