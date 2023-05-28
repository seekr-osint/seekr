package seekrd

import (
	"fmt"
	"log"
	"time"
)

func (instance *SeekrdInstance) SeekrdTicker() {
	instance.initialRun = true
	ticker := time.NewTicker(time.Duration(instance.Interval) * time.Minute)
	for range ticker.C {
		err := instance.Run()
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
	}
}

func (instance *SeekrdInstance) Run() error {
	var err error
	for _, service := range instance.Services {
		if service.Repeat || instance.initialRun { // run a service if it is a repeating service or run all if it is the initialRun
			log.Printf("Running Seekrd Service %s\n", service.Name)
			// Load the db
			err = instance.ApiConfig.LoadDBPointer()
			if err != nil {
				return err
			}

			// Modify the db
			instance.ApiConfig, err = service.Func(instance.ApiConfig)
			if err != nil {
				return err
			}

			// Save the db
			err = instance.ApiConfig.SaveDB()
			if err != nil {
				return err
			}
		}
	}
	instance.initialRun = false
	return nil
}
