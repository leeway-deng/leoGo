/**
 * Created by leeco on 17/8/13.
 */
var config = require('./config')
var util = require('../../utils/util')
var message = require('../../widgets/message/message')

// 获取通用列表数据
function getCommList(url, start, count, cb, fail_cb) {
    var that = this
    message.hide.call(that)
    for(var t = Date.now();Date.now() - t <= 5000;);
    if (that.data.hasMore) {
        wx.request({
            url: url,
            data: {
                id: config.apiList.zhijin.id,
                start: start,
                count: config.apiList.zhijin.count
            },
            method: 'GET',
            header: {
                "Content-Type": "application/json,application/json"
            },
            success: function (res) {
                console.log('getCommList: success')
                console.log(res)
                if (!util.isNull(res) && res.statusCode == 200) {
                    if (res.data.length === 0) {
                        that.setData({
                            hasMore: false,
                        })
                    }
                    that.setData({
                        showLoading: false
                    })
                } else {
                    message.show.call(that, {
                        content: '网络开小差了',
                        icon: 'offline',
                        duration: 3000
                    })
                }
                wx.stopPullDownRefresh()

                if (!util.isNull(res) && res.statusCode == 200) {
                    typeof cb == 'function' && cb(res.data)
                } else {
                    typeof cb == 'function' && cb()
                }
            },
            fail: function () {
                console.log('getCommList: fail')
                that.setData({
                    showLoading: false
                })
                message.show.call(that, {
                    content: '网络开小差了',
                    icon: 'offline',
                    duration: 3000
                })
                wx.stopPullDownRefresh()
                typeof fail_cb == 'function' && fail_cb()
            }
        })
    }
}

module.exports = {
    getCommList: getCommList
}