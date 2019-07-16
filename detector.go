package crawlerdetector

import (
	"regexp"
	"sync"
)

var w sync.WaitGroup

// Piwik struct to parse the yml
type Piwik struct {
	Name  string `yaml:"name"`
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

// NewPiwik returns a new initialized CrawlerDetector from Piwik
func NewPiwik() *CrawlerDetector {
	crw := &CrawlerDetector{
		Crawlers:   regexp.MustCompile(PiwikCrawlersList()),
		Mobiles:    regexp.MustCompile(PiwikMobilesList()),
		Exclusions: regexp.MustCompile(ShortExclusionsList()),
		Matched:    []string{},
		Browser:    false,
		Mobile:     false,
		Crawler:    false,
	}
	return crw
}

// NewShort returns a new basic initialized CrawlerDetector
func NewShort() *CrawlerDetector {
	crw := &CrawlerDetector{
		Crawlers:   regexp.MustCompile(ShortCrawlersList()),
		Mobiles:    regexp.MustCompile(ShortMobilesList()),
		Exclusions: regexp.MustCompile(ShortExclusionsList()),
		Matched:    []string{},
		Browser:    false,
		Mobile:     false,
		Crawler:    false,
	}
	return crw
}

// Parse is to perform all operations by user agent
func (cd *CrawlerDetector) Parse(userAgent string) *CrawlerDetector {
	w.Add(3)
	go func() {
		cd.Browser = (len(cd.Exclusions.ReplaceAllString(userAgent, "")) == 0)
		defer w.Done()
	}()
	go func() {
		cd.Mobile = (len(cd.Mobiles.FindAllString(userAgent, -1)) != 0)
		defer w.Done()
	}()
	go func() {
		cd.Crawler = (len(cd.Crawlers.FindAllString(userAgent, -1)) != 0)
		defer w.Done()
	}()
	w.Wait()
	return cd
}

// ParseUnsafe is to perform all browser and mobile operations by user agent but if not is a browser we asume that is a crawler
func (cd *CrawlerDetector) ParseUnsafe(userAgent string) *CrawlerDetector {
	w.Add(2)
	go func() {
		cd.IsCrawler(userAgent)
		defer w.Done()
	}()
	go func() {
		cd.IsMobile(userAgent)
		defer w.Done()
	}()
	w.Wait()
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
