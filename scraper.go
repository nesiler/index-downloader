// Description: This file contains the scraper methods for the Index Downloader.
// Author: Enes Diler | Github: @nesiler | Website: nesiler.com | mail: me@nesiler.com
package main

import (
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// Scraper holds the Colly collector and other settings.
type Scraper struct {
	collector *colly.Collector
	baseURL   *url.URL
}

// initializes a new Scraper with the given URL.
func newScraper(siteURL string) (*Scraper, error) {
	parsedURL, err := url.Parse(siteURL)
	if err != nil {
		return nil, err
	}

	c := colly.NewCollector(
		colly.AllowedDomains(parsedURL.Host),
		// colly.MaxDepth(0), // No depth limit, remove comment to set a depth limit
	)

	extensions.RandomUserAgent(c)

	s := &Scraper{
		collector: c,
		baseURL:   parsedURL,
	}

	s.setupCallbacks()

	return s, nil
}

// sets up the necessary callbacks for the collector.
func (s *Scraper) setupCallbacks() {
	// Restrict the collector to only visit URLs that are within the base URL domain and path
	baseDomain := s.baseURL.Hostname()
	basePath := s.baseURL.Path
	if basePath == "" {
		basePath = "/"
	}

	// Compile a regex that matches URLs within the base domain and path
	urlRegex, err := regexp.Compile("https?://" + regexp.QuoteMeta(baseDomain) + regexp.QuoteMeta(basePath) + ".*")
	if err != nil {
		log.Fatalf("Failed to compile URL filter regex: %v", err)
	}

	// Assign the compiled regex to the URLFilters
	s.collector.URLFilters = []*regexp.Regexp{urlRegex}

	// Find all links and visit them
	s.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.HasSuffix(link, ".pdf") {
			// enqueue PDF link for download
			s.collector.Request("GET", link, nil, e.Request.Ctx, nil)
		} else {
			// Visit link only if it's not a file (e.g., .pdf, .jpg, etc.)
			if !strings.Contains(filepath.Ext(link), ".") {
				e.Request.Visit(link)
			}
		}
	})

	// Download PDF files
	s.collector.OnResponse(func(r *colly.Response) {
		if strings.HasSuffix(r.Request.URL.String(), ".pdf") {
			s.savePDF(r)
		}
	})
}

// saves the PDF from the response.
func (s *Scraper) savePDF(r *colly.Response) {
	pdfURL := r.Request.URL.String()
	parsedPDFURL, err := url.Parse(pdfURL)
	if err != nil {
		log.Println("Invalid PDF URL:", pdfURL)
		return
	}

	// Create a subfolder based on the path of the URL
	subfolder := strings.Trim(parsedPDFURL.Path, "/")
	subfolderPath := filepath.Join("downloads", filepath.Dir(subfolder))

	// Ensure the subfolder exists
	if err := os.MkdirAll(subfolderPath, os.ModePerm); err != nil {
		log.Println("Error creating directory:", err)
		return
	}

	// Set the file name and path
	fileName := filepath.Base(parsedPDFURL.Path)
	filePath := filepath.Join(subfolderPath, fileName)

	// Check if the file already exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Save the file
		err = r.Save(filePath)
		if err != nil {
			log.Println("Error saving PDF:", err)
		} else {
			log.Println("Saved PDF:", filePath)
		}
	} else {
		log.Println("File already exists:", filePath)
	}
}

// go go go!
func (s *Scraper) Start() {
	s.collector.Visit(s.baseURL.String())
}
