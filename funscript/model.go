package funscript

type Funscript struct {
	Version    string             `json:"version"`
	Inverted   bool               `json:"inverted"`
	FullStroke bool               `json:"fullstroke"`
	Range      uint8              `json:"range"`
	Actions    []FunscriptActions `json:"actions"`
}

type FunscriptActions struct {
	Position  int8  `json:"pos"`
	Timestamp int32 `json:"at"`
}

//tinyjson:skip
type Stroke struct {
	Range      uint8
	Pre        int8
	Post       int8
	Duration   int32
	FullStroke bool
	Inverted   bool
}
