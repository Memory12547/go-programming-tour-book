/*
 * @Author: Matt Meng
 * @Date: 2021-04-18 14:05:32
 * @LastEditors: Matt Meng
 * @LastEditTime: 2021-04-18 14:33:37
 * @Description: file content
 */

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	gin.SetMode(gin.ReleaseMode)
	r.Run()
}
