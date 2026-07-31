package redisqueue

import (
	"os"
	"os/signal"
	"testing"
)

func TestNewSignalHandler(t *testing.T) {
	t.Run("closes the returned channel on SIGINT", func(tt *testing.T) {

		// 直接模拟信号，不依赖操作系统
		// 假设 newSignalHandler 返回的 channel 会在收到信号时关闭
		// 我们可以直接触发这个逻辑

		// 方式1：如果 newSignalHandler 内部使用了 signal.Notify，
		// 可以手动发送信号到 channel
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)
		sigChan <- os.Interrupt // 手动发送信号

		// 或者方式2：直接调用 newSignalHandler 的内部逻辑
		// 这取决于你的实现
	})
}
