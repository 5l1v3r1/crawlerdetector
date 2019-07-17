package crawlerdetector

var botConfig = []byte(`[
  {
      "name": "360Spider",
      "regex": "360Spider(-Image|-Video)?"
  },
  {
      "name": "Aboundexbot",
      "regex": "Aboundex"
  },
  {
      "name": "Acoon",
      "regex": "AcoonBot"
  },
  {
      "name": "AddThis.com",
      "regex": "AddThis\\.com"
  },
  {
      "name": "aHrefs Bot",
      "regex": "AhrefsBot"
  },
  {
      "name": "Alexa Crawler",
      "regex": "ia_archiver|alexabot|verifybot"
  },
  {
      "name": "Alexa Site Audit",
      "regex": "alexa site audit"
  },
  {
      "name": "Amorank Spider",
      "regex": "AmorankSpider"
  },
  {
      "name": "ApacheBench",
      "regex": "ApacheBench"
  },
  {
      "name": "Applebot",
      "regex": "Applebot"
  },
  {
      "name": "Arachni",
      "regex": "Arachni"
  },
  {
      "name": "Castro 2",
      "regex": "Castro 2, Episode Duration Lookup"
  },
  {
      "name": "Analytics SEO Crawler",
      "regex": "Curious George"
  },
  {
      "name": "archive.org bot",
      "regex": "archive\\.org_bot|special_archiver"
  },
  {
      "name": "Ask Jeeves",
      "regex": "Ask Jeeves/Teoma"
  },
  {
      "name": "Backlink-Check.de",
      "regex": "Backlink-Check\\.de"
  },
  {
      "name": "BacklinkCrawler",
      "regex": "BacklinkCrawler"
  },
  {
      "name": "Baidu Spider",
      "regex": "baiduspider(-image)?|baidu Transcoder|baidu.*spider"
  },
  {
      "name": "BazQux Reader",
      "regex": "BazQux"
  },
  {
      "name": "BingBot",
      "regex": "MSNBot|msrbot|bingbot|BingPreview|msnbot-(UDiscovery|NewsBlogs)|adidxbot"
  },
  {
      "name": "Blekkobot",
      "regex": "Blekkobot"
  },
  {
      "name": "BLEXBot Crawler",
      "regex": "BLEXBot(Test)?"
  },
  {
      "name": "Bloglovin",
      "regex": "Bloglovin"
  },
  {
      "name": "Blogtrottr",
      "regex": "Blogtrottr"
  },
  {
      "name": "Bountii Bot",
      "regex": "BountiiBot"
  },
  {
      "name": "Browsershots",
      "regex": "Browsershots"
  },
  {
      "name": "BUbiNG",
      "regex": "BUbiNG"
  },
  {
      "name": "Butterfly Robot",
      "regex": "(HTC)?[ _]Butterfly/"
  },
  {
      "name": "CareerBot",
      "regex": "CareerBot"
  },
  {
      "name": "ccBot crawler",
      "regex": "CCBot"
  },
  {
      "name": "Cliqzbot",
      "regex": "Cliqzbot"
  },
  {
      "name": "CloudFlare AMP Fetcher",
      "regex": "Cloudflare-AMP"
  },
  {
      "name": "CloudFlare Always Online",
      "regex": "CloudFlare-AlwaysOnline"
  },
  {
      "name": "Cốc Cốc Bot",
      "regex": "coccoc/"
  },
  {
      "name": "Collectd",
      "regex": "collectd"
  },
  {
      "name": "CommaFeed",
      "regex": "CommaFeed"
  },
  {
      "name": "CSS Certificate Spider",
      "regex": "CSS Certificate Spider"
  },
  {
      "name": "Datadog Agent",
      "regex": "Datadog Agent"
  },
  {
      "name": "Dataprovider",
      "regex": "Dataprovider"
  },
  {
      "name": "Daum",
      "regex": "Daum(oa)?[ /][0-9]"
  },
  {
      "name": "Dazoobot",
      "regex": "Dazoobot"
  },
  {
      "name": "Discobot",
      "regex": "discobot(-news)?"
  },
  {
      "name": "Domain Re-Animator Bot",
      "regex": "Domain Re-Animator Bot|support@domainreanimator.com"
  },
  {
      "name": "DotBot",
      "regex": "DotBot"
  },
  {
      "name": "DuckDuckGo Bot",
      "regex": "DuckDuck"
  },
  {
      "name": "Easou Spider",
      "regex": "EasouSpider"
  },
  {
      "name": "EMail Exractor",
      "regex": "EMail Exractor"
  },
  {
      "name": "evc-batch",
      "regex": "evc-batch"
  },
  {
      "name": "ExaBot",
      "regex": "Exabot(-Thumbnails|-Images)?|ExaleadCloudview"
  },
  {
      "name": "ExactSeek Crawler",
      "regex": "ExactSeek Crawler"
  },
  {
      "name": "Ezooms",
      "regex": "Ezooms"
  },
  {
      "name": "Facebook External Hit",
      "regex": "facebookexternalhit|facebookplatform"
  },
  {
      "name": "Feedbin",
      "regex": "Feedbin"
  },
  {
      "name": "FeedBurner",
      "regex": "FeedBurner"
  },
  {
      "name": "Feed Wrangler",
      "regex": "Feed Wrangler"
  },
  {
      "name": "Feedly",
      "regex": "(Meta)?Feedly(Bot|App)?"
  },
  {
      "name": "Feedspot",
      "regex": "Feedspot"
  },
  {
      "name": "Fever",
      "regex": "Fever/[0-9]"
  },
  {
      "name": "Flipboard",
      "regex": "FlipboardProxy|FlipboardRSS"
  },
  {
      "name": "Findxbot",
      "regex": "Findxbot"
  },
  {
      "name": "Genieo Web filter",
      "regex": "Genieo"
  },
  {
      "name": "Gigablast",
      "regex": "GigablastOpenSource"
  },
  {
      "name": "Gluten Free Crawler",
      "regex": "Gluten Free Crawler"
  },
  {
      "name": "Goo",
      "regex": "ichiro/mobile goo"
  },
  {
      "name": "Google Search Console",
      "regex": "Google Search Console"
  },
  {
      "name": "Google PageSpeed Insights",
      "regex": "Google Page Speed Insights"
  },
  {
      "name": "Google Partner Monitoring",
      "regex": "google_partner_monitoring"
  },
  {
      "name": "Google Structured Data Testing Tool",
      "regex": "Google-Structured-Data-Testing-Tool"
  },
  {
      "name": "Gmail Image Proxy",
      "regex": "via ggpht\\.com GoogleImageProxy"
  },
  {
      "name": "Seznam Email Proxy",
      "regex": "SeznamEmailProxy"
  },
  {
      "name": "Seznam Zbozi.cz",
      "regex": "Seznam-Zbozi-robot"
  },
  {
      "name": "Heureka Feed",
      "regex": "Heurekabot-Feed"
  },
  {
      "name": "ShopAlike",
      "regex": "ShopAlike"
  },
  {
      "name": "Googlebot",
      "regex": "Googlebot(-Mobile|-Image|-Video|-News)?|Feedfetcher-Google|Google-Test|Google-Site-Verification|Google Web Preview|AdsBot-Google(-Mobile)?|Google-Adwords-Instant|APIs-Google|Mediapartners-Google|Google.*/\\+/web/snippet|GoogleProducer|Google[ -]Publisher[ -]Plugin|Google-Shopping-Quality|Google-Adwords-DisplayAds|Google-Assess|Google-AdWords-Express|Google-speakr|Google-Read-Aloud"
  },
  {
      "name": "Heritrix",
      "regex": "heritrix"
  },
  {
      "name": "HubSpot",
      "regex": "HubSpot "
  },
  {
      "name": "HTTPMon",
      "regex": "HTTPMon"
  },
  {
      "name": "ICC-Crawler",
      "regex": "ICC-Crawler"
  },
  {
      "name": "IIS Site Analysis",
      "regex": "iisbot"
  },
  {
      "name": "IPS Agent",
      "regex": "ips-agent"
  },
  {
      "name": "IP-Guide Crawler",
      "regex": "IP-Guide\\.com"
  },
  {
      "name": "Kouio",
      "regex": "kouio"
  },
  {
      "name": "Larbin web crawler",
      "regex": "larbin"
  },
  {
      "name": "Lighthouse",
      "regex": "([A-z0-9]*)-Lighthouse"
  },
  {
      "name": "Linkdex Bot",
      "regex": "linkdexbot(-mobile)?|linkdex\\.com"
  },
  {
      "name": "LinkedIn Bot",
      "regex": "LinkedInBot"
  },
  {
      "name": "LTX71",
      "regex": "ltx71"
  },
  {
      "name": "Mail.Ru Bot",
      "regex": "Mail\\.RU(_Bot)?"
  },
  {
      "name": "Magpie-Crawler",
      "regex": "magpie-crawler"
  },
  {
      "name": "MagpieRSS",
      "regex": "MagpieRSS"
  },
  {
      "name": "masscan",
      "regex": "masscan"
  },
  {
      "name": "Mastodon Bot",
      "regex": "Mastodon/"
  },
  {
      "name": "Meanpath Bot",
      "regex": "meanpathbot"
  },
  {
      "name": "MetaJobBot",
      "regex": "MetaJobBot"
  },
  {
      "name": "MetaInspector",
      "regex": "MetaInspector"
  },
  {
      "name": "Mixrank Bot",
      "regex": "MixrankBot"
  },
  {
      "name": "MJ12 Bot",
      "regex": "MJ12bot"
  },
  {
      "name": "Mnogosearch",
      "regex": "Mnogosearch"
  },
  {
      "name": "MojeekBot",
      "regex": "MojeekBot"
  },
  {
      "name": "Munin",
      "regex": "munin"
  },
  {
      "name": "NalezenCzBot",
      "regex": "NalezenCzBot"
  },
  {
      "name": "Nagios check_http",
      "regex": "check_http/v"
  },
  {
      "name": "Netcraft Survey Bot",
      "regex": "Netcraft( Web Server Survey| SSL Server Survey|SurveyAgent)"
  },
  {
      "name": "netEstate",
      "regex": "netEstate NE Crawler"
  },
  {
      "name": "Netvibes",
      "regex": "Netvibes"
  },
  {
      "name": "NewsBlur",
      "regex": "NewsBlur .*(Fetcher|Finder)"
  },
  {
      "name": "NewsGator",
      "regex": "NewsGatorOnline"
  },
  {
      "name": "NLCrawler",
      "regex": "nlcrawler"
  },
  {
      "name": "Nmap",
      "regex": "Nmap Scripting Engine"
  },
  {
      "name": "Octopus",
      "regex": "Octopus [0-9]"
  },
  {
      "name": "Omgili bot",
      "regex": "omgili(?:bot)?"
  },
  {
      "name": "Openindex Spider",
      "regex": "OpenindexSpider"
  },
  {
      "name": "OpenLinkProfiler",
      "regex": "spbot"
  },
  {
      "name": "OpenWebSpider",
      "regex": "OpenWebSpider"
  },
  {
      "name": "Orange Bot",
      "regex": "OrangeBot|VoilaBot"
  },
  {
      "name": "PaperLiBot",
      "regex": "PaperLiBot"
  },
  {
      "name": "Phantomas",
      "regex": "phantomas/"
  },
  {
      "name": "PHP Server Monitor",
      "regex": "phpservermon"
  },
  {
      "name": "PocketParser",
      "regex": "PocketParser"
  },
  {
      "name": "PritTorrent",
      "regex": "PritTorrent"
  },
  {
      "name": "Picsearch bot",
      "regex": "psbot(-page)?"
  },
  {
      "name": "Pingdom Bot",
      "regex": "Pingdom\\.com"
  },
  {
      "name": "Quora Link Preview",
      "regex": "Quora Link Preview"
  },
  {
      "name": "RamblerMail Image Proxy",
      "regex": "RamblerMail"
  },
  {
      "name": "QuerySeekerSpider",
      "regex": "QuerySeekerSpider"
  },
  {
      "name": "Qwantify",
      "regex": "Qwantify"
  },
  {
      "name": "Rainmeter",
      "regex": "Rainmeter"
  },
  {
      "name": "Reddit Bot",
      "regex": "redditbot"
  },
  {
      "name": "Riddler",
      "regex": "Riddler"
  },
  {
      "name": "Rogerbot",
      "regex": "rogerbot"
  },
  {
      "name": "ROI Hunter",
      "regex": "ROI Hunter"
  },
  {
      "name": "SafeDNSBot",
      "regex": "SafeDNSBot"
  },
  {
      "name": "Scrapy",
      "regex": "Scrapy"
  },
  {
      "name": "Screaming Frog SEO Spider",
      "regex": "Screaming Frog SEO Spider"
  },
  {
      "name": "ScreenerBot",
      "regex": "ScreenerBot"
  },
  {
      "name": "Semrush Bot",
      "regex": "SemrushBot"
  },
  {
      "name": "Sensika Bot",
      "regex": "SensikaBot"
  },
  {
      "name": "SEOENGBot",
      "regex": "SEOENG(World)?Bot"
  },
  {
      "name": "SEOkicks-Robot",
      "regex": "SEOkicks-Robot"
  },
  {
      "name": "Seoscanners.net",
      "regex": "seoscanners\\.net"
  },
  {
      "name": "Skype URI Preview",
      "regex": "SkypeUriPreview"
  },
  {
      "name": "Seznam Bot",
      "regex": "SeznamBot|SklikBot|Seznam screenshot-generator"
  },
  {
      "name": "ShopWiki",
      "regex": "ShopWiki"
  },
  {
      "name": "SilverReader",
      "regex": "SilverReader"
  },
  {
      "name": "SimplePie",
      "regex": "SimplePie"
  },
  {
      "name": "SISTRIX Crawler",
      "regex": "SISTRIX Crawler"
  },
  {
      "name": "SiteSucker",
      "regex": "SiteSucker"
  },
  {
      "name": "Sixy.ch",
      "regex": "sixy.ch"
  },
  {
      "name": "Slackbot",
      "regex": "Slackbot|Slack-ImgProxy"
  },
  {
      "name": "Sogou Spider",
      "regex": "(Sogou (web|inst|Pic) spider)|New-Sogou-Spider"
  },
  {
      "name": "Soso Spider",
      "regex": "Sosospider|Sosoimagespider"
  },
  {
      "name": "sqlmap",
      "regex": "sqlmap/"
  },
  {
      "name": "SSL Labs",
      "regex": "SSL Labs"
  },
  {
      "name": "StatusCake",
      "regex": "StatusCake"
  },
  {
      "name": "Superfeedr Bot",
      "regex": "Superfeedr bot"
  },
  {
      "name": "Sparkler",
      "regex": "Sparkler/[0-9]"
  },
  {
      "name": "Spinn3r",
      "regex": "Spinn3r"
  },
  {
      "name": "Sputnik Bot",
      "regex": "Sputnik(Image)?Bot"
  },
  {
      "name": "Survey Bot",
      "regex": "SurveyBot"
  },
  {
      "name": "Tarmot Gezgin",
      "regex": "TarmotGezgin"
  },
  {
      "name": "TelegramBot",
      "regex": "TelegramBot"
  },
  {
      "name": "TLSProbe",
      "regex": "TLSProbe"
  },
  {
      "name": "TinEye Crawler",
      "regex": "TinEye-bot"
  },
  {
      "name": "Tiny Tiny RSS",
      "regex": "Tiny Tiny RSS"
  },
  {
      "name": "Trendiction Bot",
      "regex": "trendictionbot"
  },
  {
      "name": "TurnitinBot",
      "regex": "TurnitinBot"
  },
  {
      "name": "TweetedTimes Bot",
      "regex": "TweetedTimes Bot"
  },
  {
      "name": "Tweetmeme Bot",
      "regex": "TweetmemeBot"
  },
  {
      "name": "Twitterbot",
      "regex": "Twitterbot"
  },
  {
      "name": "UniversalFeedParser",
      "regex": "UniversalFeedParser"
  },
  {
      "name": "UkrNet Mail Proxy",
      "regex": "via secureurl\\.fwdcdn\\.com"
  },
  {
      "name": "Uptimebot",
      "regex": "Uptimebot"
  },
  {
      "name": "Uptime Robot",
      "regex": "UptimeRobot"
  },
  {
      "name": "URLAppendBot",
      "regex": "URLAppendBot"
  },
  {
      "name": "Vagabondo",
      "regex": "Vagabondo"
  },
  {
      "name": "VK Share Button",
      "regex": "vkShare; "
  },
  {
      "name": "Visual Site Mapper Crawler",
      "regex": "VSMCrawler"
  },
  {
      "name": "W3C CSS Validator",
      "regex": "Jigsaw"
  },
  {
      "name": "W3C I18N Checker",
      "regex": "W3C_I18n-Checker"
  },
  {
      "name": "W3C Link Checker",
      "regex": "W3C-checklink"
  },
  {
      "name": "W3C Markup Validation Service",
      "regex": "W3C_Validator|Validator.nu"
  },
  {
      "name": "W3C MobileOK Checker",
      "regex": "W3C-mobileOK"
  },
  {
      "name": "W3C Unified Validator",
      "regex": "W3C_Unicorn"
  },
  {
      "name": "Wappalyzer",
      "regex": "Wappalyzer"
  },
  {
      "name": "WebPageTest",
      "regex": "PTST/"
  },
  {
      "name": "WeSEE:Search",
      "regex": "WeSEE(:Search)?"
  },
  {
      "name": "WebbCrawler",
      "regex": "WebbCrawler"
  },
  {
      "name": "WebSitePulse",
      "regex": "websitepulse[+ ]checker"
  },
  {
      "name": "WordPress",
      "regex": "WordPress"
  },
  {
      "name": "Wotbox",
      "regex": "Wotbox"
  },
  {
      "name": "YaCy",
      "regex": "yacybot"
  },
  {
      "name": "Yahoo! Slurp",
      "regex": "Yahoo! Slurp|Yahoo!-AdCrawler"
  },
  {
      "name": "Yahoo! Link Preview",
      "regex": "Yahoo Link Preview|Yahoo:LinkExpander:Slingstone"
  },
  {
      "name": "Yahoo! Cache System",
      "regex": "YahooCacheSystem"
  },
  {
      "name": "Yandex Bot",
      "regex": "Yandex(SpravBot|ScreenshotBot|MobileBot|AccessibilityBot|ForDomain|Vertis|Market|Catalog|Calendar|Sitelinks|AdNet|Pagechecker|Webmaster|Media|Video|Bot|Images|Antivirus|Direct|Blogs|Favicons|ImageResizer|Verticals|News(links)?|Metrika|\\.Gazeta Bot)|YaDirectFetcher"
  },
  {
      "name": "Yeti/Naverbot",
      "regex": "Yeti"
  },
  {
      "name": "Youdao Bot",
      "regex": "YoudaoBot"
  },
  {
      "name": "Yourls",
      "regex": "YOURLS v[0-9]"
  },
  {
      "name": "Yunyun Bot",
      "regex": "YRSpider|YYSpider"
  },
  {
      "name": "zgrab",
      "regex": "zgrab"
  },
  {
      "name": "Zookabot",
      "regex": "Zookabot"
  },
  {
      "name": "ZumBot",
      "regex": "ZumBot"
  },
  {
      "name": "Yottaa Site Monitor",
      "regex": "YottaaMonitor"
  },
  {
      "name": "Yahoo Gemini",
      "regex": "Yahoo Ad monitoring.*yahoo-ad-monitoring-SLN24857.*"
  },
  {
      "name": "Outbrain",
      "regex": ".*Java.*outbrain"
  },
  {
      "name": "HubPages",
      "regex": "HubPages.*crawlingpolicy"
  },
  {
      "name": "Pinterest",
      "regex": "Pinterest/\\d\\.\\d.*www\\.pinterest\\.com.*"
  },
  {
      "name": "Site24x7 Website Monitoring",
      "regex": "Site24x7"
  },
  {
      "name": "Snapchat Proxy",
      "regex": "s~snapchat-proxy"
  },
  {
      "name": "Let's Encrypt Validation",
      "regex": "Let's Encrypt validation server"
  },
  {
      "name": "Grapeshot",
      "regex": "GrapeshotCrawler"
  },
  {
      "name": "Monitor.Us",
      "regex": "www\\.monitor\\.us"
  },
  {
      "name": "Catchpoint",
      "regex": "Catchpoint( bot)?"
  },
  {
      "name": "BitlyBot",
      "regex": "bitlybot"
  },
  {
      "name": "Zao",
      "regex": "Zao/"
  },
  {
      "name": "Lycos",
      "regex": "lycos"
  },
  {
      "name": "Inktomi Slurp",
      "regex": "Slurp"
  },
  {
      "name": "Speedy",
      "regex": "Speedy Spider"
  },
  {
      "name": "ScoutJet",
      "regex": "ScoutJet"
  },
  {
      "name": "NetResearchServer",
      "regex": "nrsbot|netresearch"
  },
  {
      "name": "Scooter",
      "regex": "scooter"
  },
  {
      "name": "Gigabot",
      "regex": "gigabot"
  },
  {
      "name": "Charlotte",
      "regex": "charlotte"
  },
  {
      "name": "Pompos",
      "regex": "Pompos"
  },
  {
      "name": "ichiro",
      "regex": "ichiro"
  },
  {
      "name": "PagePeeker",
      "regex": "PagePeeker"
  },
  {
      "name": "WebThumbnail",
      "regex": "WebThumbnail"
  },
  {
      "name": "Willow Internet Crawler",
      "regex": "Willow Internet Crawler"
  },
  {
      "name": "EmailWolf",
      "regex": "EmailWolf"
  },
  {
      "name": "NetLyzer FastProbe",
      "regex": "NetLyzer FastProbe"
  },
  {
      "name": "ADMantX",
      "regex": "AdMantX.*admantx\\.com"
  },
  {
      "name": "Server Density",
      "regex": "Server Density Service Monitoring.*"
  },
  {
      "name": "Generic Bot",
      "regex": "(A6-Indexer|nuhk|TsolCrawler|Yammybot|Openbot|Gulper Web Bot|grub-client|Download Demon|SearchExpress|Microsoft URL Control|borg|altavista|dataminr.com|tweetedtimes.com|TrendsmapResolver|teoma|blitzbot|oegp|furlbot|http%20client|polybot|htdig|mogimogi|larbin|scrubby|searchsight|seekbot|semanticdiscovery|snappy|vortex( Build)?|zeal|fast-webcrawler|converacrawler|dataparksearch|findlinks|BrowserMob|HttpMonitor|ThumbShotsBot|URL2PNG|ZooShot|GomezA|Google SketchUp|Read%20Later|Minimo|RackspaceBot)"
  },
  {
      "name": "Sentry Bot",
      "regex": "^sentry"
  },
  {
      "name": "Nutch-based Bot",
      "regex": "Nutch"
  },
  {
      "name": "Generic Bot",
      "regex": "[a-z0-9\\-_]*((cu|power[ _]|m[ _])?bot([ _]TAB|[ _]?5[0-9])?|crawler|crawl|checker|archiver|transcoder|spider)([^a-z]|$)"
  },
  {
      "name": "Curl",
      "regex": "[c|C]url"
  },
  {
      "name": "Website check",
      "regex": "wscheck"
  }
]
`)

