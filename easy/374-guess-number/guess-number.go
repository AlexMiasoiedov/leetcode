/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    var try, gs, half int

    try = n
    half = try

    for {
        gs = guess(try)

        if gs == 0 {
            break
        } else if gs < 0 {
            half = (try / 2) + (try % 2)
            try = (try / 2)
        } else {
            try = half + (try / 2)
        }
    }

    return try
}
