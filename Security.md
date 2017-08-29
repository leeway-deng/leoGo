# [Web-Security-Learning](https://chybeta.github.io/2017/08/19/Web-Security-Learning/)

在学习Web安全的过程中整合的一些资料。
该repo会不断更新，最近更新日期为：2017/8/27。

同步更新于： [chybeta: Web-Security-Learning (带目录) ](https://chybeta.github.io/2017/08/19/Web-Security-Learning/) 

8月27日更新：
+ 新增[java-web代码审计模块](https://github.com/CHYbeta/Web-Security-Learning#java-web%E4%BB%A3%E7%A0%81%E5%AE%A1%E8%AE%A1)
+ 整合[php代码审计模块](https://github.com/CHYbeta/Web-Security-Learning#php%E4%BB%A3%E7%A0%81%E5%AE%A1%E8%AE%A1)
+ 新收录文章：
	+ XSS；
		+ [Xssing Web With Unicodes](http://blog.rakeshmane.com/2017/08/xssing-web-part-2.html)
		+ [XSS Tricks - 从 Self-XSS 到登录你的账户](http://blog.fr1day.me/2016/09/07/XSS_Tricks_from_selfxss_to_your_account/)
		+ [DOMXSS Wiki](https://github.com/wisec/domxsswiki/wiki)
	+ java-web；
		+ [站在Java的视角，深度分析防不胜防的小偷——“XSS” ](http://mp.weixin.qq.com/s?__biz=MzIzMzgxOTQ5NA==&mid=100000340&idx=1&sn=6ca4ec15ef6338daf1d4a907351d7c08&chksm=68fe9e5d5f89174b44fd0cae2e3d5c0018859d3d1dc6d60a2e16dcde34499ba224d6ea17a982#rd)


<!-- more -->
# Web Security

## sql注入

### MySql
+ [MySQL False注入及技巧总结](http://bobao.360.cn/learning/detail/3804.html)
+ [MySQL 注入攻击与防御](http://bobao.360.cn/learning/detail/3758.html)
+ [sql注入学习总结 ](https://mp.weixin.qq.com/s?__biz=MzI5MDQ2NjExOQ==&mid=2247484372&idx=1&sn=ffcc51a88c9acf96c312421b75fc2a26&chksm=ec1e33fcdb69baea53838fd545a236c0deb8a42f3b341ee0879c9e4ac9427c2147fab95b6669#rd)
+ [SQL注入防御与绕过的几种姿势](http://bobao.360.cn/learning/detail/3801.html)
+ [MySQL偏门技巧](http://rcoil.me/2017/05/MySQL%E5%81%8F%E9%97%A8%E6%8A%80%E5%B7%A7/)
+ [mysql注入可报错时爆表名、字段名、库名](http://www.wupco.cn/?p=4117)
+ [高级SQL注入:混淆和绕过](http://www.cnblogs.com/croot/p/3450262.html)

### MSSQL
+ [MSSQL DBA权限获取WEBSHELL的过程 ](http://fuping.site/2017/05/16/MSSQL-DBA-Permission-GET-WEBSHELL/)
+ [MSSQL 注入攻击与防御](http://bobao.360.cn/learning/detail/3807.html)
+ [CLR在SQL Server中的利用技术分析](http://docs.ioin.in/writeup/cert.360.cn/_files_CLR_E5_9C_A8SQL_20Server_E4_B8_AD_E7_9A_84_E5_88_A9_E7_94_A8_E6_8A_80_E6_9C_AF_E5_88_86_E6_9E_90_pdf/index.pdf)

### PostgreSQL
+ [postgresql数据库利用方式 ](https://mp.weixin.qq.com/s?__biz=MzI5MDQ2NjExOQ==&mid=2247484788&idx=1&sn=8a53b1c64d864cd01bab095d97a17715&chksm=ec1e355cdb69bc4a2535bc1a053bfde3ec1838d03936ba8e44156818e91bbec9b5b04a744005#rd)
+ [PostgreSQL渗透测试指南](http://bobao.360.cn/learning/detail/4135.html)

### MongoDB
+ [MongoDB安全 – PHP注入检测](http://www.mottoin.com/94341.html)

### 技巧
+ [我的WafBypass之道（SQL注入篇）](https://xianzhi.aliyun.com/forum/read/349.html)
+ [Bypass 360主机卫士SQL注入防御](http://www.cnblogs.com/xiaozi/p/7275134.html)
+ [SQL注入之骚姿势小记](https://mp.weixin.qq.com/s/ORsciwsBGQJhFdKqceprSw)
+ [CTF比赛中SQL注入的一些经验总结 ](http://www.freebuf.com/articles/web/137094.html)
+ [如何绕过WAF/NGWAF的libinjection实现SQL注入](http://bobao.360.cn/learning/detail/3855.html)
+ [HackMe-SQL-Injection-Challenges
](https://github.com/breakthenet/HackMe-SQL-Injection-Challenges)
+ [绕过WAF注入](https://bbs.ichunqiu.com/thread-25397-1-1.html?from=sec)
+ [bypassGET和POST的注入防御思路分享](https://bbs.ichunqiu.com/thread-16134-1-1.html?from=sec)
+ [SQL注入的常规思路及奇葩技巧 ](https://mp.weixin.qq.com/s/hBkJ1M6LRgssNyQyati1ng)
+ [Beyond SQLi: Obfuscate and Bypass](https://www.exploit-db.com/papers/17934/)


### 工具
+ [使用burp macros和sqlmap绕过csrf防护进行sql注入](http://bobao.360.cn/learning/detail/3557.html)
+ [sqlmap 使用总结 ](http://www.zerokeeper.com/web-security/sqlmap-usage-summary.html)
+ [SQLmap tamper脚本注释](http://www.lengbaikai.net/?p=110)
+ [通过Burp以及自定义的Sqlmap Tamper进行二次SQL注入](http://www.4hou.com/system/6945.html)

## 前端安全
### XSS
+ [漫谈同源策略攻防](http://bobao.360.cn/learning/detail/3848.html)
+ [再谈同源策略 ](https://lightless.me/archives/review-SOP.html)
+ [跨域方法总结](https://xianzhi.aliyun.com/forum/read/1404.html?fpage=4)
+ [浅谈跨站脚本攻击与防御 ](http://thief.one/2017/05/31/1/)
+ [跨站的艺术-XSS入门与介绍](http://www.fooying.com/the-art-of-xss-1-introduction/)
+ [Content Security Policy 入门教程](https://jaq.alibaba.com/community/art/show?spm=a313e.7916646.24000001.49.ZP8rXN&articleid=518)
+ [LoRexxar-CSP](http://lorexxar.cn/tags/csp/)
+ [XSS小记](https://xianzhi.aliyun.com/forum/read/196.html?fpage=7)
+ [DOMXSS Wiki](https://github.com/wisec/domxsswiki/wiki)
+ [XSS Bypass Cookbook](https://xianzhi.aliyun.com/forum/read/536.html?fpage=7)
+ [从瑞士军刀到变形金刚--XSS攻击面拓展](https://xianzhi.aliyun.com/forum/read/1988.html)
+ [我们要在任何可能的地方测试XSS漏洞 ](http://www.freebuf.com/vuls/142751.html)
+ [Alternative to Javascript Pseudo-Protocol](http://brutelogic.com.br/blog/alternative-javascript-pseudo-protocol/)
+ [Bypassing CSP using polyglot JPEGs ](http://blog.portswigger.net/2016/12/bypassing-csp-using-polyglot-jpegs.html)
+ [Bypass unsafe-inline mode CSP](http://paper.seebug.org/91/)
+ [Chrome XSS Auditor – SVG Bypass](https://brutelogic.com.br/blog/chrome-xss-auditor-svg-bypass/)
+ [Cross site scripting payload for fuzzing](https://xianzhi.aliyun.com/forum/read/1704.html)
+ [不常见的xss利用探索](http://docs.ioin.in/writeup/wps2015.org/_2016_06_27__E4_B8_8D_E5_B8_B8_E8_A7_81_E7_9A_84xss_E5_88_A9_E7_94_A8_E6_8E_A2_E7_B4_A2_/index.html)
+ [XSS攻击另类玩法](https://bbs.ichunqiu.com/thread-25578-1-1.html?from=sec)
+ [XSS易容术---bypass之编码混淆篇+辅助脚本编写](https://bbs.ichunqiu.com/thread-17500-1-1.html?from=sec)
+ [Xssing Web With Unicodes](http://blog.rakeshmane.com/2017/08/xssing-web-part-2.html)
+ [Chrome 是怎么过滤反射型 XSS 的呢？](https://www.zhihu.com/question/20941818/answer/180842222?utm_source=qq&utm_medium=social)
+ [XSS Tricks - 从 Self-XSS 到登录你的账户](http://blog.fr1day.me/2016/09/07/XSS_Tricks_from_selfxss_to_your_account/)
+ [Electron hack —— 跨平台 XSS ](https://mp.weixin.qq.com/s?__biz=MzU2NjE2NjIxNg==&mid=2247483756&amp;idx=1&amp;sn=96ae19e53426d5088718b6d37996e700&source=41#wechat_redirect)
+ [XSS without HTML: Client-Side Template Injection with AngularJS ](http://blog.portswigger.net/2016/01/xss-without-html-client-side-template.html)
+ [一个URL跳转引发的一系列“惨案” ](https://zhchbin.github.io/2016/04/09/Problems-Caused-by-URL-Redirection/)
+ [利用反射型XSS二次注入绕过CSP form-action限制](http://blog.neargle.com/SecNewsBak/drops/%E5%88%A9%E7%94%A8%E5%8F%8D%E5%B0%84%E5%9E%8BXSS%E4%BA%8C%E6%AC%A1%E6%B3%A8%E5%85%A5%E7%BB%95%E8%BF%87CSP%20form-action%E9%99%90%E5%88%B6.html)
+ [看我如何挖到GoogleMaps XSS漏洞并获得5000刀赏金](http://bobao.360.cn/learning/detail/3592.html)
+ [Modern Alchemy: Turning XSS into RCE](https://blog.doyensec.com/2017/08/03/electron-framework-security.html)

### CSRF
+ [CRLF Injection and Bypass Tencent WAF ](https://zhchbin.github.io/2016/01/31/CRLF-Injection-and-Bypass-WAF/)
+ [用代码来细说Csrf漏洞危害以及防御](https://bbs.ichunqiu.com/thread-24127-1-1.html?from=sec)
+ [Cookie-Form型CSRF防御机制的不足与反思](https://www.leavesongs.com/PENETRATION/think-about-cookie-form-csrf-protected.html)
+ [关于JSON CSRF的一些思考](https://mp.weixin.qq.com/s?__biz=MzIzMTc1MjExOQ==&mid=2247484126&idx=1&sn=f437882b19bed8d99d0a00938accc0c8&chksm=e89e2a06dfe9a310506419467ada63bee80f10c32267d0b11ea7d1f5491c5afdb344c5dac74e&mpshare=1&scene=23&srcid=0614BOCQBHPjaS2IOtADI3PP#rd)

### 其他
+ [当代 Web 的 JSON 劫持技巧](http://paper.seebug.org/130/?from=timeline&isappinstalled=0)

## SSRF
+ [Build Your SSRF Exploit Framework SSRF](http://docs.ioin.in/writeup/fuzz.wuyun.org/_src_build_your_ssrf_exp_autowork_pdf/index.pdf)
+ [SSRF攻击实例解析](http://www.freebuf.com/articles/web/20407.html)
+ [SSRF漏洞分析与利用](http://www.4o4notfound.org/index.php/author/1/)
+ [SSRF漏洞的挖掘经验](https://www.secpulse.com/archives/4747.html)
+ [SSRF漏洞的利用与学习](http://uknowsec.cn/posts/notes/SSRF%E6%BC%8F%E6%B4%9E%E7%9A%84%E5%88%A9%E7%94%A8%E4%B8%8E%E5%AD%A6%E4%B9%A0.html)
+ [SSRF漏洞中绕过IP限制的几种方法总结](http://www.freebuf.com/articles/web/135342.html)
+ [利用ssrf漏洞获取google内部的dns信息](http://bobao.360.cn/learning/detail/3566.html)
+ [What is Server Side Request Forgery (SSRF)?](https://www.acunetix.com/blog/articles/server-side-request-forgery-vulnerability/)
+ [DNS Rebinding技术绕过SSRF/代理IP限制](http://www.mottoin.com/95734.html)
+ [Discuz ssrf漏洞利用的几个python脚本](https://phpinfo.me/2017/02/23/1438.html)
+ [Discuz X系列门户文章功能SSRF漏洞挖掘与分析](http://bobao.360.cn/learning/detail/2889.html)
+ [SSRF to GET SHELL](http://blog.feei.cn/ssrf/)
+ [SSRF Tips](http://blog.safebuff.com/2016/07/03/SSRF-Tips/)


## XXE
+ [浅谈XXE漏洞攻击与防御](http://thief.one/2017/06/20/1/)
+ [XXE漏洞分析](http://www.4o4notfound.org/index.php/archives/29/)
+ [XML实体注入漏洞攻与防](http://www.hackersb.cn/hacker/211.html)
+ [XML实体注入漏洞的利用与学习](http://uknowsec.cn/posts/notes/XML%E5%AE%9E%E4%BD%93%E6%B3%A8%E5%85%A5%E6%BC%8F%E6%B4%9E%E7%9A%84%E5%88%A9%E7%94%A8%E4%B8%8E%E5%AD%A6%E4%B9%A0.html)
+ [XXE注入:攻击与防御 - XXE Injection: Attack and Prevent](http://le4f.net/post/xxe-injection-attack_and_prevent)
+ [XXE (XML External Entity Injection) 漏洞实践](http://www.mottoin.com/101806.html)
+ [黑夜的猎杀-盲打XXE](https://xianzhi.aliyun.com/forum/read/1837.html)
+ [Hunting in the Dark - Blind XXE](https://blog.zsec.uk/blind-xxe-learning/)
+ [XMLExternal Entity漏洞培训模块](https://www.sans.org/freading-room/whitepapers/application/hands-on-xml-external-entity-vulnerability-training-module-34397)
+ [如何挖掘Uber网站的XXE注入漏洞](http://www.mottoin.com/86853.html)
+ [XXE被提起时我们会想到什么](http://www.mottoin.com/88085.html)
+ [XXE漏洞的简单理解和测试](http://www.mottoin.com/92794.html)
+ [XXE漏洞攻防之我见](http://bobao.360.cn/learning/detail/3841.html)
+ [XXE漏洞利用的一些技巧](http://www.91ri.org/17052.html)
+ [神奇的Content-Type——在JSON中玩转XXE攻击](http://bobao.360.cn/learning/detail/360.html)

## JSONP注入
+ [JSONP注入解析 ](http://www.freebuf.com/articles/web/126347.html)
+ [JSONP 安全攻防技术](http://blog.knownsec.com/2015/03/jsonp_security_technic/)
+ [一次关于JSONP的小实验与总结](http://www.cnblogs.com/vimsk/archive/2013/01/29/2877888.html)
+ [利用JSONP跨域获取信息](https://xianzhi.aliyun.com/forum/read/1571.html)
+ [关于跨域和jsonp的一些理解(新手向)](https://segmentfault.com/a/1190000009577990)

## SSTI
+ [乱弹Flask注入](http://www.freebuf.com/articles/web/88768.html)
+ [服务端模板注入攻击 （SSTI）之浅析 ](http://www.freebuf.com/vuls/83999.html)
+ [Exploring SSTI in Flask/Jinja2](https://nvisium.com/blog/2016/03/09/exploring-ssti-in-flask-jinja2/)
+ [Flask Jinja2开发中遇到的的服务端注入问题研究](http://www.freebuf.com/articles/web/136118.html)
+ [FlaskJinja2 开发中遇到的的服务端注入问题研究 II](http://www.freebuf.com/articles/web/136180.html)
+ [Exploring SSTI in Flask/Jinja2, Part II](https://nvisium.com/blog/2016/03/11/exploring-ssti-in-flask-jinja2-part-ii/)
+ [Injecting Flask](https://nvisium.com/blog/2015/12/07/injecting-flask/)
+ [Server-Side Template Injection: RCE for the modern webapp](https://www.blackhat.com/docs/us-15/materials/us-15-Kettle-Server-Side-Template-Injection-RCE-For-The-Modern-Web-App-wp.pdf)
+ [Exploiting Python Code Injection in Web Applications](https://sethsec.blogspot.jp/2016/11/exploiting-python-code-injection-in-web.html)
+ [利用 Python 特性在 Jinja2 模板中执行任意代码](http://rickgray.me/2016/02/24/use-python-features-to-execute-arbitrary-codes-in-jinja2-templates.html)
+ [Python 模板字符串与模板注入](https://virusdefender.net/index.php/archives/761/)


## 代码执行
+ [PHP Code Injection Analysis](http://www.polaris-lab.com/index.php/archives/254/)
+ [	利用环境变量LD_PRELOAD来绕过php disable_function执行系统命令](http://doc.ph0en1x.com/wooyun_drops/%E5%88%A9%E7%94%A8%E7%8E%AF%E5%A2%83%E5%8F%98%E9%87%8FLD_PRELOAD%E6%9D%A5%E7%BB%95%E8%BF%87php%20disable_function%E6%89%A7%E8%A1%8C%E7%B3%BB%E7%BB%9F%E5%91%BD%E4%BB%A4.html)
+ [Hack PHP mail additional_parameters](http://blog.nsfocus.net/hack-php-mail-additional_parameters/)
+ [详细解析PHP mail()函数漏洞利用技巧](http://bobao.360.cn/learning/detail/3818.html)
+ [在PHP应用程序开发中不正当使用mail()函数引发的血案](http://bobao.360.cn/learning/detail/3809.html)
+ [BigTree CMS - Bypass CSRF filter and execute code with PHPMailer](https://www.cdxy.me/?p=765)
+ [基于时间反馈的RCE](http://www.mottoin.com/97678.html)
+ [正则表达式使用不当引发的系统命令执行漏洞](http://bobao.360.cn/learning/detail/3609.html)

## 文件包含
+ [Turning LFI into RFI](https://l.avala.mp/?p=241)
+ [PHP文件包含漏洞总结](http://wooyun.jozxing.cc/static/drops/tips-3827.html)
+ [常见文件包含发生场景与防御](http://bobao.360.cn/learning/detail/3873.html)
+ [基于云端的本地文件包含漏洞](http://bobao.360.cn/learning/detail/3871.html)
+ [zip或phar协议包含文件](https://bl4ck.in/tricks/2015/06/10/zip%E6%88%96phar%E5%8D%8F%E8%AE%AE%E5%8C%85%E5%90%AB%E6%96%87%E4%BB%B6.html)
+ [文件包含漏洞 一](http://drops.blbana.cc/2016/08/12/e6-96-87-e4-bb-b6-e5-8c-85-e5-90-ab-e6-bc-8f-e6-b4-9e/)
+ [文件包含漏洞 二](http://drops.blbana.cc/2016/12/03/e6-96-87-e4-bb-b6-e5-8c-85-e5-90-ab-e6-bc-8f-e6-b4-9e-ef-bc-88-e4-ba-8c-ef-bc-89/)


## 文件上传 / 解析漏洞
+ [我的WafBypass之道（upload篇）](https://xianzhi.aliyun.com/forum/read/458.html)
+ [文件上传漏洞（绕过姿势） ](http://thief.one/2016/09/22/%E4%B8%8A%E4%BC%A0%E6%9C%A8%E9%A9%AC%E5%A7%BF%E5%8A%BF%E6%B1%87%E6%80%BB-%E6%AC%A2%E8%BF%8E%E8%A1%A5%E5%85%85/)
+ [服务器解析漏洞 ](http://thief.one/2016/09/21/%E6%9C%8D%E5%8A%A1%E5%99%A8%E8%A7%A3%E6%9E%90%E6%BC%8F%E6%B4%9E/)
+ [文件上传总结 ](https://masterxsec.github.io/2017/04/26/%E6%96%87%E4%BB%B6%E4%B8%8A%E4%BC%A0%E6%80%BB%E7%BB%93/)
+ [文件上传绕过姿势总结](http://www.cnnetarmy.com/%E6%96%87%E4%BB%B6%E4%B8%8A%E4%BC%A0%E7%BB%95%E8%BF%87%E5%A7%BF%E5%8A%BF%E6%80%BB%E7%BB%93/)
+ [尽最大可能分析上传源码及漏洞利用方式](https://www.hackfun.org/pentest/make-the-most-possible-analysis-of-the-source-code-and-exploit-the-vulnerability.html)
+ [从XSSer的角度测试上传文件功能](https://xianzhi.aliyun.com/forum/read/224.html)
+ [代码审计之逻辑上传漏洞挖掘](http://wooyun.jozxing.cc/static/drops/papers-1957.html)
+ [渗透测试方法论之文件上传](https://bbs.ichunqiu.com/thread-23193-1-1.html?from=sec)

## 逻辑漏洞
+ [代码审计之逻辑上传漏洞挖掘](http://wooyun.jozxing.cc/static/drops/papers-1957.html)
+ [逻辑至上——内含各种酷炫姿势](http://bobao.360.cn/learning/detail/3769.html)
+ [Web安全测试中常见逻辑漏洞解析（实战篇）](http://www.freebuf.com/vuls/112339.html)
+ [逻辑漏洞之密码重置 ](https://mp.weixin.qq.com/s/Lynmqd_ieEoNJ3mmyv9eQQ)
+ [逻辑漏洞之支付漏洞](https://mp.weixin.qq.com/s/w22omfxO8vU6XzixXWmBxg)
+ [逻辑漏洞之越权访问](https://mp.weixin.qq.com/s/ChiXtcrEyQeLkGOkm4PTog)
+ [密码找回逻辑漏洞总结](http://wooyun.jozxing.cc/static/drops/web-5048.html)
+ [一些常见的重置密码漏洞分析整理](http://wooyun.jozxing.cc/static/drops/papers-2035.html)
+ [密码逻辑漏洞小总结](http://docs.ioin.in/writeup/blog.heysec.org/_archives_643/index.html)
+ [漏洞挖掘之逻辑漏洞挖掘](https://bbs.ichunqiu.com/thread-21161-1-1.html)
+ [tom0li: 逻辑漏洞小结](https://tom0li.github.io/2017/07/17/%E9%80%BB%E8%BE%91%E6%BC%8F%E6%B4%9E%E5%B0%8F%E7%BB%93/)

## 其他漏洞
+ [Web之困笔记](http://www.au1ge.xyz/2017/08/09/web%E4%B9%8B%E5%9B%B0%E7%AC%94%E8%AE%B0/)
+ [常见Web源码泄露总结](http://www.mottoin.com/95749.html)
+ [Github信息泄露升级版案例](http://www.ms509.com/?p=718)
+ [Hacking iSCSI](https://ricterz.me/posts/Hacking%20iSCSI)

### RPO(relative path overwrite)
+ [初探 Relative Path Overwrite](https://xianzhi.aliyun.com/forum/read/1527.html?fpage=2)
+ [Detecting and exploiting path-relative stylesheet import (PRSSI) vulnerabilities](http://blog.portswigger.net/2015/02/prssi.html)
+ [RPO](http://www.thespanner.co.uk/2014/03/21/rpo/)
+ [A few RPO exploitation techniques](http://www.mbsd.jp/Whitepaper/rpo.pdf)

### Web Cache
+ [浅析 Web Cache 欺骗攻击](http://bobao.360.cn/learning/detail/3828.html)

### redis
+ [利用redis写webshell](https://www.leavesongs.com/PENETRATION/write-webshell-via-redis-server.html)
+ [Redis 未授权访问配合 SSH key 文件利用分析](http://blog.knownsec.com/2015/11/analysis-of-redis-unauthorized-of-expolit/)
+ [redis未授权访问漏洞利用总结](https://xianzhi.aliyun.com/forum/read/750.html)。

## PHP相关
### 弱类型
+ [从弱类型利用以及对象注入到SQL注入](http://bobao.360.cn/learning/detail/3486.html)
+ [PHP中“＝＝”运算符的安全问题](http://bobao.360.cn/learning/detail/2924.html)
+ [PHP弱类型安全问题总结 ](http://blog.spoock.com/2016/06/25/weakly-typed-security/)
+ [浅谈PHP弱类型安全](http://wooyun.jozxing.cc/static/drops/tips-4483.html)
+ [php比较操作符的安全问题](http://wooyun.jozxing.cc/static/drops/tips-7679.html)

### 随机数问题
+ [Cracking PHP rand()](http://www.sjoerdlangkemper.nl/2016/02/11/cracking-php-rand/)
+ [php里的随机数](http://5alt.me/2017/06/php%E9%87%8C%E7%9A%84%E9%9A%8F%E6%9C%BA%E6%95%B0/)
+ [php_mt_seed - PHP mt_rand() seed cracker](http://www.openwall.com/php_mt_seed/)
+ [The GLIBC random number generator](http://www.mscs.dal.ca/~selinger/random/)
+ [一道伪随机数的CTF题](https://github.com/wonderkun/CTF_web/blob/master/web500-2/writeup.pdf)

### 伪协议
+ [谈一谈php://filter的妙用](www.leavesongs.com/PENETRATION/php-filter-magic.html)
+ [php 伪协议](http://lorexxar.cn/2016/09/14/php-wei/)
+ [利用 Gopher 协议拓展攻击面](https://blog.chaitin.cn/gopher-attack-surfaces/)
+ [PHP伪协议之 Phar 协议（绕过包含）](https://www.bodkin.ren/?p=902)
+ [PHP伪协议分析与应用](http://www.4o4notfound.org/index.php/archives/31/)
+ [LFI、RFI、PHP封装协议安全问题学习](http://www.cnblogs.com/LittleHann/p/3665062.html)

### 序列化
+ [PHP反序列化漏洞](bobao.360.cn/learning/detail/4122.html)
+ [浅谈php反序列化漏洞 ](https://chybeta.github.io/2017/06/17/%E6%B5%85%E8%B0%88php%E5%8F%8D%E5%BA%8F%E5%88%97%E5%8C%96%E6%BC%8F%E6%B4%9E/)
+ [PHP反序列化漏洞成因及漏洞挖掘技巧与案例](http://bobao.360.cn/learning/detail/3193.html)

### php mail header injection
+ [What is Email Header Injection?](https://www.acunetix.com/blog/articles/email-header-injection/)
+ [PHP Email Injection Example](http://resources.infosecinstitute.com/email-injection/)

### 其他
+ [php & apache2 &操作系统之间的一些黑魔法](http://wonderkun.cc/index.html/?p=626)
+ [php内存破坏漏洞exp编写和禁用函数绕过](http://blog.th3s3v3n.xyz/2016/05/01/bin/2016-5-1-php%E5%86%85%E5%AD%98%E7%A0%B4%E5%9D%8F%E6%BC%8F%E6%B4%9Eexp%E7%BC%96%E5%86%99%E5%92%8C%E7%A6%81%E7%94%A8%E5%87%BD%E6%95%B0%E7%BB%95%E8%BF%87/)
+ [挖掘PHP禁用函数绕过利用姿势](http://blog.th3s3v3n.xyz/2016/11/20/web/%E6%8C%96%E6%8E%98PHP%E7%A6%81%E7%94%A8%E5%87%BD%E6%95%B0%E7%BB%95%E8%BF%87%E5%88%A9%E7%94%A8%E5%A7%BF%E5%8A%BF/)
+ [.user.ini文件构成的PHP后门](http://wooyun.jozxing.cc/static/drops/tips-3424.html)

### php代码审计
+ [论PHP常见的漏洞](http://wooyun.jozxing.cc/static/drops/papers-4544.html)
+ [浅谈代码审计入门实战：某博客系统最新版审计之旅 ](http://www.freebuf.com/articles/rookie/143554.html)
+ [ctf中的php代码审计技巧](http://www.am0s.com/ctf/200.html)
+ [PHP代码审计tips](http://docs.ioin.in/writeup/www.91ri.org/_15074_html/index.html)
+ [代码审计之文件越权和文件上传搜索技巧](http://docs.ioin.in/writeup/blog.heysec.org/_archives_170/index.html)
+ [PHP代码审计入门集合](http://wiki.ioin.in/post/group/6Rb)
+ [PHP代码审计学习](http://phantom0301.cc/2017/06/06/codeaudit/)
+ [PHP漏洞挖掘思路+实例](http://wooyun.jozxing.cc/static/drops/tips-838.html)
+ [PHP漏洞挖掘思路+实例 第二章](http://wooyun.jozxing.cc/static/drops/tips-858.html)
+ [浅谈代码审计入门实战：某博客系统最新版审计之旅 ](http://www.freebuf.com/articles/rookie/143554.html)

## java-Web
### 反序列
+ [如何攻击Java反序列化过程](http://bobao.360.cn/learning/detail/4267.html)
+ [深入理解JAVA反序列化漏洞](https://www.vulbox.com/knowledge/detail/?id=11)
+ [Attacking Java Deserialization](https://nickbloor.co.uk/2017/08/13/attacking-java-deserialization/)
+ [jackson反序列化详细分析](http://bobao.360.cn/learning/detail/4118.html)

### Struct2
+ [Struts2 命令执行系列回顾](http://www.zerokeeper.com/vul-analysis/struts2-command-execution-series-review.html)

### java-Web代码审计
+ [Java代码审计连载之—SQL注入](https://bbs.ichunqiu.com/forum.php?mod=viewthread&tid=22170&highlight=Java%E4%BB%A3%E7%A0%81%E5%AE%A1%E8%AE%A1%E8%BF%9E%E8%BD%BD)
+ [Java代码审计连载之—任意文件下载](https://bbs.ichunqiu.com/forum.php?mod=viewthread&tid=23587&highlight=Java%E4%BB%A3%E7%A0%81%E5%AE%A1%E8%AE%A1%E8%BF%9E%E8%BD%BD)
+ [Java代码审计连载之—XSS](https://bbs.ichunqiu.com/forum.php?mod=viewthread&tid=22875&highlight=Java%E4%BB%A3%E7%A0%81%E5%AE%A1%E8%AE%A1%E8%BF%9E%E8%BD%BD)
+ [Java代码审计连载之—添油加醋](https://bbs.ichunqiu.com/forum.php?mod=viewthread&tid=25475&highlight=Java%E4%BB%A3%E7%A0%81%E5%AE%A1%E8%AE%A1%E8%BF%9E%E8%BD%BD)

### 其他
+ [层层放大java审计的攻击面 ](https://mp.weixin.qq.com/s/WT1EXEryUGGqHQpSi959xw)
+ [以Java的视角来聊聊SQL注入 ](https://mp.weixin.qq.com/s?__biz=MzIzMzgxOTQ5NA==&mid=2247483954&idx=1&sn=418b7e55b16c717ee5140af990298e22&chksm=e8fe9e3bdf89172d0670690060944bf2434cc2d2e8fba4477711299a0775cf3735a2022c0778#rd)
+ [站在Java的视角，深度分析防不胜防的小偷——“XSS” ](http://mp.weixin.qq.com/s?__biz=MzIzMzgxOTQ5NA==&mid=100000340&idx=1&sn=6ca4ec15ef6338daf1d4a907351d7c08&chksm=68fe9e5d5f89174b44fd0cae2e3d5c0018859d3d1dc6d60a2e16dcde34499ba224d6ea17a982#rd)
+ [你的 Java web 配置安全吗？ ](https://mp.weixin.qq.com/s?__biz=MzIzMzgxOTQ5NA==&mid=100000318&idx=1&sn=9011af3e3968e0d87499605ef1a68291&chksm=68fe9e375f8917213297855bd9e1ab1203ae4c9b0b5ca351de7b2c0f7a7799bd1f4843cd13f4#rd)
+ [Java代码审计PPT ](https://xianzhi.aliyun.com/forum/read/1904.html)
+ [spring任意文件读取](https://github.com/ilmila/springcss-cve-2014-3625/tree/master/src)


## python-Web
+ [Python 格式化字符串漏洞（Django为例）](https://www.leavesongs.com/PENETRATION/python-string-format-vulnerability.html)
+ [format注入](http://www.venenof.com/index.php/archives/360/)
+ [Be Careful with Python's New-Style String Format](http://lucumr.pocoo.org/2016/12/29/careful-with-str-format/)
+ [Python urllib HTTP头注入漏洞](http://www.tuicool.com/articles/2iIj2eR)
+ [Hack Redis via Python urllib HTTP Header Injection](https://security.tencent.com/index.php/blog/msg/106)
+ [Python Waf黑名单过滤下的一些Bypass思路](http://www.0aa.me/index.php/archives/123/)
+ [Python Sandbox Bypass](https://mp.weixin.qq.com/s?__biz=MzIzOTQ5NjUzOQ==&mid=2247483665&idx=1&sn=4b18de09738fdc5291634db1ca2dd55a)
+ [pyt: 针对 Python 应用程序的源码静态分析工具](https://github.com/python-security/pyt)

## Node-js
+ [Pentesting Node.js Application : Nodejs Application Security(需翻墙)](http://www.websecgeeks.com/2017/04/pentesting-nodejs-application-nodejs.html)
+ [从零开始学习渗透Node.js应用程序 ](https://bbs.ichunqiu.com/thread-21810-1-1.html?from=sec)

## WAF相关
+ [牛逼牛逼的payload和bypass总结](https://github.com/swisskyrepo/PayloadsAllTheThings)
+ [WAF绕过参考资料](http://www.mottoin.com/100887.html)
+ [浅谈WAF绕过技巧](http://www.freebuf.com/articles/web/136723.html)
+ [addslashes防注入的绕过案例](https://xianzhi.aliyun.com/forum/read/753.html?fpage=6)
+ [浅谈json参数解析对waf绕过的影响](https://xianzhi.aliyun.com/forum/read/553.html?fpage=8)
+ [WAF攻防研究之四个层次Bypass WAF](http://weibo.com/ttarticle/p/show?id=2309404007261092631700)
+ [使用HTTP头去绕过WAF ](http://www.sohu.com/a/110066439_468673)
+ [会找漏洞的时光机: Pinpointing Vulnerabilities](https://www.inforsec.org/wp/?p=1993)




# 渗透测试
## Course
+ [Web Service 渗透测试从入门到精通](http://bobao.360.cn/learning/detail/3741.html)
+ [渗透标准](https://www.processon.com/view/583e8834e4b08e31357bb727)
+ [Penetration Testing Tools Cheat Sheet](https://highon.coffee/blog/penetration-testing-tools-cheat-sheet/)

## 信息收集
+ [看我如何收集全网IP的whois信息 ](https://mp.weixin.qq.com/s/qz0b42DKhgo1sfitcUKhtQ)
+ [浅谈Web渗透测试中的信息收集 ](http://www.freebuf.com/articles/web/142767.html)
+ [渗透测试教程：如何侦查目标以及收集信息？](http://www.4hou.com/penetration/6850.html)
+ [本屌的web漏洞扫描器思路 技巧总结（域名信息收集篇）](weibo.com/ttarticle/p/show?id=2309404088584863883789)
+ [子域名的艺术](http://www.91ri.org/17001.html)
+ [实例演示如何科学的进行子域名收集](http://bobao.360.cn/learning/detail/4119.html)
+ [【渗透神器系列】搜索引擎 ](http://thief.one/2017/05/19/1/)
+ [域渗透基础简单信息收集（基础篇）](https://xianzhi.aliyun.com/forum/read/805.html)
+ [内网渗透定位技术总结](http://docs.ioin.in/writeup/www.mottoin.com/_92978_html/index.html)
+ [后渗透攻防的信息收集](https://www.secpulse.com/archives/51527.html)
+ [安全攻城师系列文章－敏感信息收集](http://www.mottoin.com/99951.html)
+ [子域名枚举的艺术](http://www.mottoin.com/101362.html)
+ [论二级域名收集的各种姿势](https://mp.weixin.qq.com/s/ardCYdZzaSjvSIZiFraWGA)
+ [我眼中的渗透测试信息搜集](https://xianzhi.aliyun.com/forum/read/451.html?fpage=2)
+ [大型目标渗透－01入侵信息搜集](https://xianzhi.aliyun.com/forum/read/1675.html)

## 渗透
+ [【玩转Linux系统】Linux内网渗透 ](https://mp.weixin.qq.com/s/VJBnXq3--0HBD7eVeifOKA)
+ [渗透测试指南之域用户组的范围](http://www.4hou.com/penetration/7016.html)
+ [内网主机发现技巧补充](http://mp.weixin.qq.com/s/l-Avt72ajCIo5GdMEwVx7A)
+ [Linux 端口转发特征总结 ](https://mp.weixin.qq.com/s?__biz=MzA3Mzk1MDk1NA==&mid=2651903919&idx=1&sn=686cc53137aa9e8ec323dda1e54a2c23)
+ [内网渗透（持续更新） ](http://rcoil.me/2017/06/%E5%86%85%E7%BD%91%E6%B8%97%E9%80%8F/)
+ [实战 SSH 端口转发](https://www.ibm.com/developerworks/cn/linux/l-cn-sshforward/index.html)
+ [多重转发渗透隐藏内网](http://bobao.360.cn/learning/detail/3545.html)
+ [内网转发姿势](http://www.03sec.com/3141.shtml)
+ [内网转发的工具](https://mp.weixin.qq.com/s/EWL9-AUB_bTf7pU4S4A2zg)
+ [Linux 下多种反弹 shell 方法](http://www.03sec.com/3140.shtml)
+ [php 反弹shell](http://wolvez.club/?p=458)
+ [利用ew轻松穿透多级目标内网](https://klionsec.github.io/2017/08/05/ew-tunnel/)
+ [windows内网渗透杂谈](https://bl4ck.in/penetration/2017/03/20/windows%E5%86%85%E7%BD%91%E6%B8%97%E9%80%8F%E6%9D%82%E8%B0%88.html)
+ [Windows域横向渗透](http://docs.ioin.in/writeup/www.mottoin.com/_89413_html/index.html)
+ [内网渗透中转发工具总结](http://blog.neargle.com/SecNewsBak/drops/%E5%86%85%E7%BD%91%E6%B8%97%E9%80%8F%E4%B8%AD%E8%BD%AC%E5%8F%91%E5%B7%A5%E5%85%B7%E6%80%BB%E7%BB%93.html)
+ [内网渗透思路整理与工具使用](http://bobao.360.cn/learning/detail/3683.html)
+ [Cobalt strike在内网渗透中的使用 ](http://www.freebuf.com/sectool/125237.html)
+ [反向socks5代理(windows版)](http://x95.org/archives/reverse-socks5-proxy.html)
+ [Windows渗透基础](http://www.mottoin.com/89355.html)
+ [通过双重跳板漫游隔离内网](https://xianzhi.aliyun.com/forum/read/768.html)
+ [A Red Teamer's guide to pivoting](https://artkond.com/2017/03/23/pivoting-guide/)
+ [穿越边界的姿势 ](https://mp.weixin.qq.com/s/l-0sWU4ijMOQWqRgsWcNFA)
+ [内网端口转发及穿透](https://xianzhi.aliyun.com/forum/read/1715.html)
+ [秘密渗透内网——利用 DNS 建立 VPN 传输隧道](http://www.4hou.com/technology/3143.html)
+ [Reverse Shell Cheat Sheet](http://pentestmonkey.net/cheat-sheet/shells/reverse-shell-cheat-sheet)

## 渗透实战
+ [挖洞经验 | 看我如何综合利用4个漏洞实现GitHub Enterprise远程代码执行 ](http://www.freebuf.com/news/142680.html)
+ [Splash SSRF到获取内网服务器ROOT权限](http://bobao.360.cn/learning/detail/4113.html)
+ [Pivoting from blind SSRF to RCE with HashiCorp Consul](http://www.kernelpicnic.net/2017/05/29/Pivoting-from-blind-SSRF-to-RCE-with-Hashicorp-Consul.html)
+ [我是如何通过命令执行到最终获取内网Root权限的 ](http://www.freebuf.com/articles/web/141579.html)
+ [信息收集之SVN源代码社工获取及渗透实战](https://xianzhi.aliyun.com/forum/read/1629.html)
+ [SQL注入+XXE+文件遍历漏洞组合拳渗透Deutsche Telekom](http://paper.seebug.org/256/)
+ [渗透 Hacking Team](http://blog.neargle.com/SecNewsBak/drops/%E6%B8%97%E9%80%8FHacking%20Team%E8%BF%87%E7%A8%8B.html)
+ [由视频系统SQL注入到服务器权限](https://bbs.ichunqiu.com/thread-25827-1-1.html?from=sec)
+ [From Serialized to Shell :: Exploiting Google Web Toolkit with EL Injection](http://srcincite.io/blog/2017/05/22/from-serialized-to-shell-auditing-google-web-toolkit-with-el-injection.html)
+ [浅谈渗透测试实战](http://docs.ioin.in/writeup/avfisher.win/_archives_381/index.html)
+ [渗透测试学习笔记之案例一](http://avfisher.win/archives/741)
+ [渗透测试学习笔记之案例二](http://avfisher.win/archives/756)


## 提权
+ [提权技巧](http://www.secbox.cn/skill/5583.html)
+ [linux-kernel-exploits Linux平台提权漏洞集合](https://github.com/SecWiki/linux-kernel-exploits)
+ [windows-kernel-exploits Windows平台提权漏洞集合 ](https://github.com/SecWiki/windows-kernel-exploits)
+ [Linux MySQL Udf 提权](http://www.91ri.org/16540.html)
+ [windows提权系列上篇](http://mp.weixin.qq.com/s/uOArxXIfcI4fjqnF9BDJGA)
+ [Windows提权系列中篇](https://mp.weixin.qq.com/s/ERXOLhWo0-lJbMV143I8hA)

## 渗透技巧
+ [域渗透神器Empire安装和简单使用 ](https://mp.weixin.qq.com/s/VqrUTW9z-yi3LqNNy-lE-Q)
+ [如何将简单的Shell转换成为完全交互式的TTY ](http://www.freebuf.com/news/142195.html)
+ [60字节 - 无文件渗透测试实验](https://www.n0tr00t.com/2017/03/09/penetration-test-without-file.html)
+ [内网渗透思路探索之新思路的探索与验证](http://www.tuicool.com/articles/fMFB3mY)
+ [Web端口复用正向后门研究实现与防御 ](http://www.freebuf.com/articles/web/142628.html)
+ [端口渗透总结](http://docs.ioin.in/writeup/blog.heysec.org/_archives_577/index.html)
+ [端口扫描那些事](https://mp.weixin.qq.com/s?__biz=MzI5MDQ2NjExOQ==&mid=2247484812&idx=1&sn=7d894b50b3947142fbfa3a4016f748d5&chksm=ec1e35a4db69bcb2acfe7ecb3b0cd1d366c54bfa1feaafc62c4290b3fd2eddab9aa95a98f041#rd)
+ [渗透技巧——通过cmd上传文件的N种方法 ](http://blog.neargle.com/SecNewsBak/drops/%E6%B8%97%E9%80%8F%E6%8A%80%E5%B7%A7%E2%80%94%E2%80%94%E9%80%9A%E8%BF%87cmd%E4%B8%8A%E4%BC%A0%E6%96%87%E4%BB%B6%E7%9A%84N%E7%A7%8D%E6%96%B9%E6%B3%95.html)
+ [域渗透TIPS：获取LAPS管理员密码 ](http://www.freebuf.com/articles/web/142659.html)
+ [域渗透——Security Support Provider](http://blog.neargle.com/SecNewsBak/drops/%E5%9F%9F%E6%B8%97%E9%80%8F%E2%80%94%E2%80%94Security%20Support%20Provider.html)
+ [内网渗透随想](http://docs.ioin.in/writeup/www.91ri.org/_14390_html/index.html)
+ [域渗透之流量劫持](http://bobao.360.cn/learning/detail/3266.html)
+ [渗透技巧——快捷方式文件的参数隐藏技巧](https://3gstudent.github.io/3gstudent.github.io/%E6%B8%97%E9%80%8F%E6%8A%80%E5%B7%A7-%E5%BF%AB%E6%8D%B7%E6%96%B9%E5%BC%8F%E6%96%87%E4%BB%B6%E7%9A%84%E5%8F%82%E6%95%B0%E9%9A%90%E8%97%8F%E6%8A%80%E5%B7%A7/)
+ [后门整理](https://bbs.ichunqiu.com/thread-25119-1-1.html?from=sec)
+ [Linux后门整理合集（脉搏推荐）](https://www.secpulse.com/archives/59674.html)

## 运维
+ [黑客入侵应急分析手工排查](https://xianzhi.aliyun.com/forum/read/1655.html)
+ [企业常见服务漏洞检测&修复整理](http://www.mottoin.com/92742.html)
+ [Linux基线加固](https://mp.weixin.qq.com/s/0nxiZw1NUoQTjxcd3zl6Zg)

# CTF
## 技巧总结
+ [ctf-wiki](https://ctf-wiki.github.io/ctf-wiki/#/introduction)
+ [CTF中那些脑洞大开的编码和加密](https://www.hackfun.org/CTF/coding-and-encryption-of-those-brain-holes-in-CTF.html)
+ [CTF加密与解密 ](http://thief.one/2017/06/13/1/)
+ [CTF中图片隐藏文件分离方法总结](https://www.hackfun.org/CTF/summary-of-image-hiding-files-in-CTF.html)
+ [Md5扩展攻击的原理和应用](http://www.freebuf.com/articles/database/137129.html)
+ [CTF比赛中关于zip的总结](http://bobao.360.cn/ctf/detail/203.html)
+ [十五个Web狗的CTF出题套路](http://weibo.com/ttarticle/p/show?id=2309403980950244591011)
+ [CTF备忘录](https://827977014.docs.qq.com/Bt2v7IZWnYo?type=1&_wv=1&_bid=2517)
+ [rcoil:CTF线下攻防赛总结](http://rcoil.me/2017/06/CTF%E7%BA%BF%E4%B8%8B%E8%B5%9B%E6%80%BB%E7%BB%93/)

# 杂
+ [谈谈HTML5本地存储——WebStorage](http://syean.cn/2017/08/15/%E8%B0%88%E8%B0%88HTML5%E6%9C%AC%E5%9C%B0%E5%AD%98%E5%82%A8%E2%80%94%E2%80%94WebStorage/)
+ [Linux下容易被忽视的那些命令用法](https://segmentfault.com/p/1210000010668099/read)
+ [各种脚本语言不同版本一句话开启 HTTP 服务器的总结](http://www.mottoin.com/94895.html)
+ [WebAssembly入门：将字节码带入Web世界](http://bobao.360.cn/learning/detail/3757.html)
+ [phpwind 利用哈希长度扩展攻击进行getshell](https://www.leavesongs.com/PENETRATION/phpwind-hash-length-extension-attack.html)
+ [深入理解hash长度扩展攻击（sha1为例） ](http://www.freebuf.com/articles/web/69264.html)
+ [Joomla 框架的程序执行流程及目录结构分析](http://bobao.360.cn/learning/detail/3909.html)
+ [如何通过恶意插件在Atom中植入后门](http://bobao.360.cn/learning/detail/4268.html)