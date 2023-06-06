package funscript

import (
	"embed"
	"encoding/json"
	"io"

	"github.com/sopioman/yoi/feedback"
)

func NewFuncscipt(folder embed.FS, jsonfile string) (*Funscript, error) {
	jsonFile, err := folder.Open(jsonfile)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	//feedback.Blink(7, 200)
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		//feedback.Blink(4, 100)
		return nil, err
	}

	feedback.Blink(3, 200)
	var fs Funscript
	err = json.Unmarshal(byteValue, &fs)
	if err == nil {
		feedback.Blink(5, 100)
		if fs.Range == 0 {
			fs.Range = 100
		}
	}

	feedback.Blink(10, 200)
	return &fs, err
}

func (fs *Funscript) Process() {
	//var st Stroke

	// calculate difference between timestamps
	// to get stroke duration
	// for i := 1; i < len(fs.Actions); i++ {
	// 	if fs.Actions[i-1].Timestamp < 0 {
	// 		continue
	// 	}

	// 	// calculate stroke duration
	// 	st.Duration = fs.Actions[i].Timestamp - fs.Actions[i-1].Timestamp

	// 	if st.Duration > 0 {
	// 		st.Pre = fs.Actions[i-1].Position
	// 		st.Post = fs.Actions[i].Position

	// 		st.Range = fs.Range
	// 		st.Inverted = fs.Inverted
	// 		st.FullStroke = fs.FullStroke

	// 		st.Execute()
	// 	}
	// }

}

func (s *Stroke) Execute() {
	feedback.Blink(1, s.Duration)
}
