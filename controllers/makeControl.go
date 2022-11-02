package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"makeControlModel/models"
	"os"
	"strings"
)

type MakeControlController struct {
	beego.Controller
}

func (c *MakeControlController) Get() {
	//设置是否显示视图
	beego.BConfig.WebConfig.AutoRender = true
	tableList := models.GetDbTables()
	var arr = make([]string, len(tableList))
	for m, maps := range tableList {
		arr[m] = string(maps["TABLE_NAME"])
	}
	c.Data["TableList"] = arr
	c.TplName = "makeControl/index.tpl"
}

func (c *MakeControlController) PostMakeControl() {
	//设置是否显示视图
	beego.BConfig.WebConfig.AutoRender = false
	tableName := c.GetString("table_name")
	path := c.GetString("path")
	baseController := c.GetString("baseController")
	if len(tableName) == 0 || len(path) == 0 {
		beego.Error("参数不能为空")
	}
	tableColum := models.GetTableColum(tableName)
	var columArr = make([]map[string]interface{}, len(tableColum))
	var postField, editPostField, editField, insertField, insertIfField, editIfField, primaryFiled, upperTable, humpTable,allUpperTable = "", "", "", "", "", "", "", "", "",""
	for m, maps := range tableColum {
		var column = make(map[string]interface{}, 4)
		column["column_name"] = ""

		if maps["COLUMN_NAME"] != nil && string(maps["COLUMN_KEY"]) != "PRI" {
			postField += "$" + string(maps["COLUMN_NAME"]) + " = $this->post('" + string(maps["COLUMN_NAME"]) + "');\r\n"
			if m == len(tableColum)-1 {
				insertField += "$" + string(maps["COLUMN_NAME"])
				insertIfField += "!$" + string(maps["COLUMN_NAME"])
				editField += "'" + string(maps["COLUMN_NAME"]) + "' =>$" + string(maps["COLUMN_NAME"]) + "\r\n"
			} else {
				insertIfField += "!$" + string(maps["COLUMN_NAME"]) + "||"
				insertField += "$" + string(maps["COLUMN_NAME"]) + ","
				editField += "'" + string(maps["COLUMN_NAME"]) + "' =>$" + string(maps["COLUMN_NAME"]) + ",\r\n"
			}
		}
		editPostField += "$" + string(maps["COLUMN_NAME"]) + " = $this->post('" + string(maps["COLUMN_NAME"]) + "');\r\n"
		if m == len(tableColum)-1 {
			editIfField += "!$" + string(maps["COLUMN_NAME"])
		} else {
			editIfField += "!$" + string(maps["COLUMN_NAME"]) + "||"
		}

		if maps["COLUMN_KEY"] != nil {
			if string(maps["COLUMN_KEY"]) == "PRI" {
				primaryFiled = string(maps["COLUMN_NAME"])
			}

		}

	}

	tableSlice := strings.Split(tableName, "_")
	for i := 0; i < len(tableSlice); i++ {
		allUpperTable += strFirstToUpper(tableSlice[i])
		if i == 0 {
			humpTable = strFirstToUpper(tableSlice[i])
		}else{
			humpTable =  tableSlice[i]
		}
		upperTable += humpTable
	}

	//创建文件
	fileName := path + upperTable + ".php"

	//检测文件是否已存在
	_,err := os.Open(fileName)
	if err == nil{
		panic("当前文件已存在")
		return
	}
	f, err := os.Create(fileName)

	classFile := "<?php" +
		"/**\r\n" +
		"*模块\r\n" +
		"*\r\n" +
		"* @author libo\r\n" +
		"*/\r\n" +
		"class " + upperTable + "Controller extends " + baseController + "\r\n" +
		"{\r\n" +
		"/**\r\n" +
		"* 分页条数\r\n" +
		"*/\r\n" +
		"const LIMIT = 15;\r\n" +
		"/**\r\n" +
		"* 入口方法\r\n" +
		"*/\r\n" +
		"public function init()\r\n" +
		"{\r\n" +
		"parent::init();\r\n" +
		"}\r\n" +
		"/**\r\n" +
		"* 用户列表\r\n" +
		"*\r\n" +
		"* @return html\r\n" +
		"*/\r\n" +
		"public function " + humpTable + "ListAction()\r\n" +
		"{\r\n" +
		"$page = $this->get('p', 1);\r\n" +
		"$start = ($page - 1) * self::LIMIT;\r\n" +
		"$" + tableName + "_list = (new " + allUpperTable + "Model())->page" + allUpperTable + "List($start, self::LIMIT);\r\n" +
		"$pager = new Cola_Pager($page, self::LIMIT, $" + tableName + "_list['count'], '/Admin/" + upperTable + "/" + humpTable + "List?p=%page%');\r\n" +
		"$this->assign('pager', $pager->html());\r\n" +
		"$this->assign('" + tableName + "_list', $" + tableName + "_list);\r\n" +
		"$this->displays('" + allUpperTable + "/" + humpTable + "List.phtml');\r\n" +
		"}\r\n" +
		"\r\n" +
		"/**\r\n" +
		"* 提交添加信息\r\n" +
		"*\r\n" +
		"* @return json\r\n" +
		"*/\r\n" +
		"public function postAdd" + upperTable + "Action()\r\n" +
		"{\r\n" + postField +
		"if (" + insertIfField + ") {\r\n" +
		"$this->echoJson(['type' => 'error', 'message' => '参数不能为空']);\r\n" +
		"}\r\n" +
		"$add_status = (new " + allUpperTable + "Model())->add" + allUpperTable + "(" + insertField + ");\r\n" +
		"if ($add_status) {\r\n" +
		"$this->echoJson(['type' => 'success', 'message' => '添加成功']);\r\n" +
		"} else {\r\n" +
		"$this->echoJson(['type' => 'error', 'message' => '添加失败']);\r\n" +
		"}\r\n" +
		"}\r\n" +
		"\r\n" +
		"/**\r\n" +
		"* 编辑信息\r\n" +
		"*\r\n" +
		"* @return json\r\n" +
		"*/\r\n" +
		"public function postEdit" + upperTable + "Action()\r\n" +
		"{\r\n" + editPostField +
		"if (" + editIfField + ") {\r\n" +
		"$this->echoJson(['type' => 'error', 'message' => '参数不能为空']);\r\n" +
		"}\r\n" +
		"$param = [\r\n" + editField + "];\r\n" +
		"$update_status = (new " + allUpperTable + "Model())->update" + allUpperTable + "($" + primaryFiled + ", $param);\r\n" +
		"if ($update_status) {\r\n" +
		"$this->echoJson(['type' => 'success', 'message' => '编辑成功']);\r\n" +
		"} else {\r\n" +
		"$this->echoJson(['type' => 'error', 'message' => '编辑失败']);\r\n" +
		"}\r\n" +
		"}\r\n" +
		"\r\n" +
		"/**\r\n" +
		"* 删除\r\n" +
		"*\r\n" +
		"* @return json\r\n" +
		"*/\r\n" +
		"public function delete" + upperTable + "Action()\r\n" +
		"{\r\n" +
		"$" + primaryFiled + " = $this->post('" + primaryFiled + "');\r\n" +
		"if (!$" + primaryFiled + ") {\r\n" +
		"$this->echoJson(['type' => 'error', 'message' => '参数不能为空']);\r\n" +
		"}\r\n" +
		"$delete_status = (new " + allUpperTable + "Model())->delete" + allUpperTable + "($" + primaryFiled + ");\r\n" +
		"if ($delete_status) {\r\n" +
		"$this->echoJson(['type' => 'success', 'message' => '删除成功']);\r\n" +
		"} else {\r\n" +
		"$this->echoJson(['type' => 'error', 'message' => '删除失败']);\r\n" +
		"}\r\n" +
		"}\r\n}"

	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte(classFile))
	}

	fmt.Println(columArr)
}
