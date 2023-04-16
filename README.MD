# Simple DNA golang encoder

This experimental project is not intended for professional use.

## Features

* Convert a `sentence` to a `DNA sequence`.
* Convert a `DNA sequence` to a `sentence`.

## How to use it

* Run `go run main.go "Hello"` (remplace Hello with your sentence).
* Main code is in `main.go` file, adjut this file for your tests.

```
Test  --->  01010100011001010111001101110100  --->  CCCACGCCCTATCTCA  --->  01010100011001010111001101110100  --->  Test
CAGACGCCCGTACGTACGTT  --->  Hello
```

## DNA convertion mapping

```
{
  "00": "A",
  "01": "C",
  "10": "G",
  "11": "T",
  "A": "00",
  "C": "01",
  "G": "10",
  "T": "11",
}
```

## Project information

* This is a test project that can be improved little by little.
