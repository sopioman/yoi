package funscript

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func NewFuncscipt(jsonfile string) (*Funscript, error) {
	jsonFile, err := os.Open(jsonfile)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var fs Funscript
	err = json.Unmarshal(byteValue, &fs)
	if err == nil {
		if fs.Range == 0 {
			fs.Range = 100
		}
	}

	return &fs, err
}

func (fs *Funscript) Process() {
	var st Stroke

	// calculate difference between timestamps
	// to get stroke duration
	for i := 1; i < len(fs.Actions); i++ {
		if fs.Actions[i-1].Timestamp < 0 {
			continue
		}

		// calculate stroke duration
		st.Duration = fs.Actions[i].Timestamp - fs.Actions[i-1].Timestamp

		if st.Duration > 0 {
			st.Pre = fs.Actions[i-1].Position
			st.Post = fs.Actions[i].Position

			st.Range = fs.Range
			st.Inverted = fs.Inverted
			st.FullStroke = fs.FullStroke

			st.Execute()
		}
	}

}

func (s *Stroke) Execute() {
	fmt.Println(s)
}
