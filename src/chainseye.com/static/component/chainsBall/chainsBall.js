;(function($, window, document, undefined) {

    // 定义构造函数
    var ChainsBall = function(ele, opt) {
        this.$ele = ele;
        this.options = {
            "container": "container"
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
            var scene,
                camera,
                light,
                renderer,
                earthObject; 
            var WIDTH = 900,
                HEIGHT = 400;
            var angle = 75,
                aspect = WIDTH / HEIGHT,
                near = 0.1,
                far = 3000;
            var container = document.getElementById(that.options.container);
            
            $(window).resize(function() {
                renderer.setSize(window.innerWidth, window.innerHeight);
            });

            //场景
            renderer = new THREE.WebGLRenderer({
                antialiasing : true //抗锯齿
            });
            renderer.setSize(WIDTH, HEIGHT);
            renderer.domElement.style.position = 'relative';
            container.appendChild(renderer.domElement);
            renderer.autoClear = false;
            //打开渲染器的地图阴影
            renderer.shadowMapEnabled = false;

            //摄像仪
            camera = new THREE.PerspectiveCamera(angle, aspect, near, far);

            camera.position.set(0, 0, 0);
            scene = new THREE.Scene();
            
            // 设置聚光灯光源
            // light = new THREE.SpotLight(0xffffff, 1, 0, Math.PI / 2, 1);
            // light.position.set(4000, 4000, 1500);
            // light.target.position.set(1000, 3800, 1000);
            // scene.add(light);
            
            // 设置平行光源
            var directionalLight = new THREE.DirectionalLight(0xffffff, 1.0, 0);        
            directionalLight.position.set( 0, 0, -10000 );                             
            scene.add(directionalLight);

            //test空壳
            var ball = new THREE.SphereGeometry(35, 40, 400);
            var ballMaterial = new THREE.MeshPhongMaterial({
                transparent: true,
                color: 0xffffff,
                opacity: 0.2
            });
            var cubeBall = new THREE.Mesh(ball, ballMaterial);
            cubeBall.position.set(-100, 0, 0); 
            cubeBall.rotation.x = 0.2;
            scene.add(cubeBall);
            
            //地球
            var earthGeo = new THREE.SphereGeometry (30, 40, 400);
            //                   Phong(冯氏)材质类型 
            //          对光照也有反应，用于创建金属类明亮的物体
            //                           ||
            //                           Ｖ
            var earthMat = new THREE.MeshPhongMaterial({
                map: THREE.ImageUtils.loadTexture('/static/skin/image/earth.jpg'),
                bumpMap: THREE.ImageUtils.loadTexture('/static/skin/image/bump-map.jpg'),
                bumpScale: 10,
                opacity: 0.47
                // color: 0xff3220,
                // ambient：设置材质的环境色，和AmbientLight光源一起使用，这个颜色会与环境光的颜色相乘。即是对光源作出反应。
　　            // emissive：设置材质发射的颜色，不是一种光源，而是一种不受光照影响的颜色。默认为黑色
　　            // specular：指定该材质的光亮程度及其高光部分的颜色，如果设置成和color属性相同的颜色，则会得到另一个更加类似金属的材质，如果设置成grey灰色，则看起来像塑料
　　            // shininess：指定高光部分的亮度，默认值为30
            }); 
            var earthMesh = new THREE.Mesh(earthGeo, earthMat); 
            earthMesh.position.set(-100, 0, 0); 
            earthMesh.rotation.x = 0.2;
            scene.add(earthMesh);

            camera.lookAt(earthMesh.position);
            
            //STARS
            var starGeo = new THREE.SphereGeometry (3000, 10, 100),
                starMat = new THREE.MeshBasicMaterial();
            starMat.map = THREE.ImageUtils.loadTexture('/static/skin/image/star.png');
            starMat.side = THREE.BackSide;            
            var starMesh = new THREE.Mesh(starGeo, starMat);
                        
            scene.add(starMesh);

            var render = function () {
                requestAnimationFrame( render );
                //earthMesh.rotation.x += 0.01;
                earthMesh.rotation.y += 0.005;

                renderer.render(scene, camera);
            };

            render();
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