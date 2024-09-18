# SocialNetworkAPI
Golang backend practice project. 

### Main Pkgs:

-> Related to the API core structure:

- Main
    - exec pkg
- Router
    - routes pkgs
    - mux
- Controllers
    - fucns tath will be exec by the routes
    - comm w/ Models and Repository
- Models
    - store structs, entitties
- Repository
    - database interations


### Scnd tier pkgs:

- Config:
    - environment variables
- Database:
      opens DB connection
- Auth:
    - login, tokens
- Middleware:
    - checks if user already auth
- Security:
    - passwords
- Responses:
    - api reponse patterns