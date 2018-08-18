# toot
Command line tool for posting to Mastodon; good for bots or cron

# Setup

First, install the tool. There are binaries on the GitHub releases
page, or you can use go get:

    $ go get github.com/pteichman/toot

Create a new application in the Development section of your Mastodon
configuration. Give it "write:statuses" and "write:media" permissions.

This will give you a client key, client secret, and access token.

Then create a JSON config file containing those strings and the base
URL to your server instance. Here's an example:

```
{
  "Server": "https://mastodon.social/",
  "ClientID": "client key from application config",
  "ClientSecret": "client secret from application config",
  "AccessToken": "access token from application config"
}
```

# Usage

    $ toot -c <config file> "here is my toot"

Or if you wish to post with an attached media file:

    $ toot -c <config file> -media /path/to/image.jpg "here is my toot"
