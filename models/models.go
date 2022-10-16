package models

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Substrings struct {
	Key        string   `json:"key"`
	Length     int      `json:"length"`
	Substrings []string `json:"substrings"`
}

type AllignerSeq struct {
	Primer    string   `json:"primer"`
	Count     int      `json:"count"`
	Sequences []string `json:"data"`
}

type Primers struct {
	Key    string `json:"key"`
	Primer string `json:"primer"`
}

type KeyPrimers struct {
	Key     string    `json:"key"`
	Primers []Primers `json:"primers"`
}

type KeySequence struct {
	Key       string `json:"key"`
	Sequences []Data `json:"sequences"`
}

type PreFinal struct {
	Type     string  `json:"type"`
	Primer   Primers `json:"primer"`
	Sequence Data    `json:"sequence"`
}

type A2A struct {
	PrimerGroup   string `json:"primer_group"`
	PrimerKey     string `json:"primer_key"`
	Primer        string `json:"primer"`
	SequenceGroup string `json:"sequence_group"`
	SequenceKey   string `json:"sequence_key"`
	Sequence      string `json:"sequence"`
	Ratio         int    `json:"ratio"`
}

type Crusher struct {
	Length  int
	Massive []string
}

type SBS struct {
	Sequence string
	Primers  []PandP
}

type PandP struct {
	Primer   string
	Position int
}
