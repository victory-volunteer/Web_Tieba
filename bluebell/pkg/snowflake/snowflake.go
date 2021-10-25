//方式1：snowflake（常用）
package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	//初始化全局node
	//startTime为开始时间因子，machineID为机器id
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}
func GenID() int64 {
	//生成id并转换为int64
	return node.Generate().Int64()
}

//func main() {
//	if err := Init("2020-07-01", 1); err != nil {
//		//在本程序中是单独一个模块而不是分布式的，所以machineID传1
//		fmt.Printf("init failed, err:%v\n", err)
//		return
//	}
//	id := GenID()
//	fmt.Println(id)
//}

//方式2：sonyflake
//package main
//
//import (
//	"fmt"
//	"time"
//
//	"github.com/sony/sonyflake"
//)
//
//var (
//	sonyFlake     *sonyflake.Sonyflake //同上方的node实例
//	sonyMachineID uint16
//)
//
//func getMachineID() (uint16, error) {
//	return sonyMachineID, nil
//}
//
//// 需传⼊当前的机器ID
//func Init(startTime string, machineId uint16) (err error) {
//	sonyMachineID = machineId
//	var st time.Time
//	st, err = time.Parse("2006-01-02", startTime)
//	if err != nil {
//		return err
//	}
//	settings := sonyflake.Settings{
//		StartTime: st,
//		MachineID: getMachineID,
//	}
//	sonyFlake = sonyflake.NewSonyflake(settings)
//	return
//}
//
//// GenID ⽣成id
//func GenID() (id uint64, err error) {
//	if sonyFlake == nil {
//		err = fmt.Errorf("snoy flake not inited")
//		return
//	}
//	id, err = sonyFlake.NextID()
//	return
//}
//func main() {
//	if err := Init("2020-07-01", 1); err != nil {
//		fmt.Printf("Init failed, err:%v\n", err)
//		return
//	}
//	id, _ := GenID()
//	fmt.Println(id)
//}
