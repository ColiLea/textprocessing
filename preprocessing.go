package textprocessing

import "strings"
import "io/ioutil"
import "regexp"
//import "fmt"


type Word string

type Text []Word

func NewTextFile(filename string) (text *Text) {
	rawTextByte,_ := ioutil.ReadFile(filename)
	text = NewText(string(rawTextByte))
	return
}

func NewText(str string) (text *Text) {
	const punct = "[?!.,;:^/\\<>\\(\\)\"]*"
	reStrip := regexp.MustCompile(punct+"([ \t\n]+"+punct+")+")
	str = reStrip.ReplaceAllString(str, " ")
	str = strings.Trim(str, " \n\t?!.,;:^/\\<>()\"")
	var textStringSlice []string
	if str != "" {
		textStringSlice = strings.Split(str, " ")
	}
	slice := make(Text,len(textStringSlice))
	text = &slice
	for idx,word := range textStringSlice {
		(*text)[idx]=Word(strings.ToLower(word))
	}
	return
}

func (text *Text) Length() (int) {
	return len(*text)
}

func (suffixArray *SuffixArray) Less(m, n int) (bool) {
	text := suffixArray.r
	m,n=suffixArray.i[m],suffixArray.i[n]
	for{
		switch {
			case text.Less(m,n):
				return true
			case text.Less(n,m) :
				return false
			default :
				m++
				n++
				if n == text.Length() {
					return false
				} else if m == text.Length() {
					return true
				}
		}
	}
	panic("unreachable")
}

func (text Text) Less(m, n int) (bool) {
	return text[m] < text[n]
}

func (text *Text) SuffixString(length int) (string) {
	suffix := make([]string, length)
	for index,word := range (*text)[text.Length()-length:] {
		suffix[index] = string(word)
	}
	return strings.Join(suffix, " ")
}