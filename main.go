package main

import (
	"flag"
	"log"
	"redisTool/service/redisService"
)

func main()  {
	host := flag.String("h", "127.0.0.1", "host default 127.0.0.1")
	port := flag.Int("P", 6379, "port default 6379")
	password := flag.String("p", "", "password")
	db := flag.Int("db", 0, "database")
	c := flag.String("c", "", redisService.GetHelperStr())

	flag.Parse()

	rdb := redisService.GetRDB(*host, *port, *password, *db)

	if *c == "GetAllKeys" {
		keys := rdb.GetAllKeys(0)

		for _, k := range keys {
			log.Println(k)
		}
	}
}
