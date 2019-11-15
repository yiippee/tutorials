package main

import (
	"context"
	greeter "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-service/proto"
	"github.com/micro/go-micro"
)

type MyService struct{}

// 实现了proto中的接口
func (ms *MyService) GetInfo(ctx context.Context, empty *greeter.Empty, mi *greeter.MyInfo) error {
	// 传的是指针，需要
	*mi = greeter.MyInfo{
		Name: "lzb",
		Bo: &greeter.MyInfoBody{
			Length: 170,
			Weight: 65,
		},
		Color: greeter.Color_BLACK,
	}
	//mi.Name = "lzb"
	//mi.Bo = &greeter.MyInfoBody{
	//	Length:               170,
	//	Weight:               65,
	//}
	//mi.Color = greeter.Color_RED
	return nil
}
func main() {
	srv := micro.NewService(micro.Name("greeter.Service2"), micro.Version("latest"))
	srv.Init()

	greeter.RegisterMyServiceHandler(srv.Server(), new(MyService))

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
