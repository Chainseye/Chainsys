package conf

var (
	//ROUTERPORT 路由端口号
	ROUTERPORT string
)

var (
	//DBADMIN 数据库用户名
	DBADMIN string
	//DBPASSWORD 数据库密码
	DBPASSWORD string
	//DBCONNECTDB 数据库集合名
	DBCONNECTDB string
	//DBTYPE 数据库类型
	DBTYPE string
)

//InitConfigValue 根据Config文件赋值
func InitConfigValue() {
	ROUTERPORT = Config.String("router::Port")
	DBADMIN = Config.String("db::DBAdmin")
	DBPASSWORD = Config.String("db::DBPassword")
	DBCONNECTDB = Config.String("db::DBConnectDB")
	DBTYPE = Config.String("db::DBType")
}
