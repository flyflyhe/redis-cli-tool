package logService

import "log"

func PrintStrList(list []string)  {
	for _, k := range list {
		log.Println(k)
	}
}

func PrintStrIntMap(strMap map[string]int64)  {
	for k, v := range strMap {
		log.Println(k, ":", v)
	}
}
