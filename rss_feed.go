package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	/* It should fetch a feed from the given URL, and assuming nothing goes wrong
	return a filled-out RSSFeed struct.
	*/

	// Create request to be sent by client
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	params := RSSFeed{}
	err = xml.Unmarshal(dat, &params)
	if err != nil {
		return nil, err
	}
	params.Channel.Title = html.UnescapeString(params.Channel.Title)
	params.Channel.Description = html.UnescapeString(params.Channel.Description)

	// Iterate over the RSSFeed.Item embedded struct
	// and modify all of the embedded title and descriptions
	for i, item := range params.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		params.Channel.Item[i] = item
	}

	// Return a filled-out RSSFeed Struct
	return &params, nil
}
