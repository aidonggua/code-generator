package helper

import (
	"github.com/duke-git/lancet/v2/datetime"
	"time"
)

type DateHelper struct {
}

func (d *DateHelper) Now() string {
	return datetime.FormatTimeToStr(time.Now(), "yyyy-mm-dd hh:mm:ss")
}

func (d *DateHelper) Date() string {
	return datetime.FormatTimeToStr(time.Now(), "yyyy-mm-dd")
}

func (d *DateHelper) Time() string {
	return datetime.FormatTimeToStr(time.Now(), "hh:mm:ss")
}
