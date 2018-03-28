;
(function($, window, document, undefined) {

    // 定义构造函数
    var ChainsWebSocket = function(ele, opt) {

        this.options = {
            "wsType": "ws",
            "wsUrl": "127.0.0.1",
            "wsPort": "8099",
            "service": "ws",
            "token": "000000",
            "openFunc": function() {
                console.log("WebSocket Loading...");
            },
            "closeFunc": function(e) {
                console.log("WebSocket: " + e.code + " Connection closed")
            },
            "messageFunc": function(e) {
                console.log("Get Info: " + e.code)
            },
            "inputContainer": "#wsInput",
            "buttonContainer": "#wsButton",
            "targetContainer": "#pre",
            "trigger": "click"
        };
        this.wsItem = null;
        this.options = $.extend({}, this.options, opt)
    }

    // 定义方法
    ChainsWebSocket.prototype = {
        init: function() {
            var that = this;
            var wsLink = that.options.wsType + "://" + that.options.wsUrl + ":" + that.options.wsPort + "/" + that.options.service;
            that.wsItem = new WebSocket(wsLink);

            that.wsItem.onmessage = function(msg) {
                that.displayMessage(msg.data);
            }
            that.wsItem.onopen = function(evt) {
                that.bindSendEvent();
            }
        },
        bindSendEvent: function() {
            var that = this;
            $(that.options.inputContainer).on("keyup", function(e) {
                if (e.keyCode == 13) {
                    that.sendMessage(e.target.value);
                }
            })
            $(that.options.buttonContainer).on(that.options.trigger, function(e) {
                that.sendMessage($(that.options.inputContainer).val());
            });
        },
        sendMessage: function(value) {
            var that = this;
            if (typeof value !== "undefined" && value != "") {
                that.wsItem.send(that.options.token + value);
                $(that.options.inputContainer).val("").focus();
            }
        },
        displayMessage: function(value) {
            var that = this;
            var token = value.slice(0, 6);
            var message = value.slice(6);
            var wClass = (that.options.token == token) ? "wMessage wRight" : "wMessage wLeft";
            $(that.options.targetContainer).append("<div class='" + wClass + "'><div>" + token + "</div><div style='color: #" + token + ";'>" + message + "</div></div>\n");
        }
    };

    // 使用对象
    $.fn.ChainsWebSocket = function(options) {
        // 创建实体
        var chainsWebSocket = new ChainsWebSocket(this, options);

        // 调用初始化方法
        return chainsWebSocket.init();
    }
})(jQuery, window, document);