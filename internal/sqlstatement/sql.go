package sqlstatement

var (
	SQL_DROP_TABLE = `
		DROP TABLE IF EXISTS %s
	`

	SQL_CREATE_TABLE = `
		CREATE TABLE IF NOT EXISTS %s
		(
			ClientId UInt32,
			EventType String,
			Data String,
			Time DateTime('Etc/UTC'),
			ClientConnectionTime DateTime('Etc/UTC'),
			SessionDurationSeconds UInt64,
			StreamName String,
			StreamPluginName String,
			Provider String,
			Platform String,
			Merchant String,
			Account String,
			ProjectId String,
			Fp String,
			PlayerVersion String,
			Ip String,
			RealIp String,
			Forward String,
			BitRate Float64,
			FramesDecodedRate Float64,
			NackCountRate Float64,
			PacketsLossPercent Float64,
			Rtt Float64,
			DelayStart Float64,
			PreStart Float64,
			Ua String,
			BrowserName String,
			BrowserVersion String,
			OsName String,
			OsPlatform String,
			OsVersion String,
			DeviceType String,
			Geohash String,
			AccuracyRadius UInt64,
			Latitude Float64,
			Longitude Float64,
			Timezone String,
			CityNameEn String,
			CountryNameEn String,
			ContinentNameEn String
		)
		ENGINE = %s
		ORDER BY (Time)
	`
)
