# Mutant api
Permite detectar el ADN de una persona mediante un api ingresando su codigo genetico, como tambien muestra las estadisticas de los mutantes detectados para unirlos a la lucha encontra de los X-MEN.

## USO

Para su correcto uso se requiere poderes de telequinesis y tener el siguiente armamentos:
* GOLANG
* GIT
* DOCKER

- PASO 1:
```bash
git clone https://github.com/rafaeldinho/meli-challenge
```

- PASO 2:
```bash
importar collection postman dentro del direcorio raiz "files" del proyecto con los distintos multiversos (ambientes local/prod)
```

- PASO 3:
este ultimo paso creara un contenedor con la base de datos y el app
```bash
docker compose up
```

## TEST:
```bash
make test
```

```bash
make testcoverage
```

## TEST COVERAGE:
```bash
make testcoverage
```

## FORMAT:
```bash
make format
```

## URIS:
Estado de salud:
```
GET - {{HOST}}/health

{
    "status": "UP",
    "version": "1.0.0"
}
```

IS_MUTANT:
```
POST - {{HOST}}/mutant
body:
*dna slice string required

response:
OK-200
{
    "dna": "ATGCGA",
    "isMutant": true
}

forbbiden-403 (si no es mutante)
{
    "dna": "ATGCGA",
    "isMutant": false
}

```

Estadisticas:
```
GET - {{HOST}}/mutant/stats

{
    "CountMutantDNA": 1,
    "CountHumanDNA": 0,
    "Ratio": 2.6
}
```

## License by rafaeldinho
[MIT](https://choosealicense.com/licenses/mit/)
