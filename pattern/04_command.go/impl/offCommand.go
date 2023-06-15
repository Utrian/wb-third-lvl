package main

// Конкретная команда
type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}
