package main

import (
	"errors"
	"os"
)

type ChannelPrinter struct {
	FilePath string
}

func NewChannelPrinter(filePath string) ChannelPrinter {
	return ChannelPrinter{
		FilePath: filePath,
	}
}

func (c *ChannelPrinter) PrintToFile(channels *Channels) error {
	if channels == nil {
		return errors.New("channels is not exist")
	}

	f, err := os.Create(c.FilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(channels.String())
	if err != nil {
		return err
	}

	return nil
}
