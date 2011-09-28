package textprocessing

//import "strings"
//import "fmt"


func (suffixArray *SuffixArray) GetSubstrings(lcp []int, minOcc, minLength, maxLength int) (suffixIndices, occurrences, lengths []int){
	for length:=minLength ; length <= maxLength ; length++ {
		for idx:=1 ; idx < len(lcp) ; idx++ {
			if lcp[idx] >= length {
				startIdx := idx-1
					for ; idx < len(lcp) && lcp[idx]>=length ; idx++ {}
						if idx-startIdx >= minOcc {
							suffixIndices = append(suffixIndices, startIdx)
							occurrences = append(occurrences, idx-startIdx)
							lengths = append(lengths, length)
				}
			}
		}
	}
	return 
}