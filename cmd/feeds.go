package main

var releaseURLs = []string{
	"http://aakinshin.net/en/rss.xml",
	"http://aarvik.dk/rss/",
	"http://addyosmani.com/rss.xml",
	"http://ai.googleblog.com/feeds/posts/default",
	"http://antirez.com/rss",
	"http://artplustech.com/feed/",
	"http://artsy.github.io/feed",
	"http://bad-concurrency.blogspot.com/feeds/posts/default",
	"http://blog.8thcolor.com/feed.xml",
	"http://blog.andrewray.me/rss/",
	"http://blog.apps.npr.org/atom.xml",
	"http://blog.arkency.com/atom.xml",
	"http://blog.avenuecode.com/rss.xml",
	"http://blog.blundellapps.co.uk/feed/",
	"http://blog.claymcleod.io/atom.xml",
	"http://blog.cleancoder.com/atom.xml",
	"http://blog.codefx.org/feed/",
	"http://blog.codeship.com/feed/",
	"http://blog.fogus.me/feed/",
	"http://blog.honeybadger.io/feed.xml",
	"http://blog.jayfields.com/feeds/posts/default",
	"http://blog.joda.org/feeds/posts/default",
	"http://blog.joingrouper.com/rss",
	"http://blog.klipse.tech//feed.xml",
	"http://blog.lerner.co.il/feed/",
	"http://blog.mallow-tech.com/feed/",
	"http://blog.mandrill.com/feeds/all.atom.xml",
	"http://blog.memsql.com/feed/",
	"http://blog.pamelafox.org/feeds/posts/default",
	"http://blog.reverberate.org/feed.xml",
	"http://blog.rtwilson.com/feed/",
	"http://blog.sensible.io/rss",
	"http://blog.thislongrun.com/feeds/posts/default",
	"http://blog.venmo.com/hf2t3h4x98p5e13z82pl8j66ngcmry?format=RSS",
	"http://blog.vngrs.com/rss.xml",
	"http://blog.wittchen.biz.pl/feed.xml",
	"http://blogs.msdn.microsoft.com/pythonengineering/feed/",
	"http://code.flickr.net/feed/",
	"http://code.hireart.com/feed.xml",
	"http://code.hootsuite.com/rss",
	"http://codebeforethehorse.tumblr.com/rss",
	"http://dailytechvideo.com/feed/",
	"http://deborah-digges.github.io/atom.xml",
	"http://deliveroo.engineering/feed.xml",
	"http://development.wombatsecurity.com/feed.xml",
	"http://dtrace.org/blogs/bmc/feed/",
	"http://eng.rightscale.com/feed.xml",
	"http://eng.wealthfront.com/feed/",
	"http://engineering.blogfoster.com/rss/",
	"http://engineering.cerner.com/atom.xml",
	"http://engineering.chartbeat.com/feed/",
	"http://engineering.curalate.com/feed.xml",
	"http://engineering.grab.com/feed.xml",
	"http://engineering.hackerearth.com/rss",
	"http://engineering.harrys.com/feed.xml",
	"http://engineering.ifttt.com/feed.xml",
	"http://engineering.indeedblog.com/feed/",
	"http://engineering.khanacademy.org/rss.xml",
	"http://engineering.laterooms.com/rss/",
	"http://engineering.opensooq.com/feed/",
	"http://engineering.remind.com/feed.xml",
	"http://engineering.speedledger.com/feed/",
	"http://engineering.tripadvisor.com/rss",
	"http://engineering.vine.co/rss",
	"http://engineering.vinted.com//atom.xml",
	"http://engineering.wattpad.com/rss",
	"http://engineering.wingify.com/atom.xml",
	"http://engineroom.trackmaven.com/feeds/rss.xml",
	"http://fbrnc.net/blog.atom",
	"http://feedpress.me/jlongster",
	"http://feeds.feedburner.com/2ality",
	"http://feeds.feedburner.com/HighScalability",
	"http://feeds.feedburner.com/PivotalEngineeringJournal",
	"http://feeds.feedburner.com/buckblog",
	"http://feeds.feedburner.com/codinghorror",
	"http://feeds.feedburner.com/mishadoff",
	"http://feeds.feedburner.com/zumba_engineering",
	"http://feeds.hanselman.com/ScottHanselman",
	"http://feeds.regulargeek.com/RegularGeek",
	"http://feeds2.feedburner.com/patshaughnessy",
	"http://fuzzyblog.io//blog/feed.xml",
	"http://glennengstrand.info/blog/?feed=rss2",
	"http://huonw.github.io/blog/atom.xml",
	"http://iansommerville.com/systems-software-and-technology/feed/",
	"http://ieftimov.com/feed.xml",
	"http://intentmedia.com/feed/",
	"http://jakewharton.com/feed.xml",
	"http://jakeyesbeck.com/atom.xml",
	"http://jelv.is/blog/rss.xml",
	"http://johannesbrodwall.com/feed/",
	"http://lackingrhoticity.blogspot.com/feeds/posts/default",
	"http://lambda-the-ultimate.org/rss.xml",
	"http://lea.verou.me/feed/",
	"http://lg.io/feed.xml",
	"http://lifepluslinux.blogspot.com/feeds/posts/default",
	"http://loige.co/rss/",
	"http://lucumr.pocoo.org/feed.atom",
	"http://making.fiftythree.com/feed.xml",
	"http://manu.sporny.org/feed/",
	"http://masnun.com/feed",
	"http://matt.might.net/articles/feed.rss",
	"http://mattwarren.org/atom.xml",
	"http://mherman.org/feed.xml",
	"http://neopythonic.blogspot.com/feeds/posts/default",
	"http://nikolay.rocks/atom.xml",
	"http://nullprogram.com/feed/",
	"http://petersteinberger.com/atom.xml",
	"http://piotrpasich.com/feed/",
	"http://planet.mozilla.org/ateam/atom.xml",
	"http://planet.mozilla.org/releng/atom.xml",
	"http://preshing.com/feed",
	"http://prog21.dadgum.com/atom.xml",
	"http://radek.io/rss.xml",
	"http://raganwald.com/atom.xml",
	"http://redino.net/blog/feed/",
	"http://research.baidu.com/Publications",
	"http://rhodesmill.org/brandon/feed",
	"http://rocksdb.org/feed.xml",
	"http://rockthecode.io/feed/",
	"http://samsaffron.com/posts.rss",
	"http://semaphoreci.com/blog/engineering.xml",
	"http://sergeyzhuk.me/feed.xml",
	"http://stackabuse.com/rss/",
	"http://stdout.in/en/cat/all.rss",
	"http://sudhagar.com/feed.xml",
	"http://tech.adroll.com/feed.xml",
	"http://tech.finn.no/atom.xml",
	"http://tech.gc.com/atom.xml",
	"http://tech.gilt.com/rss",
	"http://tech.secretescapes.com/feed/",
	"http://tech.taskrabbit.com/feed.xml",
	"http://tech.wimdu.com/rss",
	"http://techblog.thescore.com/feed.xml",
	"http://teespring.engineering/index.xml",
	"http://tenderlovemaking.com/atom.xml",
	"http://teropa.info/blog/feed.xml",
	"http://thedailywtf.com/rss",
	"http://thelazylog.com/rss/",
	"http://upcoder.com/feed",
	"http://vanillajava.blogspot.com/feeds/posts/default",
	"http://wongatech.github.io/feed.xml",
	"http://www.aaronsw.com/2002/feeds/pgessays.rss",
	"http://www.afronski.pl/feed.xml",
	"http://www.billthelizard.com/feeds/posts/default",
	"http://www.blackbytes.info/rss",
	"http://www.born2data.com/feed_atom.xml",
	"http://www.brendangregg.com/blog/rss.xml",
	"http://www.catonmat.net/feed/",
	"http://www.codenameone.com/feed.xml",
	"http://www.evanjones.ca/index.rss",
	"http://www.evanmiller.org/news.xml",
	"http://www.jonkensy.com/feed/",
	"http://www.madetech.com/feed",
	"http://www.michaelgallego.fr/feed.xml",
	"http://www.mikeash.com/pyblog/rss.py",
	"http://www.nikola-breznjak.com/blog/feed/atom/",
	"http://www.norvig.com/rss-feed.xml",
	"http://www.practicallyefficient.com/feed.xml",
	"http://www.rudyhuyn.com/blog/feed/",
	"http://www.schibsted.pl/blog/feed/",
	"http://www.tjmaher.com/feeds/posts/default",
	"http://www.vertabelo.com/_rss/blog.xml",
	"http://www.wefearchange.org/feeds/all.atom.xml",
	"http://www.wilfred.me.uk/rss.xml",
	"http://www.windytan.com/feeds/posts/default",
	"http://yifan.lu/feed.xml",
	"https://0xadada.pub/feed.xml",
	"https://8thlight.com/blog/feed/atom.xml",
	"https://8thlight.com/insights/",
	"https://99designs.com/blog/engineering/",
	"https://99designs.com/tech-blog/feed.xml",
	"https://advancedweb.hu/",
	"https://advancedweb.hu/atom.xml",
	"https://aerotwist.com/blog/feed/",
	"https://airbnb.io/",
	"https://akrabat.com/feed/",
	"https://alanstorm.com/feed/feed.xml",
	"https://allegro.tech/feed.xml",
	"https://android-developers.googleblog.com/feeds/posts/default",
	"https://aphyr.com/posts.atom",
	"https://ariya.io/index.xml",
	"https://auth0.com/blog/rss.xml",
	"https://aws.amazon.com/blogs/?awsf.blog-master-category=*all&awsf.blog-master-learning-levels=*all&awsf.blog-master-industry=*all&awsf.blog-master-analytics-products=*all&awsf.blog-master-artificial-intelligence=*all&awsf.blog-master-aws-cloud-financial-management=*all&awsf.blog-master-blockchain=*all&awsf.blog-master-business-applications=*all&awsf.blog-master-compute=*all&awsf.blog-master-customer-enablement=*all&awsf.blog-master-customer-engagement=*all&awsf.blog-master-database=*all&awsf.blog-master-developer-tools=*all&awsf.blog-master-devops=*all&awsf.blog-master-end-user-computing=*all&awsf.blog-master-mobile=*all&awsf.blog-master-iot=*all&awsf.blog-master-management-governance=*all&awsf.blog-master-media-services=*all&awsf.blog-master-migration-transfer=*all&awsf.blog-master-migration-solutions=*all&awsf.blog-master-networking-content-delivery=*all&awsf.blog-master-programming-language=*all&awsf.blog-master-sector=*all&awsf.blog-master-security=*all&awsf.blog-master-storage=*all", //nolint:lll
	"https://aws.amazon.com/blogs/aws/feed/",
	"https://azure.microsoft.com/en-us/blog/",
	"https://backtrace.io/feed/",
	"https://bandcamptech.wordpress.com/feed/",
	"https://begriffs.com/atom.xml",
	"https://benchling.engineering/",
	"https://benchling.engineering/feed",
	"https://benmccormick.org/feed.json",
	"https://binary-studio.com/blog/rss",
	"https://bjornjohansen.no/feed",
	"https://blog.algolia.com/feed/",
	"https://blog.asana.com/feed/",
	"https://blog.avenuecode.com/",
	"https://blog.babbel.com/en/feed/",
	"https://blog.blakeerickson.com/feed.xml",
	"https://blog.booking.com/",
	"https://blog.chaps.io/feed.xml",
	"https://blog.chef.io/rss",
	"https://blog.cloudera.com/feed/",
	"https://blog.cloudflare.com/rss/",
	"https://blog.codelitt.com/",
	"https://blog.coursera.org/",
	"https://blog.cryptographyengineering.com/feed/",
	"https://blog.cugu.eu/index.xml",
	"https://blog.developer.bazaarvoice.com/",
	"https://blog.developer.bazaarvoice.com/feed/",
	"https://blog.digitalocean.com/rss/",
	"https://blog.discordapp.com/feed",
	"https://blog.docker.com/feed/",
	"https://blog.engineyard.com/feed.xml",
	"https://blog.evantahler.com/feed",
	"https://blog.faraday.io/rss/",
	"https://blog.fedecarg.com/feed/",
	"https://blog.filippo.io/rss/",
	"https://blog.getbootstrap.com/feed.xml",
	"https://blog.github.com/feed.xml",
	"https://blog.gojekengineering.com/feed",
	"https://blog.golang.org/feed.atom",
	"https://blog.hasura.io/feed",
	"https://blog.heroku.com/engineering",
	"https://blog.heroku.com/engineering/feed",
	"https://blog.hypriot.com/index.xml",
	"https://blog.imaginea.com/feed/",
	"https://blog.imgur.com/feed/",
	"https://blog.jessfraz.com/index.xml",
	"https://blog.jetbrains.com/kotlin/feed/",
	"https://blog.jgrossi.com/feed/",
	"https://blog.jooq.org/feed/",
	"https://blog.monstermuffin.org/feed/",
	"https://blog.moove-it.com/rss",
	"https://blog.nelhage.com/atom.xml",
	"https://blog.novoda.com/rss/",
	"https://blog.octo.com/en/feed/",
	"https://blog.pchudzik.com/index.xml",
	"https://blog.rapidapi.com/feed/",
	"https://blog.rinesi.com/feed/",
	"https://blog.risingstack.com/rss/",
	"https://blog.rust-lang.org/feed.xml",
	"https://blog.scrapinghub.com/rss.xml",
	"https://blog.shazam.com/feed",
	"https://blog.siftscience.com/feed/",
	"https://blog.sketchapp.com/feed",
	"https://blog.sleeplessbeastie.eu/feed.xml",
	"https://blog.soshace.com/en/feed/",
	"https://blog.sqreen.io/feed/",
	"https://blog.takipi.com/feed/",
	"https://blog.timescale.com/feed",
	"https://blog.twitter.com/engineering/en_us",
	"https://blog.twitter.com/engineering/feed",
	"https://blog.versioneye.com/feed/",
	"https://blog.wearewizards.io/all.atom.xml",
	"https://blogs.dropbox.com/tech/feed/",
	"https://blogs.janestreet.com/feed.xml",
	"https://blogs.msdn.microsoft.com/dotnet/feed/",
	"https://blogs.msdn.microsoft.com/oldnewthing/feed",
	"https://blogs.nvidia.com/feed/",
	"https://blogs.windows.com/msedgedev/rss",
	"https://bohops.com/feed/",
	"https://brendaneich.com/feed/",
	"https://brooker.co.za/blog/rss.xml",
	"https://buildingvts.com/feed",
	"https://bytes.swiggy.com/feed",
	"https://capgemini.github.io/",
	"https://capgemini.github.io/feed.xml",
	"https://carlosbecker.com/index.xml",
	"https://cloud.blog.csc.fi/feeds/posts/default",
	"https://code.blender.org/rss",
	"https://code.dblock.org/feed.xml",
	"https://code.fb.com/feed/",
	"https://code.oursky.com/feed/",
	"https://codeascraft.com/feed/",
	"https://codeblog.jonskeet.uk/feed/",
	"https://codewithoutrules.com/atom.xml",
	"https://codewithstyle.info/feed/",
	"https://computer.forensikblog.de/en/atom.xml",
	"https://conductofcode.io/feed.xml",
	"https://convox.com/blog/rss.xml",
	"https://crowdfire.engineering/feed",
	"https://crypt.codemancers.com/index.xml",
	"https://crystal-lang.org/feed.xml",
	"https://danluu.com/atom.xml",
	"https://data-artisans.com/feed",
	"https://databricks.com/feed",
	"https://dave.cheney.net/feed",
	"https://davidwalsh.name/feed/atom",
	"https://deanhume.com/rss/",
	"https://deezer.io/",
	"https://deezer.io/feed",
	"https://deliveroo.engineering/",
	"https://dev.firmafon.dk/blog/feed.xml",
	"https://devblog.coolblue.nl/feed/",
	"https://devblog.kogan.com/blog?format=RSS",
	"https://devblog.songkick.com/",
	"https://devdactic.com/feed/",
	"https://developer.apple.com/swift/blog/news.rss",
	"https://developer.atlassian.com/blog/feed.xml",
	"https://developer.here.com/blog/feed",
	"https://developer.ibm.com/dwblog/feed/",
	"https://developer.okta.com/feed.xml",
	"https://developer.salesforce.com/blogs",
	"https://developer.salesforce.com/blogs/feed",
	"https://developerblog.zendesk.com/feed",
	"https://developers.livechatinc.com/blog/rss",
	"https://developers.redhat.com/blog/feed/atom/",
	"https://developers.soundcloud.com/blog.rss",
	"https://developers.soundcloud.com/blog/",
	"https://domenicoluciani.com/feed.xml",
	"https://dotdev.co/feed/",
	"https://dragan.rocks/feed.xml",
	"https://drewdevault.com/feed.xml",
	"https://drivy.engineering/feed.xml",
	"https://dropbox.tech/",
	"https://eaf4.com/rss",
	"https://elegantcode.com/feed/",
	"https://eli.thegreenplace.net/feeds/all.atom.xml",
	"https://eng.datafox.com/feed.xml",
	"https://eng.localytics.com/rss/",
	"https://eng.lyft.com/",
	"https://eng.lyft.com/feed",
	"https://eng.uber.com/feed/",
	"https://engblog.nextdoor.com/",
	"https://engblog.nextdoor.com/feed",
	"https://engineering.atspotify.com/",
	"https://engineering.bittorrent.com/feed/",
	"https://engineering.brandwatch.com/rss/",
	"https://engineering.clever.com/",
	"https://engineering.clever.com/rss",
	"https://engineering.coinbase.com/feed",
	"https://engineering.creditkarma.com/feed",
	"https://engineering.doximity.com/feed",
	"https://engineering.fb.com/",
	"https://engineering.fb.com/feed/",
	"https://engineering.foursquare.com/feed",
	"https://engineering.giphy.com/rss",
	"https://engineering.gosquared.com/feed",
	"https://engineering.grab.com/",
	"https://engineering.groupon.com/feed/",
	"https://engineering.gusto.com/",
	"https://engineering.gusto.com/rss/",
	"https://engineering.hashnode.com/rss.xml",
	"https://engineering.haus.com/feed",
	"https://engineering.imvu.com/feed/",
	"https://engineering.intercom.io/feed.xml",
	"https://engineering.linecorp.com/",
	"https://engineering.linecorp.com/en/blog/rss2",
	"https://engineering.linkedin.com/blog",
	"https://engineering.mixmax.com/rss/",
	"https://engineering.mixpanel.com/",
	"https://engineering.mixpanel.com/feed/",
	"https://engineering.panoramaed.com/feed/",
	"https://engineering.prezi.com/",
	"https://engineering.prezi.com/feed",
	"https://engineering.riotgames.com/rss.xml",
	"https://engineering.semantics3.com/feed",
	"https://engineering.squarespace.com/blog?format=RSS",
	"https://engineering.tumblr.com/rss",
	"https://engineering.universe.com/feed",
	"https://engineering.upgrad.com/feed",
	"https://engineering.vena.io/rss/",
	"https://engineering.webengage.com/feed/",
	"https://engineering.wingify.com/",
	"https://engineering.zalando.com/",
	"https://engineering.zenefits.com/feed/",
	"https://engineering.zomato.com/rss",
	"https://engineeringblog.yelp.com/",
	"https://engineeringblog.yelp.com/feed.xml",
	"https://engineroom.settled.co.uk/feed",
	"https://engineroom.teamwork.com/feed",
	"https://enoent.fr/atom.xml",
	"https://envoy.engineering/feed",
	"https://ericlippert.com/feed/",
	"https://erikrunyon.com/feed.xml",
	"https://evanhahn.com/feed.xml",
	"https://evernote.com/blog/feed/",
	"https://evilmartians.com/chronicles.atom",
	"https://eviltrout.com/feed.xml",
	"https://facebook.github.io/react-native/blog/atom.xml",
	"https://feed.laravel-news.com/",
	"https://feeds.feedburner.com/GiantRobotsSmashingIntoOtherGiantRobots",
	"https://feeds.feedburner.com/JohnResig",
	"https://feeds.feedburner.com/juristrumpflohner",
	"https://feeds.feedburner.com/paul-irish",
	"https://feeds.feedburner.com/philipwalton",
	"https://feeds.feedburner.com/ponyfoo",
	"https://firstdoit.com/feed",
	"https://freeletics.engineering/feed.xml",
	"https://fullstack.info/feed/",
	"https://galois.com/feed/",
	"https://github.blog/category/engineering/",
	"https://githubengineering.com/atom.xml",
	"https://glebbahmutov.com/blog/atom.xml",
	"https://gocardless.com/blog/atom.xml",
	"https://godaddy.com/engineering/feed.xml",
	"https://grafana.com/blog/blog/index.xml",
	"https://hacks.mozilla.org/feed/",
	"https://haptik.ai/tech/feed/",
	"https://hashrocket.com/blog.rss",
	"https://haydenjames.io/rss",
	"https://heapanalytics.com/blog/category/engineering/feed",
	"https://henrikwarne.com/feed/",
	"https://hookrace.net/blog/feed/",
	"https://hyegar.com/rss.xml",
	"https://idea.popcount.org/rss.xml",
	"https://idiosyncratic-ruby.com/feed.xml",
	"https://idontgetoutmuch.wordpress.com/feed/",
	"https://infrequently.org/feed/",
	"https://instagram-engineering.com/",
	"https://instagram-engineering.com/feed",
	"https://ipfs.io/blog/index.xml",
	"https://iridakos.com/feed.xml",
	"https://ivanursul.com/feed.xml",
	"https://jack.ofspades.com/rss/index.html",
	"https://jaketrent.com/index.xml",
	"https://jeremykun.com/feed/",
	"https://jerrygamblin.com/feed/",
	"https://jes.al/atom.xml",
	"https://jobandtalent.engineering/feed",
	"https://jollygoodcode.github.io/atom.xml",
	"https://joshtronic.com/atom.xml",
	"https://jtreminio.com/atom.xml",
	"https://jvns.ca/atom.xml",
	"https://kev.inburke.com/feed/",
	"https://kickstarter.engineering/feed",
	"https://kinvolk.io/blog/index.xml",
	"https://kolosek.com/rss/",
	"https://lab.getbase.com/feed/",
	"https://labs.spotify.com/feed/",
	"https://lambda.blinkit.com/",
	"https://lambda.grofers.com/feed",
	"https://latacora.micro.blog/feed.xml",
	"https://liveramp.com/engineering/feed/",
	"https://martinfowler.com/feed.atom",
	"https://maryrosecook.com/blog/feed.xml",
	"https://matt.aimonetti.net/atom.xml",
	"https://medium.com/bbc-product-technology",
	"https://medium.com/feed/@Pinterest_Engineering",
	"https://medium.com/feed/@SkyscannerEng",
	"https://medium.com/feed/@dschmidt1992",
	"https://medium.com/feed/@elbrujohalcon",
	"https://medium.com/feed/@kirill_shevch",
	"https://medium.com/feed/airbnb-engineering",
	"https://medium.com/feed/bbc-design-engineering",
	"https://medium.com/feed/better-practices",
	"https://medium.com/feed/blablacar-tech",
	"https://medium.com/feed/criteo-labs",
	"https://medium.com/feed/dailyjs",
	"https://medium.com/feed/doordash-blog/tagged/engineering",
	"https://medium.com/feed/engineering-housing",
	"https://medium.com/feed/expedia-group-tech",
	"https://medium.com/feed/feedzaitech",
	"https://medium.com/feed/helpshift-engineering",
	"https://medium.com/feed/homeaway-tech-blog",
	"https://medium.com/feed/javascript-scene",
	"https://medium.com/feed/jettech",
	"https://medium.com/feed/jobteaser-dev-team",
	"https://medium.com/feed/netflix-techblog",
	"https://medium.com/feed/postmates-blog/tagged/engineering",
	"https://medium.com/feed/retailmenot-engineering",
	"https://medium.com/feed/shyp-engineering",
	"https://medium.com/feed/square-corner-blog",
	"https://medium.com/feed/strava-engineering",
	"https://medium.com/feed/twitch-news/tagged/engineering",
	"https://medium.com/feed/unexpected-token",
	"https://medium.com/feed/vevo-engineering",
	"https://medium.com/feed/walmartlabs",
	"https://medium.com/feed/wemake-services",
	"https://medium.com/feed/yammer-engineering",
	"https://medium.com/feed/yld-engineering-blog",
	"https://medium.com/feed/zendesk-engineering",
	"https://medium.com/feed/zoosk-engineering",
	"https://medium.engineering/",
	"https://medium.engineering/feed",
	"https://meowni.ca/atom.xml",
	"https://mesosphere.com/feed/",
	"https://mirocupak.com/feed.xml",
	"https://muffinman.io/atom.xml",
	"https://multithreaded.stitchfix.com/feed.xml",
	"https://muratbuffalo.blogspot.com/feeds/posts/default?alt=rss",
	"https://murze.be/feed",
	"https://natashatherobot.com/rss",
	"https://nativeguru.wordpress.com/feed/",
	"https://netflixtechblog.com/?gi=b888aca8d907",
	"https://neverfriday.com/feed/",
	"https://nickcraver.com/blog/feed.xml",
	"https://nickdesaulniers.github.io/atom.xml",
	"https://nordicapis.com/feed/",
	"https://nshipster.com/feed.xml",
	"https://ocramius.github.io/atom.xml",
	"https://oleb.net/blog/atom.xml",
	"https://open.nytimes.com/feed",
	"https://os.phil-opp.com/atom.xml",
	"https://petr-mitrichev.blogspot.com/feeds/posts/default",
	"https://postmarkapp.com/blog/feed.atom",
	"https://product.canva.com/feed.xml",
	"https://product.hubspot.com/blog/rss.xml",
	"https://quoraengineering.quora.com/",
	"https://rachelbythebay.com/w/atom.xml",
	"https://radimrehurek.com/feed/",
	"https://reactjs.org/feed.xml",
	"https://reactjsnews.com/feed.xml",
	"https://realm.io/feed.xml",
	"https://realpython.com/atom.xml",
	"https://research.facebook.com/publications/",
	"https://research.google/pubs/?area=distributed-systems-and-parallel-computing",
	"https://rhettinger.wordpress.com/feed/",
	"https://ruslanspivak.com/feeds/all.atom.xml",
	"https://schakko.de/feed",
	"https://security.googleblog.com/feeds/posts/default",
	"https://segment.com/blog/atom.xml",
	"https://semaphoreci.com/",
	"https://semaphoreci.com/community/tutorials.atom",
	"https://serverless.com/blog/feed.xml",
	"https://shopify.engineering/blog.atom",
	"https://slack.engineering/",
	"https://slack.engineering/feed",
	"https://snook.ca/jonathan/index.rdf",
	"https://snyk.io/blog/feed.xml",
	"https://software.intel.com/en-us/blogs/feed",
	"https://sourcecode.entelo.com/feed.xml",
	"https://spin.atomicobject.com/feed/",
	"https://stackoverflow.blog/engineering/feed/",
	"https://stackshare.io/featured-posts.atom",
	"https://steve-yegge.blogspot.com/feeds/posts/default",
	"https://stripe.com/au/guides/atlas/scaling-eng",
	"https://stripe.com/blog/engineering",
	"https://stripe.com/blog/feed.rss",
	"https://swizec.com/blog/feed",
	"https://tania.dev/rss.xml",
	"https://target.github.io/feed.xml",
	"https://tech.buzzfeed.com/",
	"https://tech.ebayinc.com/",
	"https://tech.findmypast.com/",
	"https://tech.findmypast.com/feed.xml",
	"https://tech.geoblink.com/feed/",
	"https://tech.gotinder.com/rss/",
	"https://tech.grammarly.com/feed.xml",
	"https://tech.instacart.com/",
	"https://tech.instacart.com/feed",
	"https://tech.just-eat.com/feed/",
	"https://tech.lendinghome.com/feed",
	"https://tech.nextroll.com/",
	"https://tech.olx.com/feed",
	"https://tech.pic-collage.com/feed",
	"https://tech.scrunch.com/blog/feeds/rss/",
	"https://tech.showmax.com/feed.xml",
	"https://tech.small-improvements.com/rss",
	"https://tech.transferwise.com/rss/",
	"https://tech.trivago.com/index.xml",
	"https://tech.wayfair.com/feed/",
	"https://tech.xing.com/feed",
	"https://techblog.appnexus.com/feed",
	"https://techblog.badoo.com/feed.xml",
	"https://techblog.commercetools.com/",
	"https://techblog.commercetools.com/feed",
	"https://techblog.king.com/feed/",
	"https://technology.condenast.com/feed/rss",
	"https://technology.skybettingandgaming.com/feed.xml",
	"https://thatthinginswift.com/index.xml",
	"https://thecodedself.github.io/feed.xml",
	"https://themodernlife.github.io/feed.xml",
	"https://toddmotto.com/feed.xml",
	"https://tomassetti.me/feed/",
	"https://tympanus.net/codrops/rss",
	"https://umbrella.cisco.com/blog/feed/",
	"https://una.im/feed.xml",
	"https://undocumentedmatlab.com/feed",
	"https://upday.github.io/feed.xml",
	"https://useyourloaf.com/blog/rss.xml",
	"https://vinted.engineering//",
	"https://vladmihalcea.com/feed/",
	"https://webuild.envato.com/atom.xml",
	"https://wecode.wepay.com/feed.xml",
	"https://www.9lessons.info/feeds/posts/default",
	"https://www.amitmerchant.com/feed.xml",
	"https://www.andrewcbancroft.com/feed/",
	"https://www.ardanlabs.com/blog/index.xml",
	"https://www.atlassian.com/engineering",
	"https://www.azavea.com/blog/category/software-development/rss",
	"https://www.benefitfocus.com/rss.xml",
	"https://www.bfilipek.com/feeds/posts/default",
	"https://www.bigeng.io/",
	"https://www.bigeng.io/rss/",
	"https://www.boxever.com/feed/",
	"https://www.canva.dev/blog/engineering/",
	"https://www.chenhuijing.com/feed.xml",
	"https://www.client9.com/index.xml",
	"https://www.cockroachlabs.com/blog/",
	"https://www.cockroachlabs.com/blog/index.xml",
	"https://www.codelitt.com/blog/rss",
	"https://www.codementor.io/tutorial/feed",
	"https://www.confluent.io/",
	"https://www.confluent.io/feed/",
	"https://www.cs.columbia.edu/~smb/blog/control/blog.xml",
	"https://www.ctl.io/developers/blog/rss",
	"https://www.darkcoding.net/feed/",
	"https://www.datchley.name/rss/",
	"https://www.dereuromark.de/feed/",
	"https://www.devroom.io/index.xml",
	"https://www.discovermeteor.com/feed.xml",
	"https://www.drivenbycode.com/feed/",
	"https://www.ebayinc.com/stories/blogs/tech/rss/",
	"https://www.eharmony.com/engineering/feed/",
	"https://www.elastic.co/blog/feed",
	"https://www.erlang-solutions.com/news.rss",
	"https://www.etsy.com/codeascraft",
	"https://www.eventbrite.com/engineering/feed/",
	"https://www.future-processing.com/blog/",
	"https://www.future-processing.pl/technical-blog/rss",
	"https://www.gajotres.net/feed/",
	"https://www.hashicorp.com/blog/feed.xml",
	"https://www.hostinger.com/blog/feed/",
	"https://www.igvita.com/feed/",
	"https://www.joelonsoftware.com/feed/",
	"https://www.johnwittenauer.net/rss/",
	"https://www.jointaro.com/blog/",
	"https://www.justinweiss.com/atom.xml",
	"https://www.maptiler.com/blog/feed/posts.xml",
	"https://www.mattcutts.com/blog/feed/",
	"https://www.metachris.com/index.xml",
	"https://www.michaelcrump.net/feed.xml",
	"https://www.miqu.me/atom.xml",
	"https://www.nateberkopec.com/feed.xml",
	"https://www.paypal-engineering.com/feed/",
	"https://www.previous.cloudbees.com/blog.xml",
	"https://www.pubnub.com/blog/feed/",
	"https://www.raizlabs.com/dev/feed/",
	"https://www.raywenderlich.com/rss",
	"https://www.rea-group.com/category/tech/feed/",
	"https://www.red-lang.org/feeds/posts/default",
	"https://www.reddit.com/r/RedditEng.atom",
	"https://www.reddit.com/r/golang/.rss",
	"https://www.redditinc.com/blog",
	"https://www.rosehosting.com/blog/feed/",
	"https://www.runtastic.com/blog/en/feed/",
	"https://www.sakib.ninja/rss/",
	"https://www.sharethis.com/feed/",
	"https://www.sitepoint.com/feed/",
	"https://www.stridenyc.com/blog/rss.xml",
	"https://www.surveymonkey.com/feed/",
	"https://www.theguardian.com/info/2023/aug/16/serverless-postgres-at-the-guardian",
	"https://www.theguardian.com/info/series/digital-blog/rss",
	"https://www.theguardian.com/info/series/engineering-blog",
	"https://www.thepolyglotdeveloper.com/blog/index.xml",
	"https://www.thoughtworks.com/rss/insights.xml",
	"https://www.thumbtack.com/engineering/feed/",
	"https://www.toptal.com/blog.rss",
	"https://www.twilio.com/blog/feed",
	"https://www.uber.com/blog/research/",
	"https://www.uber.com/en-AU/blog/engineering/",
	"https://www.uber.com/en-AU/blog/melbourne/engineering/",
	"https://www.yegor256.com/rss.xml",
	"https://www.zillow.com/engineering/rss",
	"https://wyeworks.com/blog/atom.xml",
	"https://yahooeng.tumblr.com/rss",
	"https://yurichev.com/blog/rss.xml",
	"https://zachholman.com/atom.xml",
	"https://zapier.com/engineering/",
	"https://zeemee.engineering/feed",
	"https://zendesk.engineering/",
	"https://zolmeister.com/feeds/posts/default",
	"https://zulily-tech.com/feed/",
}
