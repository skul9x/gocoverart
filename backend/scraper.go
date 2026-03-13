package backend

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const googleSearchURL = "https://www.google.com/search"

var googleHeaders = map[string]string{
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
	"accept-language":           "en-US,en;q=0.9",
	"cache-control":             "max-age=0",
	"sec-ch-ua":                "\"Chromium\";v=\"143\", \"Not A(Brand\";v=\"24\", \"Google Chrome\";v=\"143\"",
	"sec-ch-ua-arch":           "\"x86\"",
	"sec-ch-ua-bitness":        "\"64\"",
	"sec-ch-ua-full-version":   "\"143.0.7499.41\"",
	"sec-ch-ua-mobile":         "?0",
	"sec-ch-ua-platform":        "\"Windows\"",
	"sec-fetch-dest":           "document",
	"sec-fetch-mode":           "navigate",
	"sec-fetch-site":           "none",
	"sec-fetch-user":           "?1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36",
}

var googleRegexPatterns = []*regexp.Regexp{
	regexp.MustCompile(`"ou":"(http[^"]+?)"`),
	regexp.MustCompile(`data-src="(http[^"]+?)"`),
	regexp.MustCompile(`src="(http[^"]+?)"`),
	regexp.MustCompile(`AF_initDataCallback\((.*?)\);`),
	regexp.MustCompile(`\["(http[^"]+?)",\d+,\d+\]`),
	regexp.MustCompile(`(?i)(https?://[^"]+\.(?:jpg|jpeg|png|webp))`),
}

func decodeUnicode(input string) string {
	// Simple unescape for common Google search JSON/HTML escapes
	res := input
	res = strings.ReplaceAll(res, "\\u003d", "=")
	res = strings.ReplaceAll(res, "\\u0026", "&")
	res = strings.ReplaceAll(res, "\\u0025", "%")
	res = strings.ReplaceAll(res, "\\/", "/")

	// Basic HTML entity unescape
	res = strings.ReplaceAll(res, "&amp;", "&")
	res = strings.ReplaceAll(res, "&lt;", "<")
	res = strings.ReplaceAll(res, "&gt;", ">")
	res = strings.ReplaceAll(res, "&quot;", "\"")
	res = strings.ReplaceAll(res, "&#39;", "'")

	return res
}

// SearchGoogleImage finds top 5 cover art images for a given query
func SearchGoogleImage(query string) ([]CoverResult, error) {
	// Add "cover art" to prioritize high quality images
	searchQuery := fmt.Sprintf("%s cover art", query)
	params := url.Values{}
	params.Add("q", searchQuery)
	params.Add("tbm", "isch") // Image search
	params.Add("safe", "active")

	req, err := http.NewRequest("GET", googleSearchURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	for k, v := range googleHeaders {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("google returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	html := string(body)

	uniqueUrls := make(map[string]bool)
	var candidates []string

	for _, re := range googleRegexPatterns {
		matches := re.FindAllStringSubmatch(html, -1)
		for _, m := range matches {
			if len(m) < 2 {
				continue
			}
			rawUrl := m[1]
			decoded := decodeUnicode(rawUrl)

			if strings.HasPrefix(decoded, "http") {
				lower := strings.ToLower(decoded)
				// Filter for image extensions
				if strings.Contains(lower, ".jpg") || strings.Contains(lower, ".jpeg") ||
					strings.Contains(lower, ".png") || strings.Contains(lower, ".webp") ||
					strings.Contains(lower, "encrypted-tbn0") {

					if !uniqueUrls[decoded] {
						uniqueUrls[decoded] = true
						candidates = append(candidates, decoded)
					}
				}
			}
		}
	}

	var results []CoverResult
	// Filter out thumbnails for the first batch, but keep them as fallback
	var highRes []string
	for _, url := range candidates {
		if !strings.Contains(url, "encrypted-tbn0") && !strings.Contains(url, "gstatic.com") {
			highRes = append(highRes, url)
		}
	}

	finalList := highRes
	if len(finalList) == 0 {
		finalList = candidates
	}

	for i, url := range finalList {
		if len(results) >= 5 {
			break
		}
		results = append(results, CoverResult{
			URL:     url,
			Preview: url,
			Rank:    i + 1,
		})
	}

	return results, nil
}
