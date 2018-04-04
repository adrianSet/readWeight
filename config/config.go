package config

import (
	"flag"
	"github.com/larspensjo/config"
	"log"
)

var (
	configFile = flag.String("configfile", "config.ini", "General configuration file")
	M          = make(map[string]string)
)

const (
	SectionCom    = "COM"
	WeightMachine = "weight_machine"
)




func init() {
	flag.Parse()
	parseIni()

}

func parseIni() {

	cfg, err := config.ReadDefault(*configFile)

	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}

	if cfg.HasSection(SectionCom) {
		section, err := cfg.SectionOptions(SectionCom)
		if err == nil {
			for _, v := range section {
				options, err := cfg.String(SectionCom, v)
				if err == nil {
					M[v] = options
				}
			}
		}
	}
}