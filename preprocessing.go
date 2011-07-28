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
    const punct = "[?!.,;:^/\\<>//(//)\"\t\n]*"
    reStrip := regexp.MustCompile(punct+`[ \t\n]+`+punct)
    textStringSlice := strings.Split(reStrip.ReplaceAllString(str, " "), " ", -1)
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

func (text Text) Less(length, m, n int) (bool) {
        for{
                switch {
                        case text[m]<text[n] :
                                return true
                        case text[m]>text[n] :
                                return false
                        case text[m]==text[n] :
                                m++
                                n++
                                if n == length {
                                        return false
                                } else if m == length {
                                        return true
                                }
                }
        }
        panic("unreachable")
}

func (text *Text) SuffixString() (string){
	suffixes := make([]string, text.Length())
	for idx,_ := range *text {
		tmp := make([]string, (*text).Length()-idx)
		for index,word := range (*text)[idx:] {
			tmp[index] = string(word)
		suffixes[idx] = strings.Join(tmp," ")
		}
	}
	return strings.Join(suffixes, "\n")
}
