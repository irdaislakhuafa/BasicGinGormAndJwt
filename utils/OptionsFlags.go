package utils

import "flag"

type OptionsFlags struct {
	AppPort string

	DbUsername string
	DbPassword string
	DbName     string
	DbPort     string
	DbHost     string
}

func (option *OptionsFlags) EnableOptionFlags() {
	flag.StringVar(&option.AppPort, "app-port", "8080", "Put your custom app here!")
	flag.StringVar(&option.DbUsername, "db-username", "root", "Put your database username here!")
	flag.StringVar(&option.DbPassword, "db-password", "", "Put your database password here!")
	flag.StringVar(&option.DbName, "db-name", "test", "Put your database name here!")
	flag.StringVar(&option.DbPort, "db-port", "3306", "Put your database port here!")
	flag.StringVar(&option.DbHost, "db-host", "localhost", "Put your database port here!")
	flag.Parse()
}
