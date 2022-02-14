package event

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ClientId               uint32
	EventType              string
	Data                   string
	Time                   time.Time
	ClientConnectionTime   time.Time
	SessionDurationSeconds uint64
	StreamName             string
	StreamPluginName       string
	Provider               string
	Platform               string
	Merchant               string
	Account                string
	ProjectId              string
	Fp                     string
	PlayerVersion          string
	Ip                     string
	RealIp                 string
	Forward                string
	BitRate                float64
	FramesDecodedRate      float64
	NackCountRate          float64
	PacketsLossPercent     float64
	Rtt                    float64
	DelayStart             float64
	PreStart               float64
	Ua                     string
	BrowserName            string
	BrowserVersion         string
	OsName                 string
	OsPlatform             string
	OsVersion              string
	DeviceType             string
	Geohash                string
	AccuracyRadius         uint64
	Latitude               float64
	Longitude              float64
	Timezone               string
	CityNameEn             string
	CountryNameEn          string
	ContinentNameEn        string
}

func NewEvent() Event {
	return Event{
		ClientId:           uuid.New().ID(),
		EventType:          "test",
		Data:               "test",
		Time:               time.Now(),
		StreamName:         "test",
		StreamPluginName:   "test",
		Provider:           "default",
		Platform:           "default",
		Merchant:           "default",
		Account:            "default",
		ProjectId:          "default",
		Fp:                 "hahahah",
		PlayerVersion:      "test",
		Ip:                 "localhost",
		RealIp:             "test",
		Forward:            "test",
		BitRate:            25.2,
		FramesDecodedRate:  55.4,
		NackCountRate:      0.02,
		PacketsLossPercent: 0.0,
		Rtt:                20.1,
		DelayStart:         0.66,
		PreStart:           0.31,
		Ua:                 "Geckoengine",
		BrowserName:        "chrome",
		BrowserVersion:     "dklajfklda",
		OsName:             "linux!!",
		OsVersion:          "ooboonta",
		DeviceType:         "compukter",
		Geohash:            "jaldksj234524akhdgoahjsdf",
		AccuracyRadius:     30,
		Latitude:           2354.4524,
		Longitude:          3421.454,
		Timezone:           "EuropeMoscow",
		CityNameEn:         "Moscow",
		CountryNameEn:      "Rossiya",
		ContinentNameEn:    "Eurasia",
	}
}
