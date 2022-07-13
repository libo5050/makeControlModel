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

    header,
    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }

    header {
      padding: 100px 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .description {
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
      text-decoration: none;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }

    .item {
        margin-left:120px;
        margin-bottom:20px;
    }
    .item span {

    }
  </style>
</head>

<body>
<form name="form" action="/makeModels/postMakeModels" method="post">
  <div class="item">
  <span>表名:</span>
   <select style="width:500px;" name="table_name">
   {{range $ind, $elem := .TableList}}
       <option value="{{$elem}}">{{$elem}}</option>
   {{end}}
   <select>
 </div>
 <div class="item">
   <span>生成路径:</span>
    <input autocomplete="off" style="width:500px;" type="text" name="path">
 </div>
 <div class="item">
     <input type="submit" name="submit" value="提交">
  </div>
  </form>
</body>
</html>
