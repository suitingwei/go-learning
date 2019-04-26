package main

import "./taskManager"

func main(){
	taskManager:=taskManager.New()

	taskManager.Add("key1","value1",10)
}

