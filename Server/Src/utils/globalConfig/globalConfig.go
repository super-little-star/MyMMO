package globalConfig

import (
	"github.com/go-ini/ini"
	"log"
)

type config interface {
	Load(cfg *ini.File)
}

var ProjectCfg *ProjectConfig

type ProjectConfig struct {
	DevelopmentMode bool   `ini:"DevelopmentMode"`
	MaxPackageSize  uint32 `ini:"MaxPackageSize"`
}

func (pc *ProjectConfig) Load(cfg *ini.File) {
	LoadConfig(cfg, "project", ProjectCfg)
}

var MySQLCfg *MySQLConfig

type MySQLConfig struct {
	IP       string `ini:"IP"`
	Port     string `ini:"Port"`
	User     string `ini:"User"`
	Password string `ini:"Password"`
	DataBase string `ini:"DataBase"`
}

func (mc *MySQLConfig) Load(cfg *ini.File) {
	LoadConfig(cfg, "mysql", MySQLCfg)
}

//-------------------------------------------

func LoadConfig(cfg *ini.File, section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("[Error]config Match to [%s] is err : %v", section, err)
	}
	log.Printf("[Info]config Match to [%s] is success!!", section)
}

func LoadConfigs(cfg *ini.File, configs ...config) {
	for _, c := range configs {
		c.Load(cfg)
	}
}

func Init() {
	cfgF, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Fatalln("Fail to read config , err is :", err)
	}
	ProjectCfg = &ProjectConfig{}
	MySQLCfg = &MySQLConfig{}
	LoadConfigs(cfgF, ProjectCfg, MySQLCfg)
}
