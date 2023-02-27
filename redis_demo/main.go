package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	_, err = rdb.Ping().Result()
	return err
}

func redisExample() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed: %v\n", err)
		return
	}
	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed,err; %v\n", err)
		return
	}
	fmt.Println("score", val)
	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed .err: %v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}

}
func hgetDemo() {
	result, err := rdb.HGetAll("user").Result()
	if err != nil {
		//redis.Nil(err)
		//其他
		fmt.Printf("hgetall failed err :%v\n", err)
	}
	fmt.Println(result)
	v2 := rdb.HMGet("user", "name", "age").Val()
	fmt.Println(v2)
	v3 := rdb.HGet("user", "age").Val()
	fmt.Println(v3)
}

// doCommand go-redis基本使用示例
func doCommand() {
	//ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	//defer cancel()

	// 执行命令获取结果
	val, err := rdb.Get("key").Result()
	fmt.Println(val, err)

	// 先获取到命令对象
	cmder := rdb.Get("key")
	fmt.Println(cmder.Val()) // 获取值
	fmt.Println(cmder.Err()) // 获取错误

	// 直接执行命令获取错误
	err = rdb.Set("key", 10, time.Hour).Err()

	// 直接执行命令获取值
	value := rdb.Get("key").Val()
	fmt.Println(value)
}

// zsetDemo 操作zset示例
func zsetDemo() {
	// key
	zsetKey := "language_rank"
	// value
	languages := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	//defer cancel()

	// ZADD
	err := rdb.ZAdd(zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, *op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

// watchDemo 在key值不变的情况下将其值+1
func watchDemo(key string) error {
	return rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		// 假设操作耗时5秒
		// 5秒内我们通过其他的客户端修改key，当前事务就会失败
		time.Sleep(5 * time.Second)
		_, err = tx.TxPipelined(func(pipe redis.Pipeliner) error {
			//业务逻辑
			pipe.Set(key, n+1, time.Hour)
			return nil
		})
		return err
	}, key)
}
func main() {
	if err := initClient(); err != nil {
		fmt.Printf("init redis client failed: %v\n", err)
		return
	}
	fmt.Println("connnect redis success...")
	defer rdb.Close()
	//redisExample()
	//hgetDemo()
	zsetDemo()
}
