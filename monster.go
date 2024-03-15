package main

import (
	"fmt"
	"math/rand"
)

/*************************************************************************
 * MonsterLib
 *
 */
type MonsterTemplate struct {
	Symbol rune
	Carry  int
	XP     int
	Level  int
	AC     int
	Dmg    string
	Name   string
}

// Index is used as difficulty of the monsters
//
//	min = depth - 6
//	max = depth + 3
//
// (apparently called "vorpalness" in original Rogue source code)
// https://datadrivengamer.blogspot.com/2019/05/identifying-mechanics-of-rogue.html
var MonsterLib = []MonsterTemplate{
	{'K', 0, 2, 1, 7, "1d4", "kobold"},
	{'J', 0, 2, 1, 7, "1d2", "jackal"},
	{'B', 0, 1, 1, 3, "1d2", "bat"},
	{'S', 0, 3, 1, 5, "1d3", "snake"},
	{'H', 0, 3, 1, 5, "1d8", "hobgoblin"},
	{'E', 0, 5, 1, 9, "0d0", "floating eye"},
	{'A', 0, 10, 2, 3, "1d6", "giant ant"},
	{'O', 15, 5, 1, 6, "1d7", "orc"},
	{'Z', 0, 7, 2, 8, "1d8", "zombie"},
	{'G', 10, 8, 1, 5, "1d6", "gnome"},
	{'L', 0, 10, 3, 8, "1d1", "leprechaun"},
	{'C', 15, 15, 4, 4, "1d6/1d6", "centaur"},
	{'R', 0, 25, 5, 2, "0d0/0d0", "rust monster"},
	{'Q', 30, 35, 3, 2, "1d2/1d2/1d4", "quasit"},
	{'N', 100, 40, 3, 9, "0d0", "nymph"},
	{'Y', 30, 50, 4, 6, "1d6/1d6", "yeti"},
	{'T', 50, 55, 6, 4, "1d8/1d8/2d6", "troll"},
	{'W', 0, 55, 5, 4, "1d6", "wraith"},
	{'F', 0, 85, 8, 3, "0d0", "violet fungi"},
	{'I', 0, 120, 8, 3, "4d4", "invisible stalker"},
	{'X', 0, 120, 7, -2, "1d3/1d3/1d3/4d6", "xorn"},
	{'U', 40, 130, 8, 2, "3d4/3d4/2d5", "umber hulk"},
	{'M', 30, 140, 7, 7, "3d4", "mimic"},
	{'V', 30, 380, 8, 1, "1d10", "vampire"},
	{'D', 100, 9000, 10, -1, "1d8/1d8/3d10", "dragon"},
	{'P', 70, 7000, 15, 6, "2d12/2d4", "purple worm"},
}

// Uses public variable MonsterLib
func randomMonster(depth int) *Monster {
	min := depth - 6
	max := depth + 3
	if min < 0 {
		min = 0
	}
	if max > len(MonsterLib) {
		max = len(MonsterLib)
	}

	idx := len(MonsterLib) - 1 // Default to most difficult monster
	if min < len(MonsterLib) { // Ensure we don't go out of bounds
		idx = rand.Intn(max-min) + min
	}
	debug.Add("monster: len=%d, min=%d, max=%d, idx=%d", len(MonsterLib), min, max, idx)
	return newMonster(idx)
}

/*************************************************************************
 * MonsterList
 *
 */

type MonsterList []*Monster

func (ml *MonsterList) Add(m *Monster, x, y int) {
	m.X, m.Y = x, y
	*ml = append(*ml, m)
}

func (ml *MonsterList) Remove(idx int) {
	*ml = append((*ml)[:idx], (*ml)[idx+1:]...)
}

func (ml *MonsterList) Clear() {
	*ml = nil
}

func (ml MonsterList) MonsterAt(x, y int) *Monster {
	for _, m := range ml {
		if m.X == x && m.Y == y {
			return m
		}
	}
	return nil
}

/*************************************************************************
 * Monster
 * implements Entity interface
 */

type Monster struct {
	X, Y   int
	Symbol rune
	Name   string
	HP     int
}

func newMonster(id int) *Monster {
	mt := MonsterLib[id]
	m := &Monster{
		Name:   mt.Name,
		HP:     1,
		Symbol: mt.Symbol,
	}
	return m
}

func CreateMonster(n string, sym rune, hp int) *Monster {
	return &Monster{
		Symbol: sym,
		Name:   n,
		HP:     hp,
	}
}

func (m *Monster) DebugString() string {
	return fmt.Sprintf("%s x=%d y=%d hp=%d", m.Name, m.X, m.Y, m.HP)
}

// Implement the Entity interface

func (m *Monster) SetPos(newX, newY int) {
	m.X = newX
	m.Y = newY
}

func (m *Monster) Pos() (int, int) {
	return m.X, m.Y
}

func (m *Monster) Rune() rune {
	return m.Symbol
}
