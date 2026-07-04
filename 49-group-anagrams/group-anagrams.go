func groupAnagrams(strs []string) [][]string {
    anagramMap := make(map[string][]string)

    for _,s := range strs{
        runes := []rune(s) // list of runes are a word

        sort.Slice(runes, func(i,j int)bool{
            return runes[i] < runes[j]
        }) // sort lexicographicaly - acts as key

        key := string(runes)

        anagramMap[key] = append(anagramMap[key],s)

    }
        ans := make([][]string, 0, len(anagramMap)) // list of a list of strings, size same as map and init at 0

        for _,gr := range anagramMap{
            ans = append(ans,gr) // appends, list of values one by one, a key at a time
        }
        
    return ans
    
}