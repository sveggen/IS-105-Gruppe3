# ICA-03

**Oppgave 1**

For å kjøre filen må følgende kommando brukes:
```
$cd ICA03/Oppgave1/1A/
$go run filereader.go
```

b)

For å kjøre filen må følgende kommando brukes:
```
$cd ICA03/Oppgave1/1B/
$go run lineshift.go
```

**Oppgave 2**
```
$cd ICA03/Oppgave2/
```
Spesifiser ønsket operativsystem som
filen skal kompileres til. f.eks Mac OS Mojave 10.14.3 ved bruk av følgende kommando:

```
$GOOS=darwin GOARCH=amd64 go build fileinfo-ver2.go
```
For å kjøre filen på regnemaskinen med Mac OS, må filen først bli tildelt skrive-og-kjøre rettigheter ved bruk av følgende kommando:

```
$sudo chmod 777 fileinfo-ver2.go
```

Deretter kan filen kjøres ved å bruke kommandoen:

```
$./fileinfo-ver2 «fil-som-skal-analyseres»
```

**Oppgave 5**

a)

For å starte en simpel tcp-server og tcp-klient må følgende kommandoer kjøres. Dette må gjøres i separate terminaler.

```
$cd ICA03/Oppgave5/simpeltcp/
$go run simpletcplistener.go
$go run simpletcpdialer.go
```

## TCP

**Opprettelse av tcp-server:**
```
$cd ICA03/Oppgave5/tcp/
$go run tcplistener.go
```
**Opprettelse av tcp-klient:**
- IP = Server sin IP-adresse
- PORT = Server sin port
```
$cd ICA03/Oppgave5/tcp/
$./tcpdialer -tcp="IP:PORT"
```

## HTTP

**Opprettelse av http-server:**
```
$cd ICA03/Oppgave5/http/
$go run httplistener.go
```

**Opprettelse av http-klient:**
- IP = Server sin IP-adresse
- PORT = Server sin port
```
$cd ICA03/Oppgave5/http/
$./httpdialer -http="http://IP:PORT"
```
