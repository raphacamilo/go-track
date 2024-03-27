package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC]|\[DBG\]|\[INF\]|\[WRN]|\[ERR]|\[FTL]).*$`)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`".*\b(?i:password)\b.*?"`)
	counter := 0

	for _, log := range lines {
		if re.MatchString(log) {
			counter++
		}
	}

	return counter
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]{1,}`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s{1,}([A-Za-z0-9]{1,})`)

	for i, log := range lines {
		foundedLine := re.FindString(log)

		if len(foundedLine) > 0 {
			lines[i] = re.ReplaceAllString(foundedLine, "[USR] $1 ") + log
		}
	}

	return lines
}
