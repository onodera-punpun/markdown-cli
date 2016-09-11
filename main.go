package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	markdown "github.com/shurcooL/github_flavored_markdown"
)

func main() {
	filename := os.Args[1]
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	extension := filepath.Ext(filename)
	name := filename[0:len(filename)-len(extension)]

	markdownBytes := markdown.Markdown(fileBytes)

	outputFile, err := os.Create(name + ".html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("markdown").Parse(htmlTemplate)
	err = tmpl.Execute(outputFile, context{Title: filepath.Clean(name), Body: string(markdownBytes)})
	if err != nil {
		log.Fatal(err)
	}
}

type context struct {
	Title string
	Body  string
}

var htmlTemplate = `
<html>
	<head>
		<meta charset="utf-8">
		<title>{{.Title}}</title>
		<style>
* {
	outline: 0px;
	text-decoration: none; 
}

body {
	min-width: 200px;
	max-width: 790px;
	padding: 100px;
}

::selection {
	background: #37BF8D;
	color: #37BF8D;
}
::-moz-selection {
	background: #37BF8D;
	color: #37BF8D;
}

.markdown-body {
	-webkit-text-size-adjust: 100%;
	-ms-text-size-adjust: 100%;
	color: #021B21;
	overflow: hidden;
	font-family: Helvetica, Arial, freesans, sans-serif;
	font-size: 14px;
	line-height: 29px;
	word-wrap: break-word;
}

.markdown-body a {
	background-color: transparent;
	color: #37BF8D;
}

.markdown-body a:active,
.markdown-body a:hover {
	background: #37BF8D;
	color: #37BF8D;
}

.markdown-body strong {
	font-weight: bold;
}

.markdown-body h1 {
	font-weight: bold;
}

.markdown-body img {
	border: 0;
}

.markdown-body hr {
	box-sizing: content-box;
	height: 0;
}

.markdown-body pre {
	overflow: auto;
}

.markdown-body code,
.markdown-body kbd,
.markdown-body pre {
	font-family: monospace, monospace;
	font-size: 1em;
}

.markdown-body input {
	color: inherit;
	font: inherit;
	margin: 0;
}

.markdown-body html input[disabled] {
	cursor: default;
}

.markdown-body input {
	line-height: normal;
}

.markdown-body input[type="checkbox"] {
	box-sizing: border-box;
	padding: 0;
}

.markdown-body table {
	border-collapse: collapse;
	border-spacing: 0;
}

.markdown-body td,
.markdown-body th {
	padding: 0;
}

.markdown-body * {
	box-sizing: border-box;
}

.markdown-body input {
	font: 13px/1.4 Helvetica, arial, nimbussansl, liberationsans, freesans, clean, sans-serif, "Segoe UI Emoji", "Segoe UI Symbol";
}

.markdown-body a {
	font-weight: bold;
	text-decoration: none;
}

.markdown-body a:hover,
.markdown-body a:active {
	background: #37BF8D;
	color: #37BF8D;
}

.markdown-body hr {
	height: 0;
	margin: 15px 0;
	overflow: hidden;
	background: transparent;
	border: 0;
}

.markdown-body hr:before {
	display: table;
	content: "";
}

.markdown-body hr:after {
	display: table;
	clear: both;
	content: "";
}

.markdown-body h1,
.markdown-body h2,
.markdown-body h3,
.markdown-body h4,
.markdown-body h5,
.markdown-body h6 {
	margin-top: 15px;
	margin-bottom: 15px;
	line-height: 1.1;
}

.markdown-body h1 {
	font-size: 30px;
}

.markdown-body h2 {
	font-size: 21px;
}

.markdown-body h3 {
	font-size: 16px;
}

.markdown-body h4 {
	font-size: 14px;
}

.markdown-body h5 {
	font-size: 12px;
}

.markdown-body h6 {
	font-size: 11px;
}

.markdown-body blockquote {
	margin: 0;
}

.markdown-body ul,
.markdown-body ol {
	padding: 0;
	margin-top: 0;
	margin-bottom: 0;
}

.markdown-body ol ol,
.markdown-body ul ol {
	list-style-type: lower-roman;
}

.markdown-body ul ul ol,
.markdown-body ul ol ol,
.markdown-body ol ul ol,
.markdown-body ol ol ol {
	list-style-type: lower-alpha;
}

.markdown-body dd {
	margin-left: 0;
}

.markdown-body code {
	font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
	font-size: 12px;
}

.markdown-body pre {
	margin-top: 0;
	margin-bottom: 0;
	font: 12px Consolas, "Liberation Mono", Menlo, Courier, monospace;
}

.markdown-body .octicon {
	font: normal normal normal 16px/1 octicons-anchor;
	display: inline-block;
	text-decoration: none;
	text-rendering: auto;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	-webkit-user-select: none;
	-moz-user-select: none;
	-ms-user-select: none;
	user-select: none;
}

.markdown-body .octicon-link:before {
	content: '\f05c';
}

.markdown-body>*:first-child {
	margin-top: 0 !important;
}

.markdown-body>*:last-child {
	margin-bottom: 0 !important;
}

.markdown-body a:not([href]) {
	color: inherit;
	text-decoration: none;
}

.markdown-body .anchor {
	position: absolute;
	top: 0;
	left: 0;
	display: block;
	padding-right: 6px;
	padding-left: 30px;
	margin-left: -30px;
}

.markdown-body .anchor:focus {
	outline: none;
}

.markdown-body h1,
.markdown-body h2,
.markdown-body h3,
.markdown-body h4,
.markdown-body h5,
.markdown-body h6 {
	position: relative;
	margin-top: 1em;
	margin-bottom: 16px;
	font-weight: bold;
	line-height: 1.4;
}

.markdown-body h1 .octicon-link,
.markdown-body h2 .octicon-link,
.markdown-body h3 .octicon-link,
.markdown-body h4 .octicon-link,
.markdown-body h5 .octicon-link,
.markdown-body h6 .octicon-link {
	display: none;
	color: #000;
	vertical-align: middle;
}

.markdown-body h1:hover .anchor,
.markdown-body h2:hover .anchor,
.markdown-body h3:hover .anchor,
.markdown-body h4:hover .anchor,
.markdown-body h5:hover .anchor,
.markdown-body h6:hover .anchor {
	padding-left: 8px;
	margin-left: -30px;
	text-decoration: none;
}

.markdown-body h1:hover .anchor .octicon-link,
.markdown-body h2:hover .anchor .octicon-link,
.markdown-body h3:hover .anchor .octicon-link,
.markdown-body h4:hover .anchor .octicon-link,
.markdown-body h5:hover .anchor .octicon-link,
.markdown-body h6:hover .anchor .octicon-link {
	display: inline-block;
}

.markdown-body h1 {
	font-size: 20px;
	background: #F6F6F6;
	line-height: 25px;
	padding: 32px;
	text-transform: uppercase;
}

.markdown-body h2 {
	padding-bottom: 10px;
	font-size: 16px;
	border-bottom: 4px solid #F6F6F6;
	line-height: 25px;
	margin-top: 80px;
	margin-bottom: 20px;
}

.markdown-body h3 {
	padding-bottom: 10px;
	font-family: Georgia, Times New Roman, serif;
	font-weight: normal;
	font-size: 18px;
	color: #37BF8D;
	text-align: right;
}

.markdown-body h4 {
	font-size: 1.25em;
}

.markdown-body h5 {
	font-size: 1em;
}

.markdown-body h6 {
	font-size: 1em;
	color: #777;
}


.markdown-body p,
.markdown-body blockquote,
.markdown-body ul,
.markdown-body ol,
.markdown-body dl,
.markdown-body table,
.markdown-body pre {
	margin-top: 0;
	margin-bottom: 16px;
}

.markdown-body hr {
	height: 4px;
	padding: 0;
	margin: 16px 0;
	background-color: #F6F6F6;
	border: 0 none;
}

.markdown-body ul,
.markdown-body ol {
	padding-left: 2em;
}

.markdown-body ul ul,
.markdown-body ul ol,
.markdown-body ol ol,
.markdown-body ol ul {
	margin-top: 0;
	margin-bottom: 0;
}

.markdown-body li>p {
	margin-top: 16px;
}

.markdown-body dl {
	padding: 0;
}

.markdown-body dl dt {
	padding: 0;
	margin-top: 16px;
	font-size: 1em;
	font-style: italic;
	font-weight: bold;
}

.markdown-body dl dd {
	padding: 0 16px;
	margin-bottom: 16px;
}

.markdown-body blockquote {
	padding: 0 15px;
	color: #021B21;
	font-weight: bold;
	border-left: 4px solid #F6F6F6;
}

.markdown-body blockquote>:first-child {
	margin-top: 0;
}

.markdown-body blockquote>:last-child {
	margin-bottom: 0;
}

.markdown-body table {
	display: block;
	width: 100%;
	overflow: auto;
	word-break: normal;
	word-break: keep-all;
}

.markdown-body table th {
	font-weight: bold;
	border-bottom: 4px solid #f6f6f6;
}

.markdown-body table th,
.markdown-body table td {
	padding: 8px 14px;
}

.markdown-body table td:nth-child(1n) {
	border-bottom: 4px solid #f6f6f6;
}

.markdown-body table td:nth-child(2n) {
	border-left: 8px solid #ffffff;
	border-right: 8px solid #ffffff;
}

.markdown-body table tr:nth-child(2n) {
}

.markdown-body img {
	max-width: 100%;
	box-sizing: border-box;
	border: 1px solid #F5F5F5;
	border-radius: 3px;
}

.markdown-body code {
	padding: 0;
	padding-top: 0.2em;
	padding-bottom: 0.2em;
	margin: 0;
	font-size: 85%;
	background-color: #F6F6F6;
}

.markdown-body code:before,
.markdown-body code:after {
	letter-spacing: -0.2em;
	content: "\00a0";
}

.markdown-body pre>code {
	padding: 0;
	margin: 0;
	font-size: 100%;
	word-break: normal;
	white-space: pre;
	background: transparent;
	border: 0;
}

.markdown-body .highlight {
	margin-bottom: 16px;
}

.markdown-body .highlight pre,
.markdown-body pre {
	padding: 16px;
	overflow: auto;
	font-size: 85%;
	line-height: 1.45;
	background-color: #F6F6F6;
}

.markdown-body .highlight pre {
	margin-bottom: 0;
	word-break: normal;
}

.markdown-body pre {
	word-wrap: normal;
}

.markdown-body pre code {
	display: inline;
	max-width: initial;
	padding: 0;
	margin: 0;
	overflow: initial;
	line-height: inherit;
	word-wrap: normal;
	background-color: transparent;
	border: 0;
}

.markdown-body pre code:before,
.markdown-body pre code:after {
	content: normal;
}

.markdown-body kbd {
	display: inline-block;
	padding: 3px 5px;
	font-size: 11px;
	line-height: 10px;
	color: #555;
	vertical-align: middle;
	background-color: #fcfcfc;
	border: solid 1px #ccc;
	border-bottom-color: #bbb;
	box-shadow: inset 0 -1px 0 #bbb;
}

.markdown-body .pl-c {
	color: #969896;
}

.markdown-body .pl-c1,
.markdown-body .pl-s .pl-v {
	color: #0086b3;
}

.markdown-body .pl-e,
.markdown-body .pl-en {
	color: #795da3;
}

.markdown-body .pl-s .pl-s1,
.markdown-body .pl-smi {
	color: #333;
}

.markdown-body .pl-ent {
	color: #63a35c;
}

.markdown-body .pl-k {
	color: #a71d5d;
}

.markdown-body .pl-pds,
.markdown-body .pl-s,
.markdown-body .pl-s .pl-pse .pl-s1,
.markdown-body .pl-sr,
.markdown-body .pl-sr .pl-cce,
.markdown-body .pl-sr .pl-sra,
.markdown-body .pl-sr .pl-sre {
	color: #183691;
}

.markdown-body .pl-v {
	color: #ed6a43;
}

.markdown-body .pl-id {
	color: #b52a1d;
}

.markdown-body .pl-ii {
	background-color: #b52a1d;
	color: #f8f8f8;
}

.markdown-body .pl-sr .pl-cce {
	color: #63a35c;
	font-weight: bold;
}

.markdown-body .pl-ml {
	color: #693a17;
}

.markdown-body .pl-mh,
.markdown-body .pl-mh .pl-en,
.markdown-body .pl-ms {
	color: #1d3e81;
	font-weight: bold;
}

.markdown-body .pl-mq {
	color: #008080;
}

.markdown-body .pl-mi {
	color: #333;
	font-style: italic;
}

.markdown-body .pl-mb {
	color: #333;
	font-weight: bold;
}

.markdown-body .pl-md {
	background-color: #ffecec;
	color: #bd2c00;
}

.markdown-body .pl-mi1 {
	background-color: #eaffea;
	color: #55a532;
}

.markdown-body .pl-mdr {
	color: #795da3;
	font-weight: bold;
}

.markdown-body .pl-mo {
	color: #1d3e81;
}

.markdown-body kbd {
	display: inline-block;
	padding: 3px 5px;
	font: 11px Consolas, "Liberation Mono", Menlo, Courier, monospace;
	line-height: 10px;
	color: #555;
	vertical-align: middle;
	background-color: #fcfcfc;
	border: solid 1px #ccc;
	border-bottom-color: #bbb;
	box-shadow: inset 0 -1px 0 #bbb;
}

.markdown-body .task-list-item {
	list-style-type: none;
}

.markdown-body .task-list-item+.task-list-item {
	margin-top: 3px;
}

.markdown-body .task-list-item input {
	margin: 0 0.35em 0.25em -1.6em;
	vertical-align: middle;
}

.markdown-body :checked+.radio-label {
	z-index: 1;
	position: relative;
	border-color: #4078c0;
}

.highlight .h{color:#333;font-style:normal;font-weight:normal}
.highlight .mf,.highlight .mh,.highlight .mi,.highlight .mo,.highlight .il,.highlight .m{color:#945277}
.highlight .s,.highlight .sb,.highlight .sc,.highlight .sd,.highlight .s2,.highlight .se,.highlight .sh,.highlight .si,.highlight .sx,.highlight .s1{color:#df5000}
.highlight .kc,.highlight .kd,.highlight .kn,.highlight .kp,.highlight .kr,.highlight .kt,.highlight .k,.highlight .o{font-weight:bold}
.highlight .kt{color:#458}
.highlight .c,.highlight .cm,.highlight .c1{color:#998;font-style:italic}
.highlight .cp,.highlight .cs{color:#999;font-weight:bold}
.highlight .cs{font-style:italic}
.highlight .n{color:#333}
.highlight .na,.highlight .nv,.highlight .vc,.highlight .vg,.highlight .vi{color:#008080}
.highlight .nb{color:#0086B3}
.highlight .nc{color:#458;font-weight:bold}
.highlight .no{color:#094e99}
.highlight .ni{color:#800080}
.highlight .ne{color:#990000;font-weight:bold}
.highlight .nf{color:#945277;font-weight:bold}
.highlight .nn{color:#555}
.highlight .nt{color:#000080}
.highlight .err{color:#a61717;background-color:#e3d2d2}
.highlight .gd{color:#000;background-color:#fdd}
.highlight .gd .x{color:#000;background-color:#faa}
.highlight .ge{font-style:italic}
.highlight .gr{color:#aa0000}
.highlight .gh{color:#999}
.highlight .gi{color:#000;background-color:#dfd}
.highlight .gi .x{color:#000;background-color:#afa}
.highlight .go{color:#888}.highlight .gp{color:#555}
.highlight .gs{font-weight:bold}
.highlight .gu{color:#800080;font-weight:bold}
.highlight .gt{color:#aa0000}.highlight .ow{font-weight:bold}
.highlight .w{color:#bbb}.highlight .sr{color:#017936}
.highlight .ss{color:#8b467f}.highlight .bp{color:#999}
.highlight .gc{color:#999;background-color:#EAF2F5}
.type-csharp .highlight .k{color:#0000FF}
.type-csharp .highlight .kt{color:#0000FF}
.type-csharp .highlight .nf{color:#000000;font-weight:normal}
.type-csharp .highlight .nc{color:#2B91AF}
.type-csharp .highlight .nn{color:#000000}
.type-csharp .highlight .s{color:#A31515}
.type-csharp .highlight .sc{color:#A31515}
		</style>
	</head>
<body>
	<article class="markdown-body">
		{{.Body}}
	</article>
</htm>
`
