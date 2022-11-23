package main

type Config struct {
	Characters []struct {
		Name           string `json:"name"`
		UseAimbot      bool   `json:"useAimbot"`
		AimbotSettings struct {
			Strength int `json:"strength"`
			Radius   int `json:"radius"`
			YOffset  int `json:"y-offset"`
			Flick    int `json:"flick"`
		} `json:"aimbotSettings"`
		UseTriggerbot bool `json:"useTriggerbot"`
	} `json:"characters"`
	Version int `json:"version"`
}
