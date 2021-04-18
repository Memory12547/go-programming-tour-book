/*
 * @Author: Matt Meng
 * @Date: 2021-04-11 17:19:14
 * @LastEditors: Matt Meng
 * @LastEditTime: 2021-04-11 17:19:14
 * @Description: file content
 */
package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
}
