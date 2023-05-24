package discord

import (
	"github.com/hugolgst/rich-go/client"
	"time"
)

func Rich() error {
	err := client.Login("1110620773103697941")
	if err != nil {
		return err
	}
	timeNow := time.Now()
	err = client.SetActivity(client.Activity{
		State:      "Gathering data...",
		Details:    "Multi-purpose OSINT toolkit.",
		LargeImage: "seekr-logo",
		LargeText:  "seekr-osint",
		Buttons: []*client.Button{
			{
				Label: "GitHub",
				Url:   "https://github.com/seekr-osint/seekr",
			},
		},
		Timestamps: &client.Timestamps{
			Start: &timeNow,
		},
	})

	if err != nil {
		return err
	}
	return nil
}
