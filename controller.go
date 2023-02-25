package main

import (
	"context"
	"coredemo/framework"
	"fmt"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) err {
	finish := make(chan struct{}, 1)
	panicChain := make(chan struct{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChain <- p
			}
		}()
		time.Sleep(10 * time.Second)
		c.Json(100, "ok")
		finish <- struct{}{}
	}()

	select {
	case p := <-panicChain:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic occured")
	case <-finish:
		fmt.Println("finish ok")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out")
		c.SetHasTimeout()
	}
	return nil
}
