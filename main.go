package main

const (
	FILE_PATH = "C:\\projects\\Pioneer\\New export\\TEST\\TESTDATA\\TIMERSID_NBID.txt"
)

func main() {
	printer := NewChannelPrinter(FILE_PATH)
	maker := NewEUChannelMaker()
	controller := NewChannelController(printer, maker)
	err := controller.WriteDataFile()
	if err != nil {
		 panic(err)
	}
}
