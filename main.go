// Code generated by hertz generator.

package main

import "github.com/cloudwego/hertz/pkg/app/server"

func main() {
	h := server.Default(
		server.WithHostPorts("127.0.0.1:6789"),
	)
	register(h)
	h.Spin()
}
