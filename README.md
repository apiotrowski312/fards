![Coverage](https://img.shields.io/badge/Coverage-66.7%25-yellow)
[![Lint](https://github.com/apiotrowski312/fards/actions/workflows/lint.yml/badge.svg)](https://github.com/apiotrowski312/fards/actions/workflows/lint.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/apiotrowski312/fards)
[![Go Report Card](https://goreportcard.com/badge/github.com/apiotrowski312/fards)](https://goreportcard.com/report/github.com/apiotrowski312/fards)


# Fards

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="assets/icon-white.png">
  <img alt="fards icon" src="assets/icon.png" width="100">
</picture>

This app is called *Flash Cards* aka **Fards**. It provides all basic functions you could need from a flash card app. 

Yeah, I am not the best with names :stuck_out_tongue_winking_eye:

## Installation and usage:

You need to have those two installed on your computer:
- [Go (Golang)](https://go.dev/)
For android installation:
- [Fyne](https://fyne.io/) 

You can run it on desktop with:

```
go run main.go
```

For android devices you need to run these commands ([Full tutorial on Fyne docs](https://developer.fyne.io/started/mobile.html)):

```
fyne package -os android/arm .
adb install Fards.ap
```

## Attributions

- [Fart icons created by Freepik - Flaticon](https://www.flaticon.com/free-icons/fart)
- [Card icons created by Freepik - Flaticon](https://www.flaticon.com/free-icons/card)