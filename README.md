# stud-eng-api

## DBマイグレーション方法
詳細はsubmodule`./database`配下のREADMEを見てください。
## test実行方法
```
// makefileで実行するのでmakeが実行可能な環境を準備してください
//on Ubuntu
$ sudo apt-get update
$ sudo apt-get -y install make

$ make test

    or 

$ make test-db
$ make test-go
```

## submodule 関係
```
$ git submodule update --remote
```

## Lint チェック
```
$ golangci-lint run 
```