# bonniertv - Trivial Omdb Service 
This service lists the best films ever made sorted alphabetacally

# Warning: Risk of biased results. 

![matrix](https://m.media-amazon.com/images/M/MV5BNzQzOTk3OTAtNDQ0Zi00ZTVkLWI0MTEtMDllZjNkYzNjNTc4L2ltYWdlXkEyXkFqcGdeQXVyNjU0OTQ0OTY@._V1_SX300.jpg)

## Functional Requirements (Swedish)
- [x] Tjänsten ska nås via HTTP och data ska presenteras i json-format.
- [x] Filmerna ska sorteras i bokstavsordning.
- [x] Tjänsten ska hämta information om de olika filmerna genom omdb's api på omdbapi.com, se sektionen
OMDB nedan.
- [x] Skriv en README.md fil med instruktioner för att bygga/köra tjänsten

## Extras (Swedish)
- [*] Gör din tjänst så produktions-redo som möjligt.
- [x] Skriv en Dockerfile för att köra din app.
- [x] Tjänsten ska helst använda paket från go's standard library.
- [x] Visa din arbetsgång genom att jobba metodiskt i Git med en commit per funktionell förändring.
- [x] Tjänsten svarar endast på HTTP GET om man skickar med headern "X-Secret" med värdet "1234".
- [x] Tjänsten hanterar en intern cache så den max gör en request per minut mot omdb.
- [ ] Tjänsten hämtar hem filmerna från omdb "concurrently"
- [*] Enhetstester och/eller integrationstester

## Build and run from source
### Prerequicites
  - go1.13.8 or higher
<pre>
$ cp  .env.example .env #enter the apikey
$ source .env
$ go build
$ ./bonniertv
</pre>

## Build and run with docker
<pre>
$ docker build -t bonniertv .
$ docker run -p8080:8080 -e OMDB_BASEURL="http://www.omdbapi.com/" -e OMDB_APIKEY="{api_key}" bonniertv 
</pre>

## Run unit tests
<pre>
$ source .env
$ go test *.go -v
</pre>

## Test endpoint
<pre>
./test.sh
</pre>


# Possible improvements
- [ ] Cache Middleware, Concurrency, Integrationtests
- [ ] Implement jwt verification
- [ ] Configure for serverless deployment
- [ ] Use a redis cache
- [ ] Implement full support for omdbClient
- [ ] Use makefile or other tools to manage building and testing
- [ ] Improve test coverage
- [ ] Refactor to comply with naming conventions.  



