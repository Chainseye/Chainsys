;(function($, window, document, undefined) {

    // 定义构造函数
    var ChainsBar = function(ele, opt) {
        this.$ele = ele;
        this.options = {
            "width": "",
            "height": "20px",
            "maxCount": 5,
            "content": ["news1", "news2", "news3", "news4"],
            "direction": "left"
        };
        this.options = $.extend({}, this.options, opt)
    };

    // 定义方法
    ChainsBar.prototype = {
        init: function() {
            var that = this;
            that.createChainsBarDom();
        },
        createChainsBarDom: function() {
            var that = this;
            var newsItemStr = "";
            that.options.content.forEach(function(val, index) {
                newsItemStr += "<div class='chainsBar-newsItem'><p>" + index + "==>" + val + "</p></div>";
            });
            that.$ele.html("<div class='chainsBar-content'>" + newsItemStr + "</div>");
            that.addBaseChainsBarCss();
        },
        addBaseChainsBarCss: function() {
            var that = this;

        }
    };

    // 使用对象
    $.fn.chainsBar = function(options) {
        // 创建实体
        var chainsBar = new ChainsBar(this, options);

        // 调用初始化方法
        return chainsBar.init();
    }
})(jQuery, window, document);