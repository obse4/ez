package initial

import (
	"ez/global"
	"ez/pkg/database"
	"fmt"
	"log"
	"sync"
	"time"
)

type Task func()

type WorkerPool struct {
	maxWorkers int
	tasks      chan Task
	wg         sync.WaitGroup
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
	return &WorkerPool{
		maxWorkers: maxWorkers,
		tasks:      make(chan Task),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.maxWorkers; i++ {
		go func() {
			for task := range wp.tasks {
				task()
				wp.wg.Done()
			}
		}()
	}
}

func (wp *WorkerPool) AddTask(task Task) {
	wp.wg.Add(1)
	wp.tasks <- task
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}
func Run() {
	// 最先初始化全局配置
	InitConfig()

	// 初始化jwt
	InitJwt()

	// 初始化mysql
	InitMysql()

	// 初始化redis
	InitRedis()

	test()

	// d, _ := service.NewUserService().ApiV1Login(request.UserApiV1LoginRequest{
	// 	Username: "testedit",
	// 	Password: "123",
	// })
	// log.Printf("%#v\n", d)

	// var pData = common.UserJwt{}
	// global.Config.Jwt.Jwt.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InRlc3RlZGl0IiwiVXNlcklkIjoxMSwiVXNlclR5cGUiOjIsImV4cCI6MTY4OTg1MzEzN30.fXqPjon_4htctkCtYVla2YVFcCS01mFQktE6N7jOQys", &pData)

	// log.Printf("TRANS DATA%#v\n", pData)
	// 最后初始化http服务
	InitServer()
}

func test() {
	pool := NewWorkerPool(5)
	pool.Start()

	since := time.Now()
	rd := database.GetRedisConn(global.Config.RedisDb.Pool, 0)
	defer rd.Close()
	for i := 0; i < 10000000; i++ {
		pool.AddTask(func() {
			n := i + 563716
			var user_id = fmt.Sprintf("user_%d", n)

			rd.Do("SET", user_id, time.Now().Unix())
			rd.Do("EXPIRE", user_id, 60*60*24*7)
		})
	}
	pool.Wait()
	log.Panic(time.Since(since).Seconds())
}
