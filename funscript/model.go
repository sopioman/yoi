package funscript

type Funscript struct {
	Version         string `json:"version"`
	Inverted        bool   `json:"inverted"`
	FullStroke      bool   `json:"fullstroke"`
	Range           uint8  `json:"range"`
	PositionCurrent int8
	PositionNew     int8
	Duration        int32
}

type FunscriptActions struct {
	Position  int8  `json:"pos"`
	Timestamp int32 `json:"at"`
}

type Stroke struct {
	Range      uint8
	Pre        int8
	Post       int8
	Duration   int32
	FullStroke bool
	Inverted   bool
}
