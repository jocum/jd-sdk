package sdk

import (
	"net/http"
	"net"
	"time"
	"io/ioutil"
)


/*
	@Description 	net/http  自带的 get 请求没有超时限制  所以要自己写一个有超时限制的请求方法
	@Params 	usrStr  string 	请求的url
							res   interace{}  返回的数据
	@Return 	err 	error
	@Author  cwy
*/
func Get(urlStr string ) ([]byte,error)  {
	client := &http.Client{
        Transport: &http.Transport{
            Dial: func(netw, addr string) (net.Conn, error) {
                conn, err := net.DialTimeout(netw, addr, time.Second*5)    //设置建立连接超时
                if err != nil {
                    return nil, err
                }
                conn.SetDeadline(time.Now().Add(time.Second * 5))    //设置发送接受数据超时
                return conn, nil
            },
            ResponseHeaderTimeout: time.Second * 5,
        },
    }
	reqest,err := http.NewRequest("GET",urlStr,nil)  
	if err != nil {
		return nil,err
	}
	response,err := client.Do(reqest)
	if err != nil {
		return nil,err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}