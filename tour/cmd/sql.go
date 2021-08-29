/*
 * @Author: Matt Meng
 * @Date: 2021-08-28 20:45:41
 * @LastEditors: Matt Meng
 * @LastEditTime: 2021-08-29 11:23:00
 * @Description: file content
 */
package cmd

import (
    "log"

    "github.com/go-programming-tour-book/tour/internal/sql2struct"
    "github.com/spf13/cobra"
)
// 声明七个 cmd 全局变量，用于接收外部的命令行参数
var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
    Use:"sql",
    Short:"sql 转换与处理",
    Long:"sql 转换与处理",
    Run: func(cmd *cobra.Command, args []string){},
}

var sql2structCmd = &cobra.Command{
    Use:"struct",
    Short:"sql 转换",
    Long:"sql 转换",
    Run: func(cmd *cobra.Command, args []string){
        dbInfo := &sql2struct.DBInfo{
            DBType:dbType,
            Host:host,
            UserName:username,
            Password:password,
            Charset:charset,
        }
        dbModel := sql2struct.NewDBModel(dbInfo)
        err := dbModel.Connect()
        if err!=nil{
            log.Fatalf("dbModel.Connect err: %v", err)
        }
        columns, err := dbModel.GetColumns(dbName, tableName)
        if err!=nil{
            log.Fatalf("dbModel.GetColumns err: %v", err)
        }

        template := sql2struct.NewStructTemplate()
        templateColumns := template.AssemblyColumns(columns)
        err = template.Generate(tableName, templateColumns)
        if err != nil{
            log.Fatalf("structTemplate.Generate err: %v", err)
        }
    },
}

func init(){
    sqlCmd.AddCommand(sql2structCmd)
    sql2structCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库账号")
    sql2structCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库密码")
    sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的HOST")
    sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
    sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例类型")
    sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
    sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}