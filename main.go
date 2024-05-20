package main

import (
	"bufio"
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Config struct {
	HOST string `yaml:"HOST"`
	UID  string `yaml:"UID"`
	KEY  string `yaml:"KEY"`
}

func main() {
	// 读取config.yml文件
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("无法读取配置文件： %v", err)
	}

	// 解析YAML数据
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("无法解析配置文件： %v", err)
	}

	//fmt.Printf("UID: %s\n", cfg.UID)
	//fmt.Printf("KEY: %s\n", cfg.KEY)
	//fmt.Printf("HOST: %s\n", cfg.HOST)

	fmt.Printf("请选择你需要的操作编号：\n\t1.查询用户信息\n\t2.验证登录信息\n\t3.修改用户信息\n\t4.注册新用户\n\n请输入你需要的操作编号> ")
	reader := bufio.NewScanner(os.Stdin)
	if reader.Scan() {
		if reader.Text() == "1" {
			fmt.Printf("请输入你要查询的ID>")
			reader := bufio.NewScanner(os.Stdin)
			if reader.Scan() {
				var uid = reader.Text()
				data := url.Values{}
				data.Set("uid", cfg.UID)
				data.Set("api_key", cfg.KEY)
				data.Set("func", "GetUserInfo")
				data.Set("searchID", uid)
				post_requests(cfg.HOST, data)

			}
		}
		if reader.Text() == "2" {
			fmt.Printf("请输入你要验证的用户名>")
			reader := bufio.NewScanner(os.Stdin)
			if reader.Scan() {
				var uname = reader.Text()
				fmt.Printf("请输入你要验证的密码>")
				reader := bufio.NewScanner(os.Stdin)
				if reader.Scan() {
					var passwd = reader.Text()
					data := url.Values{}
					data.Set("uid", cfg.UID)
					data.Set("api_key", cfg.KEY)
					data.Set("func", "UserLogin")
					data.Set("LoginUsername", uname)
					data.Set("LoginPass", passwd)
					post_requests(cfg.HOST, data)
				}
			}
		}
		if reader.Text() == "3" {
			fmt.Printf("请输入需要修改的用户ID> ")
			reader := bufio.NewScanner(os.Stdin)
			if reader.Scan() {
				var UserID = reader.Text()
				// 'UserName', 'Email', 'Phone','sex','IdNumber','Name','Password','Provice','City','Leave','Unit','LoginTime'
				fmt.Printf("您需要修改？\n\t1.用户名\n\t2.邮箱\n\t3.手机号\n\t4.性别\n\t5.身份证号码\n\t6.姓名\n\t7.密码\n\t8.省份\n\t9.城市\n\t10.级别\n\t11.单位\n\t12.登录时间\n\n请输入你需要的操作编号> ")
				reader := bufio.NewScanner(os.Stdin)
				if reader.Scan() {
					if reader.Text() == "1" {
						fmt.Printf("请输入修改后的内容> ")
						reader := bufio.NewScanner(os.Stdin)
						if reader.Scan() {
							var value = reader.Text()
							data := url.Values{}
							data.Set("uid", cfg.UID)
							data.Set("api_key", cfg.KEY)
							data.Set("func", "Modify")
							data.Set("ModID", UserID)
							data.Set("ColName", "UserName")
							data.Set("Value", value)
							post_requests(cfg.HOST, data)
						}
					}
					if reader.Text() == "2" {
						fmt.Printf("请输入修改后的内容> ")
						reader := bufio.NewScanner(os.Stdin)
						if reader.Scan() {
							var value = reader.Text()
							data := url.Values{}
							data.Set("uid", cfg.UID)
							data.Set("api_key", cfg.KEY)
							data.Set("func", "Modify")
							data.Set("ModID", UserID)
							data.Set("ColName", "Email")
							data.Set("Value", value)
							post_requests(cfg.HOST, data)
						}
					}
					if reader.Text() == "3" {
						fmt.Printf("请输入修改后的内容> ")
						reader := bufio.NewScanner(os.Stdin)
						if reader.Scan() {
							var value = reader.Text()
							data := url.Values{}
							data.Set("uid", cfg.UID)
							data.Set("api_key", cfg.KEY)
							data.Set("func", "Modify")
							data.Set("ModID", UserID)
							data.Set("ColName", "Phone")
							data.Set("Value", value)
							post_requests(cfg.HOST, data)
						}
					}
					if reader.Text() == "4" {
						fmt.Printf("请输入修改后的内容> ")
						reader := bufio.NewScanner(os.Stdin)
						if reader.Scan() {
							var value = reader.Text()
							data := url.Values{}
							data.Set("uid", cfg.UID)
							data.Set("api_key", cfg.KEY)
							data.Set("func", "Modify")
							data.Set("ModID", UserID)
							data.Set("ColName", "sex")
							data.Set("Value", value)
							post_requests(cfg.HOST, data)
						}
					}
					if reader.Text() == "5" {
						fmt.Printf("请输入修改后的内容> ")
						reader := bufio.NewScanner(os.Stdin)
						if reader.Scan() {
							var value = reader.Text()
							data := url.Values{}
							data.Set("uid", cfg.UID)
							data.Set("api_key", cfg.KEY)
							data.Set("func", "Modify")
							data.Set("ModID", UserID)
							data.Set("ColName", "IdNumber")
							data.Set("Value", value)
							post_requests(cfg.HOST, data)
						}
					}
					if reader.Text() == "6" {
						fmt.Printf("请输入修改后的内容> ")
						reader := bufio.NewScanner(os.Stdin)
						if reader.Scan() {
							var value = reader.Text()
							data := url.Values{}
							data.Set("uid", cfg.UID)
							data.Set("api_key", cfg.KEY)
							data.Set("func", "Modify")
							data.Set("ModID", UserID)
							data.Set("ColName", "Name")
							data.Set("Value", value)
							post_requests(cfg.HOST, data)
						}
					}
					if reader.Text() == "7" {
						fmt.Printf("请输入修改后的内容> ")
						reader := bufio.NewScanner(os.Stdin)
						if reader.Scan() {
							var value = reader.Text()
							data := url.Values{}
							data.Set("uid", cfg.UID)
							data.Set("api_key", cfg.KEY)
							data.Set("func", "Modify")
							data.Set("ModID", UserID)
							data.Set("ColName", "Password")
							data.Set("Value", value)
							post_requests(cfg.HOST, data)
						}
					}
					if reader.Text() == "8" {
						fmt.Printf("请输入修改后的内容> ")
						reader := bufio.NewScanner(os.Stdin)
						if reader.Scan() {
							var value = reader.Text()
							data := url.Values{}
							data.Set("uid", cfg.UID)
							data.Set("api_key", cfg.KEY)
							data.Set("func", "Modify")
							data.Set("ModID", UserID)
							data.Set("ColName", "Provice")
							data.Set("Value", value)
							post_requests(cfg.HOST, data)
						}
					}
					if reader.Text() == "9" {
						fmt.Printf("请输入修改后的内容> ")
						reader := bufio.NewScanner(os.Stdin)
						if reader.Scan() {
							var value = reader.Text()
							data := url.Values{}
							data.Set("uid", cfg.UID)
							data.Set("api_key", cfg.KEY)
							data.Set("func", "Modify")
							data.Set("ModID", UserID)
							data.Set("ColName", "City")
							data.Set("Value", value)
							post_requests(cfg.HOST, data)
						}
					}
					if reader.Text() == "10" {
						fmt.Printf("请输入修改后的内容> ")
						reader := bufio.NewScanner(os.Stdin)
						if reader.Scan() {
							var value = reader.Text()
							data := url.Values{}
							data.Set("uid", cfg.UID)
							data.Set("api_key", cfg.KEY)
							data.Set("func", "Modify")
							data.Set("ModID", UserID)
							data.Set("ColName", "Leave")
							data.Set("Value", value)
							post_requests(cfg.HOST, data)
						}
					}
					if reader.Text() == "11" {
						fmt.Printf("请输入修改后的内容> ")
						reader := bufio.NewScanner(os.Stdin)
						if reader.Scan() {
							var value = reader.Text()
							data := url.Values{}
							data.Set("uid", cfg.UID)
							data.Set("api_key", cfg.KEY)
							data.Set("func", "Modify")
							data.Set("ModID", UserID)
							data.Set("ColName", "Unit")
							data.Set("Value", value)
							post_requests(cfg.HOST, data)
						}
					}
					if reader.Text() == "12" {
						fmt.Printf("请输入修改后的内容> ")
						reader := bufio.NewScanner(os.Stdin)
						if reader.Scan() {
							var value = reader.Text()
							data := url.Values{}
							data.Set("uid", cfg.UID)
							data.Set("api_key", cfg.KEY)
							data.Set("func", "Modify")
							data.Set("ModID", UserID)
							data.Set("ColName", "LoginTime")
							data.Set("Value", value)
							post_requests(cfg.HOST, data)
						}
					}

				}
			}

		}
		if reader.Text() == "4" {
			fmt.Printf("请输入新注册的用户名>")
			reader := bufio.NewScanner(os.Stdin)
			if reader.Scan() {
				var uname = reader.Text()
				fmt.Printf("请输入新注册的密码>")
				reader := bufio.NewScanner(os.Stdin)
				if reader.Scan() {
					var passwd = reader.Text()
					data := url.Values{}
					data.Set("uid", cfg.UID)
					data.Set("api_key", cfg.KEY)
					data.Set("func", "UserRegister")
					data.Set("RegUsername", uname)
					data.Set("RegPass", passwd)
					post_requests(cfg.HOST, data)
				}
			}

		}
	}
}

func post_requests(url string, data url.Values) {
	postData := data.Encode()
	// 设置请求的URL
	reqURL := url + "api/api.php" // 替换为你的服务器地址
	// 创建HTTP客户端
	client := &http.Client{}
	// 创建请求
	req, err := http.NewRequest("POST", reqURL, bytes.NewBufferString(postData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	// 打印响应内容
	fmt.Println("Response:", string(body))
}
