;(function($, window, document, undefined) {

    // 定义构造函数
    var ChainsBall = function(ele, opt) {
        this.$ele = ele;
        this.options = {
            "container": ""
        };
        this.options = $.extend({}, this.options, opt)
    };

    // 定义方法
    ChainsBall.prototype = {
        init: function() {
            var that = this;
            that.createChainsBall();
        },
        createChainsBall: function() {
            var that = this;
        }
    };

    // 使用对象
    $.fn.chainsBall = function(options) {
        // 创建实体
        var chainsBall = new ChainsBall(this, options);

        // 调用初始化方法
        return chainsBall.init();
    }
})(jQuery, window, document);