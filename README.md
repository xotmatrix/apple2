# Apple ][+, //e emulator

Portable emulator of an Apple II+ or //e. Written in Go.

[![CircleCI](https://circleci.com/gh/ivanizag/apple2/tree/master.svg?style=svg)](https://circleci.com/gh/ivanizag/apple2/tree/master)

## Features

- Models:
    - Apple ][+ with 48Kb of base RAM
    - Apple //e with 128Kb of RAM
    - Apple //e enhanced with 128Kb of RAM
    - Base64A clone with 48Kb of base RAM and paginated ROM
- Storage
    - 16 Sector diskettes in NIB format
    - 16 Sector diskettes in DSK format
    - 16 Sector diskettes in WOZ 1.0 or 2.0 format (read only)
    - Hard disk with ProDOS and SmartPort support
- Emulated extension cards:
    - DiskII controller
    - 16Kb Language Card
    - 256Kb Saturn RAM
    - 1Mb Memory Expansion Card
    - ThunderClock Plus real time clock
    - Bootable hard disk card
    - Apple //e 80 columns with 64Kb extra RAM
    - VidHd, limited to the ROM signature and SHR as used by Total Replay, only for //e models with 128Kb
    - FASTChip, limited to what Total Replay needs to set and clear fast mode
- Graphic modes:
    - Text 40 columns
    - Text 80 columns (Apple //e only)
    - Low-Resolution graphics
    - Double-Width Low-Resolution graphics (Apple //e only)
    - High-Resolution graphics
    - Double-Width High-Resolution graphics (Apple //e only)
    - Super High Resolution (VidHD only)
    - Mixed mode
- Displays:
    - Green monochrome monitor with half width pixel support
    - NTSC Color TV (extracting the phase from the mono signal)
    - RGB for Super High Resolution
    - ANSI Console, avoiding the SDL2 dependency
- Other features:
    - Sound
    - Joystick support. Up to two joysticks or four paddles.
    - Adjustable speed.
    - Fast disk mode to set max speed while using the disks. 
    - Single file executable with embedded ROMs and DOS 3.3


## Running the emulator

No installation required. [Download](https://github.com/ivanizag/apple2/releases) the single file executable `apple2xxx_xxx` for linux or Mac, SDL2 graphics or console. Build from source to get the latest features.

### Default mode

Execute without parameters to have an emulated Apple //e Enhanced with 128kb booting DOS 3.3 ready to run Applesoft:
```
casa@servidor:~$ ./apple2sdl
```

![DOS 3.3 started](doc/dos33.png)

### Play games
Download a DSK or WOZ file or use an URL ([Asimov](https://www.apple.asimov.net/images/) is an excellent source) with the `-disk` parameter:
```
casa@servidor:~$ ./apple2sdl -disk "https://www.apple.asimov.net/images/games/action/karateka/karateka (includes intro).dsk"
```
![Karateka](doc/karateka.png)

### Play the Total Replay collection
Download the excellent [Total Replay](https://archive.org/details/TotalReplay) compilation by
[a2-4am](https://github.com/a2-4am/4cade). Run it with the `-hd` parameter:
```
casa@servidor:~$ ./apple2sdl -hd "Total Replay v3.0.2mg"
```
Displays super hi-res box art as seen with the VidHD card.

![Total Replay](doc/totalreplay.png)

### Terminal mode
To run text mode right on the terminal without the SDL2 dependency, use `apple2console`. It runs on the console using ANSI escape codes. Input is sent to the emulated Apple II one line at a time: 
```
casa@servidor:~$ ./apple2console -model 2plus

############################################
#                                          #
#                APPLE II                  #
#                                          #
#     DOS VERSION 3.3  SYSTEM MASTER       #
#                                          #
#                                          #
#            JANUARY 1, 1983               #
#                                          #
#                                          #
# COPYRIGHT APPLE COMPUTER,INC. 1980,1982  #
#                                          #
#                                          #
# ]10 PRINT "HELLO WORLD"                  #
#                                          #
# ]LIST                                    #
#                                          #
# 10  PRINT "HELLO WORLD"                  #
#                                          #
# ]RUN                                     #
# HELLO WORLD                              #
#                                          #
# ]_                                       #
#                                          #
#                                          #
############################################
Line: 

```

### Keys

- Ctrl-F1: Reset button
- F5: Toggle speed between real and fastest
- Ctrl-F5: Show current speed in Mhz
- F6: Toggle between NTSC color TV and green phosphor monochrome monitor
- F7: Save current state to disk (incomplete)
- F8: Restore state from disk (incomplete)
- F10: Cycle character generator codepages. Only if the character generator ROM has more than one 2Kb page.
- F11: Toggle on and off the trace to console of the CPU execution
- F12: Save a screen snapshot to a file `snapshot.png`
- Pause: Pause the emulation

Only valid on SDL mode

### Command line options

```
  -charRom string
        rom file for the character generator (default "<default>")
  -disk string
        file to load on the first disk drive (default "<internal>/dos33.dsk")
  -disk2Slot int
        slot for the disk driver. -1 for none. (default 6)
  -diskRom string
        rom file for the disk drive controller (default "<internal>/DISK2.rom")
  -dumpChars
        shows the character map
  -fastChipSlot int
    	slot for the FASTChip accelerator card, -1 for none (default 3)        
  -fastDisk
        set fast mode when the disks are spinning (default true)
  -hd string
        file to load on the hard disk
  -hdSlot int
        slot for the hard drive if present. -1 for none. (default -1)
  -languageCardSlot int
        slot for the 16kb language card. -1 for none
  -memoryExpSlot int
    	  slot for the Memory Expansion card with 1GB. -1 for none (default 4)
  -mhz float
        cpu speed in Mhz, use 0 for full speed. Use F5 to toggle. (default 1.0227142857142857)
  -model string
        set base model. Models available 2plus, 2e, 2enh, base64a (default "2enh")
  -mono
        emulate a green phosphor monitor instead of a NTSC color TV. Use F6 to toggle.
  -panicSS
        panic if a not implemented softswitch is used
  -profile
        generate profile trace to analyse with pprof
  -rom string
        main rom file (default "<default>")
  -saturnCardSlot int
        slot for the 256kb Saturn card. -1 for none (default -1)
  -thunderClockCardSlot int
        slot for the ThunderClock Plus card. -1 for none (default 5)
  -traceCpu
        dump to the console the CPU execution. Use F11 to toggle.
  -traceHD
        dump to the console the hd commands
  -traceSS
        dump to the console the sofswitches calls
  -vidHDSlot int
    	  slot for the VidHD card, only for //e models. -1 for none (default 2)
  -woz string
    	  show WOZ file information


```

## Building from source

### apple2console

The only dependency is having a working Go installation on any platform.

Run:
```
$ go get github.com/ivanizag/apple2/apple2console 
$ go build github.com/ivanizag/apple2/apple2console 
``` 

### apple2sdl

Besides having a working Go installation, install the SDL2 developer files. Valid for any platform

Run:
```
$ go get github.com/ivanizag/apple2/apple2sdl
$ go build github.com/ivanizag/apple2/apple2sdl 
```

### Use docker to cross compile for Linux and Windows

To create executables for Linux and Windows without installing Go, SDL2 or the Windows cross compilation toosl, run:
```
$ cd docker
$ ./build.sh
```

To run in Windows, copy the file `SDL2.dll` on the same folder as `apple2sdl.exe`. The latest `SDL2.dll` can be found in the [Runtime binary for Windows 64-bit](https://www.libsdl.org/download-2.0.php).
