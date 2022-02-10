package event

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id                 uint32
	Type               string
	Data               string
	Time               time.Time
	Name               string
	Provider           string
	Platform           string
	Merchant           string
	Account            string
	ProjectId          string
	Fp                 string
	PlayerVersion      string
	Ip                 string
	RealIp             string
	Forward            string
	BitRate            float64
	FramesDecodedRate  float64
	NackCountRate      float64
	PacketsLossPercent float64
	Rtt                float64
	DelayStart         float64
	PreStart           float64
	Ua                 string
	BrowserName        string
	BrowserVersion     string
	OsName             string
	OsPlatform         string
	OsVersion          string
	DeviceType         string
	Geohash            string
	AccuracyRadius     uint64
	Latitude           float64
	Longitude          float64
	Timezone           string
	CityNameEn         string
	CountryNameEn      string
	ContinentNameEn    string
}

func NewEvent() Event {
	return Event{
		Id:                 uuid.New().ID(),
		Type:               "test",
		Data:               "test data",
		Time:               time.Now(),
		Name:               "unique name",
		Provider:           "default",
		Platform:           "default",
		Merchant:           "default",
		Account:            "default",
		ProjectId:          "default",
		Fp:                 "hahahah",
		PlayerVersion:      "boring!!1",
		Ip:                 "localhost heh",
		RealIp:             ")",
		Forward:            "i don't care...",
		BitRate:            25.2,
		FramesDecodedRate:  55.4,
		NackCountRate:      0.02,
		PacketsLossPercent: 0,
		Rtt:                20.1,
		DelayStart:         0.66,
		PreStart:           0.31,
		Ua:                 "Gecko engine",
		BrowserName:        "chrome",
		BrowserVersion:     "1.1.1.1",
		OsName:             "linux!!",
		OsVersion:          "ooboonta",
		DeviceType:         "compukter",
		Geohash:            "jaldksj234524akhdgoahjsdf",
		AccuracyRadius:     30,
		Latitude:           2354.4524,
		Longitude:          34521.454,
		Timezone:           "Europe/Moscow",
		CityNameEn:         "Moscow",
		CountryNameEn:      "Rossiya",
		ContinentNameEn:    "Evrasia",
	}
}
