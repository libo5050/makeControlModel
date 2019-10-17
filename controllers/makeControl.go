package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"makeModels/models"
	"os"
	"strings"
)

type MakeControlController struct {
	beego.Controller
}

func (c *MakeControlController) Get() {
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
	var postField, editPostField, editField, insertField, insertIfField, editIfField, primaryFiled, upperTable, humpTable = "", "", "", "", "", "", "", "", ""

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
	humpTable = tableSlice[0]
	for i := 0; i < len(tableSlice); i++ {
		upperTable += strFirstToUpper(tableSlice[i])
		if i > 0 {
			humpTable += strFirstToUpper(tableSlice[i])
		}
	}

	//创建文件
	fileName := path + upperTable + ".php"

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
		"$" + tableName + "_list = (new " + upperTable + "Model())->page" + upperTable + "List($start, self::LIMIT);\r\n" +
		"$pager = new Cola_Pager($page, self::LIMIT, $" + tableName + "_list['count'], '/" + upperTable + "/" + upperTable + "List?p=%page%');\r\n" +
		"$this->assign('pager', $pager->html());\r\n" +
		"$this->assign('" + tableName + "_list', $" + tableName + "_list);\r\n" +
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
		"$add_status = (new " + upperTable + "Model())->add" + upperTable + "(" + insertField + ");\r\n" +
		"if ($add_status) {\r\n" +
		"$this->echoJson(['type' => 'success', 'message' => '添加成功']);\r\n" +
		"} else {\r\n" +
		"$this->echoJson(['type' => 'success', 'message' => '添加失败']);\r\n" +
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
		"$update_status = (new " + upperTable + "Model())->update" + upperTable + "($" + primaryFiled + ", $param);\r\n" +
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
		"$delete_status = (new " + upperTable + "Model())->delete" + upperTable + "($" + primaryFiled + ");\r\n" +
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
