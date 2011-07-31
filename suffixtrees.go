package textprocessing

import "strings"
import "sort"

type SuffixSortable interface {
	Less(int, int) bool
	Length() int
	SuffixString(int) string
}

type SuffixArray struct{
        i []int
        r SuffixSortable
}

func NewSuffixArray(str SuffixSortable) (suffixArray *SuffixArray) {
        suffixArray = new(SuffixArray)
        suffixArray.r=str
        suffixArray.i = make([]int, str.Length())
        for idx := range suffixArray.i {
                suffixArray.i[idx] = idx
        }
        sort.Sort(suffixArray)
        return
}

func (suffixArray *SuffixArray) Len() (int) {
        return len(suffixArray.i)
}


func (suffixArray *SuffixArray) Swap(m, n int) {
        suffixArray.i[m], suffixArray.i[n] = suffixArray.i[n], suffixArray.i[m]
}

func (suffixArray *SuffixArray) Lcp() (height []int){
	rank := make([]int, len(suffixArray.i))
	height = make([]int, len(suffixArray.i))
	for idx, _ := range suffixArray.i {
		rank[suffixArray.i[idx]] = idx
	}
	h := 0
	for i, _ := range rank {
		if rank[i] > 0 {
			k := suffixArray.i[rank[i]-1]
			for i+h < suffixArray.r.Length() && k+h < suffixArray.r.Length() && !suffixArray.r.Less(k+h, i+h) {
				h++
			}
			height[rank[i]] = h
			if h > 0 {
				h--
			}
		}
	}
	return height
}


func (suffixArray *SuffixArray) String() (string){
        array := make([]string, len(suffixArray.i))
        for idx, element := range suffixArray.i {
                array[idx] = suffixArray.r.SuffixString(len(suffixArray.i)-element)
        }
        return strings.Join(array, "\n")
}