# vgstack-cli
CLI App to scaffold a fullstack web application project with Go and Vue (based on Vite scaffolding)

Implementations:

- PostgreSQL
- Vue Router
- Chi
- Docker

### Run default command
Use explicit version(@v1.x.x) for security
```
go run github.com/gonzalobosio/vgstack-cli@latest
```
#### To ignore dockerfiles
```
go run github.com/gonzalobosio/vgstack-cli@latest -nodocker
```