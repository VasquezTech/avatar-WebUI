
### Install docker
 From [docs.docker](https://docs.docker.com/engine/install/)

### Avatar-ui startup

Quick start: `docker run -p 127.0.0.1:8055:8055 mrvasquez96/avatar-ui:latest`

```
docker-compose build && docker-compose up -d
```
or 
 `apt install make`, then:
```
make docker-run # Only docker container
```
or
```
make docker-build # Uses source code
```
Open `https:/localhost/8055`

#### Avatar traits collected from: https://www.avatarsinpixels.com/
