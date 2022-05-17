package blockToHtml

type Children struct {
	Key   string   `json:"_key"`
	Type  string   `json:"_type"`
	Marks []string `json:"marks"`
	Text  string   `json:"text"`
}

type MarkDef struct {
	Key  string `json:"_key"`
	Type string `json:"_type"`
	Href string `json:"href"`
}

type Block struct {
	Key      string     `json:"_key"`
	Type     string     `json:"_type"`
	Children []Children `json:"children"`
	MarkDefs []MarkDef  `json:"markDefs"`
	Style    string     `json:"style"`
	Level    int        `json:"level"`
	ListItem string     `json:"ListItem"`
}
