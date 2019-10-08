package util

import "fmt"

func FontColor(s string) string {
	return fmt.Sprintf("\033[1;31;40m%s\033[0m", s)
}
