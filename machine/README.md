# electrify-x16: Simulated machine

This is a 16-bit simulated machine ie it runs on a host computer as a normal computer program. This  machine is equipped with five registers, which are named R1, R2, R3, COND, and PC. These registers are utilized for storing data and controlling the execution flow of instructions.

Furthermore, the machine operates on a single CPU, which can execute instructions at different clock speeds. The clock speed is adjustable by the user.

The machine has a memory capacity of 64K, which is used to store data and instructions. The machine can be safely booted up and shut down using specific instructions.

[![SEE USAGE](https://img.shields.io/badge/SEE%20USAGE-DOCS-green?labelColor=GREEN&style=flat-square&link=https://github.com/oyamo/electrify-x16#usage)](https://github.com/oyamo/electrify-x16#usage)

```text
 /\/\/\                            /  \
| \  / |                         /      \
|  \/  |                       /          \
|  /\  |----------------------|     /\     |
| /  \ |                      |    /  \    |
|/    \|                      |   /    \   |
|\    /|                      |  | (  ) |  |
| \  / |                      |  | (  ) |  |
|  \/  |                 /\   |  |      |  |   /\
|  /\  |                /  \  |  |      |  |  /  \
| /  \ |               |----| |  |      |  | |----|
|/    \|---------------|    | | /|   .  |\ | |    |
|\    /|               |    | /  |   .  |  \ |    |
| \  / |               |    /    |   .  |    \    |
|  \/  |               |  /      |   .  |      \  |
|  /\  |---------------|/        |   .  |        \|
| /  \ |              /   NASA   |   .  |  NASA    \
|/    \|              (          |      |           )
|/\/\/\|               |    | |--|      |--| |    |
------------------------/  \-----/  \/  \-----/  \--------
                        \\//     \\//\\//     \\//
                         \/       \/  \/       \/
ASCII Art copyright of asciiart.eu

```

## Architecture
![machine-arch.png](..%2Fscreenshots%2Fmachine-arch.png)

## Known issues
There are currently no known issues about this machine.