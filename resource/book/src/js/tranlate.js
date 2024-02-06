var head= document.getElementsByTagName('head')[0]; 
var script= document.createElement('script');
script.type= 'text/javascript';
script.src= 'https://res.zvo.cn/translate/inspector_v2.js';
script.onload = script.onreadystatechange = function() {
	translate.localLanguage = "en";
	translate.selectLanguageTag.languages = 'zh-CN,zh-TW,en,ko';
	translate.execute();
	document.getElementById('translate').style.position = 'fixed';
	document.getElementById('translate').style.color = 'red';
	document.getElementById('translate').style.right = '1px';
	document.getElementById('translate').style.top = '45px';
	document.getElementById('translate').style.fontSize = '14px';
	document.getElementById('translate').style.zIndex = '9999999999999';
}
head.appendChild(script);
