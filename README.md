# dockmd2redis

Execute commands or script in Docker and save results in Redis.

```
run(*image, "nodes", "docker-machine", "ls")
run(*image, "containers", "list_all_containers")
```

```
> ./dockmd2redis -image krkr/c2
INFO[0000] OK                                            key=nodes:krkr/c2
INFO[0001] OK                                            key=containers:krkr/c2
```

```
> d exec -ti banana_redis_1 redis-cli get nodes:krkr/c2
"NAME   ACTIVE   DRIVER   STATE     URL                          SWARM   DOCKER        ERRORS\nn1     *        ovh      Running   tcp://149.202.176.196:2376           v1.12.0-rc3   \n"

> d exec -ti banana_redis_1 redis-cli get containers:krkr/c2 | sed -e 's|\\n||g'  -e 's|\\||g' -e 's|^"||' -e 's|"$||' | jq '.[].Names'
[
  "/iris.1.6e67srh31ifb3vrmnheat8ezz"
]
[
  "/apish.1.3rjklddxirxqahw90pfreto64"
]
[
  "/o_metrics"
]

```