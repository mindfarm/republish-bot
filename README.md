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

# Code Layout/Architecture
Because I often point people to this repository as example code and architecture I will
add here some information about the architecture.

The root of the project holds the business logic, and the definition of
interfaces that the business logic uses. This allows the code to use the
[Dependency Inversion principle](https://en.wikipedia.org/wiki/Dependency_inversion_principle) to ensure that the business logic is loosely coupled to the services that it uses.

The [platform](platform/) directory holds implementations of the Platform
interface (which is defined in publish.go). There are currently three
implementations, one each for twitter, irc (freenode), and reddit.

The implementations of the Updates Platform interfaces do not have
to be put in those directories, they are there purely because it's easy to
distribute this project, and to make it easy for people (future me!) to find
them and know their job.

The [monitor](monitor/) directory holds implementations of the Updates
interface (which is defined in updates.go). At this point in time there is only one implementation,
[rss](monitor/rss/), which provides code that will poll an rss endpoint for new
updates. The rss code will store in its own data repository all previously
retrieved news items.

Note: The rss logic defines an interface to describe how it accesses existing
information, and there is no reason that other datastores cannot be used.

[Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection) is
used in [main](cmd/main.go) to instantiate instances of the RSS and Platform
types, and passes them to the business logic, so that the business logic can do
its job.

The [Twelve factor app](https://12factor.net/config) recommendation for
configuration to be stored in environment variables is why the credentials for
the platforms are passed to the application the way that they are.

The Dependency Injection pattern means that the business logic never knows what it is
monitoring (or how), and what it is publishing to (again, or how). Adding a new
platform to publish to (eg. Facebook) or source to monitor (eg. a Websocket) is
simply a matter of writing code that implements the appropriate interface,
instantiating an instance of the type in main.go, and passing it to the business
logic.
