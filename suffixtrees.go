package textprocessing

import "strings"
import "sort"
import "fmt"

type SuffixArrayInterfact interface {
	Less()
}

type SuffixArray struct{
        i []int
        r *Text
}

func NewSuffixArray(str *Text) (suffixArray *SuffixArray) {
        suffixArray = new(SuffixArray)
        suffixArray.r=str
        suffixArray.i = make([]int, str.Length())
        for idx := range suffixArray.i {
                suffixArray.i[idx] = idx
        }
        sort.Sort(suffixArray)
        fmt.Println(suffixArray.i)
        fmt.Println(suffixArray.r)
        return
}

func (suffixArray *SuffixArray) Len() (int) {
        return len(suffixArray.i)
}

func (suffixArray *SuffixArray) Less(m, n int) (bool) {
        return suffixArray.r.Less(len(suffixArray.i),suffixArray.i[m],suffixArray.i[n])
}

func (suffixArray *SuffixArray) Swap(m, n int) {
        suffixArray.i[m], suffixArray.i[n] = suffixArray.i[n], suffixArray.i[m]
}

// func (suffixArray *SuffixArray) SuffixOfLength(length int) (string) {
// 	return strings.Join(*suffixArray.r[suffixArray.r.Length()-length:], " ")
// }

func (suffixArray *SuffixArray) String() (string){
	suffixes := strings.Split(suffixArray.r.SuffixString(), "\n", -1)
	array := make([]string, len(suffixes))
        for idx,val := range suffixArray.i {
		array[idx] = suffixes[val]
        }
        sorted := strings.Join(array, "\n")
        return string(sorted)
}

// func (suffixArray *SuffixArray) Less(m, n int) (bool) {
//         midx:=suffixArray.i[m]
//         nidx:=suffixArray.i[n]
//         for{
//                 switch {
//                         case suffixArray.s[midx]<suffixArray.s[nidx] :
//                                 return true
//                         case suffixArray.s[midx]>suffixArray.s[nidx] :
//                                 return false
//                         case suffixArray.s[midx]==suffixArray.s[nidx] :
//                                 midx++
//                                 nidx++
//                                 if nidx == len(suffixArray.i) {
//                                         return false
//                                 } else if midx == len(suffixArray.i) {
//                                         return true
//                                 }
//                 }
//         }
//         panic("unreachable")
// }
