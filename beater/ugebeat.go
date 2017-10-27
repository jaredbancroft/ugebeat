package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/jaredbancroft/ugebeat/config"
)

// Ugebeat - Define struct
type Ugebeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New - Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Ugebeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

// Run - Run beater
func (bt *Ugebeat) Run(b *beat.Beat) error {
	logp.Info("ugebeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":    b.Info.Name,
				"counter": counter,
				"ugeroot": bt.config.Ugeroot,
				"ugecell": bt.config.Ugecell,
			},
		}
		bt.client.Publish(event)
		logp.Info("Event sent")
		counter++
	}
}

//Stop - kill beater
func (bt *Ugebeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
