func canJump(nums []int) bool {
    var toJump int = len(nums)-1

    // Start from the end
    // check if toJump is reachable from fromJump
    // if yes -- set new toJump
    for fromJump := len(nums)-2; fromJump >= 0; fromJump-- {
        if fromJump+nums[fromJump] >= toJump { toJump = fromJump }
    }

    return toJump == 0
}
