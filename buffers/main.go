package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Buffer struct {
	sync.RWMutex
	b bytes.Buffer
}

func (b Buffer) Read(p []byte) (n int, err error) {
	b.Lock()
	defer b.Unlock()
	return b.b.Read(p)
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.Lock()
	defer b.Unlock()
	return b.b.Write(p)
}

func (b *Buffer) String() string {
	b.Lock()
	defer b.Unlock()
	return b.b.String()
}

func (b *Buffer) Bytes() []byte {
	b.Lock()
	defer b.Unlock()
	return b.b.Bytes()
}

func (b *Buffer) Len() int {
	b.Lock()
	defer b.Unlock()
	return b.b.Len()
}

func (b *Buffer) WriteString(s string) (n int, err error) {
	b.Lock()
	defer b.Unlock()
	return b.b.WriteString(s)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

type counter struct {
	sync.RWMutex
	i int
}

func (c *counter) Incr() int {
	c.Lock()
	defer c.Unlock()
	c.i++
	return c.i
}

func main() {

	// create global buffer
	buf := Buffer{}
	counter := counter{}

	// create goroutines writing to buffer

	for i := 0; i < 1; i++ {
		go func(i int) {
			for {
				// t := time.Now()
				// entropy := rand.New(rand.NewSource(t.UnixNano()))
				// msg := fmt.Sprintf("%d %d %s\n", i, count, ulid.MustNew(ulid.Timestamp(t), entropy))
				// buf.WriteString(msg)
				buf.WriteString(fmt.Sprintln(counter.Incr()))
				time.Sleep(time.Second)
			}
		}(i)
	}

	// create goroutines reading from buffer
	for i := 0; i < 1; i++ {
		go func(i int) {
			for {
				fmt.Println(i, "================")
				// fmt.Println(string(buf.Next(100)))
				// b, err := buf.ReadBytes('\n')
				// if err != nil {
				// 	// log.Println(err)
				// }
				// fmt.Println(string(b), buf.Len()) // fmt.Println(string(buf.string()))
				b := buf.Bytes()
				buf.Read(b)
				fmt.Println(string(b), buf.Len())
				io.Copy(os.Stdout, buf)
				time.Sleep(time.Second)
			}
		}(i)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

}
