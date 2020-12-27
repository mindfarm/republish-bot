A simple relay bot

This bot will watch a given set of RSS feeds and will (re)publish the
information found in the RSS item to the configured platforms.

As the bot is set right now it will watch the hardcoded RSS/atom feeds and
republish new items to Twitter, saving those that have already been seen to the
local database (currently Postgres). Therefore the configuration information
described below is for that set-up.

#Configuration
```bash
# Postgres Credentials
export POSTGRES_DATABASE="user=REDACTED dbname=REDACTED sslmode=disable password=REDACTED"

# Twitter Credentials
export TWITTER_API_KEY="REDACTED"
export TWITTER_API_SECRET="REDACTED"
export TWITTER_ACCESS_TOKEN="REDACTED-REDACTED"
export TWITTER_ACCESS_TOKEN_SECRET="REDACTED"

# Reddit Credentials
export REDDIT_CLIENT_ID="REDACTED"
export REDDIT_CLIENT_SECRET="REDACTED"
export REDDIT_USERNAME="announce_bot" # change as necessary
export REDDIT_PASSWORD="REDACTED"

# Freenode credentials
export FREENODE_NICK="announce-bot"  # change as necessary
export FREENODE_PASSWORD="REDACTED"
```

#TODO
Create drivers that will allow the bot to publish to other platforms, such as
  - Slack
  - Usenet
  - Direct mailing lists

Create other repository drivers, so that the bot can be used with other SQL or
NOSQL datastores.

Create the ability to read configuration from other sources, such as
  - HCF
  - YAML
  - JSON

This code is used for [AGopls](https://twitter.com/AGopl://twitter.com/AGopls)
announce-bot on [Freenode](https://freenode.net/)
and [/u/announce_bot](https://www.reddit.com/user/announce_bot)
