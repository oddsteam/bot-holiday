package botholiday_test

import (
	"testing"
	"time"

	botholiday "github.com/oddsteam/bot-holiday/pkg/botholiday"
)

func TestCheckIsHoliday(t *testing.T) {
	dateNow := time.Date(2023, 3, 2, 12, 0, 0, 0, time.Local)
	t.Run("should return false when today is NOT contain holiday in year 2023", func(t *testing.T) {
		holidays := botholiday.HolidayResponse{
			Result: botholiday.HolidayResponseResult{
				Data: []botholiday.HolidayResponseData{{
					Date: "2023-01-01",
				}},
			},
		}
		isHoliday := botholiday.CheckHoliday(holidays, dateNow)
		if isHoliday != false {
			t.Error("isHoliday is true; expected false.")
		}
	})
	t.Run("should return true when today is contain holiday in year 2023", func(t *testing.T) {
		holidays := botholiday.HolidayResponse{
			Result: botholiday.HolidayResponseResult{
				Data: []botholiday.HolidayResponseData{{
					Date: dateNow.Format("2006-01-02"),
				}},
			},
		}
		isHoliday := botholiday.CheckHoliday(holidays, dateNow)
		if isHoliday != true {
			t.Error("isHoliday is false; expected true.")
		}
	})

	t.Run("should return true to day is week end", func(t *testing.T) {
		holidays := botholiday.HolidayResponse{
			Result: botholiday.HolidayResponseResult{
				Data: []botholiday.HolidayResponseData{{
					Date: dateNow.Format("2006-01-02"),
				}},
			},
		}
		weekend := time.Date(2023, 4, 15, 12, 0, 0, 0, time.Local)
		isHoliday := botholiday.CheckHoliday(holidays, weekend)
		if isHoliday != true {
			t.Error("isHoliday is false; expected true.")
		}
	})
}
