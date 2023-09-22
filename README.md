# API for BRØK Navnetjener

## Beskrivelse

API for BRØK Navnetjener er designet for å håndtere interaksjonen mellom navn, fødselsdato, og lommebokadresser. Den gir fagsystemer mulighet til å lese og skrive data mens offentlig lesing er tilgjengelig med begrensninger.

## Funksjonaliteter

### Nåværende:

* `POST /wallet firstname, lastename, ssn, walletAddress`: Oppretter en ny wallet-mapping.
* `GET /wallet/{walletAddress}`: Returnerer navn og fødselsdato for en gitt lommebokadresse.
* `GET /aksjeeier/{orgnr/fnr}`: Liste med alle selskapene personen eller organisasjonen eier aksjer i
* `GET /aksjebok/?page=0`: Liste med alle selskapene på Brøk, 25 foretak per side
* `GET /aksjebok/{orgnr}`: Returnerer selskapet med matchende orgnr
* `GET /aksjebok/{orgnr}/balanse/{orgnr/fnr}`: Antall aksjer en person eller organisasjon eier i et foretak
* `POST /akejebok/{orgnr}/aksjeeier`: Tar inn en liste med personer og organisasjoner, og svarer om disse eier aksjer i foretaket eller ikke
* Tester: Inkluderer tester for å validere systemets funksjonalitet, de viktigste testene er plassert i `/api` mappen. Kjør testene med `go test ./api ./model ./utils`


### Planlagt:

* Auth: Evt. integrere med BR sin API manager for autentisering.
* Integrere FM Person i Navnetjeneren

### Fremtidig arbeid

* Sletting av data.
* Oppdatering av data, hvis det er nødvendig.
* Støtte for beskyttede personer kode 6 og 7 i Folkeregisteret.

## Autentisering

* Skriveadgang er begrenset til fagsystemer.
* Lesing er åpent for alle med begrensninger.

## Beskyttelse mot datakryping

* Hvis nødvendig, kan man implementere throttling for å begrense antall spørringer. F.eks. DN vil ha behov for mange spørringer. Ongoing work med Sverre.

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
cp .env .env.local
docker run --name navnetjener -d -e POSTGRES_PASSWORD="password" -v ./database/testdata.sql:/docker-entrypoint-initdb.d/testdata.sql -p 6666:6666 postgres -p 6666
```

### Kjør i dev

* For å kjøre i dev, bruk kommandoen: `go run .`.


### Testing

* For å kjøre testene, bruk kommandoen: `go test`.


### Bygging

* For å opprette en binærfil, bruk kommandoen: `go build`.

## Arkitektur

Navnetjeneren i den overordnede arkitekturen kan sees i følgende diagram:
![image](https://github.com/brreg/brok-navnetjener/assets/18251869/4929baf9-35b6-4dea-b21c-77d57f185608)


![image](https://github.com/brreg/brok-navnetjener/assets/877417/266b0aaa-81d1-4fa6-a1f3-a463f96bcca6)
