package utility

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gookit/color"
)

func prefix() string {
	dt := time.Now()
	dtf := dt.Format("15:04:05.000")
	red := color.FgLightRed.Render
	white := color.FgLightWhite.Render
	return white("["+dtf+"]") + red(" | ")
}
func taskPrefix(id int) string {
	ids := strconv.Itoa(id)
	dt := time.Now()
	dtf := dt.Format("15:04:05.000")
	red := color.FgLightRed.Render
	white := color.FgLightWhite.Render
	return white("["+dtf+"]") + red(" | ") + white("[Task "+ids+"]") + red(" | ")

}
func red(data string) string {
	red := color.FgLightRed.Render
	return fmt.Sprintf("%s", red(data))
}

func purple(data string) string {
	purple := color.FgLightCyan.Render
	return fmt.Sprintf("%s", purple(data))
}

func cyan(data string) string {
	cyan := color.FgLightCyan.Render
	return fmt.Sprintf("%s", cyan(data))
}

func green(data string) string {
	green := color.FgLightGreen.Render
	return fmt.Sprintf("%s", green(data))
}
func yellow(data string) string {
	yellow := color.FgLightYellow.Render
	return fmt.Sprintf("%s", yellow(data))
}
func white(data string) string {
	white := color.FgLightWhite.Render
	return fmt.Sprintf("%s", white(data))
}
