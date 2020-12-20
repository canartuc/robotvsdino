package main

// LOWERBOUND and UPPERBOUND of initialized board. This is programming bounds so it always begins with 0
// so ends with logical upper - 1
const (
	LOWERBOUND = 0
	UPPERBOUND = 49
)

// REQUESTGROUP: In future, if there will be new version and deprecation notices, this parameter can be used
// for canary releases. http(s)://IP_OR_DOMAIN:PORT/REQUESTGROUP/WHATEVERTHEREST
const REQUESTGROUP = "/api/v1"

// BOARDGROUP, ROBOTGROUP and DINOSAURGROUP defines URL request grouping after REQUESTGROUP
// http(s)://IP_OR_DOMAIN:PORT/REQUESTGROUP/[BOARDGROUP||ROBOTGROUP||DINOSAURGROUP]/ACTION/WHATEVERTHEREST
const (
	BOARDGROUP    = "/board"
	ROBOTGROUP    = "/robot"
	DINOSAURGROUP = "/dinosaur"
)

// Board is global scope variable and the main unit to play on
// We define here as we will share this structure with dinosaur and robot because there is no long term storage
// in this project but temporary.
var Board = make([][]string, UPPERBOUND+1)

// In programming, row and column define as integer value but in human world, we are using letter for Column
// for games generally - for example: Chess board
var Column = map[string]int{
	"A":  0,
	"B":  1,
	"C":  2,
	"D":  3,
	"E":  4,
	"F":  5,
	"G":  6,
	"H":  7,
	"I":  8,
	"J":  9,
	"K":  10,
	"L":  11,
	"M":  12,
	"N":  13,
	"O":  14,
	"P":  15,
	"Q":  16,
	"R":  17,
	"S":  18,
	"T":  19,
	"U":  20,
	"V":  21,
	"W":  22,
	"X":  23,
	"Y":  24,
	"Z":  25,
	"AA": 26,
	"AB": 27,
	"AC": 28,
	"AD": 29,
	"AE": 30,
	"AF": 31,
	"AG": 32,
	"AH": 33,
	"AI": 34,
	"AJ": 35,
	"AK": 36,
	"AL": 37,
	"AM": 38,
	"AN": 39,
	"AO": 40,
	"AP": 41,
	"AQ": 42,
	"AR": 43,
	"AS": 44,
	"AT": 45,
	"AU": 46,
	"AV": 47,
	"AW": 48,
	"AX": 49,
}

// Face will be used for last face of robots as well as moving direction. The remove ambiguity, I replaced
// directions with cardinal directions/points
var Face = map[string]string{
	"RIGHT":  "east",
	"LEFT":   "west",
	"TOP":    "north",
	"BOTTOM": "south",
}
