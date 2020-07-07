package protein

import "errors"

// ErrStop error
var ErrStop = errors.New("err stop")
// ErrInvalidBase error
var ErrInvalidBase = errors.New("invalid base")


// FromCodon func
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG": {
		return "Methionine", nil
	}
	case "UUU", "UUC": {
		return "Phenylalanine", nil
	}	
	case "UUA", "UUG": {
		return "Leucine", nil
	}
	case "UCU", "UCC", "UCA", "UCG": {
		return "Serine", nil
	}
	case "UAU", "UAC": {
		return "Tyrosine",nil
	}
	case "UGU", "UGC": {
		return "Cysteine", nil
	}
	case "UGG": {
		return "Tryptophan", nil
	}
	case "UAA", "UAG", "UGA": {
		return "", ErrStop	
	}
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA fun
func FromRNA(protein string) ([]string, error) {
	var result []string
	for len(protein) >= 3 {
		x, err := FromCodon(protein[:3])
		if err == ErrStop {
			return result, nil
		} else if err != nil {
			return result, err
		}
		result = append(result, x)
		protein = protein[3:]
	}	
	return result, nil
}