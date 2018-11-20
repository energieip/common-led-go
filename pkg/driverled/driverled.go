package driverled

import (
	"encoding/json"
)

const (
	DbName    = "status"
	TableName = "leds"

	UrlHello   = "setup/hello"
	UrlStatus  = "status/dump"
	UrlSetup   = "setup/config"
	UrlSetting = "update/settings"
)

//Led led driver representation
type Led struct {
	ID                string  `gorethink:"id,omitempty" json:"ID"`
	Mac               string  `gorethink:"mac" json:"mac"`
	Group             int     `gorethink:"group" json:"group"`
	Protocol          string  `gorethink:"protocol" json:"protocol"`
	Topic             string  `gorethink:"topic" json:"topic"`
	SwitchMac         string  `gorethink:"switch_mac" json:"switchMac"`
	IsConfigured      bool    `gorethink:"is_configured" json:"isConfigured"`
	Version           float32 `gorethink:"version" json:"version"`
	IsBleEnabled      bool    `gorethink:"is_ble_enabled" json:"isBleEnabled"`
	Temperature       int     `gorethink:"temperature" json:"temperature"`
	Error             int     `gorethink:"error" json:"error"`
	ResetNumbers      int     `gorethink:"reset_numbers" json:"resetNumbers"`
	InitialSetupDate  float64 `gorethink:"initial_setup_date" json:"initialSetupDate"`
	LastResetDate     float64 `gorethink:"last_reset_date" json:"lastResetDate"`
	IMax              int     `gorethink:"i_max" json:"iMax"`
	SlopeStart        int     `gorethink:"slope_start" json:"slopeStart"`
	SlopeStop         int     `gorethink:"slope_stop" json:"slopeStop"`
	Duration          float64 `gorethink:"duration" json:"duration"`
	Setpoint          int     `gorethink:"setpoint" json:"setpoint"`
	ThresoldLow       int     `gorethink:"thresold_low" json:"thresoldLow"`
	ThresoldHigh      int     `gorethink:"thresold_high" json:"thresoldHigh"`
	DaisyChainEnabled bool    `gorethink:"daisy_chain_enabled" json:"daisyChainEnabled"`
	DaisyChainPos     int     `gorethink:"daisy_chain_pos" json:"daisyChainPos"`
	DevicePower       int     `gorethink:"device_power" json:"devicePower"`
	Energy            float64 `gorethink:"energy" json:"energy"`
	VoltageLed        int     `gorethink:"voltage_led" json:"voltageLed"`
	VoltageInput      int     `gorethink:"voltage_input" json:"voltageInput"`
	LinePower         int     `gorethink:"line_power" json:"linePower"`
	TimeToAuto        int     `gorethink:"time_to_auto" json:"timeToAuto"`
	Auto              bool    `gorethink:"auto" json:"auto"`
	Watchdog          int     `gorethink:"watchdog" json:"watchdog"`
}

//LedSetup initial setup send by the server when the driver is authorized
type LedSetup struct {
	Mac          string `json:"mac"`
	IMax         int    `json:"iMax"`
	Group        *int   `json:"group"`
	Auto         *bool  `json:"auto"`
	Watchdog     *int   `json:"watchdog"`
	IsBleEnabled *bool  `json:"isBleEnabled"`
	ThresoldHigh *int   `json:"thresoldHigh"`
	ThresoldLow  *int   `json:"thresoldLow"`
}

//LedConf customizable configuration by the server
type LedConf struct {
	Mac          string `json:"mac"`
	Group        *int   `json:"group"`
	Setpoint     *int   `json:"setpoint"`
	Auto         *bool  `json:"auto"`
	Watchdog     *int   `json:"watchdog"`
	IsConfigured *bool  `json:"isConfigured"`
	IsBleEnabled *bool  `json:"isBleEnabled"`
	ThresoldHigh *int   `json:"thresoldHigh"`
	ThresoldLow  *int   `json:"thresoldLow"`
}

// ToMapInterface convert led struct in Map[string] interface{}
func (led Led) ToMapInterface() map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(led)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
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
