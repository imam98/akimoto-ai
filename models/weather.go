package models

type Report struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Weather   Weather `json:"currently"`
}

type Weather struct {
	Time        int64   `json:"time"`
	Summary     string  `json:"summary"`
	Icon        string  `json:"icon"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	PrecipProb  float32 `json:"precipProbability"`
}

const MsgTmpl string = `*Weather Report [{{time .Weather.Time}}]*
*{{.Timezone}}*
==========================
{{ic2emot .Weather.Icon}} {{.Weather.Summary}}

Temp. {{.Weather.Temperature}}°C
Humidity: {{.Weather.Humidity}}%
PrecipProb: {{.Weather.PrecipProb}}
{{if and (gt .Weather.Humidity 60.0) (gt .Weather.PrecipProb 60.0)}}
It might be raining later, I suggest to bring an umbrella {{print "\xE2\x98\x94"}}
{{else}}
I think the weather would be nice for the moment 🤞
{{end}}
`

const (
	Excluded   string = "minutely,hourly,daily,alerts,flags"
	DsBasePath string = "https://api.darksky.net/forecast/"

	IconClear        string = "\xE2\x98\x80"
	IconPartialCloud string = "🌤"
	IconCloudySky    string = "🌥"
	IconHumid        string = "\xE2\x98\x81"
	IconRain         string = "🌧️"
	IconThunderstorm string = "🌩️"
)
