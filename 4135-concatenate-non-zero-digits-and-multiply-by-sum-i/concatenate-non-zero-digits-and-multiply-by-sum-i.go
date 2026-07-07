func sumAndMultiply(n int) int64 {
    s := strconv.Itoa(n)
    var new_s string
    sum := 0
    for _,ch := range s{
        digit := int(ch-'0')
        if digit != 0{
            new_s += string(ch)
            sum += digit
        }
    }
    ans,_:= strconv.Atoi(new_s)
    return int64(ans*sum)
}