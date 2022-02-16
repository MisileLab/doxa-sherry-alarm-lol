package main

import (
	"fmt"
	"time"
	"os"
	"log"
	"github.com/faiface/beep/mp3"
	gomp3 "github.com/hajimehoshi/go-mp3" 
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep"
	"encoding/json"
	"github.com/gen2brain/beeep"
)

type timestruct struct {
	hour int
	minute int
	second int
	am bool // true if am else pm
}

func (t timestruct) check_time(time time.Time) bool {
	hour := t.hour
	if !t.am {
		hour = hour + 12
	}
	if hour == time.Hour() && t.minute == time.Minute() && t.second == time.Second() {
		return true
	} else {
		return false
	}
}

func main() {
	const sampleSize = 4
	reset := os.Getenv("reset")
	if reset == "true" {
		os.Remove("config.json")
	}
	_, err := os.Open("config.json")
	if err != nil {
		_, err = os.Create("config.json")
		if err != nil {
			log.Fatal(err)
		}
	}
	data := make(map[string]interface{})
	filecontent, _ := os.ReadFile("config.json")
	json.Unmarshal([]byte(filecontent), &data)
	var sleephour int
	var sleepmin int
	var sleepam bool
	var sleeprepeat int
	var dayhour int
	var daymin int
	var dayam bool
	var dayrepeat int
	if data["sleephour"] == nil {
		println("몇 시에 자는지 알려주세요. 예시: 12시 = 12") 
		_, _ = fmt.Scan(&sleephour)
		data["sleephour"] = sleephour
	}
	if data["sleepmin"] == nil {
		println("몇 분에 자는지 알려주세요. 예시: 30분 = 30") 
		_, _ = fmt.Scan(&sleepmin)
		data["sleepmin"] = sleepmin
	}
	if data["sleepam"] == nil {
		println("자는 시간이 오전인지 오후인지 알려주세요. 예시: 오전 = true")
		_, _ = fmt.Scan(&sleepam)
		data["sleepam"] = sleepam
	}
	if data["sleeprepeat"] == nil {
		println("알람 소리가 몇 번 반복될지 알려주세요. (잘 때) 예시: 3번 = 3")
		_, _ = fmt.Scan(&sleeprepeat)
		data["sleeprepeat"] = sleeprepeat
	}
	if data["dayhour"] == nil {
		println("몇 시에 일어나는지 알려주세요. 예시: 12시 = 12") 
		_, _ = fmt.Scan(&dayhour)
		data["dayhour"] = dayhour
	}
	if data["daymin"] == nil {
		println("몇 분에 일어나는지 알려주세요. 예시: 30분 = 30") 
		_, _ = fmt.Scan(&daymin)
		data["daymin"] = daymin
	}
	if data["dayam"] == nil{
		println("일어나는 시간이 오전인지 오후인지 알려주세요. 예시: 오전 = true")
		_, _ = fmt.Scan(&dayam)
		data["dayam"] = dayam
	}
	if data["dayrepeat"] == nil {
		println("알람 소리가 몇 번 반복될지 알려주세요. (일어날 때) 예시: 3번 = 3")
		_, _ = fmt.Scan(&dayrepeat)
		data["dayrepeat"] = dayrepeat
	}
	sleephour, _ = data["sleephour"].(int)
	sleepmin, _ = data["sleepmin"].(int)
	sleepam, _ = data["sleepam"].(bool)
	sleeprepeat, _ = data["sleeprepeat"].(int)
	dayhour, _ = data["dayhour"].(int)
	daymin, _ = data["daymin"].(int)
	dayam, _ = data["dayam"].(bool)
	dayrepeat, _ = data["dayrepeat"].(int)

	jsoncontent, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile("config.json", jsoncontent, 0644)

	sleeptime := timestruct{hour: sleephour, minute: sleepmin, second: 0, am: sleepam}
	daytime := timestruct{hour: dayhour, minute: daymin, second: 0, am: dayam}

	goodnightfile, err := os.Open("./sounds/goodnight.mp3")
	if err != nil { log.Fatal(err) }
	goodnightstreamer, goodnightformat, err := mp3.Decode(goodnightfile)
	if err != nil { log.Fatal(err)}
	goodnightbuffer := beep.NewBuffer(goodnightformat)
	goodnightbuffer.Append(goodnightstreamer)
	defer goodnightstreamer.Close()
	gooddayfile, gooddayfileerror := os.Open("./sounds/goodday.mp3")
	if gooddayfileerror != nil { log.Fatal(gooddayfileerror) }
	gooddaystreamer, gooddayformat, gooddayerror := mp3.Decode(gooddayfile)
	if gooddayerror != nil { log.Fatal(gooddayerror) }
	gooddaybuffer := beep.NewBuffer(gooddayformat)
	gooddaybuffer.Append(gooddaystreamer)
	defer gooddaystreamer.Close()
	goodnightfileanother, _ := os.Open("./sounds/goodnight.mp3")
	gooddayfileanother, _ := os.Open("./sounds/goodday.mp3")
	d, err := gomp3.NewDecoder(goodnightfileanother)
	if err != nil { log.Fatal(err) }
	samples := d.Length() / sampleSize
	goodnightlength := samples / int64(d.SampleRate())
	d, err = gomp3.NewDecoder(gooddayfileanother)
	if err != nil { log.Fatal(err) }
	samples = d.Length() / sampleSize
	gooddaylength := samples / int64(d.SampleRate())

	speaker.Init(goodnightformat.SampleRate, goodnightformat.SampleRate.N(time.Second/10))

	println("알람이 성공적으로 실행되었습니다!")

	for {
		nowtime := time.Now()
		println("현재 시간: " + nowtime.Format("15:04:05"))
		if sleeptime.check_time(nowtime) {
			err := beeep.Notify("Good Night", "잘 자요.", "assets/doxa.png")
			if err != nil {
				panic(err)
			}
			for i := 1; i <= sleeprepeat; i++ {
				goodnightsound := goodnightbuffer.Streamer(0, goodnightbuffer.Len())
				speaker.Play(goodnightsound)
				time.Sleep(time.Duration(second_to_nano(goodnightlength)))
			}
		} else if daytime.check_time(nowtime) {
			err := beeep.Notify("Wake up!", "일어나!", "assets/sherry.png")
			if err != nil {
				panic(err)
			}
			for i := 1; i <= dayrepeat; i++ {
				gooddaysound := gooddaybuffer.Streamer(0, gooddaybuffer.Len())
				speaker.Play(gooddaysound)
				time.Sleep(time.Duration(second_to_nano(gooddaylength)))
			}
		}
		time.Sleep(time.Duration(second_to_nano(1)))
	}
}

func second_to_nano(integer int64) int64 {
	return integer * 1000000000
}