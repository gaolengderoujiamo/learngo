package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var cwd string

func main() {
	//s := "httsp://asdsadsadsad"
	//fmt.Println(VerifyURL(s))
	c := make(chan int)

	go func() {
		if value, ok := <-c; ok {
			fmt.Printf("#### %v\n", value)
		} else {
			fmt.Printf("channel closed[%d]", value)
		}
	}()

	time.Sleep(3 * time.Second)

	c <- 0

	time.Sleep(10 * time.Second)

}

func VerifyURL(url string) string {
	var buf strings.Builder
	if strings.Index(url, "http") != 0 &&
		strings.Index(url, "https") != 0 {
		buf.WriteString("http://")
	}
	buf.WriteString(url)
	return buf.String()
}

func d1() {
	nums := [3]int{}
	nums[0] = 1
	n := nums[0]
	n = 2
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("n: %d\n", n)
}

func demoA() {
	months := [...]string{1: "J", 2: "W"}
	fmt.Println(months)

}

func chandemo() {
	strChan := make(chan string, 3)
	syncChan := make(chan struct{})

	go func() {

		defer func() {
			syncChan <- struct{}{}
		}()

		for v := range strChan {
			fmt.Println(v)
		}
	}()

	go func() {

		defer func() {
			syncChan <- struct{}{}
		}()

		for i := 0; i < 5; i++ {
			strChan <- strconv.Itoa(i)
			if i == 3 {
				close(strChan)
				break
			}
		}
	}()

	<-syncChan
	<-syncChan
}

func deferFunc(name string, a, b int) int {
	fmt.Println(name, a, b)
	return a + b
}

func deferTest() {
	a, b := 1, 2
	defer deferFunc("one", a, deferFunc("two", a, b))
	a = 0
	defer deferFunc("three", a, deferFunc("four", a, b))
	panic("five")
}

func selectTest() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case r := <-ch:
			fmt.Println(r)
		case ch <- i:
		}
	}
}

func t5() {
	f := func() func() {
		fmt.Println("0123456789")
		return func() {
			fmt.Println("asdasdasdasd")
		}
	}
	defer f()()
	defer func() {
		fmt.Println("qwert")
	}()
	fmt.Println("0000000000")
}

func t4() {
	fmt.Printf("%.3s \n%-5d%s", "qwertyuio11111p", 123, "sd")
}

func t3() {
	var nn map[string]int
	n := make(map[string]int)
	m := map[string]int{}
	fmt.Println(nn == nil)
	fmt.Println(n == nil)
	fmt.Println(m == nil)
}

func t2() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	fmt.Print(<-c)
	c <- 4
}

func t1() {
	str1 := `<a href="http://www.zhenai.com/zhenghun/aba/nv">阿坝女士征婚</a>`
	str2 := `<span class="nickName" data-v-3c42fade>劣酒灼心</span>    </div> <div class="des f-cl" data-v-3c42fade>阿坝 | 25岁 | 中专 | 未婚 | 165cm | 5001-8000元</div> <div class="actions" data-v-3c42fade><div class="item sayHi" data-v-3c42fade>打招呼</div>`
	var genderRe = regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[\w]+/([\w]+)">[^<]*士征婚</a>`)
	var profileRe = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^\|]+)\|([^\|]+)\|([^\|]+)\|([^\|]+)\|([^\|]+)\|([^<]+)</div>`)
	submatch1 := genderRe.FindStringSubmatch(str1)
	for _, s := range submatch1 {
		fmt.Println(s)
	}

	submatch2 := profileRe.FindStringSubmatch(str2)
	for _, s := range submatch2 {
		fmt.Println(s)
	}

	str3 := " qwer  "
	fmt.Println(strings.Trim(str3, " "))

	str4 := "165力魔"
	str4rune := []rune(str4)
	fmt.Println(len(str4rune))
	fmt.Println(string(str4rune[:len(str4rune)-2]))

	ageRune := []rune("25岁")
	age, err := strconv.Atoi(string(ageRune[:len(ageRune)-1]))
	if err != nil {
		panic(err)
	}
	fmt.Println(age)

}

// 交易成功通知主体, 存于RabbitMQ中
type Notification struct {
	NotifyURL string               `json:"notify_url"`   // 交易成功通知地址
	Body      *NotificationRequest `json:"request_body"` // 参数
}

// 推送交易通知request body
type NotificationRequest struct {
	TransactionID   string `json:"transaction_id"`    // 乐惠交易单号
	OrderNO         string `json:"order_no"`          //合作伙伴订单号
	UpstreamOrderNO string `json:"upstream_order_no"` // 微信/支付宝/银联订单号
	MerchantOrderNO string `json:"merchant_order_no"` // 支付客户端显示的商户单号（由乐惠生成）
	MerchantID      string `json:"merchant_id"`       // 商户ID
	TerminalID      string `json:"terminal_id"`       // 终端ID
	AppID           string `json:"app_id"`            // 微信/支付宝公众号或服务窗ID
	BuyerID         string `json:"buyer_id"`          // 微信或支付宝的用户ID
	Amount          uint64 `json:"amount"`            // 支付金额（分）
	Status          string `json:"status"`            // 支付订单状态： 支付中 processing  支付成功 succeeded 支付失败 failed 交易关闭 closed
	ClientType      string `json:"client_type"`       // 客户端类型： 微信 weixin 支付宝 alipay银联 unionpay京东 jdpay
	TradeType       string `json:"trade_type"`        // 支付方式： JSAPI jspay 刷卡 micropay
	CreatedAt       int64  `json:"created_at"`        // 交易创建时间
	FinishedAt      int64  `json:"finished_at"`       // 交易完成时间
}
