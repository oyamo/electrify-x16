# Assembler

This is the assembler for the electrify-x16 simulated machine. It takes assembly plaintext file and converts into bytecode. 


[![SEE USAGE](https://img.shields.io/badge/SEE%20USAGE-DOCS-green?labelColor=GREEN&style=flat-square&link=https://github.com/oyamo/electrify-x16#usage)](https://github.com/oyamo/electrify-x16#usage)
## Structure of bytecode
![assembler_ins.png](..%2Fscreenshots%2Fassembler_ins.png)
## Known Issues
The Assembler truncates numbers greater than 8-bits. This means that if your assembly language source code includes numbers that are larger than 255 (the maximum value that can be represented with 8 bits), the Assembler will only store the least significant 8 bits of the number.
