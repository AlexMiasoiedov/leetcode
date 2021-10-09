func findWords(board [][]byte, words []string) []string {
    var outWords []string = make([]string, 0, len(words))

    OUTER:
    for _, word := range words {
        for y := 0; y < len(board); y++ {
            for x := 0; x < len(board[y]); x++ {
                if backtrack(board, word, x, y) {
                    outWords = append(outWords, word)
                    continue OUTER
                }
            }
        }
    }

    return outWords
}

func backtrack(board [][]byte, word string, x, y int) bool {
    var b byte

    if len(word) < 1 { return true }
    if y < 0 || y > len(board)-1 { return false }
    if x < 0 || x > len(board[0])-1 { return false }
    if board[y][x] != word[0] { return false }

    board[y][x] = b

    if backtrack(board, word[1:], x, y-1) {
        board[y][x] = word[0]
        return true
    }
    if backtrack(board, word[1:], x+1, y) {
        board[y][x] = word[0]
        return true
    }
    if backtrack(board, word[1:], x, y+1) {
        board[y][x] = word[0]
        return true
    }
    if backtrack(board, word[1:], x-1, y) {
        board[y][x] = word[0]
        return true
    }

    board[y][x] = word[0]

    return false
}
