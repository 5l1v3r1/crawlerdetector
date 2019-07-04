package crawlerdetector

import (
	"regexp"
	"strings"
)

// CrawlerDetector is crawler detector structure
type CrawlerDetector struct {
	Crawlers   *regexp.Regexp
	Mobiles    *regexp.Regexp
	Exclusions *regexp.Regexp
	Matched    []string
}

// New returns a new initialized CrawlerDetector
func New() *CrawlerDetector {
	return &CrawlerDetector{
		Crawlers:   regexp.MustCompile(CombineRegexp(CrawlersList())),
		Mobiles:    regexp.MustCompile(CombineRegexp(MobilesList())),
		Exclusions: regexp.MustCompile(CombineRegexp(ExclusionsList())),
		Matched:    []string{},
	}
}

// IsCrawler is detect crawlers/spiders/bots by user agent
func (cd *CrawlerDetector) IsCrawler(userAgent string) bool {
	if cd.IsExclusion(userAgent) {
		return false
	}
	cd.Matched = cd.Crawlers.FindAllString(userAgent, -1)

	if len(cd.Matched) != 0 {
		return true
	}

	return false
}

// IsMobile is detect mobile device by user agent
func (cd *CrawlerDetector) IsMobile(userAgent string) bool {

	cd.Matched = cd.Mobiles.FindAllString(userAgent, -1)

	if len(cd.Matched) != 0 {
		return true
	}

	return false
}

// IsExclusion is detect exclusion from user agent
func (cd *CrawlerDetector) IsExclusion(userAgent string) bool {
	isExclusion := cd.Exclusions.ReplaceAllString(userAgent, "")

	if len(isExclusion) == 0 {
		return true
	}

	return false
}

// CombineRegexp is build regex from givement patterns list
func CombineRegexp(patterns []string) string {
	return "(" + strings.Join(patterns, "|") + ")"
}

// GetMatched is getter of matched result
func (cd *CrawlerDetector) GetMatched() []string {
	return cd.Matched
}
