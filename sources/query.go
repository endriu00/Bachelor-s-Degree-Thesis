type Query struct {
	Attributes   Attributes       `json:"attributes"`
	Seismometers []string         `json:"seismometers"`
	Location     Location         `json:"location"`
	From         time.Time        `json:"from"`
	To           time.Time        `json:"to"`
	Tg           string           `json:"tg"`
	Mail         string           `json:"mail"`
	Timestamp    time.Time        `json:"timestamp"`
	Output       OutputExtensions `json:"output"`
	ID           string           `json:"id"`
}