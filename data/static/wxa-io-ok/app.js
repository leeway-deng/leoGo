//app.js
App({
    globalData: {
        userInfo: null
    },

    onLaunch: function () {
        // 初始化应用
        this.initApp();
        // 初始化缓存
        this.initStorage();
    },

    getUserInfo: function (cb) {
        var that = this;
        if (this.globalData.userInfo) {
            typeof cb == "function" && cb(this.globalData.userInfo);
        } else {
            // 调用登录接口
            wx.getUserInfo({
                withCredentials: false,
                success: function (res) {
                    that.globalData.userInfo = res.userInfo
                    typeof cb == "function" && cb(that.globalData.userInfo)
                }
            });
        }
    },

    initApp: function () {
        // 调用API从本地缓存中获取数据
        var logs = wx.getStorageSync('logs') || [];
        logs.unshift(Date.now());
        wx.setStorageSync('logs', logs);

        // 获取用户信息
        this.getUserInfo();
    },

    initStorage: function () {
        wx.getStorageInfo({
            success: function (res) {
                // 判断知金报表是否存在，没有则创建
                if (!('zhijin_report' in res.keys)) {
                    wx.setStorage({
                        key: 'zhijin_report',
                        data: []
                    })
                }
                // 个人信息默认数据
                var userInfo = {
                    name: '',
                    nickName: '',
                    gender: '',
                    age: '',
                    birthday: '',
                    company: '',
                    tel: '',
                    email: '',
                    intro: ''
                }
                // 判断个人信息是否存在，没有则创建
                if (!('user_info' in res.keys)) {
                    wx.setStorage({
                        key: 'user_info',
                        data: userInfo
                    })
                }
            }
        });
    }
});
