package utils

import (
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

func Content2progress(content string) string {
	regex := regexp.MustCompile(`\(([^)]+)\)`) // matches the value inside the first parenthesis
	match := regex.FindStringSubmatch(content)
	progress := ""
	if len(match) > 1 {
		progress = match[1]
	}
	return progress
}

func UriToHash(uri string) string {
	parts := strings.Split(uri, "_")
	filename := parts[len(parts)-1]
	extension := filepath.Ext(filename)
	return strings.TrimSuffix(filename, extension)
}

func Content2pormpt(content string) string {
	pattern := regexp.MustCompile(`\*\*(.*?)\*\*`)
	matches := pattern.FindStringSubmatch(content)
	if len(matches) > 1 {
		return matches[1]
	} else {
		log.Println("No match found.", content)
		return content
	}
}
