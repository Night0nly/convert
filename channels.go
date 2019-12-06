package main

const LINE_BREAK = "\n"

type Channels []*Channel

func NewEmptyChannels() *Channels {
	return &Channels{}
}

func (c *Channels) Add(channel *Channel) {
	if channel == nil {
		return
	}
	*c = append(*c, channel)
}

func (c *Channels) AddAll(channels Channels) {
	if channels == nil {
		return
	}
	for _, channel := range channels{
		c.Add(channel)
	}
}

func (c *Channels) String() string {
	result := ""
	for _, channel := range *c {
		result += channel.String() + LINE_BREAK
	}
	return result
}