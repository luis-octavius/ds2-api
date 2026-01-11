package main

type Pyromancy struct {
	Name        string
	Uses        string
	Attunement  string
	Description string
	Acquired    string
	Cost        string
	Type        string
}

type Hex struct {
	Name         string
	Uses         string
	Attunement   string
	Intelligence string
	Faith        string
	Description  string
	Acquired     string
	Cost         string
	Type         string
}

type Miracle struct {
	Name        string
	Slots       string
	Faith       string
	Attunement  string
	Description string
	Acquired    string
	Type        string
}

type Sorcery struct {
	Name         string
	Uses         string
	Spellslot    string
	Intelligence string
	Description  string
	Acquired     string
	Cost         string
	Type         string
}
