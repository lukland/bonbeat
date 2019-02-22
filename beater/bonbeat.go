	package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/lucaslandry/bonbeat/config"
	"github.com/shirou/gopsutil/mem"

)

// Bonbeat configuration.
type Bonbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of bonbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Bonbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts bonbeat.
func (bt *Bonbeat) Run(b *beat.Beat) error {
	logp.Info("bonbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
		v, _ := mem.VirtualMemory()

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":    b.Info.Name,
				"cpu_used_percent": v.UsedPercent,
			},
		}
		bt.client.Publish(event)
		logp.Info("Event sent")

	}
}

// Stop stops bonbeat.
func (bt *Bonbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
