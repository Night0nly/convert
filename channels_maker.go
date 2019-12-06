package main

type ChannelsMaker interface {
	CreateChannelsWithDifferentLatLongitude() (*Channels, error)
}
