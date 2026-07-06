func removeCoveredIntervals(intervals [][]int) int {
    // intervals with same starts-
    // sort acc to 2nd ele, descending

    sort.Slice(intervals, func (i,j int)bool{
        if intervals[i][0] == intervals[j][0]{
            return intervals[i][1] > intervals[j][1] // 2nd ele - desc
        }
        return intervals[i][0] < intervals[j][0] // 1st ele - asc
    })

    cnt := 0
    prevEnd := 0

    // curr interval  be > prev interval -> covered .. ++
    for i := 0; i < len(intervals); i++{
        if intervals[i][1] > prevEnd{
            cnt++
            prevEnd = intervals[i][1]
        }
    }
    return cnt
}