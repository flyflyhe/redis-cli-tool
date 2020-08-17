package redisService

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strings"
)

var FuncMap = map[string]string{
	"GetAllKeys":"获取所有key",
	"MemoryUsage":"获取所有key使用的存储",
}

func GetHelperStr() string {
	strBuilder := strings.Builder{}
	for f, comment := range FuncMap {
		strBuilder.WriteString(f)
		strBuilder.WriteString(":")
		strBuilder.WriteString(comment)
		strBuilder.WriteString("\n")
	}

	return  strBuilder.String()
}

type RDB struct {
	conn *redis.Client
}


func GetRDB (host string, port int, password string, db int) RDB  {
	rClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	})

	return RDB{conn:rClient}
}

func (rdb RDB) MemoryUsage (cursor uint64) map[string]int64  {
	keys, cursor, err := rdb.conn.Scan(context.Background(), cursor, "", 10).Result()
	if err != nil {
		log.Fatal(err)
	}

	keyMemoryMap := make(map[string]int64)
	for _, key := range keys {
		b, err := rdb.conn.MemoryUsage(context.Background(), key).Result()
		if err != nil {
			log.Fatal(err)
		}
		keyMemoryMap[key] = b
	}

	if cursor != 0 {
		for k, v := range rdb.MemoryUsage(cursor) {
			keyMemoryMap[k] = v
		}
	}

	return keyMemoryMap
}

func (rdb RDB) GetAllKeys(cursor uint64) []string {
	keys, cursor, err := rdb.conn.Scan(context.Background(), cursor, "", 10).Result()
	log.Println(cursor)
	if err != nil {
		log.Fatal(err)
	}
	var keyList []string

	keyList = append(keyList, keys...)

	if cursor != 0 {
		keyList = append(keyList, rdb.GetAllKeys(cursor)...)
	}

	return  keyList
}