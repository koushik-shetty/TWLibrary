package views

type IndexOptions struct {
	Title string
}

const Index = `	<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html lang="en">
<head>
	<meta http-equiv="content-type" content="text/html; charset=utf-8">
	<title>{{.Title}}</title>
	<link rel="icon" type="image/x-icon" href="public/styles/favicon.ico">
	<link rel="stylesheet" type="text/css" href="public/styles/main.css">
</head>

<body>
	<span>
		<h1 id="ww">Western Wall</h1>
	</span>
	<div id="root"></div>
	<script src="public/javascript/main.js"></script>
</body>
</html>`
