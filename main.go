package main

import (
	"encoding/json"
	"fmt"
	g "github.com/AllenDang/giu"
	"image/color"
	"io/ioutil"
	"os"
	"strings"
)

const (
	MouseEventAbsolute = 0x8000
	MouseEventMove     = 0x0001
)

var (
	cfg     Config
	configs []string

	useAimbot            bool
	useTriggerbot        bool
	widowmakerFullCharge bool

	dynamicConfig bool
	editConfig    bool

	aimRepetitions int32
	aimRadius      int32
	headOffset     int32
	flickOffset    int32

	triggerbotConfidence int32

	selectedConfig int32

	currentCharacter string
	currentIndex     int

	characterToIndex map[string]int
)

// exit button helper
func exit() {
	os.Exit(0)
}

// gui
func loop() {
	g.SingleWindow().Flags(g.WindowFlagsAlwaysVerticalScrollbar).Flags(g.WindowFlagsNoMove).Flags(g.WindowFlagsNoTitleBar).Layout(
		g.Label("twinkle - v2.0"),
		g.Separator(),
		g.Checkbox("aimbot", &useAimbot).OnChange(editCfg),
		g.Child().Size(g.Auto, 178).Layout(
			g.Label("aimbot settings"),
			g.Separator(),
			g.SliderInt(&aimRepetitions, 1, 100).Label("strength").OnChange(editCfg),
			g.SliderInt(&aimRadius, 50, 400).Label("radius").OnChange(editCfg),
			g.SliderInt(&headOffset, 35, 100).Label("y-offset").OnChange(editCfg),
			g.SliderInt(&flickOffset, 1, 50).Label("flick").OnChange(editCfg),
		),
		g.Checkbox("triggerbot", &useTriggerbot).OnChange(editCfg),
		g.Child().Size(g.Auto, 80).Layout(
			g.Label("triggerbot settings"),
			g.Separator(),
			g.SliderInt(&triggerbotConfidence, 0, 3).Label("confidence").OnChange(editCfg),
		),
		g.Separator(),
		g.Row(
			g.Checkbox("dynamic config |", &dynamicConfig),
			g.Combo("", configs[selectedConfig], configs, &selectedConfig).Size(220).OnChange(loadCfg),
		),
		g.Button("save config").OnClick(saveCfg).Size(g.Auto, 25),
		g.Separator(),
		g.Row(
			g.Button("exit").OnClick(exit).Size(175, 25),
			g.Label("| character: "+strings.ReplaceAll(currentCharacter, "_", " ")),
		),
	)
}

func loadCfg() {
	cfgFile, _ := os.Open("./configs/" + configs[selectedConfig])

	byteValue, _ := ioutil.ReadAll(cfgFile)

	err2 := json.Unmarshal(byteValue, &cfg)
	if err2 != nil {
		return
	}

	characterToIndex = make(map[string]int)

	for index, element := range cfg.Characters {
		characterToIndex[element.Name] = index
	}
}

func editCfg() {
	cfg.Characters[characterToIndex[currentCharacter]].UseAimbot = useAimbot
	cfg.Characters[characterToIndex[currentCharacter]].AimbotSettings.Strength = int(aimRepetitions)
	cfg.Characters[characterToIndex[currentCharacter]].AimbotSettings.Radius = int(aimRadius)
	cfg.Characters[characterToIndex[currentCharacter]].AimbotSettings.YOffset = int(headOffset)
	cfg.Characters[characterToIndex[currentCharacter]].AimbotSettings.Flick = int(flickOffset)
	cfg.Characters[characterToIndex[currentCharacter]].UseTriggerbot = useTriggerbot
}

func saveCfg() {
	cfgFile, _ := json.MarshalIndent(cfg, "", "\t")

	_ = ioutil.WriteFile("./configs/"+configs[selectedConfig], cfgFile, 0644)
}

func main() {
	// load configs
	f, err := os.Open("./configs")
	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		configs = append(configs, v.Name())
	}

	// defaults
	useAimbot = false
	useTriggerbot = false
	dynamicConfig = true
	selectedConfig = 0

	// set default values
	aimRepetitions = 25
	aimRadius = 350
	headOffset = 44
	flickOffset = 20

	triggerbotConfidence = 3

	currentCharacter = "n/a"

	loadCfg()

	go listen()
	go monitor(1200, 800)
	go characterDetection()

	g.SetDefaultFont("Consola", 12)
	wnd := g.NewMasterWindow("✨twinkle", 300, 327, g.MasterWindowFlagsFrameless)
	wnd.SetBgColor(color.Transparent)
	wnd.SetPos(0, 0)
	wnd.Run(loop)
}
