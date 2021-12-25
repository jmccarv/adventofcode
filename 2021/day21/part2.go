package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	//"sync"
)

// I needed help with this one!
// See: https://www.youtube.com/watch?v=a6ZdJEntKkk
// And: https://www.reddit.com/r/adventofcode/comments/rm7ygy/2021_day_21_part2_could_someone_explain_the/

// All we need to track is the position and score for each player
type gameState struct {
	pos   [2]int
	score [2]int
}

type answer [2]int

var seen map[gameState]answer
var ans answer // our counts of each player's wins

func main() {
	var game gameState
	seen = make(map[gameState]answer)
	input := bufio.NewScanner(os.Stdin)

	re := regexp.MustCompile(`Player *(\d) starting position: *(\d+)`)
	for input.Scan() {
		m := re.FindStringSubmatch(input.Text())
		if len(m) != 3 {
			panic("Invalid input")
		}
		game.pos[atoi(m[1])-1] = atoi(m[2]) - 1
	}
	fmt.Println(game)

	ans := game.countWin()
	fmt.Println(ans, ans.max())
}

// There are a limited number of states that the game can ever be in.
// The only variables that matter are each players' position and score
//
// So, each player can be in one of 10 positions, and have a score from 0-20 (21 is a winning score, the game is never in that state)
// = 10 * 10 * 21 * 21 = 44100
//
// This means that we can cache the winners for each possible state, as we will
// encounter each possible state many, *many* times...
func (g gameState) countWin() (ans answer) {
	if g.score[0] >= 21 {
		return [2]int{1, 0}
	} else if g.score[1] >= 21 {
		return [2]int{0, 1}
	}

	// Have we seen this state before, if so we've already computed the answer
	if ans, ok := seen[g]; ok {
		return ans
	}

	// Compute counts for all possible combinations of three rolls of a three sided die
	for d1 := 1; d1 < 4; d1++ {
		for d2 := 1; d2 < 4; d2++ {
			for d3 := 1; d3 < 4; d3++ {
				// This player takes their turn
				newGame := g
				newGame.pos[0] = (newGame.pos[0] + d1 + d2 + d3) % 10
				newGame.score[0] += newGame.pos[0] + 1

				// swapping these allows the next player to take their turn
				x := gameState{swap(newGame.pos), swap(newGame.score)}
				a := x.countWin()
				ans[0] = ans[0] + a[1]
				ans[1] = ans[1] + a[0]
			}
		}
	}
	seen[g] = ans
	return
}

func (a answer) max() int {
	if a[0] > a[1] {
		return a[0]
	}
	return a[1]
}

func swap(a [2]int) (b [2]int) {
	b[0] = a[1]
	b[1] = a[0]
	return
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		panic("Invalid input")
	}
	return i
}
