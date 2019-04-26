package taskManager

type TaskManager struct {
	container map[string]string //key -> value
}

//Add the key value pair into the taskManager's container. And set the ttl.
func (manager *TaskManager) Add(key string, value string, timeout int ) {
	_,exists := manager.container[key]

	manager.container[key] = value
}

//Create the task manager
func New() TaskManager {
	return TaskManager{}
}
