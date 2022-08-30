package model

type Mutant struct {
	DNA      string `firestore:"dna,omitempty"`
	IsMutant bool   `firestore:"isMutant,omitempty"`
}
