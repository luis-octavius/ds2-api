package main

import (
	"fmt"

	"github.com/luis-octavius/ds2-scraper/internal/database"
)

// database saving
func (s *state) SaveMiracles(miracles []Miracle) {
	fmt.Println("Starting the scraping of miracles...")

	for _, miracle := range miracles {
		createdMiracle, _ := s.db.CreateMiracle(ctx, database.CreateMiracleParams{
			Name:        miracle.Name,
			Uses:        miracle.Uses,
			Faith:       strToInt(miracle.Faith),
			Attunement:  strToInt(miracle.Attunement),
			Description: miracle.Description,
			Acquired:    miracle.Acquired,
			MiracleType: miracle.Type,
		})

		fmt.Printf("\nName: %v\nUses: %v\nFaith: %v\nAttunement: %v\nDescription: %v\nAcquired: %v\nType: %v\n", createdMiracle.Name, createdMiracle.Uses, createdMiracle.Faith, createdMiracle.Attunement, createdMiracle.Description, createdMiracle.Acquired, createdMiracle.MiracleType)
	}
}

func (s *state) SavePyromancies(pyromancies []Pyromancy) {
	fmt.Println("Starting the scraping of pyromancies...")

	for _, pyromancy := range pyromancies {
		createdPyromancy, _ := s.db.CreatePyromancy(ctx, database.CreatePyromancyParams{
			Name:          pyromancy.Name,
			Uses:          pyromancy.Uses,
			Attunement:    strToInt(pyromancy.Attunement),
			Description:   pyromancy.Description,
			Acquired:      pyromancy.Acquired,
			Cost:          pyromancy.Cost,
			PyromancyType: pyromancy.Type,
		})

		fmt.Printf("\nName: %v\nUses: %v\nAttunement: %v\nDescription: %v\n Acquired: %v\nType: %v\n", createdPyromancy.Name, createdPyromancy.Uses, createdPyromancy.Attunement, createdPyromancy.Description, createdPyromancy.Acquired, createdPyromancy.PyromancyType)

	}

}

func (s *state) SaveHexes(hexes []Hex) {
	fmt.Println("Starting the scraping of hexes\n")

	for _, hex := range hexes {
		createdHex, _ := s.db.CreateHex(ctx, database.CreateHexParams{
			Name:         hex.Name,
			Uses:         hex.Uses,
			Attunement:   strToInt(hex.Attunement),
			Intelligence: strToInt(hex.Intelligence),
			Faith:        strToInt(hex.Faith),
			Description:  hex.Description,
			Acquired:     hex.Acquired,
			Cost:         hex.Cost,
			HexType:      hex.Type,
		})

		fmt.Printf("Name: %v\nUses: %v\nAttunement: %v\nIntelligence: %v\nFaith: %v\nDescription: %v\nAcquired: %v\nCost: %v\nHexType: %v\n", createdHex.Name, createdHex.Uses, createdHex.Attunement, createdHex.Intelligence, createdHex.Faith, createdHex.Description, createdHex.Acquired, createdHex.Cost, createdHex.HexType)
	}
}

func (s *state) SaveSorceries(sorceries []Sorcery) {
	fmt.Println("Starting the scraping of the sorceries...\n")

	for _, sorcery := range sorceries {
		createdSorcery, _ := s.db.CreateSorcery(ctx, database.CreateSorceryParams{
			Name:         sorcery.Name,
			Uses:         sorcery.Uses,
			Spellslot:    strToInt(sorcery.Spellslot),
			Intelligence: strToInt(sorcery.Intelligence),
			Description:  sorcery.Description,
			Acquired:     sorcery.Acquired,
			Cost:         sorcery.Cost,
			SorceryType:  sorcery.Type,
		})

		fmt.Printf("Name: %v\nUses: %v\nSpellslot: %v\nIntelligence: %v\nDescription: %v\nAcquired: %v\nCost: %v\nSorceryType: %v\n", createdSorcery.Name, createdSorcery.Uses, createdSorcery.Spellslot, createdSorcery.Intelligence, createdSorcery.Description, createdSorcery.Acquired, createdSorcery.Cost, createdSorcery.SorceryType)
	}
}
