;(function($, window, document, undefined) {

    // 定义构造函数
    var ChainsWebSocket = function(ele, opt) {
    
        this.options = {
            "wsType": "ws",
            "wsUrl": "127.0.0.1",
            "wsPort": "8099",
            "service": "ws",
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
        this.options = $.extend({}, this.options, opt)
    }

    // 定义方法
    ChainsWebSocket.prototype = {
        init: function() {
            var that = this;
            var wsLink = that.options.wsType + "://" + that.options.wsUrl + ":" + that.options.wsPort + "/" + that.options.service;
            var wsItem = new WebSocket(wsLink);
 
            wsItem.onmessage = function(msg) {
                $(that.options.targetContainer).append("<== " + msg.data + "\n");
            }
            wsItem.onopen = function() {
                $(that.options.buttonContainer).on(that.options.trigger, function() {
                    wsItem.send($(that.options.inputContainer).val());
                    $(that.options.targetContainer).append("==> " + $(that.options.inputContainer).val() + "\n");
                });
            }
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