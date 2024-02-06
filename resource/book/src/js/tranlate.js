var head= document.getElementsByTagName('head')[0]; 
var script= document.createElement('script');
script.type= 'text/javascript';
script.src= 'https://res.zvo.cn/translate/inspector_v2.js';
script.onload = script.onreadystatechange = function() {
	translate.localLanguage = "en";
	translate.selectLanguageTag.languages = 'zh-CN,zh-TW,en,ko';
	document.getElementById('translate').style.zIndex = '9999999999999';
	translate.execute();
}
head.appendChild(script);
