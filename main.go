package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"redisTool/service/logService"
	"redisTool/service/redisService"
	"strings"
)

func main()  {
	host := flag.String("h", "127.0.0.1", "host default 127.0.0.1")
	port := flag.Int("P", 6379, "port default 6379")
	password := flag.String("p", "", "password")
	db := flag.Int("db", 0, "database")
	c := flag.String("c", "", redisService.GetHelperStr())

	flag.Parse()

	if *c == "" {
		fmt.Print(redisService.GetHelperStr())
		os.Exit(0)
	}

	if _, ok := redisService.FuncMap[*c]; !ok {
		log.Fatal(*c + "未实现")
	}

	rdb := redisService.GetRDB(*host, *port, *password, *db)

	switch *c {
	case "GetAllKeys":
		logService.PrintStrList(rdb.GetAllKeys(0))
	case "Info":
		info := rdb.Info()
		logService.PrintStrList(strings.Split(info, "\\r\\n"))
	case "MemoryUsage":
		logService.PrintStrIntMap(rdb.MemoryUsage())
	}
}
