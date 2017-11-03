package asciicast

type CastV1 struct {
	Version uint `json:"version"` //always 1
	Width uint `json:"width"`   //#: 80,
	Height uint `json:"height"`   //#: 24,
	Duration float64 `json:"duration"`   //#: 1.515658,
	Command string `json:"command"`   //#: "/bin/zsh",
	Title string `json:"title"`   //#: "",
	Env *map[string]string  `json:"env"`
	Stdout [][]interface{}  `json:"stdout"`// [[ts in float, data],[ts in float,data]
}
// NewCastV1() emits v1 formatted data which is just a big fat json.
// after pushing all frames via PushFrame() just json.Marshal(cast) to get the output stream
//
//	        asc := asciicast.NewCastV1(asciicast.CastMetadata{
//              Width: 80,
//              Height: 24,
//          })
//	        asc.PushFrame(0.2, "test")
//	        b, err := json.Marshal(asc)
//	        fmt.Printf("json:\n%s\n---\n%s\n)",string(b),err

func NewCastV1(meta CastMetadata) (CastV1, error){
	var c CastV1
	c.Version = 1
	c.Stdout = make([][]interface{},0)
	if meta.Env != nil {
		c.Env = &meta.Env
	}
	c.Duration = meta.Duration.Seconds()
	c.Width = meta.Width
	c.Height = meta.Height
	c.Title = meta.Title
	c.Command = meta.Command
	return c, nil
}

func (a *CastV1)PushFrame(ts float64, data []byte) {
	frame := make([]interface{},2)
	frame[0] = ts
	frame[1] = string(data)
	a.Stdout = append(a.Stdout,frame)
}
