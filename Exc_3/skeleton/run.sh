#!/bin/sh
docker compose up #if you wanna stop it docker compose down
# todo
# docker build
# docker run db
# docker run orderservice


# docker image pull postgres:18

# docker volume create SBD3

# docker run -ti -name db --env-file debug.env -p 5432:5432 -v SBD3:/var/lib/postgresql/18/docker postgres:18

# docker build -t orderservice .

# docker run -d --name orderservice --env-file debug.env -p 3000:3000 --link db:db orderservice