///Mobile settings

var mobilesConfig = []byte(`[
  {
      "name": "Tunisie Telecom",
      "regex": "StarTrail TT[);/ ]"
  },
  {
      "name": "SFR",
      "regex": "StarShine|StarTrail|STARADDICT|StarText|StarNaute|StarXtrem|StarTab"
  },
  {
      "name": "HTC",
      "regex": "HTC|Sprint (?:APA|ATP)|ADR(910L)?[a-z0-9]+|NexusHD2|Amaze[ _]4G[);/ ]|(Desire|Sensation|Evo ?3D|IncredibleS|Wildfire|Butterfly)[ _]?([^;/]+) Build|(Amaze[ _]4G|One ?[XELSV\\+]+)[);/ ]|SPV E6[05]0|One M8|X525a|PG86100|PC36100|XV6975|PJ83100[);/ ]|2PYB2|0PJA10|0PJA2"
  },
  {
      "name": "Microsoft",
      "regex": "KIN\\.(One|Two)|RM-(?:1031|106[57]|109[02]|1096|1099|1109|1114|1127|1141|1154)|Microsoft; Lumia"
  },
  {
      "name": "Nokia",
      "regex": "Nokia|Lumia|Maemo RX|portalmmm/2\\.0 N7|portalmmm/2\\.0 NK|nok[0-9]+|Symbian.*\\s([a-z0-9]+)$|RX-51 N900|TA-[0-9]{4}"
  },
  {
      "name": "CnM",
      "regex": "CnM"
  },
  {
      "name": "RIM",
      "regex": "BB10;|BlackBerry|rim[0-9]+|PlayBook|STV100-[1234]|STH100-[12]|BBA100-[12]|BBB100-[1234567]|BBC100-1|BBD100-[126]|BBE100-[123456789]|BBF100-[123456789]|BBG100-1|BBH100-1"
  },
  {
      "name": "Palm",
      "regex": "(?:Pre|Pixi)/(\\d+)\\.(\\d+)|Palm|Treo|Xiino"
  },
  {
      "name": "HP",
      "regex": "TouchPad/\\d+\\.\\d+|hp-tablet|HP ?iPAQ|webOS.*P160U|HP Slate|HP [78]|Compaq [7|8]|HP; [^;/)]+"
  },
  {
      "name": "TiPhone",
      "regex": "TiPhone ?([a-z0-9]+)"
  },
  {
      "name": "Apple",
      "regex": "(?:iTunes-)?Apple[ _]?TV|(?:Apple-|iTunes-)?(?:iPad|iPhone)|iPh[0-9],[0-9]|CFNetwork"
  },
  {
      "name": "MicroMax",
      "regex": "(?:MicroMax[ \\-\\_]?[a-z0-9]+|Q327)|P70221 Build"
  },
  {
      "name": "3Q",
      "regex": "(AC0731B|AC1024C|AC7803C|BC9710AM|EL72B|LC0720C|LC0723B|LC0725B|LC0804B|LC0808B|LC0809B|LC0810C|LC0816C|LC0901D|LC1016C|MT0724B|MT0729B|MT0729D|MT0811B|OC1020A|RC0709B|RC0710B|RC0718C|RC0719H|RC0721B|RC0722C|RC0726B|RC0734H|RC0743H|RC0817C|RC1018C|RC1019G|RC1025F|RC1301C|RC7802F|RC9711B|RC9712C|RC9716B|RC9717B|RC9724C|RC9726C|RC9727F|RC9730C|RC9731C|TS0807B|TS1013B|VM0711A|VM1017A|RC0813C|QS9719D|QS9718C|QS9715F|QS1023H|QS0815C|QS0730C|QS0728C|QS0717D|QS0716D|QS0715C|MT7801C)"
  },
  {
      "name": "4Good",
      "regex": "4Good[ _]|(?:4Good)?(S450m [43]G|S555m 4G|S501m 3G|T700i_3G|Light A103)"
  },
  {
      "name": "AllCall",
      "regex": "AllCall_|Heat[34]|Hot ?[125][^;/]+ Build"
  },
  {
      "name": "Bravis",
      "regex": "Bravis|A501 Bright|NB(?:10[56]|751|7[145])|NP101"
  },
  {
      "name": "Acer",
      "regex": "acer|(ZTE BLADE |ImSmart )?a(?:101|110|2[10]0|211|50[10]|51[10]|70[10])[);/ ]|Android.*V3[67]0[);/ ]|Android.*Z1[23456]0 Build|Android.*Z5\\d{2} Build|Android.*T0[234678] Build|A1-830|A1-81[01]|A3-A[1234][01]|B1-7[1235678][01]|B1-7[23]3|B1-8[1235]0|B1-A71|B3-A[12]0|B3-A3[02]|E39 Build|S5[12]0 Build|DA[0-9]+HQ?L[);/ ]"
  },
  {
      "name": "Advan",
      "regex": "(i7U|S45E|5061|S50H|I7D|i55[KD]|i4U|S7D|S4Z|i5[KE]) Build"
  },
  {
      "name": "Advance",
      "regex": "Hollogram|HL6246|IntroTr3544|Tr3845"
  },
  {
      "name": "AGM",
      "regex": "AGM ([^/;\\)]+)(?: Build|\\))"
  },
  {
      "name": "Airness",
      "regex": "AIRNESS-([\\w]+)"
  },
  {
      "name": "Aiwa",
      "regex": "AW790|M300"
  },
  {
      "name": "Akai",
      "regex": "Akai[ _-]|Eco[ _]E2|Glory[ _](?:G[1235]|L[123]|O[125])|TAB-[79]8[03]0Q?|X6 METAL|AKTB-703MZ"
  },
  {
      "name": "Alcatel",
      "regex": "Alcatel|Alc[a-z0-9]+|One[ _]?Touch|idol3|(?:4003[AJ]|4009[ADEFIKMSX]|4013[DEJKMX]|4014[ADEKMX]|4015[ADNTX]|4016[ADX]|4017[ADEFSX]|4018[ADEFMX]|4024[EDX]|4027[ADNX]|4028[AEJS]|4032[ADEX]|4034[ADEFGX]|4035[ADXY]|4045[ADEX]|4047[ADFGNX]|4049[DEGMX]|4060[SW]|4114E|5009[AD]|5010[DEGSUX]|5011A|5012[DFG]|5015[ADEX]|5016[AXJ]|5017[ABDEOX]|5019D|5022[EDX]|5023[EF]|5025[DG]|5026[AD]|5027B|5033[AX]|5033D[ _]RU|5038[ADEX]|5041C|5042[ADEFGWX]|5044[ADGIKOPSTY]|5045[ADFGIJTXY]|5046[ADGIJSTUY]|5047[DIUY]|5049[EGSWZ]|5050[ASXY]|5051[ADEJMTWX]|5052[AD]|5054[ADNSWX]|5056[ADEGIJMNTUWX]|5057M|5058[AIY]|5059[ADXYZ]|5065[ADNWX]|5070D|5080[ADFQUX]|5085[ABCDGHIJNOQY]|5086[ADY]|5090[AIY]|5095[IKY]|5098[OS]|5099[ADYUI]|5116J|5145A|6016[ADEX]|6036[AXY]|6037[BKY]|6039[AHJKY]|6043[AD]|6044D|6045[BFIKOY]|6050[AFY]|6055[ABDHIKPUYZ]|6058[ADX]|6060S|6062W|6070K|7040[ADEFKRT]|7041[DX]|7043[AEKY]|7044[AX]|7045Y|7048[ASWX]|7055A|7070X|8030Y|8050[DEGX]|8063|9001[DIX]|9002X|9003[AX]|9024O|9005X|9007[ATX]|9008[ADIJNTUX]|9010X|9022X|9203A|A570BL|I213|I216X)[);/ ]|TIMXL"
  },
  {
      "name": "Allview",
      "regex": "Allview|A4You|A5_(?:Easy(?:_TM)?|Ready(?:_TM)?|Quad|Quad_Plus_TM)|A[56789]_Lite|A5Smiley|A6_Duo|AX4Nano|C6_Duo|E[23]_Living|E3_(?:Jump|Sign)|E4_Lite|M9_Connect|P(?:43_Easy|[5689]_Energy|6_Energy_Lite|[68]_Energy_Mini(?:_TM)?|(41|[458])_eMagic(?:_TM)?|[589](?:_)?Life(?:_TM)?|[567]_Lite(?:_TM)?|6_plus|[45678]_Pro|7_Seon|10_Style|7_Xtreme|6_Qmax|4_Quad)|V(?:[13]_Viper|1_Viper_I|2_Viper_(?:E|I|S|X|X_plus|Xe))|X(?:[1234]_Soul_Xtreme|[12345]_Soul|3_Soul_Lite|[24]Soul_Mini(?:_TM)?|4_Soul_Mini_S(?:_TM)?|[234]_Soul_Style(?:_TM)?|2_Soul_Style_Plus|2_Twin)|Viva_H801"
  },
  {
      "name": "altron",
      "regex": "AL-555|GI-626"
  },
  {
      "name": "AMGOO",
      "regex": "AM(350|355|40[257]|41[05]|450|50[89]|52[037]|51[58]|53[05])"
  },
  {
      "name": "Amoi",
      "regex": "Amoi|A862W"
  },
  {
      "name": "Ainol",
      "regex": "Numy|novo[0-9]"
  },
  {
      "name": "Archos",
      "regex": "(?:YL-)?Archos"
  },
  {
      "name": "Axxion",
      "regex": "Axxion ATAB-[0-9]+[);/ ]"
  },
  {
      "name": "MEU",
      "regex": "MEU ([a-z0-9]+) Build"
  },
  {
      "name": "Arnova",
      "regex": "arnova|ARCHM901|AN7CG2|AN7G2(DTE|I)?|AN7[BCDFH]?G3|A80KSC|AN8G2|AN8[BC]?G3|AN9G2I?|AN9G3|A101[BC]|AN10G2|AN10BG2(DT|I)?|AN10BG3(DT)?|AN10[CD]G3"
  },
  {
      "name": "ARRIS",
      "regex": "ARRIS"
  },
  {
      "name": "Ask",
      "regex": "ASK "
  },
  {
      "name": "Ark",
      "regex": "(ARK[_ -])?Benefit[_ -]([^/;]*) Build|EDGE A5HD|ICON (R40\\+|R45)|Impulse[ _](P[12])"
  },
  {
      "name": "Garmin-Asus",
      "regex": "Garmin-Asus|Garminfone"
  },
  {
      "name": "Asus",
      "regex": "Asus|Transformer|TF300T|Slider SL101|PadFone|ME302(?:C|KL)|ME301T|ME371MG|ME17(?:1|2V|3X)|(?:K0[01][0-9a-z]|P(?:00[18ACIL]|01[MTVWYZ]|01MA|01T_1|02[13478])(?: Build|\\))|X015D|X018D|X003|X00[7ABT]D|Z00D|Z00[MTY]D|Z01[7FGHKMR]D)[);/ ]|ZB(602|555)KL|ZC55[34]KL|ZA550KL|ZE(520|620|55[2345])KL"
  },
  {
      "name": "Audiovox",
      "regex": "Audiovox|CDM|UTS(?:TARCOM)?\\-|audio[a-z0-9\\-]+"
  },
  {
      "name": "AVH",
      "regex": "Excer_10_PRO"
  },
  {
      "name": "Avvio",
      "regex": "Avvio[ _]?([a-z0-9\\-]+)"
  },
  {
      "name": "Azumi Mobile",
      "regex": "Azumi[_ ]"
  },
  {
      "name": "Barnes \u0026 Noble",
      "regex": "Nook|BN[TR]V[0-9]+"
  },
  {
      "name": "BGH",
      "regex": "(Positivo )?BGH ([^/;\\)]+)(?: Build|\\))"
  },
  {
      "name": "Bitel",
      "regex": "Bitel[ _-]([^/;]+) Build|B(840[5789]|841[0456]|850[2346]|860[146]|9401|950[12345])"
  },
  {
      "name": "Blu",
      "regex": "BLU |(?:blu|Dash)[ _]([^/;]+) Build|Studio (5.5|View XL|Mega|C 8\\+8|C HD|C|G|Selfie LTE|Touch|M HD|M5 Plus|J[1258]|X|X8 HD)|Advance (4.0 ?[LM]|5.0(?: HD)?|A4)|ENERGY (DIAMOND|XL)|LIFE XL|Dash XL|PURE XL|Life One X2|GRAND 5.5 HD|R1 (HD|PLUS)|Tank Xtreme 5.0|Tank Xtreme Pro"
  },
  {
      "name": "Bluegood",
      "regex": "BLUEGOOD"
  },
  {
      "name": "Blackview",
      "regex": "(BV([2456789]000|(58|9[56])00)(?:[ _](?:PRO))?|(P10000(?:[ _](?:PRO))?)|omega[ _]pro|Alife[ _][PS]1|Heatwave|DM550)"
  },
  {
      "name": "Bluboo",
      "regex": "Bluboo|Xfire|Maya Build"
  },
  {
      "name": "bogo",
      "regex": "BO-(FRSP4|LFSPBS5)"
  },
  {
      "name": "Boway",
      "regex": "(?:sprd-)?(BOWAY)"
  },
  {
      "name": "Brondi",
      "regex": "(?:CENTURION|GLADIATOR| GLORY|LUXURY|SENSUELLE|VICTORY)(?:[ _-]?[2-6])?[);/ ]|Surfing Tab"
  },
  {
      "name": "Vivo",
      "regex": "((?:VIV-|BBG-)?vivo)|V1813BA|V1818T|V1731CA|V1809T|V1813T"
  },
  {
      "name": "Vinsoc",
      "regex": "(XA Pro) Build"
  },
  {
      "name": "Bird",
      "regex": "BIRD[\\-. _]([^;/]+)"
  },
  {
      "name": "Becker",
      "regex": "Becker-([a-z0-9]+)"
  },
  {
      "name": "Beeline",
      "regex": "Beeline"
  },
  {
      "name": "Beetel",
      "regex": "Beetel ([a-z0-9]+)"
  },
  {
      "name": "BenQ-Siemens",
      "regex": "BENQ-SIEMENS - ([a-z0-9]+)"
  },
  {
      "name": "BenQ",
      "regex": "BENQ(?:[ \\-])?([a-z0-9]+)"
  },
  {
      "name": "Blaupunkt",
      "regex": "Blaupunkt|Atlantis[_ ](?:1001A|1010A|A10\\.G402)|Discovery[_ ](?:111C|1000C|1001A?)|Endeavour[_ ](?:785|101[GLM]|1000|1001|101[03]|1100)|Polaris[_ ]803"
  },
  {
      "name": "Bmobile",
      "regex": "Bmobile[_ ]|AX-?([1-9][0-9]{2,3}[eEO+]?|7OO)"
  },
  {
      "name": "bq",
      "regex": "Aquaris|bq [^/;]+ Build|BQS-(400[37]|4701|5515|505[05]|5500|5020|5065)|(BQ(ru)?[_ -]5(000|202|211|03[357]|20[34]|504|521|510|05[79]|590))|BQ(ru)?-4(072)"
  },
  {
      "name": "Bush",
      "regex": "Bush[ _-]"
  },
  {
      "name": "CAGI",
      "regex": "CAGI-"
  },
  {
      "name": "Capitel",
      "regex": "Capitel-([a-z0-9]+)"
  },
  {
      "name": "Captiva",
      "regex": "Captiva[ _-]([^;/]+) Build"
  },
  {
      "name": "Casio",
      "regex": "(C811|C7[57]1)(?: 4G)?[);/ ]"
  },
  {
      "name": "Cat",
      "regex": "Cat ?(tablet|stargate|nova)|B15Q"
  },
  {
      "name": "Carrefour",
      "regex": "CT(?:10[0123]0|7[12]0|820)(?:W|FR)?[);/ ]"
  },
  {
      "name": "Celkon",
      "regex": "Celkon"
  },
  {
      "name": "Changhong",
      "regex": "Changhong"
  },
  {
      "name": "Cherry Mobile",
      "regex": "Cherry|Flare([ _])?(2X|4|A[123]|J[1235]|P[13]|S[456]|X2)|Flare S Play|Flare_HD_MAX|FLARE LITE|Flare XL Plus|Fusion Bolt|OMEGA HD 4"
  },
  {
      "name": "China Mobile",
      "regex": "(M812C|1501_M02|9930i|A1303|A309W|M651CY) Build"
  },
  {
      "name": "Comio",
      "regex": "Comio|CT701G PLUS|CT701W|GT100"
  },
  {
      "name": "Compal",
      "regex": "Compal-[a-z0-9]+"
  },
  {
      "name": "ComTrade Tesla",
      "regex": "Tesla|Impulse 7.85 3G|Smartphone_6.1|SP(?:6.2|6.2_Lite)|TTL(?:713G|8)"
  },
  {
      "name": "ConCorde",
      "regex": "ConCorde ([^/;]+) Build"
  },
  {
      "name": "Condor",
      "regex": "(?:PGN\\-?[456][012][0-9]|PHS\\-601|Plume L1)[;/) ]|CTAB[^/;]+ Build"
  },
  {
      "name": "Coolpad",
      "regex": "(?:YL-)?Coolpad|8190Q[ ;/\\)]|(8295|5860S) Build|CP8676_I0[23] Build|CP8298_I00 Build|REVVLPLUS C3701A|VCR-I0"
  },
  {
      "name": "Cricket",
      "regex": "Cricket-([a-z0-9]+)"
  },
  {
      "name": "Crius Mea",
      "regex": "(Q7A\\+?)[);/ ]"
  },
  {
      "name": "Crosscall",
      "regex": "Crosscall|ODYSSEY_Plus|Odyssey S1|Trekker-[MSX][1234]|Action-X3"
  },
  {
      "name": "Cube",
      "regex": "(\u003c!myPhone ?)Cube|(U[0-9]+GT|K8GT)"
  },
  {
      "name": "CUBOT",
      "regex": "CUBOT"
  },
  {
      "name": "Concord",
      "regex": "Concord[ _-]|Flyfix 6|C-721[);/ ]"
  },
  {
      "name": "Crescent",
      "regex": "WING 5|WING9|Venus[ _][14]"
  },
  {
      "name": "Cyrus",
      "regex": "Cyrus[ _-]"
  },
  {
      "name": "Datang",
      "regex": "DATANG"
  },
  {
      "name": "Datsun",
      "regex": "DATSUN"
  },
  {
      "name": "Danew",
      "regex": "Dslide ?([^;/]+) Build|Konnect_(504|601)"
  },
  {
      "name": "Denver",
      "regex": "(TA[CDQ]-[0-9]+)"
  },
  {
      "name": "Dell",
      "regex": "Dell[^a-z]|Venue|XCD35"
  },
  {
      "name": "Desay",
      "regex": "Desay"
  },
  {
      "name": "DEXP",
      "regex": "DEXP|(?:Ursus|Ixion)[ _]([^;/]+) Build|H135[ _]Build"
  },
  {
      "name": "Dbtel",
      "regex": "DBTEL(?:[\\-/])?([a-z0-9]+)"
  },
  {
      "name": "Dialog",
      "regex": "Dialog ?([^;/]+) Build"
  },
  {
      "name": "Dicam",
      "regex": "DICAM-([a-z0-9]+)"
  },
  {
      "name": "Digicel",
      "regex": "DIGICEL"
  },
  {
      "name": "Digiland",
      "regex": "DL1008M"
  },
  {
      "name": "Digma",
      "regex": "Digma[_ ]([^;/]+) Build|HIT HT707[10]MG|CITI 1902 3G|(CITI (?:[A-Z0-9]+) 3G) C[ST]500[67]PG|iDrQ10 3G|iDxD8 3G|iDnD7|HIT 4G HT7074ML|IDX5|(iDx10|iDx7)|MVM900H(?:WZ|C)|MVM908HCZ|iDxD10 3G|(iDxD[45]|iDxQ5)|iDxD7[_ ]3G|PS604M|PT452E|Linx A400 3G LT4001PG|Linx C500 3G LT5001PG|Linx PS474S|NS6902QL|NS9797MG|((?:Optima|Platina|Plane)[ _](?:[EM])?(?:[0-9\\.ST]+|Prime)(?:[ _][43]G)?)|(VOX[ _](?:[0-9\\.A-Z]+)[_ ][43]G)"
  },
  {
      "name": "DoCoMo",
      "regex": "DoCoMo|\\;FOMA|KGT/1\\.0"
  },
  {
      "name": "Doogee",
      "regex": "DOOGEE[ _-]|Valencia2_Y100pro|X5max(?:[_ ]PRO)?|(BIGBOY|COLLO[23]?|DAGGER|DISCOVERY2?|FIND|HOTWIND|LATTE|MAX|MINT|MOON|PIXELS|RAINBOX|TURBO|VALENCIA|VOYAGER2?|TITANS2?)[ -_](DG[0-9]+C?)[);/ ]|BL[57]000[);/ ]|BL12000"
  },
  {
      "name": "Doov",
      "regex": "(?:Doov-)?Doov[ _]"
  },
  {
      "name": "Dopod",
      "regex": "Dopod(?: )?([a-z0-9]+)"
  },
  {
      "name": "Doro",
      "regex": "Doro"
  },
  {
      "name": "Dune HD",
      "regex": "DuneHD"
  },
  {
      "name": "DNS",
      "regex": "(?:DNS|AirTab)[ _\\-]([^;/]+)Build|S4505M[);/ ]"
  },
  {
      "name": "E-Boda",
      "regex": "E-Boda|Eclipse_(?:G400M|G500)"
  },
  {
      "name": "Easypix",
      "regex": "EasyPad|EasyPhone"
  },
  {
      "name": "EBEST",
      "regex": "EBEST"
  },
  {
      "name": "Echo Mobiles",
      "regex": "ECHO[ _](HOLI|HORIZON|MAX|MOSS|NOTE|POWER|SMART)"
  },
  {
      "name": "ECS",
      "regex": "(?:TA10CA3|TM105A?|TR10CS1)[);/ ]"
  },
  {
      "name": "EE",
      "regex": "Hawk_from_EE"
  },
  {
      "name": "Eks Mobility",
      "regex": "S5LS|X4UPlus"
  },
  {
      "name": "Elephone",
      "regex": "Elephone|P6000( ?Pro| ?Plus|\\+| ?02)? Build"
  },
  {
      "name": "Energy Sistem",
      "regex": "ENERGY[ _-]?([^;/]+) Build"
  },
  {
      "name": "Ericy",
      "regex": "Ericy-([a-z0-9]+)"
  },
  {
      "name": "Essential",
      "regex": "PH-1[ )]"
  },
  {
      "name": "E-tel",
      "regex": "E-tel_i250"
  },
  {
      "name": "Evercoss",
      "regex": "A75A\\* Build"
  },
  {
      "name": "Eton",
      "regex": "Eton"
  },
  {
      "name": "Essentielb",
      "regex": "(Wooze[_ ](?:I5|L|XL)) Build"
  },
  {
      "name": "Rikomagic",
      "regex": "MK(80[28][^/;]*) Build"
  },
  {
      "name": "Riviera",
      "regex": "RIVIERA "
  },
  {
      "name": "Senseit",
      "regex": "SENSEIT[ _]?([^;/)]+) Build"
  },
  {
      "name": "Sony",
      "regex": "Sony(?: ?Ericsson)?|SGP|Xperia|([456]0[12]|701)SO|C1[569]0[45]|C2[01]0[45]|C230[45]|C530[236]|C550[23]|C6[56]0[236]|C6616|C68(?:0[26]|[34]3)|C69(?:0[236]|16|43)|D200[45]|D21(?:0[45]|14)|D22(?:0[236]|12|43)|D230[2356]|D240[36]|D25(?:02|33)|D510[236]|D530[36]|D5316|D5322|D5503|D58[03]3|D65(?:0[23]|43|63)|D66[0345]3|D66[14]6|D6708|E(?:20[0345]3|2006|210[45]|2115|2124|230[36]|2312|23[356]3|530[36]|53[3456]3|5506|55[356]3|56[46305][36]|58[02]3|65[35]3|66[0358]3|68[358]3)|F311[12356]|F331[13]|F321[12356]|F5[13]21|F5122|F813[12]|F833[12]|G312[135]|G311[26]|G322[136]|G3212|G331[123]|G3412|G3416|G342[136]|G823[12]|G834[123]|G8[14]4[12]|H(?:3113|3123|3133|3213|3223|3311|3321|4113|4133|4213|4233|4311|4331|4413|4433|82[167]6|83[12]4|8416|9436)|(?:WT|LT|SO|ST|SK|MK)[0-9]+[a-z]+[0-9]*(?: Build|\\))|X?L39H|XM50[ht]|W960|portalmmm/2\\.0 K|S3[69]h|SOL2[246]|SOV3[1234567]|X10[ia]v?|E1[05][ai]v?|MT[0-9]{2}[a-z]? Build|SO-0(?:[12]C|[345]D|[234]E|[1235]F|[1234]G|[134]H|[1234]J|2K)|R800[aix]|LiveWithWalkman"
  },
  {
      "name": "Ericsson",
      "regex": "(?:Ericsson(?:/ )?[a-z0-9]+)|(?:R380 2.0 WAP1.1)"
  },
  {
      "name": "Energizer",
      "regex": "PowerMax"
  },
  {
      "name": "eTouch",
      "regex": "eTouch ?([a-z0-9]+)"
  },
  {
      "name": "Storex",
      "regex": "eZee[^a-z]*Tab ?([^;/]*) Build|STOREX LinkBox"
  },
  {
      "name": "Evertek",
      "regex": "(Ever(?:Glory|Shine|Miracle|Mellow|Classic|Trendy|Fancy|Vivid|Slim|Glow|Magic|Smart|Star)[^/;]*) Build|E70[25]0HD|E805[01]HD|E9054HD|E8050HG|E7914HG"
  },
  {
      "name": "Evolio",
      "regex": "Evolio|M4MAGIC"
  },
  {
      "name": "Evolveo",
      "regex": "EVOLVEO[ _]([^;/]*) Build"
  },
  {
      "name": "Explay",
      "regex": "Explay|ActiveD[ _]|Atlant |Informer[ _][0-9]+|CinemaTV 3G|Surfer[ _][0-9\\.]|sQuad[ _][0-9\\.]|Onliner[1-3]|RioPlay|Leader"
  },
  {
      "name": "EKO",
      "regex": "EKO"
  },
  {
      "name": "Extrem",
      "regex": "Rock X9+"
  },
  {
      "name": "Ezze",
      "regex": "EZZE-|EZ[a-z0-9]+"
  },
  {
      "name": "Ezio",
      "regex": "EZIO-([a-z0-9]+)"
  },
  {
      "name": "Forstar",
      "regex": "Amosta 3G5"
  },
  {
      "name": "Foxconn",
      "regex": "InFocus M[0-9]+[a-z]?"
  },
  {
      "name": "Fairphone",
      "regex": "(FP(?:1U?|2))[);/ ]"
  },
  {
      "name": "FiGO",
      "regex": "ATRIUM II F55L2|TRIO F40LT|EPIC F50G|Gravity X55L|Orion M50L|ULTRA M50G|M405B|VIRTUE3|F55L Build"
  },
  {
      "name": "FNB",
      "regex": "ConeXis (A[12]|X[12])|SP5045V"
  },
  {
      "name": "Fujitsu",
      "regex": "F-01[FHJK]|F-02[EFGH]|F-03[DEFGH]|F-04[EG]|F-05[DEFJ]|F-06E|F-07[DE]|F-08D|F-09[DE]|F-10D|F-11D|F-12C|M532|FARTM933KZ[);/ ]"
  },
  {
      "name": "G-TiDE",
      "regex": "G-TiDE"
  },
  {
      "name": "Gemini",
      "regex": "(GEM[0-9]+[a-z]*)"
  },
  {
      "name": "Geotel",
      "regex": "GEOTEL"
  },
  {
      "name": "Ghia",
      "regex": "GHIA|QS702"
  },
  {
      "name": "Gigabyte",
      "regex": "GSmart [a-z0-9 ]+ Build|Gigabyte-[a-z0-9]+"
  },
  {
      "name": "Gigaset",
      "regex": "Gigaset"
  },
  {
      "name": "Gionee",
      "regex": "(?:GIO-)?GIONEE[ _-]?[a-z0-9]+|(?:Dream_D1|V188S?|GN[0-9]{3,4}[a-z]?)[);/ ]|F103(?: Pro)? Build|P5 mini Build|M7 Power "
  },
  {
      "name": "Grape",
      "regex": "(?:Grape[ _])?GTM-5([^;/]+) Build"
  },
  {
      "name": "Vizio",
      "regex": "Vizio|VAP430|VTAB1008"
  },
  {
      "name": "GOCLEVER",
      "regex": "GOCLEVER|QUANTUM_[0-9]{3}|QUANTUM [0-9]{1}|QUANTUM 700N|QUANTUM_(1010N|1010M|900)|ARIES|INSIGNIA|ORION_|ELIPSO|LIBRA[ _]97"
  },
  {
      "name": "GoMobile",
      "regex": "GO Onyx|GO[0-9]{3,4}"
  },
  {
      "name": "Goly",
      "regex": "Goly[ _-]"
  },
  {
      "name": "Google",
      "regex": "Nexus|GoogleTV|Glass|CrKey[^a-z0-9]|Pixel Build|Pixel (?:XL|C|[23]|[23] XL)"
  },
  {
      "name": "Gradiente",
      "regex": "GRADIENTE"
  },
  {
      "name": "Grundig",
      "regex": "GR?-TB[0-9]+[a-z]*|GRUNDIG|portalmmm/2\\.0 G"
  },
  {
      "name": "Hafury",
      "regex": "HAFURY"
  },
  {
      "name": "Haier",
      "regex": "Haier|(?:HW-)?W(?:716|757|860|970)[);/ ]|(?:HM-)?(G(?:152|303|353|552|70[01])?-FL|G(?:303)?-W|I(?:557)?-FL)|(?:PAD[ _](D71|G781|d85))"
  },
  {
      "name": "HannSpree",
      "regex": "(HSG[0-9]{4})|SN97T41W|SN1AT71W\\(B\\)|SN70T51A|SN70T31?|T7-QC"
  },
  {
      "name": "Hasee",
      "regex": "Hasee"
  },
  {
      "name": "Hisense",
      "regex": "(?:HS-)?Hisense ([^;/]+) Build|Hisense [^);/]+|HS-(?:G|U|EG?|I|L|T|X)[0-9]+[a-z0-9\\-]*|E270BSA|M470BS[AE]|E2281|EG680|HLTE700T"
  },
  {
      "name": "Hi-Level",
      "regex": "HLV-T([a-z0-9]+)"
  },
  {
      "name": "Homtom",
      "regex": "Homtom|(?:HT[0-9]{1,2})(?: ?Pro)? Build"
  },
  {
      "name": "Hosin",
      "regex": "HOSIN[ _-]"
  },
  {
      "name": "Huawei",
      "regex": "(HW-)?(?:Huawei|Ideos|Honor[ _]?|(?:(?:AGS2|AGS|ALE|ALP|ANE|ARE|ATH|ATU|AUM|BAC|BAH|BG2|BGO||BKK|BKL|BLA|BLL|BLN|BND|BTV|CAG|CAM|CAN|CAZ|CHC|CHE[12]?|CHM|CLT|CMR|COL|COR|CPN|CRO|CRR|CUN|DIG|DLI|DRA|DUA|DUB|DUK|EDI|ELE|EML|EVA|EVR|FDR|FIG|FLA|FRD|G(?:527|620S|621|630|735)|GRA|H[36]0|HMA|HRY|INE|JAT|JKM|JMM|JSN|KIW|KOB|LDN|LLD|LND|LON|LUA|LYA|MHA|MYA|NEM|NEO|NXT|PAR|PCT|PIC|PLE|PLK|POT|PRA|RIO|RNE|SCC|SCL|SHT|SLA|SNE|STF|TAG|TIT|TNY|TRT|VCE|VEN|VIE|VKY|VNS|VOG|VTR|WAS|Y(?:221|330|550|6[23]5))-[A-Z]{0,2}[0-9]{1,4}[A-Z]?|H1711|U(?:8230|8500|8661|8665|8667|8800|8818|8860|9200|9508))[);/ ])|hi6210sft|DIG-L21HN"
  },
  {
      "name": "Hyundai",
      "regex": "Hyundai|Ultra (?:Active|Air|Charm|Dream|Energy|Latitude|Link|Live|Shadow|Shine|Storm|Style|Sync|Trend|Vision|Wave)|G(24027|2552[34])K|W25042L"
  },
  {
      "name": "i-Joy",
      "regex": "i-Joy|i-Call|(?:Neon[79]|Sygnus|Deox|Elektra (?:M|L|XL|XXL))[);/ ]"
  },
  {
      "name": "iBall",
      "regex": "iBall[ _]([^/;]*)[ _]Build"
  },
  {
      "name": "iRola",
      "regex": "iRola ([^/;]+) Build|(DX758|DX752|DX752|DX758[ _]?Pro) Build"
  },
  {
      "name": "iBerry",
      "regex": "AUXUS ([^/;]+) Build"
  },
  {
      "name": "iHunt",
      "regex": "iHunt|One_ Love_|TITAN_P11000_PRO|Like_(3|4U|3_Pro)|X300 Elite"
  },
  {
      "name": "IconBIT",
      "regex": "NT-(3506M|0704S|0909T|1009T|10[01]1T|1017T|3601P/3602P|3603P|3702S|3805C|3905T|3701S|0701S|0805C|0902S|370[23]M)"
  },
  {
      "name": "LYF",
      "regex": "(LYF[ _])?LS-[456][0-9]{3}|LYF/[^/]+/"
  },
  {
      "name": "M4tel",
      "regex": "M4 SS[0-9]{4}(?:-R)?|M4_B[23]"
  },
  {
      "name": "IMO Mobile",
      "regex": "IMO Q2"
  },
  {
      "name": "iNew",
      "regex": "iNew|(?:V7A|V3 Plus|V3-?E|V3C)(?: Build|[;/\\)])"
  },
  {
      "name": "iPro",
      "regex": "Kylin[ _]5.[05]"
  },
  {
      "name": "Infinix",
      "regex": "Infinix"
  },
  {
      "name": "InFocus",
      "regex": "(IF9007) Build"
  },
  {
      "name": "InnJoo",
      "regex": "InnJoo"
  },
  {
      "name": "Inkti",
      "regex": "intki[ _]([^/;]*)[ _]Build"
  },
  {
      "name": "Innostream",
      "regex": "INNO([a-z0-9]+)"
  },
  {
      "name": "INQ",
      "regex": "INQ[/ ]"
  },
  {
      "name": "Intex",
      "regex": "Intex|(Aqua|Cloud[ _][XY][0-9]{1})"
  },
  {
      "name": "i-mate",
      "regex": "i-mate ([a-z0-9]+)"
  },
  {
      "name": "i-mobile",
      "regex": "i-mobile ?[a-z0-9]+|i-Style|IQ ?(9\\.1|5\\.5|5\\.6A?)"
  },
  {
      "name": "iKoMo",
      "regex": "iKoMo ([a-z0-9]+)"
  },
  {
      "name": "iOcean",
      "regex": "iOCEAN|M6752|W180|X7[ _]?S|X8 mini"
  },
  {
      "name": "iView",
      "regex": "733TPC"
  },
  {
      "name": "Impression",
      "regex": "(?:Impression[ _\\-]?([^/;]+)|A502|I10\\-LE|ImPad[ _]?(.*)|ImSmart[ _]?(.*)) Build"
  },
  {
      "name": "iTel",
      "regex": "itel|iNote"
  },
  {
      "name": "JAY-Tech",
      "regex": "TPC-[A-Z0-9]+"
  },
  {
      "name": "Jiayu",
      "regex": "(JY-[a-z0-9]+)[);/ ]"
  },
  {
      "name": "Jolla",
      "regex": "Jolla"
  },
  {
      "name": "Just5",
      "regex": "JUST5(-)?SPACER(s)?|COSMO[_ ](L707|L808)|BLASTER|FREEDOM[ _](C100|M303|X1)|Freedom"
  },
  {
      "name": "Kalley",
      "regex": "ELEMENT Q"
  },
  {
      "name": "Kazam",
      "regex": "Kazam ([^;/]+) Build|Trooper_X[0-9][0-9][);/ ]|Tornado 348"
  },
  {
      "name": "Kempler \u0026 Strauss",
      "regex": "Alumini_?3_?Plus|Alumini ?[23]|KEMPLER_(TV|X)"
  },
  {
      "name": "Kiano",
      "regex": "Kiano|Elegance_5_0|CORE 10.1 DUAL 3G|Intelect ?(7|8|10)|Slim ?Tab ?\\d+(?:[_ ]3GR?)?[);/ ]|Cavion[_ ]"
  },
  {
      "name": "Kingsun",
      "regex": "Kingsun[ _-]"
  },
  {
      "name": "Kocaso",
      "regex": "Kocaso|M(?:729|7[357][026][hw]?|76[01236][bw]?|83[016]|8[567]0|1050s|106[012368]w?|1070|X736(?:-kit)?|X780)[);/ ]"
  },
  {
      "name": "Kogan",
      "regex": "Kogan"
  },
  {
      "name": "Komu",
      "regex": "Komu[ -]"
  },
  {
      "name": "Koobee",
      "regex": "Koobee"
  },
  {
      "name": "Kumai",
      "regex": "(?:KM-)?Kumai"
  },
  {
      "name": "KT-Tech",
      "regex": "(KM-[a-z0-9]+|EV-[a-z0-9]+)[);/ ]"
  },
  {
      "name": "KDDI",
      "regex": "kddi-([a-z0-9]+)"
  },
  {
      "name": "Kodak",
      "regex": "KODAK"
  },
  {
      "name": "KOPO",
      "regex": "KOPO[ _-]"
  },
  {
      "name": "Koridy",
      "regex": "KORIDY[ _-]([^/;]+) Build"
  },
  {
      "name": "KRONO",
      "regex": "NET1100"
  },
  {
      "name": "K-Touch",
      "regex": "K-?Touch[ _][a-z0-9]+"
  },
  {
      "name": "Kyocera",
      "regex": "Kyocera|KWC-|QC-|C5120|C5170|C5155|C5215|C6730|C6750|C6522N?|C6530N|C6740|C6743|E6810|KYL21|KYY23|S2151|KC-S701"
  },
  {
      "name": "Krüger\u0026Matz",
      "regex": "Kruger Matz|MOVE_"
  },
  {
      "name": "LAIQ",
      "regex": "LAIQ"
  },
  {
      "name": "Lava",
      "regex": "Lava[ _]|iris[ _]?([^/;]+)(?:\\)| Build)|A(?:67|76) Build|X1 Selfie Build|X41 Plus Build|Flair Z1|PixelV1|Z61_2GB"
  },
  {
      "name": "Landvo",
      "regex": "XM[123]00|(Blade )?S6 Build"
  },
  {
      "name": "Lanix",
      "regex": "LANIX-([a-z0-9]+)|Ilium[ _]|X120C"
  },
  {
      "name": "LCT",
      "regex": "LCT_([a-z0-9]+)"
  },
  {
      "name": "Le Pan",
      "regex": "Le ?Pan|TC970 ([^;/]+)Build"
  },
  {
      "name": "Leagoo",
      "regex": "M5 EDGE|KIICAA POWER|Leapad[ _]7s"
  },
  {
      "name": "Ledstar",
      "regex": "Trendy 5.0"
  },
  {
      "name": "LeEco",
      "regex": "LeEco|Le ?X[0-9][^;/]+|Le Max|Le [12]|X900 Build"
  },
  {
      "name": "Leotec",
      "regex": "LESPH501[14]"
  },
  {
      "name": "Lenco",
      "regex": "Lenco ([^/;]*) Build"
  },
  {
      "name": "Lenovo",
      "regex": "(?:LNV-)?Lenovo|IdeaTab|IdeaPad|Thinkpad|Yoga Tablet|Tab2A[0-9]-[0-9]{2}[a-z]?|YT3-X50L|ZUK[ -_]|K50a40|TB-(8504F|X304F|X704V|X704A)|A2107A-H|S6000[ ;)]"
  },
  {
      "name": "Lexand",
      "regex": "(?:SC7 PRO HD|S5A[1-4]|S4A[1-5]|S6A1|A811|A802)[);/ ]|Lexand ([^;/]+) Build"
  },
  {
      "name": "Lexibook",
      "regex": "(MFC[0-9]{3}[a-z]{2,})"
  },
  {
      "name": "LG",
      "regex": "LG|portalmmm/2\\.0 (?:KE|KG|KP|L3)|(?:VX[0-9]+|L-0[12]D|L-07C|P713|LM-Q710(?:\\(FGN\\)|\\.FGN?)?|LM-X210(?:\\(G\\)|CM)?|LM-G710N?|LM-V405|LM-X410(?:\\(FGN?\\)|\\.F(?:GN?)?))[);/ ]|NetCast|RS987|VS(?:501|988|99[56]|(410|415|425|500)PP)|Nexsus 5|LML211BL|LML212VL"
  },
  {
      "name": "Lingwin",
      "regex": "Lingwin[ _\\-]"
  },
  {
      "name": "Logicom",
      "regex": "(TAB950|TAB1062|E731|E812|E912|E1031|POWER BOT|(?:B|C|M|ID|VR)[ _]?BOT[ _]?(?:TAB[ _])?([0-9]+\\+?)(?:PLUS)?|KT712A_4\\.4|L-?IXIR[0-9]*|L-ITE|L-?EMENT|Le Lift)[_);/ ]"
  },
  {
      "name": "Konka",
      "regex": "KONKA[_ ]([a-z0-9]+)"
  },
  {
      "name": "Konrow",
      "regex": "(?:BIGCOOL|COOLFIVE|COOL-K|Just5|Link5)[);/ ]"
  },
  {
      "name": "Karbonn",
      "regex": "(?:Browser-)?Karbonn|K9 Kavach 4G|K9 Music 4G|K9 Smart|K9 VIRAAT 4G"
  },
  {
      "name": "Sagem",
      "regex": "SAGEM|portalmmm/2.0 (?:SG|my)"
  },
  {
      "name": "Coby Kyros",
      "regex": "(MID(?:1024|1125|1126|1045|1048|1060|1065|4331|7012|7015A?|7016|7022|7032|7035|7036|7042|7047|7048|7052|7065|7120|8024|8042|8048|8065|8125|8127|8128|9724|9740|9742))[);/ ]"
  },
  {
      "name": "Mpman",
      "regex": "(?:MPQC|MPDC)[0-9]+|PH(?:150|340|350|360|451|500|520)|(?:MID(?:7C|74C|82C|84C|801|811|701|711|170|77C|43C|102C|103C|104C|114C)|MP(?:843|717|718|1010|7007|7008|843|888|959|969)|MGP7)[);/ ]"
  },
  {
      "name": "Majestic",
      "regex": "MAJESTIC[ _-]|CRONO ?[0-9]+"
  },
  {
      "name": "Manta Multimedia",
      "regex": "((/)?MID713|MID(?:06[SN]|08[S]?|12|13|14|15|701|702|703|704|705(?:DC)?|706[AS]?|707|708|709|711|712|714|717|781|801|802|901|1001|1002|1003|1004(?: 3G)?|1005|1009|1010|7802|9701|9702))[);/ ]"
  },
  {
      "name": "Maxwest",
      "regex": "Astro_5N_LTE|Virtue Z5"
  },
  {
      "name": "Mediacom",
      "regex": "(?:SmartPad7503G|SmartPad970s2(?:3G)?|M[_-][MPS]P[0-9a-z]+|M-IPRO[0-9a-z]+)[);/ ]"
  },
  {
      "name": "Medion",
      "regex": "Medion|(?:MD_)?LIFETAB|X5001[);/ ]|[SPX]10[0-9]{2}X[);/ ]"
  },
  {
      "name": "MEEG",
      "regex": "MEEG[ _-]"
  },
  {
      "name": "Meizu",
      "regex": "Meizu|MZ-[a-z]|(M04[05]|MZ-16s|M35[1356]|M6T|MX[ -]?[2345](?: Pro)?|(?:MZ-)?m[1-6] note|M57[18]C|M3[ESX]|M031|m1 metal|M1 E|M2|M2 E|M5s Build|PRO 6|PRO 7-H)[);/ ]"
  },
  {
      "name": "Memup",
      "regex": "SlidePad ?([^;/]*) Build|SPNG?[0-9]{2,4}[a-z]{0,2}[ ;/)]|SP704CE?"
  },
  {
      "name": "Mecer",
      "regex": "(M785|800P3[12]C|101P51C|X1010)[);/ ]"
  },
  {
      "name": "Mio",
      "regex": "MIO(?:/)?([a-z0-9]+)"
  },
  {
      "name": "Miray",
      "regex": "MIDM[_-]|MPM[_-]"
  },
  {
      "name": "Mitsubishi",
      "regex": "MITSU|portalmmm/[12]\\.0 M"
  },
  {
      "name": "MIXC",
      "regex": "(M35|M20|G7106|G7108)[ );/]"
  },
  {
      "name": "Mobiistar",
      "regex": "Mobiistar|Zumbo_|PRIME X MAX"
  },
  {
      "name": "MSI",
      "regex": "Primo ?(?:91|76)|Enjoy 7 Plus"
  },
  {
      "name": "MLLED",
      "regex": "MLLED[ _]"
  },
  {
      "name": "Mobistel",
      "regex": "(Cynus[ _][^/;]+) Build"
  },
  {
      "name": "Modecom",
      "regex": "Modecom|Free(?:Way )?Tab|xino z[\\d]+ x[\\d]+"
  },
  {
      "name": "Mofut",
      "regex": "Mofut"
  },
  {
      "name": "Motorola",
      "regex": "MOT|(AN)?DROID ?(?:Build|[a-z0-9]+)|portalmmm/2.0 (?:E378i|L6|L7|v3)|XOOM [^;/]*Build|(?:XT|MZ|MB|ME)[0-9]{3,4}[a-z]?(?:\\(Defy\\)|-0[1-5])?(?: Build|\\))"
  },
  {
      "name": "Movic",
      "regex": "MOVIC"
  },
  {
      "name": "MTN",
      "regex": "MTN-"
  },
  {
      "name": "MyPhone",
      "regex": "(?:MyPhone|MyPad|MyTab)[ _][^;/]+ Build|Cube_LTE|myTab10II|HAMMER[ _]ENERGY|HAMMER_AXE_M_LTE|Hammer Titan 2|C-Smart_pix"
  },
  {
      "name": "MyWigo",
      "regex": "MyWigo "
  },
  {
      "name": "Myria",
      "regex": " (Myria_[^/;]+|502M|Cozy|Easy tab 9|Grand) Build|Myria_FIVE|Myria_Grand_4G|Myria_Wide_2|Myria_Wide_4G"
  },
  {
      "name": "M.T.T.",
      "regex": "M.T.T."
  },
  {
      "name": "Navon",
      "regex": "Superme_Max"
  },
  {
      "name": "NEC",
      "regex": "NEC[ _\\-]|KGT/2\\.0|portalmmm/1\\.0 (?:DB|N)|(?:portalmmm|o2imode)/2.0[ ,]N"
  },
  {
      "name": "Neffos",
      "regex": "Neffos"
  },
  {
      "name": "Nextbit",
      "regex": "(Robin) Build"
  },
  {
      "name": "Newgen",
      "regex": "NEWGEN\\-([a-z0-9]+)"
  },
  {
      "name": "Netgear",
      "regex": "GTV100"
  },
  {
      "name": "NeuImage",
      "regex": "NIM-"
  },
  {
      "name": "NextBook",
      "regex": "Next[0-9]|NX785QC8G|NXM900MC|NX008HD8G|NX010HI8G|NXM908HC|NXM726"
  },
  {
      "name": "NGM",
      "regex": "NGM[_ ][a-z0-9]+|(Forward|Dynamic)[ _]?[^/;]+(?:Build|/)"
  },
  {
      "name": "Nexian",
      "regex": "S?Nexian"
  },
  {
      "name": "Noain",
      "regex": "NOAIN"
  },
  {
      "name": "Noblex",
      "regex": "Noblex|NBX-"
  },
  {
      "name": "Nomi",
      "regex": "Nomi[ _-]"
  },
  {
      "name": "NUU Mobile",
      "regex": "N5702L"
  },
  {
      "name": "NYX Mobile",
      "regex": "NYX[ _]"
  },
  {
      "name": "Nous",
      "regex": "NS([356]|500[346]|5511) Build"
  },
  {
      "name": "Nvidia",
      "regex": "(SHIELD Tablet K1) Build"
  },
  {
      "name": "O+",
      "regex": "O\\+[ _](Air|Crunch|Ultra|Upsized|Venti)"
  },
  {
      "name": "O2",
      "regex": "Xda|O2[ \\-]|COCOON"
  },
  {
      "name": "Odys",
      "regex": "ODYS[ _-]|IEOS[_ ]([^/;]+)|NOON(?:_PRO)? Build|SPACE10_(?:PLUS|PRO)_3G|THOR_?10|TAO_X10|RAPID_?(?:10|7)_?LTE|MAVEN_?10_|CONNECT[78]|ELEMENT10_PLUS_3G|XELIO[_0-9P]|ADM816HC|ADM8000KP|NEO6_LTE|EOS10|AEON|FALCON_10_PLUS_3G|FUSION"
  },
  {
      "name": "Obi",
      "regex": "Obi[ _-]|(SJ1\\.5|SJ2\\.6|S400|S452\\+|S451|S453|S501|S502|S503\\+?|S507|S520|S550|S551|falcon)[ _]"
  },
  {
      "name": "Onda",
      "regex": "Onda"
  },
  {
      "name": "OnePlus",
      "regex": "(?:du_)?ONEPLUS|(?:A0001|A200[135]|A300[03]|A3010|A5000|A600[03]|A601[03]|E100[135]|GM1911|GM1903)(?: Build|\\))"
  },
  {
      "name": "OPPO",
      "regex": "(?:OB-)?OPPO[ _]?([a-z0-9]+)|N1T|(?:X90[07][0679]|U707T?|X909T?|R(?:10[01]1|2001|201[07]|6007|7005|7007|80[13579]|81[13579]|82[01379]|83[013]|800[067]|8015|810[679]|811[13]|820[057])[KLSTW]?|N520[79]|N5117|A33f|A33fw|A37fw?|PAAM00|PAC[T|M]00|PBAM00)[);/ ]|R7kf|R7plusf|R7Plusm|A1601|CPH[0-9]{4}|PCAM10|PADM00|RMX18(0[1579]|11|3[13])"
  },
  {
      "name": "Opsson",
      "regex": "Opsson|IUSAI"
  },
  {
      "name": "Orange",
      "regex": "SPV[ \\-]?([a-z0-9]+)|Orange[ _-]([^/;]+) Build|SC/IHD92|FunTab ([^/;]+) Build"
  },
  {
      "name": "Ouki",
      "regex": "OUKI|OK[AU][0-9]+[a-z]* Build"
  },
  {
      "name": "Overmax",
      "regex": "OV-[a-z]+(?:[^;(/]*)[();/ ]|Qualcore 1010"
  },
  {
      "name": "Oysters",
      "regex": "Oysters|T84ERI[ _]3G"
  },
  {
      "name": "öwn",
      "regex": "OWN"
  },
  {
      "name": "Panacom",
      "regex": "T-i708D"
  },
  {
      "name": "Panasonic",
      "regex": "Panasonic|PANATV[0-9]+|Viera/|P902i[);/ ]|Eluga[ _]|FZ-N1|P55 Novo 4G"
  },
  {
      "name": "PCBOX",
      "regex": "Art-PCB-V116|Bee-PCB-V216|Clap-PCB-I316|PCB-T(103|715)"
  },
  {
      "name": "PCD",
      "regex": "PH4001"
  },
  {
      "name": "PCD Argentina",
      "regex": "PCD[ ]?50[689]"
  },
  {
      "name": "Pentagram",
      "regex": "Pentagram|Quadra|Monster X5"
  },
  {
      "name": "Philips",
      "regex": "Philips|AND1E[);/ ]|NETTV/|PI3210G"
  },
  {
      "name": "phoneOne",
      "regex": "phoneOne[ \\-]?([a-z0-9]+)"
  },
  {
      "name": "Primepad",
      "regex": "PD\\-(3127NC|3127) Build"
  },
  {
      "name": "Pioneer",
      "regex": "Pioneer|.*; R1 Build"
  },
  {
      "name": "PULID",
      "regex": "PULID[ _]|F1[01357]\\+? Build"
  },
  {
      "name": "Point of View",
      "regex": "POV_TV|POV_TAB|MOB-5045"
  },
  {
      "name": "Pomp",
      "regex": "POMP[ _-]"
  },
  {
      "name": "PPTV",
      "regex": "(KING 7S?)"
  },
  {
      "name": "ProScan",
      "regex": "PLT([^;/]+) Build"
  },
  {
      "name": "RCA Tablets",
      "regex": "RCT([^;/]+) Build|RCA RLTP4028"
  },
  {
      "name": "Readboy",
      "regex": "Readboy[ _-]"
  },
  {
      "name": "Roku",
      "regex": "Roku/DVP"
  },
  {
      "name": "Rokit",
      "regex": "IO Pro"
  },
  {
      "name": "Rover",
      "regex": "Rover ([0-9]+)"
  },
  {
      "name": "RT Project",
      "regex": "Shock 5"
  },
  {
      "name": "Safaricom",
      "regex": "NEON_RAY"
  },
  {
      "name": "Santin BiTBiZ",
      "regex": "BiTBiZ"
  },
  {
      "name": "Siemens",
      "regex": "SIEMENS|SIE-|portalmmm/2\\.0 SI|S55|SL45i"
  },
  {
      "name": "Samsung",
      "regex": "SAMSUNG( ?Browser)?|Maple |SC-(?:02[CH]|04E|01F)|N[57]100|N5110|N9100|S(?:CH|GH|PH|EC|AM|HV|HW|M)-|SMART-TV|GT-|Galaxy|(?:portalmmm|o2imode)/2\\.0 [SZ]|sam[rua]|vollo Vi86[);/ ]|(?:OTV-)?SMT-E5015|ISW11SC|SCV3[1-9]|404SC"
  },
  {
      "name": "Sanei",
      "regex": "(?:8312D|G(60[25]|70[3568]G?|785|803)|N83(?:-2cpu)?|N91) Build"
  },
  {
      "name": "Selfix",
      "regex": "SELFIX"
  },
  {
      "name": "Sencor",
      "regex": "Sencor|ELEMENT[ _]?(?:7|8|9\\.7|10[ _]1)(?:[ _]?V[23])?[);/ ]|ELEMENT[ _]?(?:P[0-9]+)[);/ ]"
  },
  {
      "name": "Senwa",
      "regex": "Senwa|(?:S\\-?(?:471|7[12]5|6[01]5|915|905TL|1000)|V705B)[ /;\\)]"
  },
  {
      "name": "Sky",
      "regex": "(SKY|Elite|Fuego)[_ ][3-7]\\.[05]([A-Z]{1,2})?|Elite_5_5_Octa|Platinum_(5.0M|M5)"
  },
  {
      "name": "Skyworth",
      "regex": "Sky_?worth"
  },
  {
      "name": "Smartfren",
      "regex": "Smartfren|Androtab|Andromax|PD6D1J|AD682J|AD68[89]G|AD6B1H|AD9A1H|AD682H|AD683G"
  },
  {
      "name": "Smartisan",
      "regex": "SM(?:70[15]|801|901|919)|OD10[35]|YQ60[1357][ /;\\)]"
  },
  {
      "name": "STF Mobile",
      "regex": "Fractal|AERIAL PLUS"
  },
  {
      "name": "STK",
      "regex": "STK[_ ]"
  },
  {
      "name": "Stonex",
      "regex": "STX[ -]([^;/]+)"
  },
  {
      "name": "SuperSonic",
      "regex": "(SC-[0-9]+[a-z]+)"
  },
  {
      "name": "Supra",
      "regex": "SUPRA ([^;/]+) Build|NVTAB 7.0 3G"
  },
  {
      "name": "Sumvision",
      "regex": "(Cyclone [^/;]+) Build"
  },
  {
      "name": "SunVan",
      "regex": "SUNVAN[ _\\-]"
  },
  {
      "name": "SWISSMOBILITY",
      "regex": "SCHR9GR|ZUR722M"
  },
  {
      "name": "Thomson",
      "regex": "(?:Primo(7|8)|QM73[45]-8G|QM734-2|QM736-8G\\(HD\\)|8950|MID1002)[);/ ]|(?:Tlink|Every)[0-9]+"
  },
  {
      "name": "Pantech",
      "regex": "Pantech|P[GN]-|PT-[a-z0-9]{3,}|TX[T]?[0-9]+|IM-[a-z0-9]+[);/ ]|ADR910L"
  },
  {
      "name": "Ployer",
      "regex": "MOMO([0-9]|miniS)"
  },
  {
      "name": "Plum",
      "regex": "PLUM "
  },
  {
      "name": "Polaroid",
      "regex": "Polaroid|(?:PMID|MIDC)[0-9a-z]+[);/ ]|MID(?:1014|0714)|PRO[VG]?(?:[0-9]{3,}[a-z]*|[0-9]{2}[a-z])|P(400[56]|4526|500[56]|502[56]|504[67]|552[56])A|PSPC(505|550)|PSPCK21NA|PSPCL20A0|PSPCM20A0|PSPCZ20A0"
  },
  {
      "name": "PolyPad",
      "regex": "POLY ?PAD"
  },
  {
      "name": "Positivo",
      "regex": "YPY_S450"
  },
  {
      "name": "Prestigio",
      "regex": "(?:PMP|PAP|PMT|PSP)[0-9]+[a-z0-9_]+[);/ ]"
  },
  {
      "name": "Sanyo",
      "regex": "Sanyo|MobilePhone[ ;]"
  },
  {
      "name": "Qilive",
      "regex": "Qilive [0-9][^;/]*|Q8S55IN4G2"
  },
  {
      "name": "QMobile",
      "regex": "QMobile|QTab|Q-Smart|Noir X1S"
  },
  {
      "name": "Quantum",
      "regex": "Quantum (Fit|Mini|V|You E)"
  },
  {
      "name": "Quechua",
      "regex": "Quechua Phone 5"
  },
  {
      "name": "Ramos",
      "regex": "Ramos ?([^/;]+) Build"
  },
  {
      "name": "Sendo",
      "regex": "Sendo([a-z0-9]+)"
  },
  {
      "name": "Turbo-X",
      "regex": "Turbo-X"
  },
  {
      "name": "Silent Circle",
      "regex": "Blackphone 2"
  },
  {
      "name": "Spice",
      "regex": "Spice"
  },
  {
      "name": "Sharp",
      "regex": "SHARP|SBM|SH-?[0-9]+[a-z]?[);/ ]|SH-Z10|AQUOS|506SH|SHL22|FS8010[);/ ]|TG-L900S"
  },
  {
      "name": "Softbank",
      "regex": "Softbank|J-PHONE"
  },
  {
      "name": "Sonim",
      "regex": "Sonim[ -]|XP[67]700[);/ ]"
  },
  {
      "name": "Star",
      "regex": "N(9[5678]00|8800|9000|9977)[);/ ]"
  },
  {
      "name": "Amazon",
      "regex": "KF(?:OT|TT|JWI|JWA|[DFS]OWI|A[PRSU]WI|T[BH]WI|SAW[IA]|GIWI)[);/ ]|Kindle|Silk/\\d+\\.\\d+|Amazon (?:Tate|Jem)|AFT[BMST]|SD4930UR"
  },
  {
      "name": "Symphony",
      "regex": "SYMPHONY[ \\_]([a-z0-9]+)"
  },
  {
      "name": "Qtek",
      "regex": "Qtek[ _]?([a-z0-9]+)"
  },
  {
      "name": "T-Mobile",
      "regex": "T-Mobile[ _][a-z0-9 ]+|REVVL 2 PLUS"
  },
  {
      "name": "TB Touch",
      "regex": "Ignis [89]"
  },
  {
      "name": "TCL",
      "regex": "TCL[ -][a-z0-9]+|(TCL[_ -][^;/]+|7040N) Build|TCLGalaG60|A502DL"
  },
  {
      "name": "Teclast",
      "regex": "Teclast|Tbook"
  },
  {
      "name": "Tecno Mobile",
      "regex": "Tecno|Phantom6|Phantom6-Plus"
  },
  {
      "name": "TechPad",
      "regex": "Tech ?pad|XTAB[ _-]|Dual C1081HD|S813G"
  },
  {
      "name": "Tesco",
      "regex": "Hudl ([^/;]+) Build|W032i-C3[);/ ]"
  },
  {
      "name": "teXet",
      "regex": "Texet|(NaviPad [^/;]*) Build|TM-(?:1046|1058|1057|1067|3000|3200R|3500|4003|4071|450[34]|451[30]|4377|4082R|550[83]|5010|5017|5[23]77|6003|6906|702[13]|7055HD|709[69]|9747BT|9758|9751HD|9767|5513|520[410]|5505|507[13]|5581|7859|8044|9748[ _]3G|9720|7047HD[ _]3G|9757|9740|4982|4515|4083|500[3567]|5571|3204R|5016|8043|7047HD[ _]3G|8041HD|8066|705[23]|7041|8041HD|8051|8048|974[96]|973[78]W|974[10]|9743W|9750HD|7043XD|7049|7887|7037W|702[46])|TB-(771A|711A)|X-Force[\\(-]?TM-5009\\)?|X-Plus[\\(-]?TM-5577\\)?|X-pad (?:AIR 8|iX 7) 3G"
  },
  {
      "name": "Telefunken",
      "regex": "TEL-1013GIQA"
  },
  {
      "name": "Telego",
      "regex": "TELEGO"
  },
  {
      "name": "Telenor",
      "regex": "Telenor"
  },
  {
      "name": "Telit",
      "regex": "Telit"
  },
  {
      "name": "ThL",
      "regex": "ThL[ _-]"
  },
  {
      "name": "TIANYU",
      "regex": "TIANYU"
  },
  {
      "name": "Timovi",
      "regex": "Tmovi|Infinit_Lite_2"
  },
  {
      "name": "Tolino",
      "regex": "Tolino Tab ([^/;]+) Build"
  },
  {
      "name": "Top House",
      "regex": "Coto[_ ]|X1013"
  },
  {
      "name": "Toplux",
      "regex": "Toplux ([a-z0-9]+)"
  },
  {
      "name": "Touchmate",
      "regex": "TOUCHMATE|(TM-(?:MID1020A|MID710|MID798|MID792|MID788D|SM500N|SM410))"
  },
  {
      "name": "TrekStor",
      "regex": "(?:TrekStor|Surftab) ([^/;]+) Build|Surftab[^;\\)]*(?:[;\\)]|$)|ST10216-2A|VT10416-[12]"
  },
  {
      "name": "Trevi",
      "regex": "Trevi[ _]|TAB[ _]10[ _]3G[ _]V16|TAB[ _](7|8)[ _]3G[ _]V8|TAB9 3G|MINITAB 3GV|Phablet[ _](?:4|4\\.5|5|5\\,3|6)[ _]?[CSQ]|REVERSE[ _]5\\.5[ _]?Q"
  },
  {
      "name": "TVC",
      "regex": "(NuclearSX-SP5)"
  },
  {
      "name": "U.S. Cellular",
      "regex": "USCC-"
  },
  {
      "name": "Uhappy",
      "regex": "Uhappy|UP?580|UP350|UP[35679]20"
  },
  {
      "name": "Unimax",
      "regex": "U(67[013]|680)C"
  },
  {
      "name": "Unowhy",
      "regex": "QOOQ "
  },
  {
      "name": "UTStarcom",
      "regex": "utstar[ _-]?([a-z0-9]+)"
  },
  {
      "name": "Ulefone",
      "regex": "Ulefone|Power_[35]|(Armor(?:[ _]2)?|U00[78][ _]Pro|Be[ _]X|Be[ _]Touch(?:[ _][23])?|Be[ _](?:One|Pure)(?:[ _]Lite)?) Build"
  },
  {
      "name": "UMIDIGI",
      "regex": "UMI(?:DIGI)?[ _]"
  },
  {
      "name": "Uniscope",
      "regex": "Uniscope"
  },
  {
      "name": "Unnecto",
      "regex": "Unnecto|(?:U513|U5151|U61[1356]|U7[12]0|U-830|U90[35])[);/ ]"
  },
  {
      "name": "Unonu",
      "regex": "Unonu[ _-]"
  },
  {
      "name": "Vastking",
      "regex": "M910A[);/ ]"
  },
  {
      "name": "ViewSonic",
      "regex": "ViewSonic|VSD[0-9]+[);/ ]|ViewPad|ViewPhone"
  },
  {
      "name": "Vitelcom",
      "regex": "Vitelcom|portalmmm/[12].0 TSM"
  },
  {
      "name": "VK Mobile",
      "regex": "VK(share)?[\\-]?([a-z0-9 ]+)[ ;\\)/]"
  },
  {
      "name": "Vernee",
      "regex": "Vernee|Mars Pro|Apollo (?:Lite|X)|Thor (?:E|Plus)"
  },
  {
      "name": "Vertu",
      "regex": "Vertu[ ]?([a-z0-9]+)"
  },
  {
      "name": "Verykool",
      "regex": "verykool"
  },
  {
      "name": "Vestel",
      "regex": "Venus[ _](V[1-9]|Z[1-9]0)|V_?TAB|VP74|VT97PRO|VSP145M|VSP250G|VSP355[GS]"
  },
  {
      "name": "Videocon",
      "regex": "Videocon[_ \\-]|VT75C"
  },
  {
      "name": "Vodafone",
      "regex": "Vodafone[ _-]|Smart ?Tab ?(?:III? ?)?(?:|4G|7|10)[);/ -]|VodafoneSmartChat|VFD[ _][0-9]+[;/) ]|VFD320|VF6[89]5|VF-795"
  },
  {
      "name": "Vonino",
      "regex": "Vonino|Zun[ _]X[OS]?|Xylo[ _][XSP]|Volt[ _][XS]|Gyga[ _](?:X|S|QS|XS)|Pluri[ _](?:[BCMQ]7|[CQ]8)|Magnet[ _]M[19]|Jax[ _](?:Q|QS|S|X|Mini)|Navo_QS|Xavy_T7"
  },
  {
      "name": "Vorago",
      "regex": "VORAGO"
  },
  {
      "name": "Voto",
      "regex": "VOTO[ _\\-]|VT8[89]8[;/) ]"
  },
  {
      "name": "Voxtel",
      "regex": "Voxtel_([a-z0-9]+)"
  },
  {
      "name": "Vulcan",
      "regex": "VP5004A"
  },
  {
      "name": "Walton",
      "regex": "Walton|Walpad|Primo[ _](C[1234]?|D[12345]|E[12345]|EF[23]?|EM|F[1234567]i?|G[123457]|GF[234]?|GH[23]?|GM|H[2346]|HM|N|N1|NF|NX2?|R[1234]|RH2?|RM2?|RX[23]?|S[123456]|V1|X[1234]|Z|ZX)"
  },
  {
      "name": "WellcoM",
      "regex": "WELLCOM[ _\\-/]([a-z0-9]+)"
  },
  {
      "name": "Wexler",
      "regex": "Wexler|TAB[ _]10Q[);/ ]|ZEN[ _](?:4\\.5|4\\.7|5)"
  },
  {
      "name": "Wiko",
      "regex": "(?:WIKO[_ -])|(?:WIKO-)?CINK_[a-z0-9_]+|(?:WIKO-)?HIGHWAY_[a-z0-9_]+|(?:WIKO-)?(?:Cink([^/;]*)|Highway([^/;]*)|Iggy|Stairway|Rainbow ?(Jam|Lite|Up)?|Darkside|Darkmoon|Darkfull|Darknight|Freddy|FEVER|Jerry[2]?|Sublim|Ozzy|JIMMY|Barry|Birdy|Bloom|Getaway|Goa|Kite|Robby|Lenny[2-9]?|Slide|Sunset[2-9]?|Sunny[2-9]?|Tommy[2-9]?|PULP(?: Fab)?|Wax|HARRY|WIM Lite|Ridge(?: Fab)?|U FEEL(?: Prime| Lite)?|U PULSE(?: LITE)?|View XL|View Prime)(?: Plus)?(?: 4G)?(?: Build|$)|W_(?:C800|K[346]00)|LENNY[234]|SUNSET2|JERRY|Tommy3|PULP FAB 4G|WC300"
  },
  {
      "name": "Wink",
      "regex": "Wink_"
  },
  {
      "name": "Weimei",
      "regex": "weplus_3"
  },
  {
      "name": "Wileyfox",
      "regex": "(?:Wileyfox [^/]+)|Swift([^/]+) Build"
  },
  {
      "name": "Wolder",
      "regex": "miSmart|miTab|WIAM \\#24"
  },
  {
      "name": "Wolfgang",
      "regex": "AT[ _-]AS([0-9A-Z]+)[);/ ]"
  },
  {
      "name": "Wonu",
      "regex": "Wonu ([a-z0-9]+)"
  },
  {
      "name": "Woo",
      "regex": "SP5026i-Scorpio"
  },
  {
      "name": "Woxter",
      "regex": "Woxter[ _]([^/;]+) Build"
  },
  {
      "name": "X-View",
      "regex": "ZEN_U5+"
  },
  {
      "name": "Xiaomi",
      "regex": "Xiaomi(/MiuiBrowser)?|(?:MI [a-z0-9]+|Mi-4c|MI-One[ _]?[a-z0-9]+|MIX(?: 2S?)?)[);/ ]|HM (?:[^/;]+) (?:Build|MIUI)|(?:2014501|2014011|201481[138]|201302[23]|2013061) Build|Redmi|POCOPHONE|SKR-H0"
  },
  {
      "name": "Xion",
      "regex": "XI-CE(?:655|U[48])"
  },
  {
      "name": "Xolo",
      "regex": "Xolo|(?:Q600|Q700s?|Q800|Q1000s|Q1000[ _]Opus|Q1010i|Q2000|Omega[ _]5.[05])[);/ ]|BLACK-1XM"
  },
  {
      "name": "Yarvik",
      "regex": "Xenta[ \\-]Tab|Luna TAB|TAB09-410|TAB10-410|TAB07-485|TAB0[78]-200|TAB08-201-3G|TAB9-200|TAB09-211|TAB10-2[01]1|TAB13-201"
  },
  {
      "name": "Yes",
      "regex": "M631Y"
  },
  {
      "name": "Yezz",
      "regex": "Yezz|ANDY[ _]|4E4|A5EI"
  },
  {
      "name": "Yu",
      "regex": "YU5010A|AO5510"
  },
  {
      "name": "Yuandao",
      "regex": "N101[ _]DUAL(?:[ _]CORE)?(?:[ _]?2|\\|\\|)?(?:[ _]V11)?[);/ ]"
  },
  {
      "name": "Yusun",
      "regex": "Yusun|LA2-T"
  },
  {
      "name": "Ytone",
      "regex": "YTONE[ _\\-]"
  },
  {
      "name": "Zonda",
      "regex": "(ZM(?:CK|EM|TFTV|TN)[a-z0-9]+)|ZA409"
  },
  {
      "name": "Toshiba",
      "regex": "Toshiba|TSBNetTV/|portalmmm/[12].0 TS|T-01C|T-0[12]D|IS04|IS11T|AT1S0|AT300SE|AT(10-A|10[PL]E-A|100|200|270|300|330|374|400|470|500|503|570|703|830)"
  },
  {
      "name": "Fly",
      "regex": "Fly(Flow|touch)?|FS50[1-9]|FS511|FS551|FS40[1-7]|FS452|FS451|FS454|4FS06|MERIDIAN-|(?:IQ[0-9]{3,})[ _]?(?:Quad|Firebird|Quattro|Turbo|Magic)?(?: Build|[;/\\)])"
  },
  {
      "name": "Freetel",
      "regex": "(FTJ152[ABCD]|FT141B|FT142D_LTEXM|FT142A?|FTJ161B|FTJ152E|FTJ162D) Build"
  },
  {
      "name": "Zeemi",
      "regex": "ZEEMI[ _-]"
  },
  {
      "name": "Zenek",
      "regex": "Z5007|Z551[679]|Z6001"
  },
  {
      "name": "Zopo",
      "regex": "(?:ZOPO[_ ])?(ZP[0-9]{2,}[^/;]+) Build"
  },
  {
      "name": "ZTE",
      "regex": "ZTE|AxonPhone|([a-z0-9]+)_USA_Cricket|(?:Blade (?:L110|L[2357]|L7A|S6|V[6789]|V8Q|V8 SE|V9 VITA|X7|A(310|460|465|475|520|530|602)|V580)|N9[15]8St|NX(?:403A|406E|40[X2]|507J|503A|505J|506J|508J|510J|511J|513J|521J|523J_V1|529J|531J|541J|549J|551J|563J|569[HJ]|573J|575J|591J|595J|597J|601J|907J)|N951[0579]|N9180|N9101|N913[67]|N952[01]|N9560|N9810|N799D|[UV]9180|[UV]9815|Z(?:233V|331|667T|768G|792|81[25]|82[08]|83[12359]|85[125]|851M|(557|717|798|836|861|916)BL|986DL|(232|718|828)TL|(233|353|558|717|799|837|862|899|917|963)VL|955A|95[678]|965|97[018]|98[1237]|986U|999)|Apex2|G (LTE|Lux)|Vec)[);/ ]|KIS II Max|Kis 3"
  },
  {
      "name": "Zuum",
      "regex": "Zuum[ _-]|COVET|MAGNO|ONIX S|STEDI|STELLAR Z"
  },
  {
      "name": "Zen",
      "regex": "(Admire[_ ][^/;]+|Cinemax[^/;)]+)(?:Build|\\))"
  },
  {
      "name": "Zync",
      "regex": "ZYNC|(Cloud Z5|Z1000|Z18|Z99 Dual Core|Z99[_ ][23]G|Z99|Z900[_ ]Plus|Z909|Z930[+]|Z930)"
  },
  {
      "name": "Lemhoov",
      "regex": "Lemhoov"
  },
  {
      "name": "MTC",
      "regex": "MTC[ _](978|982[OT]|982)|SMART[ _](Race|Sprint)[ _]4G"
  },
  {
      "name": "MegaFon",
      "regex": "MFLoginPh"
  },
  {
      "name": "Inoi",
      "regex": "INOI"
  },
  {
      "name": "Vertex",
      "regex": "Lion_Dual|V709X|Baccara|Eagle[ _]4G|Impress[ _]([^;/]+)(?:Build|\\))"
  },
  {
      "name": "Unknown",
      "regex": "WebTV/(\\d+\\.\\d+)"
  },
  {
      "name": "Savio",
      "regex": "TB-PO1 Build"
  },
  {
      "name": "X-TIGI",
      "regex": "X-TIGI"
  }
]
`)
