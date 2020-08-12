Running NGINX Open Source in a Docker Container

```
$ docker run --name mynginx1 -p 80:80 -d nginx
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS                NAMES
2274a1aa2f4f        nginx               "/docker-entrypoint.…"   About a minute ago   Up About a minute   0.0.0.0:80->80/tcp   mynginx1
```
- https://docs.nginx.com/nginx/admin-guide/installing-nginx/installing-nginx-docker/

