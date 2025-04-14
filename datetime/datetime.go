package datetime

import (
	"github.com/duke-git/lancet/v2/datetime"
	"time"
)

func FormatTimeToStr(t time.Time, format string, timezone ...string) string {
	return datetime.FormatTimeToStr(t, format, timezone...)
}
