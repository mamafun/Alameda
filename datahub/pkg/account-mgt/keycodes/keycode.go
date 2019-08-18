package keycodes

import (
	"encoding/json"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	InternalLdap "github.com/containers-ai/alameda/internal/pkg/database/ldap"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"time"
)

var (
	scope                           = log.RegisterScope("account-mgt", "keycode", 0)
	KeycodeCliPath                  = defaultCliPath
	KeycodeDuration      int64      = defaultRefreshInterval
	KeycodeLicenseStatus            = ""
	KeycodeAesKey                   = []byte("")
	KeycodeTimestamp     int64      = 0
	KeycodeList          []*Keycode = nil
	KeycodeSummary       *Keycode   = nil
	KeycodeTM            time.Time
	InfluxConfig         *InternalInflux.Config
	LdapConfig           *InternalLdap.Config
)

type Keycode struct {
	Keycode          string           `json:"keycode"          example:"A5IMH-KBAFI-XTEDK-G4OQM-QMM67-4TEST"`
	KeycodeType      string           `json:"keycodeType"      example:"Regular/Trial"`
	KeycodeVersion   int              `json:"keycodeVersion"   example:"2"`
	ApplyTimestamp   int64            `json:"applyTimestamp"   example:"1546271999"`
	ExpireTimestamp  int64            `json:"expireTimestamp"  example:"1546271999"`
	LicenseState     string           `json:"licenseState"     example:"Valid/Invalid/Expired"`
	Registered       bool             `json:"registered"       example:"false"`
	Capacity         Capacity         `json:"capacity"         example:"capacity"`
	Functionality    Functionality    `json:"functionality"    example:"functionality"`
	Retention        Retention        `json:"retention"        example:"retention"`
	ServiceAgreement ServiceAgreement `json:"serviceAgreement" example:"service agreement"`
	Description      string           `json:"description"      example:"your-description"`
}

type Capacity struct {
	Users int `json:"users" example:"-1"`
	Hosts int `json:"hosts" example:"20"`
	Disks int `json:"disks" example:"200"`
}

type Functionality struct {
	Diskprophet bool `json:"diskprophet" example:"true"`
	Workload    bool `json:"workload"    example:"true"`
}

type Retention struct {
	ValidMonth int `json:"validMonth" example:"0"`
	Years      int `json:"years"      example:"1"`
}

type ServiceAgreement struct {
}

func NewKeycode(keycode string) *Keycode {
	key := Keycode{}
	if keycode != "" {
		err := json.Unmarshal([]byte(keycode), &key)
		if err != nil {
			scope.Errorf("failed to unmarshal keycode: %v", err)
			return nil
		}
	}
	return &key
}

func KeycodeInit(config *Config) error {
	KeycodeCliPath = config.CliPath
	KeycodeDuration = config.RefreshInterval
	KeycodeAesKey = config.AesKey
	InfluxConfig = config.InfluxDB
	LdapConfig = config.Ldap

	return nil
}
