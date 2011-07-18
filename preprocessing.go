package textprocessing

//import "fmt"
import "strings"
//import "sort"
import "io/ioutil"
import "regexp"
// import "math"


type Word string

type Text []Word

func NewTextFile(filename string) (text Text) {
    rawTextByte,_ := ioutil.ReadFile(filename)
    text = NewText(string(rawTextByte))
    return
}

func NewText(str string) (text Text) {
    const punct = "[?!.,;:^/\\<>//(//)\"\t\n]*"
    reStrip := regexp.MustCompile(punct+`[ \t\n]+`+punct)
    textStringSlice := strings.Split(reStrip.ReplaceAllString(str, " "), " ", -1)
    text = make([]Word,len(textStringSlice))
    for idx,word := range textStringSlice {
        text[idx]=Word(strings.ToLower(word))
        }
    return
}