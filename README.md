# electrify-x16

Custom Instruction Set Assembler and Simulator built with **Golang**

The goal of this project is to create an assembler that can take text assembly code written for our custom instruction set and encode it into a bytecode format consisting of 16-bit numbers. This will allow us to more efficiently execute programs on our custom machine.

In addition to the assembler, there is a simulator for the custom machine. This simulator takes the output of the assembler and executes it correctly, allowing us to test our programs and ensure they are functioning as intended.

## Features
1. ### Assembler  - [View](./assembler)
    - Tokeniser - Converts source file to tokens
    - Symbol table to map symbol/labels to address
    - Syntax error checker
    - Saves bytecode to binfile
2. ### Simulator - [View](./machine)
   - 5 Registers - R1,R2,R3, COND and PC
   - Single CPU with clockspeed modifiable
   - 64K memory
   - Booting & Shutdown
   - Safe shutdown on error

## Screenshots
### Assembler
![assembler.png](screenshots%2Fassembler.png)
### Simulator
![machine.png](screenshots%2Fmachine.png)

### ISA Datasheet
The following are the instructions supported by the architecture of the simulation

| Opcode | Instruction | Description                                                                                                    |
|--------|-------------|----------------------------------------------------------------------------------------------------------------|
| 0x00   | halt        | Terminate program                                                                                              |
| 0x01   | nop         | Do nothing                                                                                                     |
| 0x02   | li          | Load Immediate: `li R1 0x00000000`<br>Load `0x00000000` into `R1`                                              |
| 0x03   | lw          | Load Word: `lw R1 R2`<br>Load the contents of the memory location pointed to by `R2` into `R1`                 |
| 0x04   | sw          | Store Word: `sw R1 R2`<br>Store the contents of `R2` in the memory location pointed to by `R1`                 |
| 0x05   | add         | Add: `add R3 R1 R2`<br>Add `R1` to `R2` and store the result in `R3`                                           |
| 0x06   | sub         | Subtract: `sub R3 R1 R2`<br>Subtract `R2` from `R1` and store the result in `R3`                               |
| 0x07   | mult        | Multiply: `mult R3 R1 R2`<br>Multiply `R1` by `R2` and store the result in `R3`                                |
| 0x08   | div         | Divide: `div R3 R1 R2`<br>Divide `R1` by `R2` and store the result in `R3`                                     |
| 0x09   | j           | Unconditional Jump: `j 0x00000000`<br>Jump to memory location `0x00000000`                                     |
| 0x0A   | jr          | Unconditional Jump (Register): `jr R1`<br>Jump to memory location stored in `R1`                               |
| 0x0B   | beq         | Branch if Equal: `bne R1 R2 R3`<br>Branch to memory location stored in `R3` if `R1` and `R2` are equal         |
| 0x0C   | bne         | Branch if Not Equal: `beq R1 R2 R3`<br>Branch to memory location stored in `R3` if `R1` and `R2` are not equal |
| 0x0D   | inc         | Increment Register: `inc R1`<br>Increment `R1`                                                                 |
| 0x0E   | dec         | Decrement Register: `dec R1`<br>Decrement `R1`                                                                 |

