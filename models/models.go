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
