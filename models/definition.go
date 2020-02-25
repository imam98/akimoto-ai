package models

const (
	OdBasePath string = "https://od-api.oxforddictionaries.com/api/v2/"
	OdFields   string = "pronunciations,definitions,examples"
)

type sense struct {
	Definitions []string `json:"definitions"`
	Examples    []struct {
		Text string `json:"text"`
	} `json:"examples"`
}

type entry struct {
	Senses []sense `json:"senses"`
}

type pronunciation struct {
	Dialects         []string `json:"dialects"`
	PhoneticNotation string   `json:"phoneticNotation"`
	PhoneticSpelling string   `json:"phoneticSpelling"`
}

type lexicalEntry struct {
	Text            string          `json:"text"`
	Entries         []entry         `json:"entries"`
	Pronunciations  []pronunciation `json:"pronunciations"`
	LexicalCategory struct {
		Text string `json:"text"`
	} `json:"lexicalCategory"`
}

type headwordEntry struct {
	LexicalEntries []lexicalEntry `json:"lexicalEntries"`
}

type metadata struct {
	Operation string `json:"operation"`
	Provider  string `json:"provider"`
	Schema    string `json:"schema"`
}

type DictionaryEntry struct {
	Metadata metadata        `json:"metadata"`
	Results  []headwordEntry `json:"results"`
}

const DefTmpl string = `📃 {{.Text}}, _{{.LexicalCategory.Text}}_
{{$sense := (index (index .Entries 0).Senses 0)}}
❓ Definition: {{index $sense.Definitions 0}}`
