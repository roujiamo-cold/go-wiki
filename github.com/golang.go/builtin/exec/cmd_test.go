package exec

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"sync"
	"testing"
)

// 简单执行
func Test1(t *testing.T) {
	cmd := exec.Command("./testcmd.exe", "-s")

	// 使用CombinedOutput 将stdout stderr合并输出
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("test1 failed %s\n", err)
	}
	log.Println("test1 output ", string(out))
}

// 	stdout & stderr 分开输出
func Test2(t *testing.T) {
	cmd := exec.Command("./testcmd.exe", "-s", "-e")
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := stderr.Read(buf)

			if n > 0 {
				log.Printf("read err %s", string(buf[:n]))
			}

			if n == 0 {
				break
			}

			if err != nil {
				log.Printf("read err %v", err)
				return
			}
		}
	}()

	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := stdout.Read(buf)

			if n == 0 {
				break
			}

			if n > 0 {
				log.Printf("read out %s", string(buf[:n]))

			}

			if n == 0 {
				break
			}

			if err != nil {
				log.Printf("read out %v", err)
				return
			}

		}
	}()

	err := cmd.Wait()
	if err != nil {
		log.Printf("cmd wait %v", err)
		return
	}
}

// 按行读输出的内容
func Test3(t *testing.T) {
	cmd := exec.Command("./testcmd.exe", "-s", "-e")
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	oReader := bufio.NewReader(stdout)
	eReader := bufio.NewReader(stderr)

	cmd.Start()

	go func() {
		for {
			line, err := oReader.ReadString('\n')

			if line != "" {
				log.Printf("read line %s", line)
			}

			if err != nil || line == "" {
				log.Printf("read line err %v", err)
				return
			}

		}
	}()

	go func() {
		for {
			line, err := eReader.ReadString('\n')

			if line != "" {
				log.Printf("read err %s", line)
			}

			if err != nil || line == "" {
				log.Printf("read err %v", err)
				return
			}

		}
	}()

	err := cmd.Wait()
	if err != nil {
		log.Printf("cmd wait %v", err)
		return
	}
}

// 持续输入
func Test5(t *testing.T) {
	cmd := exec.Command("openssl")

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	stdin, _ := cmd.StdinPipe()

	cmd.Start()

	// 读
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		for {
			buf := make([]byte, 1024)
			n, err := stderr.Read(buf)

			if n > 0 {
				fmt.Println(string(buf[:n]))
			}

			if n == 0 {
				break
			}

			if err != nil {
				log.Printf("read err %v", err)
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			buf := make([]byte, 1024)
			n, err := stdout.Read(buf)

			if n == 0 {
				break
			}

			if n > 0 {
				fmt.Println(string(buf[:n]))
			}

			if n == 0 {
				break
			}

			if err != nil {
				log.Printf("read out %v", err)
				return
			}

		}
	}()

	// 写
	go func() {
		stdin.Write([]byte("version\n\n"))
		stdin.Write([]byte("ciphers -v\n\n"))
		//stdin.Write([]byte("s_client -connect razeencheng.com:443"))
		stdin.Close()
		wg.Done()
	}()

	wg.Wait()
	err := cmd.Wait()
	if err != nil {
		log.Printf("cmd wait %v", err)
		return
	}

}

//
// https://razeencheng.com/post/simple-use-go-exec-command.html
