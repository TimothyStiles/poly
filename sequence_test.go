package poly

import (
	"testing"
)

func TestGetSequenceMethods(t *testing.T) {

	gbk := ReadGbk("data/t4_intron.gb")

	// Check to see if GetSequence method works on Annotated struct
	if gbk.GetSequence() != gbk.Sequence {
		t.Errorf(" Sequence GetSequence method has failed'. Got this:\n%s instead of \n%s", gbk.GetSequence(), gbk.Sequence)
	}

	// Check to see if GetSequence method works on Features struct
	feature := gbk.Features[1].GetSequence()
	seq := "atgagattacaacgccagagcatcaaagattcagaagttagaggtaaatggtattttaatatcatcggtaaagattctgaacttgttgaaaaagctgaacatcttttacgtgatatgggatgggaagatgaatgcgatggatgtcctctttatgaagacggagaaagcgcaggattttggatttaccattctgacgtcgagcagtttaaagctgattggaaaattgtgaaaaagtctgtttga"
	if feature != seq {
		t.Errorf("Feature GetSequence method has failed. Got this:\n%s instead of \n%s", feature, seq)
	}

	// Check to see if GetSequence method works on Sequence struct
	if gbk.GetSequence() != gbk.Sequence {
		t.Errorf("Sequence GetSequence method has failed.. Got this:\n%s instead of \n%s", gbk.GetSequence(), gbk.Sequence)
	}

}

