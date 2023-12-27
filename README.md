# SendBird

Ce projet est juste pour essayer les TCP, il permet d'envoyer des commandes à distance de client à serveur

## Compilation
Tu auras besoin de GoLang
```shell
go build -o ./sendbird
```

## Execution
Pour le client
```shell
./sendbird client (Ip du serveur):(PORT)
```
Exemple:
```shell
./sendbird client 127.0.0.1:1234
```

Pour le serveur
```shell
./sendbird server (PORT)
```
Exemple:
```shell
./sendbird server 1234
```