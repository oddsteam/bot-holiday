package botholiday

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HolidayConfig struct {
	Url         string
	AccessToken string
}

type HolidayResponse struct {
	Result HolidayResponseResult `json:"result"`
}

type HolidayResponseResult struct {
	Data []HolidayResponseData `json:"data"`
}

type HolidayResponseData struct {
	HolidayWeekDay         string `json:"HolidayWeekDay"`
	HolidayWeekDayThai     string `json:"HolidayWeekDayThai"`
	Date                   string `json:"Date"`
	DateThai               string `json:"DateThai"`
	HolidayDescription     string `json:"HolidayDescription"`
	HolidayDescriptionThai string `json:"HolidayDescriptionThai"`
}

type IHolidayConfig interface {
	GetBOTHoliday() HolidayResponse
}

func Initialize(token string) IHolidayConfig {
	return &HolidayConfig{
		Url:         "https://gateway.api.bot.or.th/financial-institutions-holidays/",
		AccessToken: token,
	}
}

func CheckHoliday(holidayRes HolidayResponse, t time.Time) bool {
	if t.Weekday() == time.Saturday || t.Weekday() == time.Sunday {
		return true
	}
	isHoliday := false
	layout := "2006-01-02"
	timeNow, _ := time.Parse(layout, time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).Format(layout))
	for _, value := range holidayRes.Result.Data {
		date, _ := time.Parse(layout, value.Date)
		if timeNow.Equal(date) {
			isHoliday = true
			fmt.Printf("Today is holiday: %s\n", value.HolidayDescriptionThai)
			break
		}
	}
	return isHoliday
}

func (conf HolidayConfig) GetBOTHoliday() HolidayResponse {
	url := fmt.Sprintf("%s?year=%v", conf.Url, time.Now().Local().Year())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conf.AccessToken))
	req.Header.Set("accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	holidayRes, err := parseHttpResponse(res)
	if err != nil {
		fmt.Println(err)
	}
	return holidayRes
}

func parseHttpResponse(res *http.Response) (HolidayResponse, error) {
	holidays := HolidayResponse{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return HolidayResponse{}, err
	}
	defer res.Body.Close()
	err = json.Unmarshal(body, &holidays)
	if err != nil {
		return HolidayResponse{}, err
	}
	return holidays, nil
}
