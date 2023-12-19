package main

import (
	"flag"
	strategyMatrix "lab3/strategyMatrix"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	N   uint64  `yaml:"n" env-required:"true"`
	C   []int   `yaml:"c" env-required:"true"`
	Eps float64 `yaml:"eps" env-required:"true"`
}

func main() {
	cfgPath := flag.String("cfg", "./config.yml", "path to yaml config file")
	flag.Parse()
	cfg := loadCfg(*cfgPath)
	m := strategyMatrix.New(cfg.C, cfg.N)
	t := time.Now()
	res := m.Run(cfg.Eps)
	log.Println("Elapsed time (in seconds):", time.Since(t).Seconds())
	/* оптимальные стратегии */
	log.Println("Oптимaльныe cтpaтeгии диcпeтчepa:", m.P())
	log.Println("Oптимaльныe cтpaтeгии ВЦ:", m.Q())
	log.Println("Цена игры:", res)
	// log.Println((m.Max() + m.Min()) / 2)
}

func loadCfg(cfgPath string) *config {
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("config file is not found in the specified path: %s", cfgPath)
	}
	cfg := new(config)
	if err := cleanenv.ReadConfig(cfgPath, cfg); err != nil {
		log.Fatalf("cannot load config: %s", err)
	}
	cfg.validate()
	return cfg
}

func (c *config) validate() {
	if c.N == 0 {
		log.Fatal("количество элементарных машин должно быть положительным")
	}
	if len(c.C) != 3 {
		log.Fatal("указано некорректное количество платежей/штрафов")
	}
	if c.C[0] < 1 || c.C[0] > 3 {
		log.Fatal("некорректное значение платежа за использование одной машины в течение ед. времени")
	}
	if c.C[1] < 4 || c.C[1] > 6 {
		log.Fatal("некорректное значение штрафа в ед. времени за простой одной машины (i-j) машин")
	}
	if c.C[2] < 4 || c.C[2] > 6 {
		log.Fatal("некорректное значение штрафа за недостающие машины (j-i) машин")
	}
}
