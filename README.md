# GoChessEngine

### Summary

A chess engine built in Go that communicates using the [Universal Chess Interface](https://www.chessprogramming.org/UCI) with chess GUIs like [En Criossant](https://encroissant.org/).

### How to Run

Must have go installed. Then:

On Mac run

> go build -o output/Luna
> ./output/Luna

### Notes for Myself

For building to Windows on Mac run

> GOOS=windows GOARCH=386 go build -o output/Luna.exe
> wine output/Luna.exe

For Loading cutechess on Mac

> open /Users/hunterbowie/.wine/drive_c/Program\ Files\ \(x86\)/Cute\ Chess
