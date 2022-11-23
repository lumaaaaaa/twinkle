package main

import (
	"github.com/go-vgo/robotgo"
	"time"
)

// dynamic config
func characterDetection() {
	for {
		if dynamicConfig {
			color := robotgo.GetPixelColor(131, 1400)
			switch color {
			case "e69294":
				if currentCharacter != "d.va" {
					currentCharacter = "d.va"
				}
			case "715848":
				if currentCharacter != "doomfist" {
					currentCharacter = "doomfist"
				}
			case "c1906f":
				if currentCharacter != "junker_queen" {
					currentCharacter = "junker_queen"
				}
			case "907757":
				if currentCharacter != "orisa" {
					currentCharacter = "orisa"
				}
			case "d6d4cd":
				if currentCharacter != "reinhardt" {
					currentCharacter = "reinhardt"
				}
			case "38302a":
				if currentCharacter != "roadhog" {
					currentCharacter = "roadhog"
				}
			case "91615a":
				if currentCharacter != "sigma" {
					currentCharacter = "sigma"
				}
			case "584653":
				if currentCharacter != "winston" {
					currentCharacter = "winston"
				}
			case "cdbaa8":
				if currentCharacter != "wrecking_ball" {
					currentCharacter = "wrecking_ball"
				}
			case "c4856c":
				if currentCharacter != "zarya" {
					currentCharacter = "zarya"
				}
			case "d19a7c":
				if currentCharacter != "ashe" {
					currentCharacter = "ashe"
				}
			case "d4d8d7":
				if currentCharacter != "bastion" {
					currentCharacter = "bastion"
				}
			case "875541":
				if currentCharacter != "cassidy" {
					currentCharacter = "cassidy"
				}
			case "5699c1":
				if currentCharacter != "echo" {
					currentCharacter = "echo"
				}
			case "918c82":
				if currentCharacter != "genji" {
					currentCharacter = "genji"
				}
			case "c89471":
				if currentCharacter != "hanzo" {
					currentCharacter = "hanzo"
				}
			case "5d352a":
				if currentCharacter != "junkrat" {
					currentCharacter = "junkrat"
				}
			case "a7755a":
				if currentCharacter != "pharah" {
					currentCharacter = "pharah"
				}
			case "282624":
				if currentCharacter != "reaper" {
					currentCharacter = "reaper"
				}
			case "6d4838":
				if currentCharacter != "sojourn" {
					currentCharacter = "sojourn"
				}
			case "906960":
				if currentCharacter != "soldier_76" {
					currentCharacter = "soldier_76"
				}
			case "924536":
				if currentCharacter != "sombra" {
					currentCharacter = "sombra"
				}
			case "6e352c":
				if currentCharacter != "symmetra" {
					currentCharacter = "symmetra"
				}
			case "ceb691":
				if currentCharacter != "torbjorn" {
					currentCharacter = "torbjorn"
				}
			case "d98637":
				if currentCharacter != "tracer" {
					currentCharacter = "tracer"
				}
			case "9b8bb2":
				if currentCharacter != "widowmaker" {
					currentCharacter = "widowmaker"
				}
			case "9c6b52":
				if currentCharacter != "ana" {
					currentCharacter = "ana"
				}
			case "703f31":
				if currentCharacter != "baptiste" {
					currentCharacter = "baptiste"
				}
			case "c47156":
				if currentCharacter != "brigitte" {
					currentCharacter = "brigitte"
				}
			case "cc8a75":
				if currentCharacter != "kiriko" {
					currentCharacter = "kiriko"
				}
			case "9f7051":
				if currentCharacter != "lucio" {
					currentCharacter = "lucio"
				}
			case "e5c1aa":
				if currentCharacter != "mercy" {
					currentCharacter = "mercy"
				}
			case "ab6756":
				if currentCharacter != "moira" {
					currentCharacter = "moira"
				}
			case "835f0f":
				if currentCharacter != "zenyatta" {
					currentCharacter = "zenyatta"
				}
			}

			useAimbot = cfg.Characters[characterToIndex[currentCharacter]].UseAimbot
			aimRepetitions = int32(cfg.Characters[characterToIndex[currentCharacter]].AimbotSettings.Strength)
			aimRadius = int32(cfg.Characters[characterToIndex[currentCharacter]].AimbotSettings.Radius)
			headOffset = int32(cfg.Characters[characterToIndex[currentCharacter]].AimbotSettings.YOffset)
			flickOffset = int32(cfg.Characters[characterToIndex[currentCharacter]].AimbotSettings.Flick)
			useTriggerbot = cfg.Characters[characterToIndex[currentCharacter]].UseTriggerbot

		} else {
			currentCharacter = "n/a"
		}
		time.Sleep(250 * time.Millisecond)
	}
}
