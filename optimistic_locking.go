package main

import (
	"time"
	"go.etcd.io/etcd/clientv3"
		"fmt"
	"context"
	"log"
)

func main() {
	var (
		kv clientv3.KV
		leaseResp *clientv3.LeaseGrantResponse
		leaseid clientv3.LeaseID
		keepResp  *clientv3.LeaseKeepAliveResponse
		keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
		ctx context.Context
		cancelFunc context.CancelFunc
		err error
		txn clientv3.Txn
		txnResp *clientv3.TxnResponse
	)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.246.132:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect wrong:", err)
	}
	defer cli.Close()
	// lease实现锁自动过期
	//op操作
	//txn事务:if else then
	//1.上锁(创建租约，自动续租，拿着租约抢占一个key)
	//申请一个lease租约
	lease := clientv3.NewLease(cli)
	//准备一个用于
	//申请5s的租约
	if leaseResp,err=lease.Grant(context.TODO(),5);err !=nil{
		log.Println(err)
		return
	}
	//拿到租约id
	leaseid = leaseResp.ID
	//准备一个用于取消自动续租的context
	ctx,cancelFunc = context.WithCancel(context.TODO())
	//确保函数退出后，自动续租会停止
	defer cancelFunc()
	//释放租约
	defer lease.Revoke(context.TODO(),leaseid)
	//五秒后会取消自动续租
	//ctx,_:=context.WithTimeout(context.TODO(),5*time.Second)
	//自动续租
	if keepRespChan,err =lease.KeepAlive(ctx,leaseid);err !=nil{
		log.Println(err)
		return
	}
	go func() {
		for  {
			select {
			case keepResp = <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已经失效")
					goto END
				}else{//每秒续租一次，所以就会受到一次应答
					fmt.Println("收到自动续租应答")
				}
			}
		END:
		}
	}()
	//if 不存在key then设置它，else抢锁失败
	//获得kv api子集
	kv = clientv3.NewKV(cli)
	txn = kv.Txn(context.TODO())
	//定义事物
	//如果key不存在
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job9"),"=",0)).Then(clientv3.OpPut("/cron/lock/job9","xxx",clientv3.WithLease(leaseid))).Else(clientv3.OpGet("/cron/lock/job9"))//否则抢锁失败
	//提交事务
	if txnResp,err = txn.Commit();err != nil{
		fmt.Println(err)
		return
	}
	//判断是否抢到了锁
	if !txnResp.Succeeded{
		fmt.Println("锁被占用",string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
	}
	//2.处理业务
	fmt.Println("处理任务")
	time.Sleep(50*time.Second)
	//在锁内很安全

	//释放锁(取消自动续租，释放租约)
	//defer 会把租约释放掉，关联的kv就被删除了

}
