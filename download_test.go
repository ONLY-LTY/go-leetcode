package leetcode

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestDownload(t *testing.T) {
	urls := []string{
		"https://static.zybuluo.com/bethunebtj/1mm2eqnuop9rcccap915qrzx/image_1cfl05d2f1ub879c1lc5qsq14n9m.png",
	}
	for _, imageURL := range urls {
		download(imageURL)
	}
}

func download(imageURL string) {
	saveDirectory := "/Users/tingyang/hexo-blog/source/img" // 图片保存的目录
	imageName := filepath.Base(imageURL)                    // 获取URL中的文件名作为保存的文件名
	savePath := filepath.Join(saveDirectory, imageName)     // 构建完整的保存路径

	// 确保保存目录存在
	if err := os.MkdirAll(saveDirectory, 0755); err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		return
	}

	err := downloadImage(imageURL, savePath)
	if err != nil {
		fmt.Printf("Error downloading image: %v\n", err)
	} else {
		fmt.Printf("Image downloaded successfully and saved to: %s\n", savePath)
	}
}
