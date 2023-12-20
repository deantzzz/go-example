package test

import (
	"context"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"shorturl/rpc/transform/transformer"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestRpcSpeed(t *testing.T) {
	client := transformer.NewTransformer(zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"localhost:2379"},
			Key:   "transform.rpc",
		},
	}))

	loop := 20

	count := 30

	cDif := make(chan int64, count*loop)

	wg := &sync.WaitGroup{}
	for i := 0; i < loop; i++ {
		time.Sleep(time.Millisecond * 10)
		for i := 0; i < count; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				send := time.Now().UnixNano()
				//t.Log("send time: " + strconv.FormatInt(send, 10))
				resp, err := client.SpeedTest(context.Background(), &transformer.SpeedTestReq{
					T: send,
				})
				if err != nil {
					panic(err)
				}
				//t.Log("rec time: " + strconv.FormatInt(resp.T, 10))
				dif := resp.T - send
				cDif <- dif
				t.Log("dif: " + strconv.FormatInt(dif, 10) + "ns")
			}()
		}
	}

	wg.Wait()
	t.Log("close channel")
	close(cDif)
	var total int64
	for dif := range cDif {
		total += dif
	}
	avg := total / int64(count)

	t.Log("avg: " + strconv.FormatInt(avg, 10) + "ns")
}
