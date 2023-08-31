package republishbot

import (
	"log"
	"time"
)

func Republish(platforms map[string]Platform, monitor Updates) {
	for {
		releases, err := monitor.GetReleases()
		if err != nil {
			log.Printf("WARNING GetReleases() returned error %v", err)
		}

		for idx := range releases {
			for k := range platforms {
				if err := platforms[k].PublishContent(releases[idx]); err != nil {
					log.Printf("WARNING PublishContent() returned error %v for %s", err, releases[idx])
				}
			}
		}
		// Sleep and check again in 30 seconds
		//nolint:gomnd
		time.Sleep(time.Second * 30)
	}
}
