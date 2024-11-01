package leetcode

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// downloadImage 下载图片并保存到指定路径
func downloadImage(url, savePath string) error {
	// 发起GET请求
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download image, status code: %download", resp.StatusCode)
	}

	// 创建文件
	out, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// 将响应体内容写入文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
