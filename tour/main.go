/*
 * @Author: Matt Meng
 * @Date: 2021-04-11 11:53:50
 * @LastEditors: Matt Meng
 * @LastEditTime: 2021-08-28 16:34:29
 * @Description: file content
 */

package main

import (
    "log"

    "github.com/go-programming-tour-book/tour/cmd"
)

func main() {
    err := cmd.Execute()
    if err != nil {
        log.Fatalf("cmd.Execute err: %v", err)
    }
}
