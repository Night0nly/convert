package main

type ChannelController struct {
	Printer ChannelPrinter
	Maker   ChannelsMaker
}

func NewChannelController(printer ChannelPrinter, maker ChannelsMaker) *ChannelController {
	return &ChannelController{
		Printer: printer,
		Maker: maker,
	}
}

func (c *ChannelController) WriteDataFile() error {
	channels, err := c.Maker.CreateChannelsWithDifferentLatLongitude()
	if err != nil {
		return err
	}

	err = c.Printer.PrintToFile(channels)
	if err != nil {
		return err
	}

	return nil
}