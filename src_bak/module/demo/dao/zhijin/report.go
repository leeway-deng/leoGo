package zhijin

import (
	"encoding/json"
	"strings"

	"github.com/henrylee2cn/faygo"
	"github.com/henrylee2cn/faygo/ext/db/directsql"
)

func GetDataBySqlId4Map(modelId, sqlId string, params faygo.Map) ([]map[string]interface{}, error) {
	//执行sql获取结果
	result, err := directsql.SelectMapToMap(modelId, sqlId, params)
	if err != nil {
		faygo.Error(err.Error())
		return nil, err
	}
	faygo.Debug("GetDataBySqlId4Map result:", result)
	return result, nil
}

/**
单个查询 参数 map[string]interface{} 返回 []map[string]interface{}
tpl
   {% for o in GetDataBySqlId("biz/zhijin","select_tbProvince","{`para`:`sql参数值`}") %}
       <p>名称:{{ o.cnname }}</p>
   {% endfor %}
 */
func GetDataBySqlId(modelId, sqlId string, params string) ([]map[string]interface{}, error) {
	//参数处理
	para := make(faygo.Map)
	if len(params) > 0 {
		if err := json.Unmarshal([]byte(strings.Replace(params, "`", `"`, -1)), &para); err != nil {
			faygo.Error(err.Error())
			return nil, err
		}
	}
	faygo.Debug("GetDataBySqlId params:", para)
	return GetDataBySqlId4Map(modelId, sqlId, para)
}
