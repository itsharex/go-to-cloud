package conf

import "sync"

var agentImage *string

var onceAgent sync.Once

func GetAgentImage() *string {
	if agentImage == nil {
		onceAgent.Do(func() {
			if agentImage == nil {
				j := getConf().Agent
				agentImage = &j.Image
			}
		})
	}
	return agentImage
}
