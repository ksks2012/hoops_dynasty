package setting

import (
	"time"
)

type AppSettingS struct {
	RunMode      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	LogSavePath  string
	LogFileName  string
	LogFileExt   string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         []string
	SocketPath   string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// AttributeRange defines a range for generating random attributes
type AttributeRange struct {
	Min int
	Max int
}

type BasicSettingS struct {
	HeightRange AttributeRange
	WeightRange AttributeRange
	AgeRange    AttributeRange
}

type DefenseSettingS struct {
	PerimeterDefenseRange AttributeRange
	InteriorDefenseRange  AttributeRange
	BlockRange            AttributeRange
	StealRange            AttributeRange
	ReboundRange          AttributeRange
	DefensiveIQRange      AttributeRange
}

type MentalSettingS struct {
	CourtVisionRange          AttributeRange
	DefensivePositioningRange AttributeRange
	WorkRateRange             AttributeRange
	LeadershipRange           AttributeRange
	FocusRange                AttributeRange
}

type OffenseSettingS struct {
	ShootingRange         AttributeRange
	ShootingDistanceRange AttributeRange
	FinishingRange        AttributeRange
	PassingRange          AttributeRange
	DribblingRange        AttributeRange
}

type PhysicalSettingS struct {
	SpeedRange    AttributeRange
	StaminaRange  AttributeRange
	StrengthRange AttributeRange
	VerticalRange AttributeRange
	AgilityRange  AttributeRange
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}

	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
