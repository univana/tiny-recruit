## BookStack 项目结构分析

### 服务器启动函数 run()

#### commands.ResolveCommand(d.config.Arguments) ：分析命令行输入参数，进行相关路径配置，注册部分内容

- flag 包：golang编写命令行的包
- 进行命令行相关参数设定
- gocaptcha 验证码相关设置
- beego.LoadAppConfig： 加载 beego 配置文件
- uploads := filepath.Join(WorkingDirectory, "uploads")： 设置上传路径
- os.MkdirAll(uploads, 0666)： 创建上传目录
- beego.BConfig：保存了 beego 框架的默认参数
- RegisterDataBase()：注册数据库
- RegisterModel()：注册模型
- RegisterLogger(LogFile)：注册日志

#### commands.RegisterFunction()：添加模板函数

- beego.AddFuncMap("config", models.GetOptionValue)：获得选项值函数
- 添加 cdn 等系列模板中要用到的函数

#### beego.ErrorController(&controllers.ErrorController{})：定义错误页面

#### models.Init()：model 内初始化函数

- initAPI()：设置静态域（static_domain）
- initOptionCache()：读取默认选项并缓存

#### beego.Run()：启动 Beego 服务



