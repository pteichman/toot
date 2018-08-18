// Copyright (c) 2018 Peter Teichman

package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/context"

	mastodon "github.com/mattn/go-mastodon"
)

// config file is a json struct that can be unmarshaled to a mastodon.Config
//
// {
//   "Server": "https://mastodon.social/",
//   "ClientID": "from the web ui",
//   "ClientSecret": "from the web ui",
//   "AccessToken": "from the web ui"
// }

var (
	configFile = flag.String("c", "", "config file (json struct)")
	media      = flag.String("media", "", "path to media file")
)

func main() {
	flag.Parse()

	if *configFile == "" || len(flag.Args()) == 0 {
		flag.PrintDefaults()
		log.Fatal("Missing config or text")
	}

	bytes, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("Reading %s: %s", *configFile, err)
	}

	var conf mastodon.Config
	if err = json.Unmarshal(bytes, &conf); err != nil {
		log.Fatalf("Unmarshaling config: %s", err)
	}

	text := strings.Join(flag.Args(), " ")
	toot := mastodon.Toot{
		Status: text,
	}

	client := mastodon.NewClient(&conf)
	ctx := context.Background()

	if media != nil && *media != "" {
		//		buf, err := mastodon.Base64EncodeFileName(*media)
		//		if err != nil {
		//			log.Fatalf("Reading media: %s: %v", *media, err)
		//		}

		attachment, err := client.UploadMedia(ctx, *media)
		if err != nil {
			log.Fatalf("Uploading media: %v", err)
		}

		toot.MediaIDs = append(toot.MediaIDs, attachment.ID)
	}

	status, err := client.PostStatus(ctx, &toot)
	if err != nil {
		log.Fatalf("Posting toot: %s", err)
	}

	log.Println(status.URL)
}
