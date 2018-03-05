// Package parser implements functions to parse Spectator article text.
package parser

import (
	"fmt"
	"regexp"
	"strings"
)

var slugPattern *regexp.Regexp = regexp.MustCompile(`(?i)(outquote(\(s\))?s?:)|(focus\s+sentence:)|(word(s)?:?\s\d{2,4})|(\d{2,4}\swords[^\.])|(word count:?\s?\d{2,4})|focus:|article:`)


// Paddings are patterns we want to remove from the desired value
// (e.g. "Title: ", "Outquote(s): ")
var bylinePadding *regexp.Regexp = regexp.MustCompile(`By:/\s+`)
var focusPadding *regexp.Regexp = regexp.MustCompile(`(?i)Focus Sentence:?\s+`)
var titlePadding *regexp.Regexp = regexp.MustCompile(`Title:\s+`)

// ArticleAttributes finds the articles of an article for posting.
// It returns attributes.
func ArticleAttributes(text string) (attrs map[string]interface{}) {
	attrs = make(map[string]interface{})
	text = strings.TrimSpace(text)

	rawLines := strings.Split(text, "\n")
	content := make([]string, 0)

	attrs["title"] = titlePadding.ReplaceAllString(rawLines[0], "")

	// Start from the end of the article and add lines of content until we reach
	// the slug (header).
	i := len(rawLines) - 1
	for ; i >= 0; i-- {
		line := strings.TrimSpace(rawLines[i])
		if len(line) == 0 {
			continue
		}
		if len(slugPattern.FindStringSubmatch(line)) > 0 {
			break
		}
		// Prepend line to content
		content = append([]string{line}, content...)
	}

	attrs["content"] = fmt.Sprintf("<p>%s</p>", strings.Join(content, "</p><p>"))

	// After we've found the last index of the slug, read article information from
	// the top to the end of the slug.
	for j := 1; j <= i; j++ {
		line := strings.TrimSpace(rawLines[j])
		if len(line) == 0 {
			continue
		}

		if len(bylinePadding.FindStringSubmatch(line)) > 0 {
			contributors := ArticleContributors(line)
		}
	}

	return
}

// ArticleContributors finds the contributors in a byline.
// It returns the contributors.
func ArticleContributors(byline string) (contributors []map[string]string) {
	contributors = make([]map[string]string, 0)
	byline = bylinePadding.ReplaceAllString(byline, "")

	cutoff := 0
	for i, c := range byline {
		if c == // need to find all string for the byline match in backup.python
	}
    for i in range(0, len(byline)):
        if byline[i] in ',&' or byline[i] == 'and':
            name = clean_name(' '.join(byline[cutoff:i]))
            contributors.append(name)
            cutoff = i + 1
    contributors.append(clean_name(' '.join(
        byline[cutoff:])))
}

// nameVariables splits a name of variable length into a first name and a last
// name and removes nickname formatting (e.g. Ying Zi (Jessy) Mei).
// It returns the formatted name as a map with a first_name and last_name.
func nameVariables(name string) map[string]string {
	variables := make(map[string]string)

	// Remove redundant spaces and nicknames
	name = strings.Join(strings.Fields(name), " ")
	// name = nicknamePattern.ReplaceAllString(name, "")

	var first_name, last_name string
	components := strings.Split(name, " ")
	if len(name) == 0 {
		log.Fatalf("No name given or whole name faulty.")
	} else	if len(components) == 1 {
		first_name, last_name = name, name
	} else if len(components) > 2 {
		first_name = strings.Join(components[0:len(components)-1], " ")
		last_name = components[len(components)-1]
	} else {
		first_name, last_name = components[0], components[1]
	}

	variables["first_name"], variables["last_name"] = first_name, last_name

	return variables
}
