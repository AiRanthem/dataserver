package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type Config map[string]any

func (c Config) MustGet(key string) any {
	val, ok := c[key]
	if !ok {
		log.WithField("key", key).Fatalln("miss key")
	}
	return val
}

func (c Config) MustString(key string) string {
	val := c.MustGet(key)
	s, ok := val.(string)
	if !ok {
		log.WithField("key", key).Fatalln("cannot cast to string")
	}
	return s
}

func (c Config) MustInt(key string) int {
	val := c.MustGet(key)
	s, ok := val.(int)
	if !ok {
		log.WithField("key", key).Fatalln("cannot cast to int")
	}
	return s
}

type SectionHandler func(sec *ini.Section)

type KeyParser func(sec *ini.Section, key string) any

var ConfMap = make(map[string]Config)

var HandlerMap = map[string]SectionHandler{
	"xmu-daily-report": func(sec *ini.Section) {
		c := make(Config)
		AllStringHandlerUtil(sec, c, "mysql-ip", "mysql-db", "mysql-user", "mysql-passwd")
		AllIntHandlerUtil(sec, c, "mysql-port")
		ConfMap[sec.Name()] = c
	},
}

func StringKeyParser(sec *ini.Section, key string) any {
	return sec.Key(key).String()
}

func IntKeyParser(sec *ini.Section, key string) any {
	i, err := sec.Key(key).Int()
	if err != nil {
		log.WithError(err).WithField("key", key).Fatalln("fail to load config")
	}
	return i
}

func CheckAndParseSection(sec *ini.Section, c Config, parser KeyParser, keys ...string) {
	for _, key := range keys {
		if !sec.HasKey(key) {
			log.WithField("section", sec.Name()).WithField("key", key).Fatalln("config key miss")
		}
		c[key] = parser(sec, key)
	}
}

func AllStringHandlerUtil(sec *ini.Section, c Config, keys ...string) {
	CheckAndParseSection(sec, c, StringKeyParser, keys...)
}

func AllIntHandlerUtil(sec *ini.Section, c Config, keys ...string) {
	CheckAndParseSection(sec, c, IntKeyParser, keys...)
}

func CommonSectionHancler(sec *ini.Section) {
	c := make(Config)
	AllStringHandlerUtil(sec, c, sec.KeyStrings()...)
	ConfMap[sec.Name()] = c
}

func MustGetConfig(sec string) Config {
	config, ok := ConfMap[sec]
	if !ok {
		log.WithField("section", sec).Fatalln("try to load a config not exist")
	}
	return config
}

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.WithError(err).Fatalln("fail to read config file")
	}
	for _, sec := range cfg.Sections() {
		handler, ok := HandlerMap[sec.Name()]
		if !ok {
			handler = CommonSectionHancler
		}
		handler(sec)
		log.WithField("section", sec.Name()).Infoln("config section loaded")
	}
}