func TestLocationParser(t *testing.T) {
	gbk := ReadGbk("data/t4_intron.gb")

	// Read 1..243
	feature := gbk.Features[1].GetSequence()
	seq := "atgagattacaacgccagagcatcaaagattcagaagttagaggtaaatggtattttaatatcatcggtaaagattctgaacttgttgaaaaagctgaacatcttttacgtgatatgggatgggaagatgaatgcgatggatgtcctctttatgaagacggagaaagcgcaggattttggatttaccattctgacgtcgagcagtttaaagctgattggaaaattgtgaaaaagtctgtttga"
	if feature != seq {
		t.Errorf("Feature sequence parser has changed on test '1..243'. Got this:\n%s instead of \n%s", feature, seq)
	}

	// Read join(893..1441,2459..2770)
	featureJoin := gbk.Features[6].GetSequence()
	seqJoin := "atgaaacaatccaagatttaattaaagacatttttgaaaatggttatgaaaccgatgatcgtacaggcacaggaacaattgctctgttcggatctaaattacgctgggatttaactaaaggttttcctgcggtaacaactaagaagctcgcctggaaagcttgcattgctgagctaatatggtttttatcaggaagcacaaatgtcaatgatttacgattaattcaacacgattcgttaatccaaggcaaaacagtctgggatgaaaattacgaaaatcaagcaaaagatttaggataccatagcggtgaacttggtccaatttatggaaaacagtggcgtgattttggtggtgtagaccaaattatagaagttattgatcgtattaaaaaactgccaaatgataggcgtcaaattgtttctgcatggaatccagctgaacttaaatatatggcattaccgccttgtcatatgttctatcagtttaatgtgcgtaatggctatttggatttgcagtggtatcaacgctcagtagatgttttcttgggtctaccgtttaatattgcgtcatatgctacgttagttcatattgtagctaagatgtgtaatcttattccaggggatttgatattttctggtggtaatactcatatctatatgaatcacgtagaacaatgtaaagaaattttgaggcgtgaacctaaagagctttgtgagctggtaataagtggtctaccttataaattccgatatctttctactaaagaacaattaaaatatgttcttaaacttaggcctaaagatttcgttcttaacaactatgtatcacaccctcctattaaaggaaagatggcggtgtaa"
	if featureJoin != seqJoin {
		t.Errorf("Feature sequence parser has changed on test 'join(893..1441,2459..2770)'. Got this:\n%s instead of \n%s", featureJoin, seqJoin)
	}

	// Read complement(2791..3054)
	featureComplement := gbk.Features[10].GetSequence()
	seqComplement := "ttattcactacccggcatagacggcccacgctggaataattcgtcatattgtttttccgttaaaacagtaatatcgtagtaacagtcagaagaagttttaactgtggaaattttattatcaaaatactcacgagtcattttatgagtatagtattttttaccataaatggtaataggctgttctggtcctggaacttctaactcgcttgggttaggaagtgtaaaaagaactacaccagaagtatctttaaatcgtaaaatcat"
	if featureComplement != seqComplement {
		t.Errorf("Feature sequence parser has changed on test 'complement(2791..3054)'. Got this:\n%s instead of \n%s", featureComplement, seqComplement)
	}

	// Read join(complement(315..330),complement(339..896))
	// Note: it is known that some software, like Snapgene, assumes that since both strands are in the reverse direction
	// that the first sequence should be appended to the reverse sequence, instead of the second sequence
	// getting appended to the first. Biopython appends the second sequence to the first, and that is logically
	// the most obvious thing to do, so we are implementing it that way.
	featureJoinComplement := gbk.Features[3].GetSequence()
	seqJoinComplement := "ataccaatttaatcattcatttatatactgattccgtaagggttgttacttcatctattttataccaatgcgtttcaaccatttcacgcttgcttatatcatcaagaaaacttgcgtctaattgaactgttgaattaacacgatgccttttaacgatgcgagaaacaactacttcatctgcataaggtaatgcagcatataacagagcaggcccgccaattacacttactttagaattctgatcaagcatagtttcgaatggtgcattagggcttgacacttgaatttcgccgccagaaatgtaagttatatattgctcccaagtaatatagaaatgtgctaaatcgccgtctttagttacaggataatcacgcgcaaggtcacacaccacaatatggctacgaccaggaagtaatgtaggcaatgactggaacgttttagcacccataatcataattgtgccttcagtacgagctttaaaattctggaggtcctttttaactcgtccccatggtaaaccatcacctaaaccgaatgctaattcattaaagccgtcgaccgttttagttggaga"
	if featureJoinComplement != seqJoinComplement {
		t.Errorf("Feature sequence parser has changed on test 'join(complement(315..330),complement(339..896))'. Got this:\n%s instead of \n%s", featureJoinComplement, seqJoinComplement)
	}

	// Read complement(join(893..1098,1101..2770))
	featureComplementJoin := gbk.Features[5].GetSequence()
	seqComplementJoin := "ttacaccgccatctttcctttaataggagggtgtgatacatagttgttaagaacgaaatctttaggcctaagtttaagaacatattttaattgttctttagtagaaagatatcggaatttataaggtagaccacttattaccagctcacaaagctctttaggttcacgcctcaaaatttctttacattgttctacgtgattcatatagatatgagtattaccaccagaaaatatcaaatcccctggaataagattacacatcttagctacaatatgaactaacgtagcatatgacgcaatattaaacggtagcattatgttcagataaggtcgttaatcttaccccggaattatatccagctgcatgtcaccatgcagagcagactatatctccaacttgttaaagcaagttgtctatcgtttcgagtcacttgaccctactccccaaagggatagtcgttaggcatttatgtagaaccaattccatttatcagattttacacgataagtaactaatccagacgaaattttaaaatgtctagctgcatctgctgcacaatcaaaaataaccccatcacatgaaatctttttaatattactaggctttttacctttcatcttttctgatattttagatttagttatgtctgaatgcttatgattaaagaatgaattattttcacctgaacgatttctgcatttactacaagtataagcagaagtttgtatgcgaacaccgcacttacaaaacttatgggtttctggattccaacgcccgtttttacttccgggtttactgtaaagagctttccgaccatcaggtccaagtttaagcatcttagctttaacagtttcagaacgtttcttaataatttcttcttttaatggatgcgtagaacatgtatcaccaaacgttgcatcagcaatattgtatccattaattttagaattaagctctttaatccaaaaattttctcgttcaataatcaaatctttctcatatggaatttcttccaaaatagaacattcaaacacattaccatgtttgttaaaagacctctgaagttttatagaagaatggcatcctttttctaaatctttaaaatgcctcttccatctcttttcaaaatctttagcacttcctacatatactttattgtttaaagtatttttaatctgataaattccgcttttcataaatacctctttaaatatagaagtatttattaaagggcaagtcctacaatttagcacgggattgtctactagagaggttccccgtttagatagattacaagtataagtcaccttatactcaggcctcaattaacccaagaaaacatctactgagcgttgataccactgcaaatccaaatagccattacgcacattaaactgatagaacatatgacaaggcggtaatgccatatatttaagttcagctggattccatgcagaaacaatttgacgcctatcatttggcagttttttaatacgatcaataacttctataatttggtctacaccaccaaaatcacgccactgttttccataaattggaccaagttcaccgctatggtatcctaaatcttttgcttgattttcgtaattttcatcccagactgttttgccttggattaacgaatcgtgttgaattaatcgtaaatcatacatttgtgcttcctgataaaaaccatattagctcagcaatgcaagctttccaggcgagcttcttagttgttaccgcaggaaaacctttagttaaatcccagcgtaatttagatccgaacagagcaattgttcctgtgcctgtacgatcatcggtttcataaccattttcaaaaatgtctttaattaaatcttggtattgtttcat"
	if featureComplementJoin != seqComplementJoin {
		t.Errorf("Feature sequence parser has changed on test 'complement(join(893..1098,1101..2770))'. Got this:\n%s instead of \n%s", featureComplementJoin, seqComplementJoin)
	}

}

func TestLocationParser(t *testing.T) {
	testSeq := "ATN"
	testVariants := []string{"ATG", "ATA", "ATT", "ATC"}
	testVariantsIUPAC := AllVariantsIUPAC(testSeq)
	if featureComplementJoin != seqComplementJoin {
		t.Errorf("IUPAC variant has changed on test 'allIUPAC('ATN')'. Got this:\n%s instead of \n%s", testVariantsIUPAC, testVariants)
	}
}
