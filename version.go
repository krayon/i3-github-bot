package githubbot

import (
	"log"
	"regexp"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

var (
	reMajorVersion = regexp.MustCompile(`(i3|i3status|i3lock):?\s*(?:version|v|vers|ver)?:?\s*(3\.[a-e]|3\.\p{Greek}|[0-9]\.[0-9]+)`)
)

// extractVersion extracts all (i3|i3status|i3lock) versions out of |body| and
// returns the highest version (numerically sorted).
func extractVersion(body string) []string {
	allmatches := reMajorVersion.FindAllStringSubmatch(body, -1)
	if len(allmatches) == 0 {
		return []string{}
	}
	versions := make([]string, len(allmatches))
	firstProgram := allmatches[0][1]
	for idx, match := range allmatches {
		log.Printf("match = %v\n", match)
		if match[1] != firstProgram {
			// |body| contains versions for multiple programs (e.g. i3
			// and i3lock). Just return the first one for now.
			return allmatches[0]
		}
		versions[idx] = match[2]
	}
	collate.New(language.Und, collate.Numeric).SortStrings(versions)
	return []string{"", firstProgram, versions[len(versions)-1]}
}
