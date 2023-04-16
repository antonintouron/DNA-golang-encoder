package convertor

import (
    "fmt"
    "bytes"
    "strconv"
)

/*
	00 to adenine
	01 to cytosine
	10 to guanine
	11 to thymine
*/
var mapping = map[string]string{
  "00": "A",
  "01": "C",
  "10": "G",
  "11": "T",
  "A": "00",
  "C": "01",
  "G": "10",
  "T": "11",
}

type Convertor struct {}

/*
	Convert a text sentence to a binary sentence

	@return: string
*/
func (convertor Convertor) TextSentenceToBinary(sentenceToConvert string) string {
	var binarySentence string = ""

 	for _, character := range sentenceToConvert {
	    binarySentence = fmt.Sprintf("%s%.8b", binarySentence, character)
	}

	return binarySentence
}

/*
	Convert a binary sentence to a dna sequence

	Ex : 01100001 (letter -> a)

		-> Get 01 for loop 1 and get DNA base type (in mapping map var) -> C
		-> Get 10 for loop 2 and get DNA base type (in mapping map var) -> G
		-> Get 00 for loop 3 and get DNA base type (in mapping map var) -> A
		-> Get 01 for loop 4 and get DNA base type (in mapping map var) -> C
		-> CGAC

	@return: string
*/
func (convertor Convertor) BinarySentenceToDna(binarySentence string) string {
	var buffer bytes.Buffer

  for i := 0; i < len(binarySentence); i += 2 {
     buffer.WriteString(mapping[binarySentence[i:i+2]])
  }

  return buffer.String()
}

/*
	Convert a DNA sequence to binary

	@return: string
*/
func (convertor Convertor) DnaToBinary(dnaSequence string) string {
	var buffer bytes.Buffer

 	for i := 0; i < len(dnaSequence); i++ {
     buffer.WriteString(mapping[string(dnaSequence[i])])
  }

  return buffer.String()
}

/*
	Convert a binary sentence (8 by 8) to decimal and to string

	@return: string
*/
func (convertor Convertor) BinaryToSring(binarySentence string) string {
	var buffer bytes.Buffer

	for i := 0; i < len(binarySentence); i += 8 {
		character, _ := strconv.ParseInt(binarySentence[i:i+8], 2, 64)
		buffer.WriteString(string(character))
	}

  return buffer.String()
}
