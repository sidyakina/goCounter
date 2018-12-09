package main

import "sync"

type Tasker struct {
	packet taskPacket
	wg     MyWaitGroup
	total  int
	mux    sync.Mutex
}

type taskPacket struct {
	tasks []string
	mux   sync.Mutex
}

func InitTasker(maxRoutines int) *Tasker {
	return &Tasker{packet: taskPacket{tasks: make([]string, 0, maxRoutines)},
		wg: MyWaitGroup{maxRoutines: maxRoutines},
	}
}

func (t *Tasker) execute() {
	if !t.wg.CanNext() {
		return
	}
	if newTask, valid := t.packet.get(); valid {
		err := t.wg.Next()
		if err != nil {
			println("error!!!", err)
			t.Add(newTask)
			return
		}
		go func() {
			goCount := Count(newTask)
			t.mux.Lock()
			t.total += goCount
			t.mux.Unlock()
			t.wg.Done()
			t.execute()
		}()
	}

}

func (t *Tasker) Close() int {
	t.wg.Wait()
	return t.total
}

func (t *Tasker) Add(newTask string) {
	t.packet.add(newTask)
	t.mux.Lock()
	defer t.mux.Unlock()
	t.execute()
}

func (p *taskPacket) get() (string, bool) {
	p.mux.Lock()
	defer p.mux.Unlock()
	numberTasks := len(p.tasks)
	if numberTasks > 0 {
		newTask := p.tasks[numberTasks-1]
		p.tasks = p.tasks[:numberTasks-1]
		return newTask, true
	}
	return "", false
}

func (p *taskPacket) add(newTask string) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.tasks = append(p.tasks, newTask)
}

func (p *taskPacket) existTask() bool {
	p.mux.Lock()
	defer p.mux.Unlock()
	println("exist", len(p.tasks))
	return len(p.tasks) > 0
}
