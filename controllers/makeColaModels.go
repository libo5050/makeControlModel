package controllers

import (
	"github.com/astaxie/beego"
	"makeModels/models"
	"strings"
	"os"
	"fmt"
)

type MakeColaModelsController struct {
	beego.Controller
}

func (c *MakeColaModelsController) GetColaPretreatment() {
	//设置是否显示视图
	beego.BConfig.WebConfig.AutoRender = true
	tableList := models.GetDbTables()
	var arr = make([]string, len(tableList))
	for m, maps := range tableList {
		arr[m] = string(maps["TABLE_NAME"])
	}
	c.Data["TableList"] = arr
	c.TplName = "makeColaModels/index.tpl"
}

func (c *MakeColaModelsController) PostMakeColaModels() {
	//设置是否显示视图
	beego.BConfig.WebConfig.AutoRender = false
	tableName := c.GetString("table_name")
	path := c.GetString("path")
	if len(tableName) == 0 || len(path) == 0 {
		beego.Error("参数不能为空")
	}
	//获取表结构
	tableColum := models.GetTableColum(tableName)
	var columArr = make([]map[string]interface{}, len(tableColum))
	var pkField, insertFieldList = "", ""
	var insertField = "[\r\n"
	var insertFieldDescript = "/**\r\n* 添加数据\r\n*\r\n"
	for m, maps := range tableColum {
		var column = make(map[string]interface{}, 4)
		column["column_name"] = ""
		column["data_type"] = ""
		column["column_type"] = ""
		column["column_key"] = ""

		if maps["COLUMN_NAME"] != nil && string(maps["COLUMN_KEY"]) != "PRI" {
			if m == len(tableColum)-1 {
				insertField += "'" + string(maps["COLUMN_NAME"]) + "' =>$" + string(maps["COLUMN_NAME"]) + "\r\n"
				insertFieldList += "$" + string(maps["COLUMN_NAME"])
			} else {
				insertField += "'" + string(maps["COLUMN_NAME"]) + "' =>$" + string(maps["COLUMN_NAME"]) + ",\r\n"
				insertFieldList += "$" + string(maps["COLUMN_NAME"]) + ","
			}
			column["column_name"] = string(maps["COLUMN_NAME"])
		}

		if maps["DATA_TYPE"] != nil && string(maps["COLUMN_KEY"]) != "PRI" {
			if string(maps["DATA_TYPE"]) == "datetime" || string(maps["DATA_TYPE"]) == "varchar" || string(maps["DATA_TYPE"]) == "text" || string(maps["DATA_TYPE"]) == "longtext" {
				insertFieldDescript += "* @param string $" + string(maps["COLUMN_NAME"]) + " " + string(maps["COLUMN_COMMENT"]) + "\r\n"
			} else {
				insertFieldDescript += "* @param int $" + string(maps["COLUMN_NAME"]) + " " + string(maps["COLUMN_COMMENT"]) + "\r\n"
			}
			column["data_type"] = string(maps["DATA_TYPE"])
		}

		if maps["COLUMN_TYPE"] != nil {
			column["column_type"] = string(maps["COLUMN_TYPE"])
		}

		if maps["COLUMN_KEY"] != nil {
			if string(maps["COLUMN_KEY"]) == "PRI" {
				pkField = string(maps["COLUMN_NAME"])
			}
			column["column_key"] = string(maps["COLUMN_KEY"])
		}
		columArr[m] = column
	}
	insertField += "];\r\n"
	insertFieldDescript += "*\r\n* @return int\r\n*/\r\n"
	//go func() {
	//	f, _ := os.Create(fileName1)
	//	_, err = f.Write([]byte("测试多协程写数据"))
	//	defer f.Close()
	//}()
	upperTable := ""
	tableSlice := strings.Split(tableName, "_")
	for i := 0; i < len(tableSlice); i++ {
		upperTable += strPretreatmentFirstToUpper(tableSlice[i])
	}

	//创建文件
	fileName := path + upperTable + ".php"
	//fileName1 := path + tableName + "1.php"
	f, err := os.Create(fileName)

	classFile := "<?php " +
		"class " + upperTable + "Model extends Cola_Model\r\n{\r\n" +
		"/** 构造方法\r\n" +
		"*\r\n" +
		"* @return \r\n" +
		"*/\r\n" +
		"public function __construct()\r\n" +
		"{\r\n" +
		"$this->_table = '"+tableName+"';\r\n" +
		"$this->_pk = '"+pkField+"';\r\n" +
		"}\r\n" +
		insertFieldDescript +
		"public function add" + upperTable + "(" + insertFieldList + ")\r\n" +
		"{\r\n" +
		"$param =  " + insertField +
		"return $this->insert($param);\r\n" +
		"}\r\n" +
		"/**\r\n" +
		"* 更新数据\r\n" +
		"*\r\n" +
		"* int $id 主键id\r\n" +
		"*\r\n" +
		"* param array data 更新的数据\r\n" +
		"*\r\n" +
		"* @return int\r\n" +
		"*/\r\n" +
		"public function update" + upperTable + "($id, $data)\r\n" +
		"{\r\n" +
		"return $this->update($id,$data);\r\n" +
		"}\r\n" +
		"/**\r\n" +
		"* 获取数据\r\n" +
		"*\r\n" +
		"* @param int id 主键id\r\n" +
		"*\r\n" +
		"* @return []\r\n" +
		"*/\r\n" +
		"public function get" + upperTable + "($id)\r\n" +
		"{\r\n" +
		"return $this->db()->row(\"select * from {$this->_table} where {$this->_pk} ={$id}\");\r\n" +
		"}\r\n" +
		"/**\r\n" +
		"* 删除数据\r\n" +
		"*\r\n" +
		"* @param int $id 主键id\r\n" +
		"*\r\n" +
		"* @return int\r\n" +
		"*/\r\n" +
		"public function delete" + upperTable + "($id)\r\n" +
		"{\r\n" +
		"return $this->delete($id);\r\n" +
		"}\r\n" +
		"/**\r\n" +
		"* 分页获取数据\r\n" +
		"*\r\n" +
		"* @param int $start 开始条数\r\n" +
		"* @param int $limit 每页条数\r\n" +
		"*\r\n" +
		"* @return []\r\n" +
		"*/\r\n" +
		"public function page" + upperTable + "List($start,$limit)\r\n" +
		"{\r\n" +
		"$count = $this->count(\"1= 1\");\r\n" +
		"$data = [];\r\n" +
		"if ($count) {\r\n" +
		"$sql = \"select * from `{$this->_table}` order by `add_time` desc,`{$this->_pk}` desc limit {$start},{$limit}\";\r\n" +
		"$data = $this->sql($sql);\r\n" +
		"}\r\n" +
		"return ['count' => $count, 'data' => $data];\r\n" +
		"}\r\n" +
		"}"
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte(classFile))
	}

	fmt.Println(columArr)
}

