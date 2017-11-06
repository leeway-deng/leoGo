#!/bin/sh

g_cur_dir=$(cd "$(dirname "${0}")"; pwd)
g_sessionid='4628C8C055B344C9729395755EE9D2D9'
g_personid='128956'
g_projectid='10756987'
g_scriptid='DDA22089B250350D1433A8B0327E8265280'
g_cookie="JSESSIONID=$g_sessionid"
g_ua='Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36'
g_refer_home="http://www.tjxzxk.gov.cn/jzz/regJfProject.jsp?personid=$g_personid"
g_url_home="http://www.tjxzxk.gov.cn/jzz/tabs.jsp?jfprojectid=$g_projectid&step=mark&personid=$g_personid"
g_url_code='http://www.tjxzxk.gov.cn/jzz/image2.jsp'
g_url_date='http://www.tjxzxk.gov.cn/dwr/call/plaincall/jfxzpersonservice.hasJfXzPerson.dwr'
g_mail_addrs='dengliwei@le.com, wanghui@zhijin.com, 470429617@qq.com, 787431972@qq.com'
g_file_home="$g_cur_dir/tjxzxk_home.html"
g_file_code="$g_cur_dir/tjxzxk_code.jpg"
g_file_date="$g_cur_dir/tjxzxk_date.txt"
g_file_log="$g_cur_dir/tjxzxk.log"
g_dates_str='2017-10-25,WebOrderAm,10;2017-10-25,WebOrderPm,10;2017-10-26,WebOrderAm,10;2017-10-26,WebOrderPm,10'
g_data_log=''
echo -e `date` >> $g_file_log

g_dates=$(echo $g_dates_str | sed 's/[[:space:]]//g' | sed 's/;/ /g' | sed 's/[ \t]*$//g')

curl -s -H "Cookie: $g_cookie" -H "Referer: $g_refer_home" -H "User-Agent: $g_ua" "$g_url_home" > $g_file_home
curl_home_data=$(cat $g_file_home | iconv -f gb2312 | grep '预约名额<br>剩余' | grep -v 0 | sed s/[[:space:]]//g)
#curl_home_dates=$(cat $g_file_home | iconv -f gb2312 | grep -E '\<div class=\"seat2\".*\>' | sed s/[[:space:]]//g)

#sleep 1s

curl_code_data=''
while (true)
    do
    curl -s -H "Cookie: $g_cookie" -H "Referer: $g_refer_home" -H "User-Agent: $g_ua" -H 'Accept-Language: zh-CN,zh;q=0.8' -H 'Content-Type: image/jpeg;charset=GBK' "$g_url_code" > $g_file_code
#curl_code_data=$(/usr/local/bin/tesseract $g_file_code stdout -l eng | sed 's/[[:space:]]//g')
curl_code_data=$(/usr/local/bin/tesseract $g_file_code stdout -l eng -psm 7 -c tessedit_char_whitelist='ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789' | sed 's/[[:space:]]//g')
   #if [ "$curl_code_data" =~ '^[a-zA-Z0-9]\+$' ]; then
   if echo "$curl_code_data" | grep -q '^[a-zA-Z0-9]\+$'; then
       if [ `echo $curl_code_data|wc -L` -eq 4 ]; then
           break
       fi
   fi
done
echo -e "curl_code_data=$curl_code_data\n" >> $g_file_log

if [ -n "$curl_home_data" ]; then
    result_mail="$curl_home_data"
    echo -e "curl_home_data=$curl_home_data" >> $g_file_log
    echo "" > $g_file_date
    if [ -n "$curl_code_data" ]; then
        echo -e "curl_code_data=$curl_code_data" >> $g_file_log
        for date in $g_dates; do
            date_data="callCount=1&page=/dwr/test/jfxzpersonservice&httpSessionId=$g_sessionid&scriptSessionId=$g_scriptid&c0-scriptName=jfxzpersonservice&c0-methodName=hasJfXzPerson&c0-id=0&c0-param0=boolean:false&c0-param1=string:$curl_code_data&c0-param2=string:$date&c0-param3=string:$g_projectid&c0-param4=string:$g_personid&batchId=2"
            curl_date_data=$(curl -s -H "Cookie: $g_cookie" -H "Referer: $g_refer_home" -H "User-Agent: $g_ua" -H 'Content-Type:text/plain' -H 'Origin:http://www.tjxzxk.gov.cn' -H 'Accept-Language:zh-CN,zh;q=0.8' -d "$date_data" "$g_url_date")
            echo "$date<br> $curl_date_data" >> $g_file_date
        done
    fi
    curl_date_data=$(cat $g_file_date |iconv -f gb2312 | sed s/[[:space:]]//g)
    echo -e "curl_date_data=$curl_date_data" >> $g_file_log

    result_mail="$curl_home_data <br> $curl_date_data"
    #echo -e "result_mail=$result_mail" >> $g_file_log
    echo "$result_mail" | mail -s "tjyy: $curl_home_data" $g_mail_addrs