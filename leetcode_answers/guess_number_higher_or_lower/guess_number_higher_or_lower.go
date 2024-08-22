package main

func main() {}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	num := n
	for {
		result := guessNumber(num)
		if result == -1 {
			num--
			continue
		}
		if result == 1 {
			num++
			continue
		}
		if result == 0 {
			return num
		}
	}
}
