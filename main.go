package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 设置一个定时器，每隔5秒触发一次，用于定期轮询
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	i := 0
	for range ticker.C {
		// 使用http.Get发送GET请求到酶标仪服务的get_measurement接口
		resp, err := http.Get("http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement")
		if err != nil {
			// 检查请求是否成功，如果失败则输出错误并继续下一次循环
			fmt.Println("Error:", err)
			continue
		}
		defer resp.Body.Close()

		// 使用io.ReadAll读取响应体的全部内容
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}
		// 这里测试直接输出body是一个字节切片，将其转换为string类型得到具体的数据，然后将数据的内容按行分割，得到一个string类型的切片rows
		rows := strings.Split(string(body), "\n")
		// 定义一个data切片接收得到的所有值
		var data [][]float64
		for _, row := range rows {
			// 将每行数据按空白字符分割成一个字符串切片values。每个value代表一个数据点。
			values := strings.Fields(row)
			var rowValues []float64
			for _, value := range values {
				// 使用strconv.ParseFloat函数尝试将字符串value转换为64位浮点数num
				num, err := strconv.ParseFloat(value, 64)
				if err != nil {
					fmt.Println("Error converting value to float64:", err)
					continue
				}
				rowValues = append(rowValues, num)
			}
			data = append(data, rowValues)
		}

		var sum float64
		for _, row := range data {
			for _, value := range row {
				sum += value
			}
		}

		average := sum / float64(len(data)*len(data[0]))
		fmt.Printf("第 %d 次获取到的平均值: %f\n", i, average)
		i++
	}
}
