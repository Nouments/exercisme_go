package logs

import "strings"

// Application identifies the application emitting the given log.
func Application(log string) string {
    runes := []rune(log)
    if runes[0]=='❗'{
        return "recommendation"
    }
    if strings.Contains(log,"🔍"){
        return "search"
    }
    if strings.Contains(log,"☀"){
        return "weather"
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	a:=[]rune(log)
    for i:=0; i<len(a);i++{
        if a[i]==oldRune{
            a[i]=newRune
        }
    }
    return string(a)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log))<=limit
}
