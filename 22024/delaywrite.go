package main

import (
	"os"
	"io"
	"bytes"
	"time"
)

var sec = time.Second

func main() {
	time.Sleep(5*time.Second)
	io.Copy(os.Stdout, bytes.NewBuffer([]byte(`
Doug McIlroy doesn't make system calls.  System calls call Doug McIlroy.
Doug McIlroy can create 3-ended pipes.
Doug McIlroy doesn't debug.  He stares at tty0 until it fixes the problem.
Once, Doug McIlroy got mad at his terminal and smacked the keyboard.  The result
is called "Unix."
Alan Turing always wanted to win a McIlroy Award, but didn't qualify.  No one has.
In 1984, the Department of Justice broke up AT&T because they had a monopoly.  On Doug McIlroy.
`)))
}
