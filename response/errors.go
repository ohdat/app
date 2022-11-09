package response

type ErrCode int

func (e ErrCode) Error() string {
	return e.String()
}

func (e ErrCode) Code() int {
	return int(e)
}

// go get golang.org/x/tools/cmd/stringer
//go:generate stringer -type=ErrCode -linecomment
const (
	ErrRecordNotFound ErrCode = iota + 1002 // record not found
	// ErrNeedWait 需要等待 重复请求
	ErrNeedWait // need wait
	// ErrSendVerificationCodeFailed 验证码发送失败
	ErrSendVerificationCodeFailed // send verification code failed
	// ErrVerificationCodeWrong 验证码发送失败
	ErrVerificationCodeWrong //verification code wrong
)

//登录错误
const (
	ErrNeedLogin ErrCode = iota + 1101 // need login
	// ErrTokenExpire TOKEN过期
	ErrTokenExpire //token expired
	// ErrTokenVerificationFail token验证失败
	ErrTokenVerificationFail //token verification failed
	// ErrAccountNotExist 用户不存在
	ErrAccountNotExist //account not exist
	// ErrOtherLogin 在其他设备登录
	ErrOtherLogin //other login
	// ErrPasswordWrong 密码错误
	ErrPasswordWrong //password wrong
	// ErrNeedBindInfo 需要绑定信息
	ErrNeedBindInfo //need bind info
	// ErrNicknameExist 昵称已存在
	ErrNicknameExist //nickname exist
	// ErrSignFail 签名验证失败
	ErrSignFail //signature failed
	// ErrHashVerificationFail hash验证失败
	ErrHashVerificationFail //hash verification failed
	ErrAuthExpired          //authorization expired.
)

//支付错误
const (
	// ErrPaymentFailed 支付失败
	ErrPaymentFailed ErrCode = iota + 1201 //payment failed
	// ErrAlreadyPaid 已支付，请勿重复支付
	ErrAlreadyPaid //already paid
	// ErrNeedPay  需要支付
	ErrNeedPay //need pay
)

//售卖相关错误
const (
	// ErrNotQualify 用户无资格购买
	ErrNotQualify ErrCode = iota + 1301 // not qualify
	// ErrGoods 商品信息错误
	ErrGoods // goods not available
)

//游戏相关错误
const (
	//ErrNoRunningGame  没有正在运行的游戏场次
	ErrNoRunningGame    ErrCode = iota + 1401 //no running game
	ErrRunningGameExist                       //running game exist
)
