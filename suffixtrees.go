package textprocessing

import "strings"
import "sort"

type SuffixArrayInterface interface {
	Less()
}

type SuffixArray struct{
        i []int
        r *Text
        lcp []int
}

func NewSuffixArray(str *Text) (suffixArray *SuffixArray) {
        suffixArray = new(SuffixArray)
        suffixArray.r=str
        suffixArray.i = make([]int, str.Length())
        suffixArray.lcp = make([]int, str.Length())
        for idx := range suffixArray.i {
                suffixArray.i[idx] = idx
        }
        sort.Sort(suffixArray)
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

func (suffixArray *SuffixArray) String() (string){
        array := make([]string, len(suffixArray.i))
        for idx, element := range suffixArray.i {
                array[idx] = suffixArray.r.SuffixString(len(suffixArray.i)-element)
        }
        return strings.Join(array, "\n")
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

// func (suffixArray *SuffixArray) LCP() (lcp []int) {
// 	lcp[0] = 0
//         for idx,val := range suffixArray.i {
// 		if *suffixArray.r[suffixArray.i[idx+1]:].Length() <= *suffixArray.r[suffixArray.i[idx]:].Length() {
// 			shortestSeq := *suffixArray.r[suffixArray.i[idx+1]:]
// 		} else {
// 			shortestSeq := *suffixArray.r[suffixArray.i[idx]:]
// 		}
// 		for index, position := range  shortestSeq {
// 			if *suffixArray.r[suffixArray.i[val]][position] == *suffixArray.r[suffixArray.i[val-1]][position] {
// 					lcp[idx+1]++
// 			}
// 		}
// 	}
// }