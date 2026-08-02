package iputils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-admin/core/utils/log"
	"io"
	"net"
	"net/http"
)

// GetLocation 获取外网ip地址
func GetLocation(ip, key string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "inner ip"
	}
	url := "https://restapi.amap.com/v5/ip?ip=" + ip + "&type=4&key=" + key
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("restapi.amap.com failed: %s", err)
		return "unknown ip"
	}
	defer resp.Body.Close()
	s, err := io.ReadAll(resp.Body)

	m := make(map[string]string)

	err = json.Unmarshal(s, &m)
	if err != nil {
		log.Errorf("Umarshal failed: %s", err)
	}
	//if m["province"] == "" {
	//	return "未知位置"
	//}
	return m["country"] + "-" + m["province"] + "-" + m["city"] + "-" + m["district"] + "-" + m["isp"]
}

// GetLocaHost 获取局域网ip地址
func GetLocaHost() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		log.Errorf("net.Interfaces failed, err: %s", err)
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}

	}
	return ""
}

// GetClientIP 获取客户端真实 IP。
// 依赖 gin 的可信代理配置（SetTrustedProxies）：
// - 未配置可信代理时返回直连对端 IP（RemoteAddr），客户端伪造的 X-Forwarded-For 无效；
// - 配置了可信代理（如 nginx）时，返回代理链中最左侧的真实客户端 IP。
func GetClientIP(c *gin.Context) string {
	return c.ClientIP()
}
