package main

import (
	"github.com/go-vgo/robotgo"
	"math"
	"syscall"
	"time"
)

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")
	procMouse = moduser32.NewProc("mouse_event")
)

// aimbot
func listen() {
	for {
		if useAimbot {
			x, y := robotgo.FindColorCS(0xff00ff, 1200-int(aimRadius), 800-int(aimRadius), int(aimRadius)*2, int(aimRadius)*2)
			if x != -1 && y != -1 && y > 225 {
				x += 1200 - int(aimRadius) + 40
				y += 800 - int(aimRadius) + int(headOffset)
				var intype uint32 = MouseEventMove
				//procMouse.Call(uintptr(intype), uintptr(1), uintptr(1), 0, 0)
				xOff := 0
				yOff := 0
				if math.Abs(float64(x-1200)) < 2 {
					xOff = 0
				} else if math.Abs(float64(x-1200)) < 5 {
					if x > 1200 {
						xOff = 1
					} else {
						xOff = -1
					}
				} else if math.Abs(float64(x-1200)) < 25 {
					if x > 1200 {
						xOff = 2
					} else {
						xOff = -2
					}
				} else if math.Abs(float64(x-1200)) < 75 {
					if x > 1200 {
						xOff = 3
					} else {
						xOff = -3
					}
				} else if math.Abs(float64(x-1200)) < 100 {
					if x > 1200 {
						xOff = 5
					} else {
						xOff = -5
					}
				} else {
					if x > 1200 {
						xOff = int(flickOffset)
					} else {
						xOff = int(-flickOffset)
					}
				}

				if math.Abs(float64(y-800)) < 2 {
					yOff = 0
				} else if math.Abs(float64(y-800)) < 5 {
					if y > 800 {
						yOff = 1
					} else {
						yOff = -1
					}
				} else if math.Abs(float64(y-800)) < 25 {
					if y > 800 {
						yOff = 2
					} else {
						yOff = -2
					}
				} else if math.Abs(float64(y-800)) < 75 {
					if y > 800 {
						yOff = 3
					} else {
						yOff = -3
					}
				} else if math.Abs(float64(y-800)) < 100 {
					if y > 800 {
						yOff = 5
					} else {
						yOff = -5
					}
				} else {
					if y > 800 {
						yOff = int(flickOffset)
					} else {
						yOff = int(-flickOffset)
					}
				}

				for i := 0; i < int(aimRepetitions); i++ {
					intype |= MouseEventAbsolute
					_, _, err := procMouse.Call(uintptr(intype), uintptr(1200+xOff), uintptr(800+yOff), 0, 0)
					if err != nil {
						// ignore!
					}
				}

				if float64(xOff) < math.Abs(float64(3-int(triggerbotConfidence))) && float64(yOff) < math.Abs(float64(3-int(triggerbotConfidence))) && useTriggerbot {
					if currentCharacter == "widowmaker" {
						color := robotgo.GetPixelColor(1200, 911)
						var hex = Hex(color)
						rgb, _ := Hex2RGB(hex)
						if rgb.Red > 225 && rgb.Green > 225 && rgb.Blue > 225 {
							robotgo.Click()
						}
					} else {
						robotgo.Click()
					}
				}
			}
			time.Sleep(10 * time.Nanosecond)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

// triggerbot
func monitor(x int, y int) {
	color := robotgo.GetPixelColor(1200, 800)
	for {
		if useTriggerbot {
			color = robotgo.GetPixelColor(x, y)
			var hex = Hex(color)
			rgb, _ := Hex2RGB(hex)
			if (rgb.Red > 190 && rgb.Green < 48 && rgb.Blue > 190) && (rgb.Red != 255 && rgb.Blue != 255) && (rgb.Red != rgb.Blue) && (rgb.Red < 250 && rgb.Blue < 250 && rgb.Green > 5) {
				robotgo.Click()
				time.Sleep(1 * time.Second)
			}
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
