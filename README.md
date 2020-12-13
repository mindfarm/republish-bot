A simple relay bot

This bot will watch a given set of RSS feeds and will (re)publish the
information found in the RSS item to the configured platforms.

As the bot is set right now it will watch the supplied RSS/atom feed and
republish new items to Twitter, saving those that have already been seen to the
local database (currently Postgres). Therefore the configuration information
described below is for that set-up.

#Configuration
```bash
export TWITTER_API_KEY="REDACTED"
export TWITTER_API_SECRET="REDACTED"
export TWITTER_ACCESS_TOKEN="REDACTED-REDACTED"
export TWITTER_ACCESS_TOKEN_SECRET="REDACTED"
export POSTGRES_DATABASE="user=bot dbname=REDACTED sslmode=disable password=REDACTED"
export RELEASE_URL="https://github.com/golang/tools/releases.atom"
```

#TODO
Create drivers that will allow the bot to publish to other platforms, such as
  - Slack
  - Usenet
  - Facebook
  - IRC
  - Direct mailing lists

Create other repository drivers, so that the bot can be used with other SQL or
NOSQL datastores.

Create the ability to read configuration from other sources, such as
  - HCF
  - YAML
  - JSON
