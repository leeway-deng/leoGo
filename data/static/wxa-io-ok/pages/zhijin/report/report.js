var net_apis = require('../../../com/net/api')
var net_config = require('../../../com/net/config')
var util = require('../../../utils/util')

Page({
    data: {
        cells: [],
        start: 1,
        hasMore: true,
        showLoading: true,
        minaVersion: 'v1.0'
    },
    onLoad: function (options) {
        var that = this
        // wx.getSystemInfo({
        //     success: function (res) {
        //         var cells = [[]]
        //         var resolution = res.windowWidth * res.pixelRatio + '*' + res.windowHeight * res.pixelRatio
        //         cells[0].push({title: '手机型号', text: res.model, access: false, fn: ''})
        //         cells[0].push({title: '分辨率', text: resolution, access: false, fn: ''})
        //         cells[0].push({title: '系统语言', text: res.language, access: false, fn: ''})
        //         cells[0].push({title: '微信版本', text: res.version, access: false, fn: ''})
        //         cells[0].push({title: '小程序版本', text: that.data.minaVersion, access: false, fn: ''})
        //         that.setData({
        //             cells: cells
        //         })
        //     }
        // })

        net_apis.getCommList.call(that, net_config.apiList.zhijin.report, that.data.start,
            net_config.apiList.zhijin.count,
            function (res) {
                var cells = [[]]
                if (!util.isNull(res) && !util.isNull(res.data)) {
                    for (var i = 0; i < res.data.length; i++) {
                        for(var key in res.data[i]) {
                            cells[i].push({title: key, text: res.data[i][key], access: false, fn: ''})
                        }
                    }
                }
                that.setData({
                    cells: cells
                })
            })
    }
})