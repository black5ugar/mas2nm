package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mas2nm/lib"
	"os/exec"
)

type Masscan struct {
	Bin    string `yaml:"bin"`  // masscan 二进制文件位置（指定这个是发现centos自带的masscan输出格式有点不同)
	Port   string `yaml:"port"` // 扫描端口
	IP     string // 要扫描的 IP 从数据库读取
	Input  string `yaml:"input"`
	Output string // 输出文件
	Rate   int    `yaml:"rate"` // 最大速率
}

// Masscan 扫描调用
func (m *Masscan) Scan(input string) error {
	m.Input = input
	m.Output = "mas-output.txt"
	cmd := exec.Command(m.Bin,
		"-p",
		m.Port,
		"-iL",
		m.Input,
		"-oJ",
		m.Output,
		"--max-rate",
		fmt.Sprintf("%d", m.Rate))
	log.Println(cmd.Args)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	return nil
}

// Masscan 扫描结果的Json格式
type MasscanResult struct {
	IP    string  `json:"ip"`
	Ports []Ports `json:"ports"`
}

type Ports struct {
	Port   int    `json:"port"`
	Status string `json:"status"`
}

// 初始化 Masscan
func NewMasscan() Masscan {
	var mas Masscan
	mas.Bin = lib.ReadConfig("masscan.bin").(string)
	mas.Port = lib.ReadConfig("masscan.port").(string)
	mas.Rate = int(lib.ReadConfig("masscan.rate").(float64))

	return mas
}

// 读取 Masscan 的扫描结果
// path 为 Masscan 扫描结果的目录
func GetMasscanResult() ([]MasscanResult, error) {

	var results []MasscanResult

	var result []MasscanResult
	content, err := ioutil.ReadFile("mas-output.txt")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	results = append(results, result...)
	return results, nil
}
