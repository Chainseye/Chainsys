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
            $(that.options.targetContainer).append("<div class='" + wClass + "'>" + that.addCube(token) + "<div style='color: #" + token + ";'>" + message + "</div></div>\n");
        },
        addCube: function(color) {
            return '<div class="cube-wrap" title="' + color + '" style="border-color: #' + color + ';box-shadow:0 0 10px #' + color + '"><div class="cube"><div class="front-pane" style="box-shadow:0 0 10px #' + color + ';"></div><div class="back-pane" style="box-shadow:0 0 10px #' + color + ';"></div><div class="left-pane" style="box-shadow:0 0 10px #' + color + ';"></div><div class="right-pane" style="box-shadow:0 0 10px #' + color + ';"></div><div class="top-pane" style="box-shadow:0 0 10px #' + color + ';"></div><div class="bottom-pane" style="box-shadow:0 0 10px #' + color + ';"></div></div></div>'
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