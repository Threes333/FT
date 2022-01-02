package emsg

const (
	Success = 10000
	Error   = 20000
	//用户错误
	UsernameExist    = 10001
	UsernameNoExist  = 10002
	PasswordWrong    = 10003
	UserLogoutFailed = 10004
	//回复错误
	GetReplyFailed    = 20001
	CreateReplyFailed = 20002
	ReplyNoExist      = 20003
	DeleteReplyFailed = 20004
	//话题错误
	GetTopicFailed    = 30001
	CreateTopicFailed = 30002
	TopicNoExist      = 30003
	DeleteTopicFailed = 30004
	UpdateTopicFailed = 30005
	//token错误
	GenerateAccessTokenFailed  = 40001
	GenerateRefreshTokenFailed = 40002
	TokenErrorMalformed        = 40003
	TokenErrorExpired          = 40004
	TokenErrorNotValidYet      = 40005
	TokenCannotRecognized      = 40006
	TokenInvalid               = 40007
	AccessTokenNoExist         = 40008
	RefreshTokenNoExist        = 40009
	TokenNoPermission          = 40010
	//课堂相关
	CreateClassFailed         = 50001
	DeleteClassFailed         = 50002
	UpdateClassFailed         = 50003
	ClassNoExist              = 50004
	GetClassFailed            = 50005
	JoinClassFailed           = 50006
	OutClassFailed            = 50007
	ReleaseMaterialFailed     = 50008
	DeleteMaterialFailed      = 50009
	MaterialNoExist           = 50010
	GetMaterialFailed         = 50011
	ClassHasNoStudent         = 50012
	CreateAttendanceFailed    = 50013
	CheckInFailed             = 50014
	AttendanceNoExist         = 50015
	GetAttendanceStatusFailed = 50016
	//成绩相关
	CreateScoreFailed = 60001
	UpdateScoreFailed = 60002
	ScoreNoExist      = 60003
	GetScoreFailed    = 60004
)

var ErrorMsg map[int]string

func init() {
	ErrorMsg = make(map[int]string)
	ErrorMsg[Success] = "操作成功"
	ErrorMsg[Error] = "操作失败"
	ErrorMsg[UsernameExist] = "用户名已存在"
	ErrorMsg[UsernameNoExist] = "用户名不存在"
	ErrorMsg[PasswordWrong] = "密码错误"
	ErrorMsg[UserLogoutFailed] = "用户注销失败"
	ErrorMsg[GetReplyFailed] = "获取回复信息失败"
	ErrorMsg[CreateReplyFailed] = "创建回复失败"
	ErrorMsg[ReplyNoExist] = "回复不存在"
	ErrorMsg[DeleteReplyFailed] = "删除回复失败"
	ErrorMsg[GetTopicFailed] = "获取分类信息失败"
	ErrorMsg[CreateTopicFailed] = "创建话题失败"
	ErrorMsg[TopicNoExist] = "话题不存在"
	ErrorMsg[DeleteTopicFailed] = "删除话题失败"
	ErrorMsg[UpdateTopicFailed] = "更新话题信息失败"
	ErrorMsg[GenerateAccessTokenFailed] = "生成AccessToken失败"
	ErrorMsg[GenerateRefreshTokenFailed] = "生成RefreshToken失败"
	ErrorMsg[TokenErrorMalformed] = "Token格式错误"
	ErrorMsg[TokenErrorExpired] = "Token已过期"
	ErrorMsg[TokenErrorNotValidYet] = "Token未生效"
	ErrorMsg[TokenCannotRecognized] = "无法辨认该Token"
	ErrorMsg[TokenInvalid] = "非法Token"
	ErrorMsg[AccessTokenNoExist] = "AccessToken不存在"
	ErrorMsg[RefreshTokenNoExist] = "RefreshToken不存在"
	ErrorMsg[TokenNoPermission] = "该Token无权限访问"
	ErrorMsg[CreateClassFailed] = "创建课堂失败"
	ErrorMsg[DeleteClassFailed] = "删除课堂失败"
	ErrorMsg[UpdateClassFailed] = "更新课堂信息失败"
	ErrorMsg[ClassNoExist] = "课堂不存在"
	ErrorMsg[GetClassFailed] = "获取课堂信息失败"
	ErrorMsg[JoinClassFailed] = "加入课堂失败"
	ErrorMsg[OutClassFailed] = "退出课堂失败"
	ErrorMsg[ReleaseMaterialFailed] = "发布资料失败"
	ErrorMsg[DeleteMaterialFailed] = "删除资料失败"
	ErrorMsg[MaterialNoExist] = "资料不存在"
	ErrorMsg[GetMaterialFailed] = "获取资料失败"
	ErrorMsg[ClassHasNoStudent] = "该课堂无学生"
	ErrorMsg[CreateAttendanceFailed] = "创建考勤失败"
	ErrorMsg[CheckInFailed] = "签到失败"
	ErrorMsg[AttendanceNoExist] = "考勤不存在"
	ErrorMsg[GetAttendanceStatusFailed] = "获取考勤情况失败"
	ErrorMsg[CreateScoreFailed] = "创建成绩失败"
	ErrorMsg[UpdateScoreFailed] = "更新成绩失败"
	ErrorMsg[ScoreNoExist] = "成绩不存在"
	ErrorMsg[GetScoreFailed] = "获取成绩失败"
}

func GetErrorMsg(code int) string {
	return ErrorMsg[code]
}
