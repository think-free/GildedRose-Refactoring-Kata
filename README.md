# Instructions 
[FR](https://github.com/emilybache/GildedRose-Refactoring-Kata/blob/main/GildedRoseRequirements_fr.md)
[ES](https://github.com/emilybache/GildedRose-Refactoring-Kata/blob/main/GildedRoseRequirements_es.md)
[EN](https://github.com/emilybache/GildedRose-Refactoring-Kata/blob/main/GildedRoseRequirements.txt)

# Usage

- Run :

```shell
LOGLEVEL=info DAYS=200 go run cmd/texttest_fixtures/main.go
```

- Run tests :

```shell
go test ./...
```

- Run tests and coverage :

```shell
go test ./... -coverprofile=coverage.out

go tool cover -html=coverage.out
```

# Valid log levels 
	debug,info,warn,error,fatal,panic