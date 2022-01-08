// Code generated by elves, DO NOT EDIT.

package main

type register int

const (
	w register = iota
	x
	y
	z
)


func getBlocks() (blocks []func(inp int, s,n []state)) {
blocks = make([]func(inp int, s,n []state), 0, 14)
blocks = append(blocks, func(inp int, states, ns []state) {

		for i := range states {
			ns[i] = states[i]
			ns[i].regs[w] = inp
			ns[i].max = ns[i].max*10 + inp
			ns[i].min = ns[i].min*10 + inp
		}
for st := range ns {
    ns[st].regs[x] *= 0
    ns[st].regs[x] += ns[st].regs[z]
    ns[st].regs[x] %= 26
    ns[st].regs[z] /= 1
    ns[st].regs[x] += 13

				if ns[st].regs[x] == ns[st].regs[w] {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}

				if ns[st].regs[x] == 0 {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}
    ns[st].regs[y] *= 0
    ns[st].regs[y] += 25
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[y] += 1
    ns[st].regs[z] *= ns[st].regs[y]
    ns[st].regs[y] *= 0
    ns[st].regs[y] += ns[st].regs[w]
    ns[st].regs[y] += 6
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[z] += ns[st].regs[y]
}
})
blocks = append(blocks, func(inp int, states, ns []state) {

		for i := range states {
			ns[i] = states[i]
			ns[i].regs[w] = inp
			ns[i].max = ns[i].max*10 + inp
			ns[i].min = ns[i].min*10 + inp
		}
for st := range ns {
    ns[st].regs[x] *= 0
    ns[st].regs[x] += ns[st].regs[z]
    ns[st].regs[x] %= 26
    ns[st].regs[z] /= 1
    ns[st].regs[x] += 11

				if ns[st].regs[x] == ns[st].regs[w] {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}

				if ns[st].regs[x] == 0 {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}
    ns[st].regs[y] *= 0
    ns[st].regs[y] += 25
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[y] += 1
    ns[st].regs[z] *= ns[st].regs[y]
    ns[st].regs[y] *= 0
    ns[st].regs[y] += ns[st].regs[w]
    ns[st].regs[y] += 11
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[z] += ns[st].regs[y]
}
})
blocks = append(blocks, func(inp int, states, ns []state) {

		for i := range states {
			ns[i] = states[i]
			ns[i].regs[w] = inp
			ns[i].max = ns[i].max*10 + inp
			ns[i].min = ns[i].min*10 + inp
		}
for st := range ns {
    ns[st].regs[x] *= 0
    ns[st].regs[x] += ns[st].regs[z]
    ns[st].regs[x] %= 26
    ns[st].regs[z] /= 1
    ns[st].regs[x] += 12

				if ns[st].regs[x] == ns[st].regs[w] {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}

				if ns[st].regs[x] == 0 {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}
    ns[st].regs[y] *= 0
    ns[st].regs[y] += 25
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[y] += 1
    ns[st].regs[z] *= ns[st].regs[y]
    ns[st].regs[y] *= 0
    ns[st].regs[y] += ns[st].regs[w]
    ns[st].regs[y] += 5
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[z] += ns[st].regs[y]
}
})
blocks = append(blocks, func(inp int, states, ns []state) {

		for i := range states {
			ns[i] = states[i]
			ns[i].regs[w] = inp
			ns[i].max = ns[i].max*10 + inp
			ns[i].min = ns[i].min*10 + inp
		}
for st := range ns {
    ns[st].regs[x] *= 0
    ns[st].regs[x] += ns[st].regs[z]
    ns[st].regs[x] %= 26
    ns[st].regs[z] /= 1
    ns[st].regs[x] += 10

				if ns[st].regs[x] == ns[st].regs[w] {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}

				if ns[st].regs[x] == 0 {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}
    ns[st].regs[y] *= 0
    ns[st].regs[y] += 25
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[y] += 1
    ns[st].regs[z] *= ns[st].regs[y]
    ns[st].regs[y] *= 0
    ns[st].regs[y] += ns[st].regs[w]
    ns[st].regs[y] += 6
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[z] += ns[st].regs[y]
}
})
blocks = append(blocks, func(inp int, states, ns []state) {

		for i := range states {
			ns[i] = states[i]
			ns[i].regs[w] = inp
			ns[i].max = ns[i].max*10 + inp
			ns[i].min = ns[i].min*10 + inp
		}
for st := range ns {
    ns[st].regs[x] *= 0
    ns[st].regs[x] += ns[st].regs[z]
    ns[st].regs[x] %= 26
    ns[st].regs[z] /= 1
    ns[st].regs[x] += 14

				if ns[st].regs[x] == ns[st].regs[w] {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}

				if ns[st].regs[x] == 0 {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}
    ns[st].regs[y] *= 0
    ns[st].regs[y] += 25
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[y] += 1
    ns[st].regs[z] *= ns[st].regs[y]
    ns[st].regs[y] *= 0
    ns[st].regs[y] += ns[st].regs[w]
    ns[st].regs[y] += 8
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[z] += ns[st].regs[y]
}
})
blocks = append(blocks, func(inp int, states, ns []state) {

		for i := range states {
			ns[i] = states[i]
			ns[i].regs[w] = inp
			ns[i].max = ns[i].max*10 + inp
			ns[i].min = ns[i].min*10 + inp
		}
for st := range ns {
    ns[st].regs[x] *= 0
    ns[st].regs[x] += ns[st].regs[z]
    ns[st].regs[x] %= 26
    ns[st].regs[z] /= 26
    ns[st].regs[x] += -1

				if ns[st].regs[x] == ns[st].regs[w] {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}

				if ns[st].regs[x] == 0 {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}
    ns[st].regs[y] *= 0
    ns[st].regs[y] += 25
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[y] += 1
    ns[st].regs[z] *= ns[st].regs[y]
    ns[st].regs[y] *= 0
    ns[st].regs[y] += ns[st].regs[w]
    ns[st].regs[y] += 14
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[z] += ns[st].regs[y]
}
})
blocks = append(blocks, func(inp int, states, ns []state) {

		for i := range states {
			ns[i] = states[i]
			ns[i].regs[w] = inp
			ns[i].max = ns[i].max*10 + inp
			ns[i].min = ns[i].min*10 + inp
		}
for st := range ns {
    ns[st].regs[x] *= 0
    ns[st].regs[x] += ns[st].regs[z]
    ns[st].regs[x] %= 26
    ns[st].regs[z] /= 1
    ns[st].regs[x] += 14

				if ns[st].regs[x] == ns[st].regs[w] {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}

				if ns[st].regs[x] == 0 {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}
    ns[st].regs[y] *= 0
    ns[st].regs[y] += 25
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[y] += 1
    ns[st].regs[z] *= ns[st].regs[y]
    ns[st].regs[y] *= 0
    ns[st].regs[y] += ns[st].regs[w]
    ns[st].regs[y] += 9
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[z] += ns[st].regs[y]
}
})
blocks = append(blocks, func(inp int, states, ns []state) {

		for i := range states {
			ns[i] = states[i]
			ns[i].regs[w] = inp
			ns[i].max = ns[i].max*10 + inp
			ns[i].min = ns[i].min*10 + inp
		}
for st := range ns {
    ns[st].regs[x] *= 0
    ns[st].regs[x] += ns[st].regs[z]
    ns[st].regs[x] %= 26
    ns[st].regs[z] /= 26
    ns[st].regs[x] += -16

				if ns[st].regs[x] == ns[st].regs[w] {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}

				if ns[st].regs[x] == 0 {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}
    ns[st].regs[y] *= 0
    ns[st].regs[y] += 25
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[y] += 1
    ns[st].regs[z] *= ns[st].regs[y]
    ns[st].regs[y] *= 0
    ns[st].regs[y] += ns[st].regs[w]
    ns[st].regs[y] += 4
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[z] += ns[st].regs[y]
}
})
blocks = append(blocks, func(inp int, states, ns []state) {

		for i := range states {
			ns[i] = states[i]
			ns[i].regs[w] = inp
			ns[i].max = ns[i].max*10 + inp
			ns[i].min = ns[i].min*10 + inp
		}
for st := range ns {
    ns[st].regs[x] *= 0
    ns[st].regs[x] += ns[st].regs[z]
    ns[st].regs[x] %= 26
    ns[st].regs[z] /= 26
    ns[st].regs[x] += -8

				if ns[st].regs[x] == ns[st].regs[w] {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}

				if ns[st].regs[x] == 0 {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}
    ns[st].regs[y] *= 0
    ns[st].regs[y] += 25
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[y] += 1
    ns[st].regs[z] *= ns[st].regs[y]
    ns[st].regs[y] *= 0
    ns[st].regs[y] += ns[st].regs[w]
    ns[st].regs[y] += 7
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[z] += ns[st].regs[y]
}
})
blocks = append(blocks, func(inp int, states, ns []state) {

		for i := range states {
			ns[i] = states[i]
			ns[i].regs[w] = inp
			ns[i].max = ns[i].max*10 + inp
			ns[i].min = ns[i].min*10 + inp
		}
for st := range ns {
    ns[st].regs[x] *= 0
    ns[st].regs[x] += ns[st].regs[z]
    ns[st].regs[x] %= 26
    ns[st].regs[z] /= 1
    ns[st].regs[x] += 12

				if ns[st].regs[x] == ns[st].regs[w] {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}

				if ns[st].regs[x] == 0 {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}
    ns[st].regs[y] *= 0
    ns[st].regs[y] += 25
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[y] += 1
    ns[st].regs[z] *= ns[st].regs[y]
    ns[st].regs[y] *= 0
    ns[st].regs[y] += ns[st].regs[w]
    ns[st].regs[y] += 13
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[z] += ns[st].regs[y]
}
})
blocks = append(blocks, func(inp int, states, ns []state) {

		for i := range states {
			ns[i] = states[i]
			ns[i].regs[w] = inp
			ns[i].max = ns[i].max*10 + inp
			ns[i].min = ns[i].min*10 + inp
		}
for st := range ns {
    ns[st].regs[x] *= 0
    ns[st].regs[x] += ns[st].regs[z]
    ns[st].regs[x] %= 26
    ns[st].regs[z] /= 26
    ns[st].regs[x] += -16

				if ns[st].regs[x] == ns[st].regs[w] {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}

				if ns[st].regs[x] == 0 {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}
    ns[st].regs[y] *= 0
    ns[st].regs[y] += 25
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[y] += 1
    ns[st].regs[z] *= ns[st].regs[y]
    ns[st].regs[y] *= 0
    ns[st].regs[y] += ns[st].regs[w]
    ns[st].regs[y] += 11
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[z] += ns[st].regs[y]
}
})
blocks = append(blocks, func(inp int, states, ns []state) {

		for i := range states {
			ns[i] = states[i]
			ns[i].regs[w] = inp
			ns[i].max = ns[i].max*10 + inp
			ns[i].min = ns[i].min*10 + inp
		}
for st := range ns {
    ns[st].regs[x] *= 0
    ns[st].regs[x] += ns[st].regs[z]
    ns[st].regs[x] %= 26
    ns[st].regs[z] /= 26
    ns[st].regs[x] += -13

				if ns[st].regs[x] == ns[st].regs[w] {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}

				if ns[st].regs[x] == 0 {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}
    ns[st].regs[y] *= 0
    ns[st].regs[y] += 25
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[y] += 1
    ns[st].regs[z] *= ns[st].regs[y]
    ns[st].regs[y] *= 0
    ns[st].regs[y] += ns[st].regs[w]
    ns[st].regs[y] += 11
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[z] += ns[st].regs[y]
}
})
blocks = append(blocks, func(inp int, states, ns []state) {

		for i := range states {
			ns[i] = states[i]
			ns[i].regs[w] = inp
			ns[i].max = ns[i].max*10 + inp
			ns[i].min = ns[i].min*10 + inp
		}
for st := range ns {
    ns[st].regs[x] *= 0
    ns[st].regs[x] += ns[st].regs[z]
    ns[st].regs[x] %= 26
    ns[st].regs[z] /= 26
    ns[st].regs[x] += -6

				if ns[st].regs[x] == ns[st].regs[w] {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}

				if ns[st].regs[x] == 0 {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}
    ns[st].regs[y] *= 0
    ns[st].regs[y] += 25
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[y] += 1
    ns[st].regs[z] *= ns[st].regs[y]
    ns[st].regs[y] *= 0
    ns[st].regs[y] += ns[st].regs[w]
    ns[st].regs[y] += 6
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[z] += ns[st].regs[y]
}
})
blocks = append(blocks, func(inp int, states, ns []state) {

		for i := range states {
			ns[i] = states[i]
			ns[i].regs[w] = inp
			ns[i].max = ns[i].max*10 + inp
			ns[i].min = ns[i].min*10 + inp
		}
for st := range ns {
    ns[st].regs[x] *= 0
    ns[st].regs[x] += ns[st].regs[z]
    ns[st].regs[x] %= 26
    ns[st].regs[z] /= 26
    ns[st].regs[x] += -6

				if ns[st].regs[x] == ns[st].regs[w] {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}

				if ns[st].regs[x] == 0 {
					ns[st].regs[x] = 1
				} else {
					ns[st].regs[x] = 0
				}
    ns[st].regs[y] *= 0
    ns[st].regs[y] += 25
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[y] += 1
    ns[st].regs[z] *= ns[st].regs[y]
    ns[st].regs[y] *= 0
    ns[st].regs[y] += ns[st].regs[w]
    ns[st].regs[y] += 1
    ns[st].regs[y] *= ns[st].regs[x]
    ns[st].regs[z] += ns[st].regs[y]
}
})
return
}
func getPreSorts() (sorts []func(ns stateList)) {
sorts = append(sorts, func(ns stateList) {
				for i := range ns {
					ns[i].regs[w] = 0
				}
		})
sorts = append(sorts, func(ns stateList) {
				for i := range ns {
					ns[i].regs[w] = 0
				}
		})
sorts = append(sorts, func(ns stateList) {
				for i := range ns {
					ns[i].regs[w] = 0
				}
		})
sorts = append(sorts, func(ns stateList) {
				for i := range ns {
					ns[i].regs[w] = 0
				}
		})
sorts = append(sorts, func(ns stateList) {
				for i := range ns {
					ns[i].regs[w] = 0
				}
		})
sorts = append(sorts, func(ns stateList) {
				for i := range ns {
					ns[i].regs[w] = 0
				}
		})
sorts = append(sorts, func(ns stateList) {
				for i := range ns {
					ns[i].regs[w] = 0
				}
		})
sorts = append(sorts, func(ns stateList) {
				for i := range ns {
					ns[i].regs[w] = 0
				}
		})
sorts = append(sorts, func(ns stateList) {
				for i := range ns {
					ns[i].regs[w] = 0
				}
		})
sorts = append(sorts, func(ns stateList) {
				for i := range ns {
					ns[i].regs[w] = 0
				}
		})
sorts = append(sorts, func(ns stateList) {
				for i := range ns {
					ns[i].regs[w] = 0
				}
		})
sorts = append(sorts, func(ns stateList) {
				for i := range ns {
					ns[i].regs[w] = 0
				}
		})
sorts = append(sorts, func(ns stateList) {
				for i := range ns {
					ns[i].regs[w] = 0
				}
		})
sorts = append(sorts, func(ns stateList) {
				for i := range ns {
					ns[i].regs[w] = 0
				}
		})
return
}
