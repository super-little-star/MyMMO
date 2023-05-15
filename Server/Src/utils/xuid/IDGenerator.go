package xuid

import (
	"errors"
	"sync"
	"time"
)

// |时间戳|
const (
	workerBits  uint8 = 10                      //机器码位数
	numberBits  uint8 = 12                      //序列号位数
	workerMax   int64 = -1 ^ (-1 << workerBits) //机器码最大值（即1023）
	numberMax   int64 = -1 ^ (-1 << numberBits) //序列号最大值（即4095）
	timeShift   uint8 = workerBits + numberBits //时间戳偏移量
	workerShift uint8 = numberBits              //机器码偏移量
	epoch       int64 = 1656856144              //起始常量时间戳（秒）

)

type IGenerator interface {
	NextId() int64
}

type GGenerator struct {
	mu        sync.Mutex
	timeStamp int64
	workerId  int64
	number    int64
}

var Generator IGenerator

func Init(workerId int64) error {
	if workerId < 0 || workerId > workerMax {
		return errors.New("workerId is overflow")
	}
	Generator = &GGenerator{
		timeStamp: 0,
		workerId:  workerId,
		number:    0,
	}
	return nil
}

func (w *GGenerator) NextId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	//当前时间的毫秒时间戳
	now := time.Now().Unix()
	//如果时间戳与当前时间相同，则增加序列号
	if w.timeStamp == now {
		w.number++
		//如果序列号超过了最大值，则更新时间戳
		if w.number > numberMax {
			for now <= w.timeStamp {
				now = time.Now().Unix()
			}
		}
	} else { //如果时间戳与当前时间不同，则直接更新时间戳
		w.number = 0
		w.timeStamp = now
	}
	//ID由时间戳、机器编码、序列号组成
	ID := (now-epoch)<<timeShift | (w.workerId << workerShift) | (w.number)
	return ID
}
