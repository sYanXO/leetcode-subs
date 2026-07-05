
func topKFrequent(nums []int, k int) []int {

    type Pair struct{
    element int
    cnt int
    }
    freqMap := make(map[int]int)
    for _,n := range nums{
        freqMap[n]++
    }
    pairs := make([]Pair,0, len(freqMap)) 

    for element,cnt := range freqMap{
        pairs = append(pairs,Pair{element,cnt})
    } 

    sort.Slice(pairs,func (x,y int)bool{
        return pairs[x].cnt > pairs[y].cnt // descending
    })

    ans := make([]int,0,k)
    for i:=0; i<k; i++{
        ans = append(ans,pairs[i].element)
    }
    return ans
}