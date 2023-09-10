package Client

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type License struct {
	Username   string
	Key        string
	Experation string
	ip         string
	webport    string
	port       string
	serverip   string
	Version    string
	serverport string
}

var Lcs License
var LcsFileName string
var ServerIP string

func decodeLicense(license string, host string, port string) {
	decoded, err := base64.StdEncoding.DecodeString(license)
	decryptedText := string(decoded[:len(decoded)-len("BetterThanNothingOfCourseOrIsIt")])
	if err != nil {
		panic(err)
	}
	stripped := strings.Split(decryptedText, "\r\n")
	Lcs.Username = stripped[0]
	Lcs.Key = stripped[1]
	Lcs.ip = host
	Lcs.port = port
	Lcs.Experation = stripped[2]
	Lcs.port = stripped[4]
	Lcs.webport = stripped[5]
	Lcs.serverip = stripped[6]
	return
}

func Load(host string, port string) {
	ServerIP = GetServerIP()
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	for _, f := range files {
		stripped := strings.Split(f.Name(), ".")
		if len(stripped) != 2 {
			continue
		}
		if stripped[1] == "key" {
			LcsFileName = f.Name()
			_f, err := ioutil.ReadFile(f.Name())
			if err != nil {
				panic(err)
			}
			decodeLicense(string(_f), host, port)
			validate(true, host, port)
			go func() {
				for {
					time.Sleep(25 * time.Minute)
					validate(false, host, port)
				}
			}()
			return
		}
	}
	fmt.Println("No License File Found")
	os.Exit(0)
}

func GetServerIP() string {
	resp, err := http.Get("https://ipinfo.io/ip")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func validate(first bool, host string, port string) {
	serverAddr, err := net.ResolveUDPAddr("udp", host+":"+port)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	err = conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer conn.Close()
	if first {
		_, err = conn.Write([]byte(Lcs.Key + ";" + ServerIP))
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	} else {
		_, err = conn.Write([]byte(Lcs.Key + ";" + ServerIP + ";" + "0"))
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
	buf := make([]byte, 99999)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[0:n]))
	if string(buf[0:n]) != "[Key Was Vaild] [INFO] [VAILD]" {
		fmt.Println("License Key is invalid, please enter a valid one in the config file")
		os.Exit(0)
	}
}
