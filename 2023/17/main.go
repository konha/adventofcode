package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("filename is empty")
		return
	}
	filename := os.Args[1]

	lines, err := readFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Least Heat Loss Part 1:", leastHeatLoss(lines))
	fmt.Println("Least Heat Loss Part 2:", leastHeatLoss2(lines))
}

type point struct {
	x int
	y int
}

func (p point) move(dir direction) point {
	switch dir {
	case up:
		return point{p.x, p.y - 1}
	case down:
		return point{p.x, p.y + 1}
	case left:
		return point{p.x - 1, p.y}
	case right:
		return point{p.x + 1, p.y}
	}
	return p
}

type grid map[point]int

type direction int

const (
	up direction = iota
	down
	left
	right
)

func (d direction) turn(dir direction) direction {
	switch dir {
	case left:
		switch d {
		case up:
			return left
		case down:
			return right
		case left:
			return down
		case right:
			return up
		}
	case right:
		switch d {
		case up:
			return right
		case down:
			return left
		case left:
			return up
		case right:
			return down
		}
	}
	return d
}

func leastHeatLoss(lines []string) int {
	g, end := parse(lines)
	return bfs(g, 0, 3, end)
}

func leastHeatLoss2(lines []string) int {
	g, end := parse(lines)
	return bfs(g, 4, 10, end)
}

type queueItem struct {
	p        point
	dir      direction
	heatLoss int
	straight int
	prev     *queueItem
}

type cacheItem struct {
	p        point
	dir      direction
	straight int
}

func bfs(g grid, minStraight, maxStraight int, end point) int {
	q := []queueItem{}
	q = append(q, queueItem{
		p:        point{1, 0},
		dir:      right,
		straight: 1,
	})
	q = append(q, queueItem{
		p:        point{0, 1},
		dir:      down,
		straight: 1,
	})
	cache := make(map[cacheItem]int)

	// while queue is not empty
	for len(q) > 0 {

		// get an item from the queue
		sort.Slice(q, func(i, j int) bool {
			return q[i].heatLoss < q[j].heatLoss
		})
		item := q[0]
		q = q[1:]

		// is this on the grid? check the map for an entry.
		if _, ok := g[item.p]; !ok {
			continue
		}

		heat := item.heatLoss + g[item.p]

		// are we at the end?
		if item.p == end && item.straight >= minStraight {
			return heat
		}

		cacheItem := cacheItem{
			p:        item.p,
			dir:      item.dir,
			straight: item.straight,
		}
		// already in the cache?
		if value, ok := cache[cacheItem]; ok {
			if value <= heat {
				continue
			}
		}
		cache[cacheItem] = heat

		if item.straight < maxStraight {
			// move forward
			qi := queueItem{
				p:        item.p.move(item.dir),
				dir:      item.dir,
				heatLoss: heat,
				straight: item.straight + 1,
				prev:     &item,
			}
			q = append(q, qi)
		}

		if item.straight >= minStraight {
			// turn left
			l := item.dir.turn(left)
			qi := queueItem{
				p:        item.p.move(l),
				dir:      l,
				heatLoss: heat,
				straight: 1,
				prev:     &item,
			}
			q = append(q, qi)

			// turn right
			r := item.dir.turn(right)
			qi = queueItem{
				p:        item.p.move(r),
				dir:      r,
				heatLoss: heat,
				straight: 1,
				prev:     &item,
			}
			q = append(q, qi)
		}

	}

	return 0
}

func parse(lines []string) (grid, point) {
	g := make(grid)
	for y, line := range lines {
		for x, v := range strings.Split(line, "") {
			g[point{x, y}], _ = strconv.Atoi(v)
		}
	}
	return g, point{len(lines[0]) - 1, len(lines) - 1}
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
