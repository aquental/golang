package protein

import (
	"errors"
)
var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
    "UUC": "Phenylalanine",
	"UUA": "Leucine",
    "UUG": "Leucine",
	"UCU": "Serine",
    "UCC": "Serine",
	"UCA": "Serine",
    "UCG": "Serine",
	"UAU": "Tyrosine",
    "UAC": "Tyrosine",
	"UGU": "Cysteine",
    "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "",
    "UAG": "",
    "UGA": "",
}
var (
	ErrStop        = errors.New("reached stop Codone")
	ErrInvalidBase = errors.New("unexpected protein sequence")
)

func FromRNA(rna string) ([]string, error) {
    l := len(rna)
	result := make([]string, 0, l/3)
	for i := 0; i < l; i += 3 {
		poly, err := FromCodon(rna[i : i+3])
		if err == ErrStop {
			return result, nil
		}
		if err == ErrInvalidBase {
			return result, err
		}
		result = append(result, poly)
	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	if poly, ok := codons[codon]; ok {
		if poly == "" {
			return "", ErrStop
		}
		return poly, nil
	}
	return "", ErrInvalidBase
}

