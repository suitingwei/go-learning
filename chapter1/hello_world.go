package main

import (
	"fmt"
	"my-go-learnings/chapter1/structs"
)

func init() {
	fmt.Println("The init function has been triggered.")
}

func main() {

	////调用了内部其他包的方法、结构体。注意这里只有首字母大写的才是public 的，而且和 private 的方法还是不一样的
	////可以同时有 public 和 private 两个名字一样的方法
	//father := structs.User{"John", "John@email.com"}
	//
	//structs.LearnNestingStruct()
	//
	//father.Notify()

	//father2:=structs.PublicFather{
	//	Age:10,
	//}
	//father2.user.Name = "hello"
	//father2.user.Email ="anosiendo"
	//
	//fmt.Printf("%v\n",father2)

	//goroutine.LearnGoRoutine()
	//goroutine.LearnRoutine()

	//普通的长时间的 routine，会出现输出是交叉的
	//goroutine.LearnLongRoutines()

	//使用锁来解决并发
	//goroutine.LearnRareRoutinesWithLock()

	//使用临界区来解决并发
	//goroutine.LearnRareRoutinesWithMutex()

	//使用通道传递数据
	//channel.LearnRoutineWithChannel()
	//channel.LearnChannelWithRaceGame()

	//学习数据库操作
	//database.LearnDatabase()

	//log.LearnLog()

	//funcs.LearnTimer()

	//funcs.LearnTicker()

	//simpleEcho:= func(str string){
	//	fmt.Println("Simply echo the word:"+str)
	//}
	//newEcho := structs.FuncWrapper(simpleEcho)
	//
	//newEcho("I am the king of the night!")

	//_defer.LearnBasic()

	//http.Run()

	//sync.LearnSync()

	//my_go_learnings.LearnReflection()
	//panic.LearnPanic()
	//structs.LearnRawString()
	structs.LearnByteAndRune()
}

func printArray(arr *[5]int) {
	arrLen := len(*arr)

	for i := 0; i < arrLen; i++ {
		fmt.Print(" ", (*arr)[i])
	}
}
