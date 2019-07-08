package crawlerdetector

import (
	"regexp"
)

// Piwik struct to parse the yml
type Piwik struct {
	Regex string `yaml:"regex"`
}

// String function to dump the Regex
func (m *Piwik) String() string {
	return m.Regex
}

// CrawlerDetector is crawler detector structure
type CrawlerDetector struct {
	Crawlers   *regexp.Regexp
	Mobiles    *regexp.Regexp
	Exclusions *regexp.Regexp
	Matched    []string
	Browser    bool
	Crawler    bool
	Mobile     bool
}

// New returns a new initialized CrawlerDetector
func New() *CrawlerDetector {
	crw := &CrawlerDetector{
		Crawlers:   regexp.MustCompile(CrawlersList()),
		Mobiles:    regexp.MustCompile(MobilesList()),
		Exclusions: regexp.MustCompile(ExclusionsList()),
		Matched:    []string{},
		Browser:    false,
		Mobile:     false,
		Crawler:    false,
	}
	return crw
}

// Parse is to  perform all operations by user agent
func (cd *CrawlerDetector) Parse(userAgent string) *CrawlerDetector {
	isExclusion := cd.Exclusions.ReplaceAllString(userAgent, "")
	cd.Browser = (len(isExclusion) == 0)

	cd.Matched = cd.Mobiles.FindAllString(userAgent, -1)
	cd.Mobile = (len(cd.Matched) != 0)

	if !cd.Browser {
		cd.Matched = cd.Crawlers.FindAllString(userAgent, -1)
		cd.Crawler = (len(cd.Matched) != 0)
	}

	return cd
}

// IsCrawler is detect crawlers/spiders/bots by user agent
func (cd *CrawlerDetector) IsCrawler(userAgent string) bool {
	cd.Crawler = false
	if cd.IsExclusion(userAgent) {
		return false
	}
	cd.Matched = cd.Crawlers.FindAllString(userAgent, -1)

	if len(cd.Matched) != 0 {
		cd.Crawler = true
	}

	return cd.Crawler
}

// IsMobile is detect mobile device by user agent
func (cd *CrawlerDetector) IsMobile(userAgent string) bool {
	cd.Mobile = false
	cd.Matched = cd.Mobiles.FindAllString(userAgent, -1)

	if len(cd.Matched) != 0 {
		cd.Mobile = true
	}

	return cd.Mobile
}

// IsExclusion is detect exclusion from user agent
func (cd *CrawlerDetector) IsExclusion(userAgent string) bool {
	cd.Browser = false
	isExclusion := cd.Exclusions.ReplaceAllString(userAgent, "")

	if len(isExclusion) == 0 {
		cd.Browser = true
	}

	return cd.Browser
}

// GetMatched is getter of matched result
func (cd *CrawlerDetector) GetMatched() []string {
	return cd.Matched
}
