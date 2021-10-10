func rangeBitwiseAnd(left int, right int) int {
    out := left
    
    for i := left + 1; i <= right; i++ {
        out &= i
    }

    return out
}

