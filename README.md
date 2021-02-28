# Rater (r8r)

Work in progress...

```mermaid
graph LR;
    gw[Gateway]
    r8r[R8R Microservice]
    authService[OAuth2 Microservice];
    dbService[Database Microservice];
    postgres[PostgreSQL];

    gw --> authService
    gw --> r8r
    r8r --> dbService
    authService --> r8r
    dbService --> postgres
```