// Bad solution
func findDisappearedNumbers(nums []int) []int {
    var out []int = intRange(0, len(nums))

    for i := 0; i < len(nums); i++ {
        out[nums[i]-1] = 0
    }

    i := 0
    for {
        if i >= len(out) { break }
        
        if out[i] == 0 {
            out = append(out[:i], out[i+1:]...)
        } else {
            i++
        }
    }

    return out
}

func intRange (start, end int) []int {
    var out []int = make([]int, end - start)

    for i := 0; i < end - start; i++ {
        out[i] = i+1
    }

    return out
}
