package main

import (
	"flag"
	"os/exec"

	log "github.com/Sirupsen/logrus"
	rds "github.com/garyburd/redigo/redis"
)

var (
	image    = flag.String("image", "", "Image")
	redisURL = flag.String("redis-url", ":6379", "Redis URL")

	redis rds.Conn
)

func init() {
	flag.Parse()

	if *image == "" {
		log.Fatal("dockmd2redis: flag '-image' is required.")
	}

	c, err := rds.Dial("tcp", *redisURL)
	if err != nil {
		log.Fatal(err)
	}
	redis = c

	_, err = redis.Do("PING")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	run(*image, "nodes", "docker-machine", "ls")
	run(*image, "containers", "list_all_containers")
}

func run(image string, prefix string, cmd ...string) {
	args := append([]string{"run", "--rm", image}, cmd...)
	out, err := exec.Command("docker", args...).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	key := prefix + ":" + image
	n, err := redis.Do("SET", key, out)
	if err != nil {
		log.Fatal(err)
	}
	log.WithField("key", key).Println(n)
}
