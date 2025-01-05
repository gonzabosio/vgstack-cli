# vgstack
CLI App to bootstrap a fullstack web application project with Go and Vue (based on Vite scaffolding)

Implementations:
- PostgreSQL
- Vue Router
- Chi
- Docker

### Run default command
```
go run github.com/gonzabosio/vgstack-cli@latest
```
### Options
#### Folder names (frontend, backend)
```
go run github.com/gonzabosio/vgstack-cli@latest -f 'front' -b 'back'
```
#### To ignore dockerfiles
```
go run github.com/gonzabosio/vgstack-cli@latest -nodocker
```