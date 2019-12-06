package main

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
)

const (
	LAT_LONG_URL = "https://www.latlong.net/category/cities-76-15.html"
	LAT = 52.520
	LONG = 13.405
	FREQ_URL = "http://radiomap.eu/de/berlin"
)

type EUChannelMaker struct{}

func NewEUChannelMaker() EUChannelMaker {
	return EUChannelMaker{}
}

func (e EUChannelMaker) CreateChannelsWithDifferentLatLongitude() (*Channels, error) {
	channels := NewEmptyChannels()

	doc, err := e.fetchFromInternet(LAT_LONG_URL)
	if err != nil {
		return nil, err
	}

	areaInfos := doc.Find("table").First().Find("tr").Slice(1, goquery.ToEnd)

	areaInfos.Each(
		func(i int, selection *goquery.Selection) {
			info := selection.Find("td").Slice(1, goquery.ToEnd)
			latitude, err := strconv.ParseFloat(info.First().Text(), 32)
			if err != nil {
				panic(err)
			}
			longitude, err := strconv.ParseFloat(info.Next().Text(), 32)
			if err != nil {
				panic(err)
			}

			channels.Add(NewChannel(AM, ANALOG, DATANULL, float32(latitude), float32(longitude), 0, DATANULL, DATANULL, DATANULL))
			channels.Add(NewChannel(AM, ALL, DATANULL, float32(latitude), float32(longitude), 0, DATANULL, DATANULL, DATANULL))
			channels.Add(NewChannel(FM, ANALOG, DATANULL, float32(latitude), float32(longitude), 0, DATANULL, DATANULL, DATANULL))
			channels.Add(NewChannel(FM, ALL, DATANULL, float32(latitude), float32(longitude), 0, DATANULL, DATANULL, DATANULL))
			channels.Add(NewChannel(DAB, ALL, DATANULL, float32(latitude), float32(longitude), 0, DATANULL, DATANULL, DATANULL))
		})

	return channels, nil
}

func (e EUChannelMaker) CreateChannelsWithDifferentFrequency() (*Channels, error) {
	channels := NewEmptyChannels()

	doc, err := e.fetchFromInternet(FREQ_URL)
	if err != nil {
		return nil, err
	}

	areaInfos := doc.Find("table").First().Find("tr").Slice(1, goquery.ToEnd)

	areaInfos.Each(
		func(i int, selection *goquery.Selection) {
			info := selection.Find("td").Slice(1, goquery.ToEnd)
			latitude, err := strconv.ParseFloat(info.First().Text(), 32)
			if err != nil {
				panic(err)
			}
			longitude, err := strconv.ParseFloat(info.Next().Text(), 32)
			if err != nil {
				panic(err)
			}

			channels.Add(NewChannel(AM, ANALOG, DATANULL, float32(latitude), float32(longitude), 0, DATANULL, DATANULL, DATANULL))
			channels.Add(NewChannel(AM, ALL, DATANULL, float32(latitude), float32(longitude), 0, DATANULL, DATANULL, DATANULL))
			channels.Add(NewChannel(FM, ANALOG, DATANULL, float32(latitude), float32(longitude), 0, DATANULL, DATANULL, DATANULL))
			channels.Add(NewChannel(FM, ALL, DATANULL, float32(latitude), float32(longitude), 0, DATANULL, DATANULL, DATANULL))
			channels.Add(NewChannel(DAB, ALL, DATANULL, float32(latitude), float32(longitude), 0, DATANULL, DATANULL, DATANULL))
		})

	return channels, nil
}

func (e *EUChannelMaker) fetchFromInternet(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Error making request to server")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New("Status code is " + strconv.Itoa(res.StatusCode) + " expect 200")
	}

	doc, erro := goquery.NewDocumentFromReader(res.Body)
	if erro != nil {
		return nil, errors.New("Error fetching song")
	}
	return doc, nil
}
