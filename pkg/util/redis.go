package util

import (
	"context"
	"fmt"
	"log"
	redisMdm "yqsl.com/scm/data-task/pkg/redis/mdm"
)

func CleanFacadeMdmRedisKey(env, mathStr string) (err error) {

	ctx := context.Background()

	if env == "dev" || env == "" {
		env = "null"
	}

	if env != "prod" && env != "null" && env != "uat" {
		panic(fmt.Errorf("env must be prod or dev or uat"))
	}

	fmt.Println("CleanFacadeMdmRedisKey  mathStr : ", fmt.Sprintf("%s:MdmRedisPre:%s", env, mathStr))

	keys, err := redisMdm.Client.Keys(ctx, fmt.Sprintf("%s:MdmRedisPre:%s", env, mathStr)).Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(keys) == 0 {
		log.Println(fmt.Sprintf("%s:MdmRedisPre:%s", env, mathStr), "fetch redis keys is valid, 0")
		return
	}

	var result int64
	result, err = redisMdm.Client.Del(ctx, keys...).Result()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if result == 0 {
		fmt.Printf("delete key: %s, result: %s", keys, "error")
	} else {
		fmt.Printf("delete key: %s, result: %s, delete num is %d", keys, "successfully", result)
	}

	return
}
