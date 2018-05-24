package mwHtml

import "regexp"

var reTitle = regexp.MustCompile(`<title>.*</title>`)

//Method que agrega informacion de poca importancia al body, funciona como spam
//agrega console.log saludando y un buscado de google con al busqueda 'Porque no usar un Proxy'
func AddSpamInfo(body *string) {
	*body = reTitle.ReplaceAllString(*body, "<title>Emma Proxy</title><script type='text/javascript'>(function(){console.log('Hello, My Name is Emma')}())</script><center><FORM style=' width: 650px;margin-top:30px;border: 1px solid black;background-color: white;' method=GET action='http://www.google.com/search'><TABLE bgcolor='#FFFFFF'><tr><td><A HREF=' http://www.google.com/'><IMG SRC='http://www.google.com/logos/Logo_40wht.gif' border='0' ALT='Google' align='absmiddle'></A><INPUT TYPE=text name=q size=31 maxlength=255 value='Porque no usar un Proxy =)'><INPUT TYPE=hidden name=hl value=es><INPUT type=submit style='background-color: #0099cc;border: none;color: white;padding: 10px 16px;margin:5;text-align: center;text-decoration: none;display: inline-block;font-size: 12px;' name=btnG VALUE='Búsqueda Google'></td></tr></TABLE></FORM></center>") //strings.Replace(htmlString, "</head>", "<title>Emma Proxy</title></head>", 1)

}
