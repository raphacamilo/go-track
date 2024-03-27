package logs

import "unicode/utf8"

var applications = map[rune]string{
	10071:  "recommendation", // ❗ U+2757
	128269: "search",         // 🔍 U+1F50D
	9728:   "weather",        // ☀ U+2600
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		label, exists := applications[char]

		if exists {
			return label
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	modifiedLog := ""

	for _, char := range log {
		if char == oldRune {
			modifiedLog += string(newRune)
		} else {
			modifiedLog += string(char)
		}
	}

	return modifiedLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
