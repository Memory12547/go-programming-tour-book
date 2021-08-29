/*
 * @Author: Matt Meng
 * @Date: 2021-08-28 20:47:50
 * @LastEditors: Matt Meng
 * @LastEditTime: 2021-08-29 11:14:36
 * @Description: file content
 */
package sql2struct

import (
    "fmt"
    "os"
    "text/template"
    "github.com/go-programming-tour-book/tour/internal/word"
)

const structTpl = `type {{.TableName | ToCamelCase}} struct{
{{range .Columns}}    {{$length := len .Comment}} {{if gt $length 0}}// {{.Comment}} {{else}} //{{.Name}} {{end}}
    {{$typeLen := len .Type}} {{if gt $typeLen 0}}{{.Name | ToCamelCase}}    {{.Type}}    {{.Tag}}{{else}}{{.Name}}{{end}}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
    return "{{.TableName}}"
}`

type StructTemplate struct{
    structTpl string
}

type StructColumn struct{
    Name    string
    Type    string
    Tag     string
    Comment string
}

type StructTemplateDB struct{
    TableName string
    Columns []*StructColumn
}

func NewStructTemplate() *StructTemplate{
    return &StructTemplate{structTpl:structTpl}
}


//数据库类型到Go 结构体的转换，JSON tag的处理
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn{
    tplColumns := make([]*StructColumn, 0, len(tbColumns))
    for _, column := range tbColumns{
        tag := fmt.Sprintf("`" +"json:" + "\"%s\"" + "`", column.ColumnName)
        tplColumns = append(tplColumns, &StructColumn{
            Name: column.ColumnName,
            Type: DBTypeToStructType[column.DataType],
            Tag:  tag,
            Comment: column.ColumnComment,
        })
    }

    return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
    //声明一个名为 sql2struct 的模板对象，定义自定义函数 ToCamelCase，并与 word.UnderscoreToUpperCamelCase 绑定
    tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
        "ToCamelCase": word.UnderscoreToUpperCamelCase,
    }).Parse(t.structTpl))

    tplDB := StructTemplateDB{
        TableName: tableName,
        Columns:   tplColumns,
    }
    err := tpl.Execute(os.Stdout, tplDB)
    if err != nil {
        return err
    }

    return nil
}