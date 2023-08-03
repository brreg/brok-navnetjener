# API for BRØK Navnetjener

## Beskrivelse

API for BRØK Navnetjener er designet for å håndtere interaksjonen mellom navn, fødselsdato, og lommebokadresser. Den gir fagsystemer mulighet til å lese og skrive data mens offentlig lesing er tilgjengelig med begrensninger.

## Funksjonaliteter

### Nåværende:

* `Create(firstname, lastename, ssn, walletAddress)`: Oppretter en ny post.
* `GET /wallet/{walletAddress}`: Returnerer navn og fødselsdato for en gitt lommebokadresse.
* `GET /wallet`: Lister ut alle registrerte navn og fødselsdato.
* Tester: Inkluderer tester for å validere systemets funksjonalitet.

### Planlagt:

* Implementering av enkel one-liner for oppsett inkludert skjema.
* Teste "Setup dev".
* Logging: For å overvåke og feilsøke systemets aktiviteter.
* Auth: Evt. integrere med BR sin API manager for autentisering.
* `Read(SSN)`: Returnerer alle lommebokadresser tilhørende brukeren.

### Fremtidig arbeid

* Sletting av data.
* Oppdatering av data, hvis det er nødvendig.
* Støtte for beskyttede personer kode 6 og 7 i Folkeregisteret.

## Autentisering

* Skriveadgang er begrenset til fagsystemer.
* Lesing er åpent for alle med begrensninger.

## Beskyttelse mot datakryping

* Hvis nødvendig, kan man implementere throttling for å begrense antall spørringer. F.eks. DN vil ha behov for mange spørringer. Ongoing work med Sverre.

## Produksjon

* Vurder å slå av `GET /wallet`.

## Oppsett for utvikling

### Krav

* PostgreSQL database
* Go 1.20+

### Instruksjoner

```bash
git clone
cd brok-navnetjener
copy .env .env.local
```

* Rediger .env.local med DB-informasjon.
* Opprett skjemaet i databasen.
* Start databasen.
* Last ned nødvendige pakker: `go mod download`.
* Kjør applikasjonen: `go run .`.

### Testing

* For å kjøre testene, bruk kommandoen: `go test`.

### Bygging

* For å opprette en binærfil, bruk kommandoen: `go build`.

## Arkitektur

Navnetjeneren i den overordnede arkitekturen kan sees i følgende diagram:

![image](https://github.com/brreg/brok-navnetjener/assets/877417/266b0aaa-81d1-4fa6-a1f3-a463f96bcca6)
