<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="shortcut icon" type="image/x-icon" />

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }



    .content{
            position:absolute;
            left:30%;
            top:20%;
        }

        .item {
            width:650px;
            margin-top:30px;
        }
        .item span {
            width:80px;
        }

        input:focus{
             border:none;
             outline:none;
        }

        .sub{
            width:180px;
            height:35px;
            line-height:35px;
            margin-left:85px;
            background: #cd4a40;
            border:1px solid #cd4a40;
            border-radius: 30px;
            text-align: center;
        }
  </style>
</head>

<body>
<form name="form" action="/makeColaModels/postMakeColaModels" method="post">
<div class="content">
  <div class="item">
  <span>表名:</span>
   <select style="width:300px;margin-left:28px;height:35px;border: 1px solid #d9d9d9;border-radius: 15px;" name="table_name">
   {{range $ind, $elem := .TableList}}
       <option value="{{$elem}}">{{$elem}}</option>
   {{end}}
   <select>
 </div>
 <div class="item">
   <span>生成路径:</span>
    <input autocomplete="off" style="width:300px;height:35px;border: 1px solid #d9d9d9;border-radius: 15px;" type="text" name="path">
 </div>
 <div class="item">
     <input type="submit" class="sub" name="submit" value="提交">
  </div>
  </div>
  </form>
</body>
</html>
