package textprocessing

//import "strings"
//import "fmt"


func (suffixArray *SuffixArray) GetSubstrings(lcp []int, length, minOcc int) (suffixIndices, occurrences []int){
	for idx:=1 ; idx < len(lcp) ; idx++ {
		if lcp[idx] >= length {
			startIdx := idx-1
			for ; idx < len(lcp) && lcp[idx]>=length ; idx++ {}
			if idx-startIdx >= minOcc {
				suffixIndices = append(suffixIndices, startIdx)
				occurrences = append(occurrences, idx-startIdx)
			}
		}
	}
	return 
}