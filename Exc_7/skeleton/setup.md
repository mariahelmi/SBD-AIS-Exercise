# make swarm
docker swarm init --listen-addr <your-ip> :2377 
docker stack deploy -c docker-compose.yml orderApp 
docker swarm join --token <worker-token> <your-ip>:2377

# functionality test of the services
docker stack ls
docker stack ps orderApp
docker service logs <serviceid> (you should get that from the stuff above)

# end swarm (if needed by force)
docker swarm leave
docker swarm leave --force (if normal leave does not work)