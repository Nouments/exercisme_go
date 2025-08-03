package pangram

import "strings"

func IsPangram(input string) bool {
	alphabet:="abcdefghijklmnopqrstuvwxyz"
    for _,i:=range(alphabet){
        if !strings.Contains(strings.ToLower(input),string(i)){
            return false
        }
    }
    return true
}
