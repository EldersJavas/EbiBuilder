# [EbiBuilder](https://github.com/EldersJavas/EbiBuilder)


![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/EldersJavas/EbiBuilder?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/EldersJavas/EbiBuilder?style=flat-square)](https://github.com/EldersJavas/EbiBuilder)
[![GoDoc](https://godoc.org/github.com/EldersJavas/EbiBuilder?status.svg)](https://pkg.go.dev/github.com/EldersJavas/EbiBuilder)
[![Go Report Card](https://goreportcard.com/badge/github.com/EldersJavas/EbiBuilder)](https://goreportcard.com/report/github.com/EldersJavas/EbiBuilder)
[![wakatime](https://wakatime.com/badge/user/251739d5-2666-4202-9df0-c3b0c64457e4/project/70119925-9677-4119-a10d-db940e271e6a.svg)](https://wakatime.com/badge/user/251739d5-2666-4202-9df0-c3b0c64457e4/project/70119925-9677-4119-a10d-db940e271e6a)


The builder for [Ebitengine](https://github.com/hajimehoshi/ebiten).

A tool for managing & building the ebitengine game.

## Usage

```bash 
A tool for managing & building the ebitengine game. (Version: 1.0)                                      
Usage:                                                                                                  
  EbiBuilder [global options...] COMMAND [--options ...] [arguments ...]                            
  EbiBuilder [global options...] COMMAND [--options ...] SUBCOMMAND [--options ...] [arguments ...] 
                                                                                                        
Global Options:                                                                                                                             
  -h, --help                Display the help information                                                                                    
  --ishell                  Run in an interactive shell environment(TODO)                                                                   
  --nc, --no-color          Disable color when outputting message                                                                           
  --ni, --no-interactive    Disable interactive confirmation operation
  --np, --no-progress       Disable display progress message
  --verbose                 Set logs reporting level(quiet 0 - 5 crazy) (default 1=error)
  -V, --version             Display app version information

Available Commands:
  build        Build Ebitengine game (alias: Build,BUILD,buildgame)
  clean        Clean build (alias: Clean,CLEAN)
  genac        Generate auto complete scripts for current application (alias: gen-ac)
  self-update  Self-update (alias: Self-update,selfupdate,Selfupdate)
  help         Display help information

Use "EbiBuilder COMMAND -h" for more information about a command

```

## Feature

- [x] **build**
  - [x] Debug
  - [x] Release
- [ ] **publish**
  - [ ] **WASM**
  - [ ] Mobile
- [x] **clean**
- [ ] **create**
- [ ] config
  - [x] init
  - [ ] print
- [ ] check
  - [ ] test
  - [ ] upgrade
- [x] self-update
