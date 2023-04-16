package main

import (
	"fmt"
	"os"
	"dna_golang_encoder/convertor"
)

func main() {
	var sentenceToConvert string = os.Args[1]

	convertor := convertor.Convertor{}
	binarySentence := convertor.TextSentenceToBinary(sentenceToConvert)
	dnaSequence := convertor.BinarySentenceToDna(binarySentence)
	binarySentenceReconverted := convertor.DnaToBinary(dnaSequence)
	sentenceReconverted := convertor.BinaryToSring(binarySentenceReconverted)
	fmt.Println(os.Args[1], " ---> ", binarySentence, " ---> ", dnaSequence, " ---> ", binarySentenceReconverted, " ---> ", sentenceReconverted)

	var dnaHello = "CAGACGCCCGTACGTACGTT"
	binaryHello := convertor.DnaToBinary(dnaHello)
	sentenceHello := convertor.BinaryToSring(binaryHello)
	fmt.Println(dnaHello, " ---> ", sentenceHello)
}
