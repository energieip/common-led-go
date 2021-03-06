package driverled

import (
	"encoding/json"
)

const (
	DbStatus  = "status"
	DbConfig  = "config"
	TableName = "leds"

	UrlHello   = "setup/hello"
	UrlStatus  = "status/dump"
	UrlSetup   = "setup/config"
	UrlSetting = "update/settings"
)

//Led led driver representation
type Led struct {
	ID                string  `json:"ID,omitempty"`
	Mac               string  `json:"mac"`
	IP                string  `json:"ip"`
	Group             int     `json:"group"`
	Protocol          string  `json:"protocol"`
	SwitchMac         string  `json:"switchMac"`
	IsConfigured      bool    `json:"isConfigured"`
	SoftwareVersion   float32 `json:"softwareVersion"`
	HardwareVersion   float32 `json:"hardwareVersion"`
	IsBleEnabled      bool    `json:"isBleEnabled"`
	Error             int     `json:"error"`
	IMax              int     `json:"iMax"`
	PMax              int     `json:"pMax"`
	Duration          float64 `json:"duration"`
	Setpoint          int     `json:"setpoint"`
	ThresholdLow      int     `json:"thresholdLow"`
	ThresholdHigh     int     `json:"thresholdHigh"`
	DaisyChainEnabled bool    `json:"daisyChainEnabled"`
	DaisyChainPos     int     `json:"daisyChainPos"`
	DevicePower       int     `json:"devicePower"`
	Energy            float64 `json:"energy"`
	VoltageLed        int     `json:"voltageLed"`
	VoltageInput      int     `json:"voltageInput"`
	LinePower         int     `json:"linePower"`
	TimeToAuto        int     `json:"timeToAuto"`
	Auto              bool    `json:"auto"`
	Watchdog          int     `json:"watchdog"`
	FriendlyName      string  `json:"friendlyName"`
	DumpFrequency     int     `json:"dumpFrequency"`
	SlopeStart        int     `json:"slopeStart"`
	SlopeStop         int     `json:"slopeStop"`
	DefaultSetpoint   *int    `json:"defaultSetpoint,omitempty"` // when the switch is not responding
}

//LedSetup initial setup send by the server when the driver is authorized
type LedSetup struct {
	Mac             string  `json:"mac"`
	IMax            int     `json:"iMax"`
	PMax            int     `json:"pMax"`
	Group           *int    `json:"group,omitempty"`
	Auto            *bool   `json:"auto,omitempty"`
	Watchdog        *int    `json:"watchdog,omitempty"`
	IsBleEnabled    *bool   `json:"isBleEnabled,omitempty"`
	ThresholdHigh   *int    `json:"thresholdHigh,omitempty"`
	ThresholdLow    *int    `json:"thresholdLow,omitempty"`
	FriendlyName    *string `json:"friendlyName,omitempty"`
	SwitchMac       string  `json:"switchMac"`
	IsConfigured    *bool   `json:"isConfigured,omitempty"`
	DumpFrequency   int     `json:"dumpFrequency"`
	DefaultSetpoint *int    `json:"defaultSetpoint,omitempty"` // when the switch is not responding
	SetpointManual  *int    `json:"setpointManual,omitempty"`
	SetpointAuto    *int    `json:"setpointAuto,omitempty"`
	SlopeStart      *int    `json:"slopeStart,omitempty"`
	SlopeStop       *int    `json:"slopeStop,omitempty"`
}

//LedConf customizable configuration by the server
type LedConf struct {
	Mac            string  `json:"mac"`
	Group          *int    `json:"group,omitempty"`
	SetpointManual *int    `json:"setpointManual,omitempty"`
	SetpointAuto   *int    `json:"setpointAuto,omitempty"`
	Auto           *bool   `json:"auto,omitempty"`
	Watchdog       *int    `json:"watchdog,omitempty"`
	IsConfigured   *bool   `json:"isConfigured,omitempty"`
	IsBleEnabled   *bool   `json:"isBleEnabled,omitempty"`
	ThresholdHigh  *int    `json:"thresholdHigh,omitempty"`
	ThresholdLow   *int    `json:"thresholdLow,omitempty"`
	FriendlyName   *string `json:"friendlyName,omitempty"`
	DumpFrequency  *int    `json:"dumpFrequency,omitempty"`
	SlopeStart     *int    `json:"slopeStart,omitempty"`
	SlopeStop      *int    `json:"slopeStop,omitempty"`
}

//ToLed convert map interface to Led object
func ToLed(val interface{}) (*Led, error) {
	var light Led
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &light)
	return &light, err
}

//ToLedConf convert map interface to Led object
func ToLedConf(val interface{}) (*LedConf, error) {
	var light LedConf
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &light)
	return &light, err
}

//ToLedSetup convert map interface to Led object
func ToLedSetup(val interface{}) (*LedSetup, error) {
	var light LedSetup
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &light)
	return &light, err
}

// ToJSON dump led struct
func (led Led) ToJSON() (string, error) {
	inrec, err := json.Marshal(led)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}

// ToJSON dump led setup struct
func (led LedSetup) ToJSON() (string, error) {
	inrec, err := json.Marshal(led)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}

//ToJSON dump struct in json
func (led LedConf) ToJSON() (string, error) {
	inrec, err := json.Marshal(led)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}
