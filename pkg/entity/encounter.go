package entity

type LocationArea struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource      `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod         NamedAPIResource         `json:"encounter_method"`
	EncounterVersionDetails []EncounterVersionDetail `json:"version_details"`
}

type EncounterVersionDetail struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type EncounterDetail struct {
	Chance          int              `json:"chance"`
	ConditionValues []any            `json:"condition_values"`
	MaxLevel        int              `json:"max_level"`
	Method          NamedAPIResource `json:"method"`
	MinLevel        int              `json:"min_level"`
}

type VersionEncounterDetail struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int               `json:"max_chance"`
	Version          NamedAPIResource  `json:"version"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource         `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}
