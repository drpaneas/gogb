# Resources and References

## Essential Links for Game Boy Emulator Development

---

This document collects the most valuable resources for Game Boy emulator development. Bookmark these—you'll return to them constantly.

---

## Table of Contents

- [Documentation](#documentation)
- [Test ROMs](#test-roms)
- [Reference Emulators](#reference-emulators)
- [Development Tools](#development-tools)
- [Community](#community)
- [Video Tutorials](#video-tutorials)
- [Academic Papers](#academic-papers)
- [Books](#books)
- [Related Projects](#related-projects)

---

## Documentation

### The Pan Docs
**https://gbdev.io/pandocs/**

The definitive Game Boy technical reference. If you have one bookmark, make it this one.

Topics covered:
- CPU instruction set
- Memory map
- I/O registers
- PPU (graphics) operation
- APU (audio) operation
- Timer and interrupts
- Cartridge types (MBCs)
- Hardware quirks and edge cases

**Quality:** Excellent. Community-maintained, constantly updated.

### Game Boy CPU Manual
**http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf**

Gekkio's comprehensive CPU manual with cycle-accurate timing information.

### Game Boy Complete Technical Reference
**https://gekkio.fi/files/gb-docs/gbctr.pdf**

Detailed hardware reference from hardware reverse-engineering.

### Opcode Tables

| Resource | URL | Description |
|----------|-----|-------------|
| Pastraiser | https://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html | Clean opcode table |
| GB Dev Optables | https://gbdev.io/gb-opcodes/optables/ | Interactive with flag effects |
| Imran Nazar | http://imrannazar.com/Gameboy-Z80-Opcode-Map | Color-coded by category |

---

## Test ROMs

### Blargg's Test ROMs
**https://github.com/retrio/gb-test-roms**

The gold standard for CPU testing. Created by Shay Green.

| ROM | Tests |
|-----|-------|
| cpu_instrs.gb | All CPU instructions |
| instr_timing.gb | Instruction cycle counts |
| mem_timing.gb | Memory access timing |
| oam_bug.gb | OAM corruption bug |
| halt_bug.gb | HALT instruction edge cases |

### dmg-acid2
**https://github.com/mattcurrie/dmg-acid2**

PPU rendering accuracy test by Matt Currie. Your emulator should display the reference image exactly.

### Mooneye Test Suite
**https://github.com/Gekkio/mooneye-test-suite**

Exhaustive test suite with focus on hardware edge cases. Many tests require high accuracy to pass.

### Wilbert Pol's Tests
**https://github.com/wilbertpol/mooneye-gb**

Additional accuracy tests.

### Age Test ROMs
**https://github.com/c-sp/age-test-roms**

Audio and timing tests.

---

## Reference Emulators

### For Debugging

#### BGB
**https://bgb.bircd.org/**

Windows emulator with the best debugging features:
- Step-by-step execution
- Memory viewer
- Register inspection
- Breakpoints
- Trace logging
- VRAM viewer

*Runs on Linux/Mac via Wine.*

#### SameBoy
**https://sameboy.github.io/**

Highly accurate open-source emulator. Good for comparing behavior.

- Source: https://github.com/LIJI32/SameBoy

### For Accuracy Reference

#### Gambatte
**https://github.com/sinamas/gambatte**

Cycle-accurate emulator, commonly used as accuracy reference.

#### mGBA
**https://mgba.io/**

Multi-platform emulator supporting GBA and GB/GBC.

- Source: https://github.com/mgba-emu/mgba

### For Learning

#### Coffee GB
**https://github.com/nickthecoder/coffee-gb**

Java Game Boy emulator with clean, readable code.

#### PyBoy
**https://github.com/Baekalfen/PyBoy**

Python Game Boy emulator. Slower but very readable.

#### Gameboy-Crust
**https://github.com/mattbruv/Gameboy-Crust**

Rust Game Boy emulator with documentation.

---

## Development Tools

### Assemblers

#### RGBDS
**https://rgbds.gbdev.io/**

The standard Game Boy assembler. Used by most homebrew developers.

```bash
# Install (macOS)
brew install rgbds

# Assemble
rgbasm -o main.o main.asm
rgblink -o game.gb main.o
rgbfix -v -p 0 game.gb
```

#### WLA-DX
**https://github.com/vhelin/wla-dx**

Alternative assembler supporting multiple platforms.

### ROM Editors

#### Tilemap Studio
**https://github.com/Rangi42/tilemap-studio**

Visual tile map editor.

#### Game Boy Tile Designer (GBTD)
Classic tile editor (Windows).

#### Game Boy Map Builder (GBMB)
Classic map editor (Windows).

### Debuggers

#### bgb (see above)
Best-in-class debugging.

#### Emulicious
**https://emulicious.net/**

Java-based emulator with debugging.

### Hex Editors

Any hex editor works for ROM inspection:
- HxD (Windows)
- Hex Fiend (macOS)
- xxd (command line)

---

## Community

### gbdev Discord
**https://gbdev.io/chat**

The most active Game Boy development community. Channels for:
- #emulator-development
- #programming
- #help

Highly recommended for getting unstuck.

### Reddit r/EmuDev
**https://www.reddit.com/r/EmuDev/**

General emulator development subreddit. Beginner-friendly.

### Reddit r/Gameboy
**https://www.reddit.com/r/Gameboy/**

General Game Boy community (hardware, games, collecting).

### nesdev.org Forums
**https://forums.nesdev.org/**

Despite the name, has Game Boy sections. Older discussions, valuable archives.

### GitHub Discussions

Many emulator repositories have active discussions:
- https://github.com/LIJI32/SameBoy/discussions
- https://github.com/mgba-emu/mgba/discussions

---

## Video Tutorials

### The Ultimate Game Boy Talk
**https://www.youtube.com/watch?v=HyzD8pNlpwI**

Michael Steil's excellent one-hour overview of Game Boy hardware. Essential viewing.

### Game Boy Emulator Development Series
Various YouTubers have created tutorial series:

| Creator | Language | URL |
|---------|----------|-----|
| Low Level Devel | Rust | Search YouTube |
| javidx9 | C++ | Search YouTube |
| RSSS | JavaScript | Search YouTube |

### The Cherno - Emulator Basics
**https://www.youtube.com/thecherno**

General emulator concepts (not Game Boy specific).

---

## Academic Papers

### Cycle-Accurate Game Boy Emulation
Various blog posts document the research that went into cycle-accurate emulators:

- Gekkio's hardware research
- SameBoy development notes
- AntonioND's analysis

### Understanding Low-Level Systems
While not Game Boy specific, these help with general concepts:

- Patterson & Hennessy: *Computer Organization and Design*
- Hennessy & Patterson: *Computer Architecture: A Quantitative Approach*

---

## Books

### Game Boy-Related

There's no "definitive" Game Boy emulator book, but these help:

| Book | Topic |
|------|-------|
| *The Manga Guide to Microprocessors* | CPU basics |
| *Code: The Hidden Language* | Computing fundamentals |
| *But How Do It Know?* | CPU design |

### Emulator Development

| Book | Author | Notes |
|------|--------|-------|
| *Classic Computer Science Problems in Python* | David Kopec | Includes CHIP-8 emulator |
| *Writing Interpreters and Compilers* | Ronald Mak | Related concepts |

### Assembly Language

| Book | Platform |
|------|----------|
| *Programming the Z80* by Rodnay Zaks | Z80 (Game Boy CPU is similar) |
| *Assembly Language Step by Step* by Jeff Duntemann | x86 but concepts transfer |

---

## Related Projects

### Homebrew Games

Test your emulator with homebrew games (freely available, no legal concerns):

| Game | URL | Complexity |
|------|-----|------------|
| 2048 | https://github.com/nickthecoder/2048-gb | Simple |
| Lumberjack | https://itch.io search | Medium |
| μCity | https://github.com/nickthecoder/ucity | Complex |

### GB Studio
**https://www.gbstudio.dev/**

Visual game maker for Game Boy. Games made with GB Studio are good test cases.

### Homebrew Hub
**https://hh.gbdev.io/**

Directory of Game Boy homebrew games.

### awesome-gbdev
**https://github.com/gbdev/awesome-gbdev**

Curated list of Game Boy development resources.

---

## Quick Reference Card

Print this and keep it handy:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                          QUICK REFERENCE                                     │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Documentation:                                                             │
│    Pan Docs ────────── gbdev.io/pandocs                                     │
│    Opcodes ─────────── pastraiser.com/cpu/gameboy/gameboy_opcodes.html     │
│                                                                             │
│  Test ROMs:                                                                 │
│    cpu_instrs ──────── github.com/retrio/gb-test-roms                      │
│    dmg-acid2 ───────── github.com/mattcurrie/dmg-acid2                     │
│                                                                             │
│  Debugging:                                                                 │
│    BGB ──────────────── bgb.bircd.org                                       │
│    SameBoy ─────────── sameboy.github.io                                   │
│                                                                             │
│  Community:                                                                 │
│    Discord ─────────── gbdev.io/chat                                        │
│    Reddit ──────────── reddit.com/r/EmuDev                                 │
│                                                                             │
│  Video:                                                                     │
│    Ultimate GB Talk ── youtube.com/watch?v=HyzD8pNlpwI                     │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Contributing to the Community

Once your emulator works, consider giving back:

1. **Document your journey** — Write about problems you solved
2. **Report test ROM results** — Help track emulator accuracy
3. **Create test ROMs** — If you find edge cases
4. **Answer questions** — Help newcomers on Discord/Reddit
5. **Open source your emulator** — Others can learn from your code

The Game Boy development community is welcoming and collaborative. We all benefit from shared knowledge.

---

---

## Cycle-Exact Emulation Resources

For those pursuing 100% Mooneye accuracy (see Chapter 26), these resources are essential:

### Reference Implementations

| Emulator | Accuracy | Source | Notes |
|----------|----------|--------|-------|
| SameBoy | 100% Mooneye | github.com/LIJI32/SameBoy | Reference quality |
| Gambatte | 100% Mooneye | github.com/sinamas/gambatte | Well-documented timing |
| mooneye-gb | Mooneye author | github.com/Gekkio/mooneye-gb | Written by test author |

### Key Technical Documents

| Document | Topic | Where to Find |
|----------|-------|---------------|
| Gekkio's hardware notes | Sub-cycle timing | gekkio.fi |
| SameBoy blog posts | PPU FIFO implementation | sameboy.github.io |
| TCAGBD | Complete technical reference | github.com/AntonioND/giibiiadvance |

### Specific Timing References

For the 4 tests that require architectural changes:

| Test | What to Study |
|------|--------------|
| `intr_2_mode0_timing_sprites` | SameBoy's FIFO sprite penalty calculation |
| `lcdon_timing-GS` | Gambatte's LCD enable handling |
| `lcdon_write_timing-GS` | mooneye-gb test source code comments |
| `stat_lyc_onoff` | SameBoy's STAT line implementation |

### Hardware Analysis Tools

| Tool | Purpose |
|------|---------|
| Logic analyzer | Capture real hardware signals |
| Oscilloscope | Measure timing relationships |
| Test ROM + BGB | Compare expected vs actual behavior |

### Academic Approaches

Cycle-exact emulation is essentially reverse engineering. Techniques include:

1. **Black-box testing** — Write test ROMs that probe specific behaviors
2. **Comparative analysis** — Compare multiple emulators on edge cases
3. **Hardware measurement** — Use logic analyzers on real hardware
4. **Differential testing** — Find inputs that cause emulators to diverge

---

*Emulator development is a journey. These resources are your map.*

