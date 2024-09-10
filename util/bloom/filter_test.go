package bloom

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func TestBloom(t *testing.T) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.3.58:16379",
	})
	f, err := NewFilter(ctx, client, "test_bloom_test", 1000, 0.001)
	if err != nil {
		t.Fatal(err)
	}
	err = f.Clear(ctx)
	if err != nil {
		t.Fatal(err)
	}
	addStart := time.Now()
	for i := 0; i < 1000; i++ {
		err := f.Add(ctx, []byte(fmt.Sprintf("%d", i)))
		if err != nil {
			t.Fatal(err)
		}
	}
	existStart := time.Now()
	fmt.Println("add time:", existStart.Sub(addStart))
	for i := 0; i < 1000; i++ {
		exist, err := f.Exist(ctx, []byte(fmt.Sprintf("%d", i)))
		if err != nil {
			t.Fatal(err)
		}
		if !exist {
			t.Fatal("wring")
		}
	}
	cnt := 0
	for i := 1000; i < 2000; i++ {
		exist, err := f.Exist(ctx, []byte(fmt.Sprintf("%d", i)))
		if err != nil {
			t.Fatal(err)
		}
		if exist {
			cnt++
		}
	}
	fmt.Println(cnt)
	fmt.Println(time.Since(existStart))
}

func TestBatchAdd(t *testing.T) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.3.58:16379",
	})
	f, err := NewFilter(ctx, client, "test_bloom_test", 1000, 0.001)
	if err != nil {
		t.Fatal(err)
	}
	err = f.Clear(ctx)
	if err != nil {
		t.Fatal(err)
	}
	addStart := time.Now()
	dataList := make([][]byte, 0, 1000)
	for i := 0; i < 1000; i++ {
		dataList = append(dataList, []byte(fmt.Sprintf("%d", i)))
	}
	err = f.BatchAdd(ctx, dataList...)
	if err != nil {
		t.Fatal(err)
	}
	existStart := time.Now()
	fmt.Println("add time:", existStart.Sub(addStart))
	for i := 0; i < 1000; i++ {
		exist, err := f.Exist(ctx, []byte(fmt.Sprintf("%d", i)))
		if err != nil {
			t.Fatal(err)
		}
		if !exist {
			t.Fatal("wrong")
		}
	}
	cnt := 0
	for i := 1000; i < 2000; i++ {
		exist, err := f.Exist(ctx, []byte(fmt.Sprintf("%d", i)))
		if err != nil {
			t.Fatal(err)
		}
		if exist {
			cnt++
		}
	}
	fmt.Println(cnt)
	fmt.Println(time.Since(existStart))
}
