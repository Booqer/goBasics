package main

import (
	"fmt"
	"sort"
)

// Player struct defines the player properties.
type Player struct {
	name       string
	territory  int
	aggression int
}

// By is the type of a "less" function that defines the ordering of its Player arguments.
type By func(p1, p2 *Player) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(players []Player) {
	playerSlice := &playerSorter{
		players: players,
		by:      by,
	}
	sort.Sort(playerSlice)
}

type playerSorter struct {
	players []Player
	by      func(p1, p2 *Player) bool
}

func (s *playerSorter) Len() int {
	return len(s.players)
}

func (s *playerSorter) Swap(i, j int) {
	s.players[i], s.players[j] = s.players[j], s.players[i]
}

func (s *playerSorter) Less(i, j int) bool {
	return s.by(&s.players[i], &s.players[j])
}

var players = []Player{
	{"ToadLicker", 18, 55},
	{"Bucky", 3, 0},
	{"LordMonkeyButtox", 77, 83},
	{"HashtagSag", 8, 255},
	{"Jelatinous", 42, 5},
}

func main() {
	name := func(p1, p2 *Player) bool {
		return p1.name > p2.name
	}
	territory := func(p1, p2 *Player) bool {
		return p1.territory > p2.territory
	}
	aggression := func(p1, p2 *Player) bool {
		return p1.aggression > p2.aggression
	}

	By(name).Sort(players)
	fmt.Println("By name: ", players)

	By(territory).Sort(players)
	fmt.Println("By territory: ", players)

	By(aggression).Sort(players)
	fmt.Println("By aggression: ", players)
}
