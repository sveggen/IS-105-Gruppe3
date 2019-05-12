# ICA-04

## Oppgave 3


**1) gspeech**

For å analysere lydfilen "Jorgen.flac" må følgende kommando brukes:

```
$cd ICA04/Oppgave3/gspeech/
$./speech_rec.sh -i Jorgen.flac --rate 16000
```



**2) gospeak**

Naviger til riktig mappe ved bruk av følgende kommando:

```
$cd ICA04/Oppgave3/gospeak/
```
For å kjøre gospeak filen må det gjøres en endring i kildekoden. WiT API nøkkelen som er lagt inn i "gospeak.go" filen må byttes til din egen. Deretter kan filen kjøres med følgende kommando:
```
$ go run gospeak.go
```


**3) Google Cloud STT**

For å kunne kjøre speech-to-text.go filen må følgende erstattes med path til egen JSON-prosjektfil

```
option.WithCredentialsFile("/Users/magnusneergaard/Documents/Skole/My-First-Project-89ab7741a22f.json")
```

og filename må erstattes av path til egen nedlastet .flac fil

```
// Viser hvilken fil som skal bli transkribert
	filename := "/Users/magnusneergaard/Desktop/ICA04-Lydopptak/Jorgen.flac"
```



**4) IBM Speech to Text**

For å kunne transkribere tekst med IBM sin plattform må følgene API key endres til din egen:

```
$ curl -X POST -u "apikey:AVJRAHLH4cWm3iUsWg0bdjHq-FqHPbZv_HcJMHQRtqNJ" \
```
samt "--data-binary" path må endres til egen path for .flac fil
```
$ --data-binary @/Users/magnusneergaard/Desktop/ICA04-Lydopptak/Jorgen.flac \
```
