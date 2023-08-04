# API for BRØK Navnetjener

## Beskrivelse

API for BRØK Navnetjener er designet for å håndtere interaksjonen mellom navn, fødselsdato, og lommebokadresser. Den gir fagsystemer mulighet til å lese og skrive data mens offentlig lesing er tilgjengelig med begrensninger.

## Funksjonaliteter

### Nåværende:

* `POST /wallet firstname, lastename, ssn, walletAddress`: Oppretter en ny wallet-mapping.
* `GET /wallet/{walletAddress}`: Returnerer navn og fødselsdato for en gitt lommebokadresse.
* `GET /wallet`: Lister ut registrert navn og fødselsdato.
* Tester: Inkluderer tester for å validere systemets funksjonalitet.

### Planlagt:

* `GET /company/{orgnr}`: Lister ut alle registrerte navn og fødselsdato.
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

* Docker
* Go 1.20+

### Instruksjoner

```bash
git clone git@github.com:brreg/brok-navnetjener.git
cd brok-navnetjener
git checkout go
go mod download
docker run --name navnetjener -e POSTGRES_PASSWORD="password" -p 5432:5432 postgres -p 5432
docker exec -it navnetjener psql -h localhost -p 5432 -U postgres -c "CREATE DATABASE brok;"
```

### Kjør i dev

* For å kjøre i dev, bruk kommandoen: `go run .`.


### Testing

* For å kjøre testene, bruk kommandoen: `go test`.


### Bygging

* For å opprette en binærfil, bruk kommandoen: `go build`.

## Arkitektur

Navnetjeneren i den overordnede arkitekturen kan sees i følgende diagram:

![image](https://github.com/brreg/brok-navnetjener/assets/877417/266b0aaa-81d1-4fa6-a1f3-a463f96bcca6)
