package response

type ErrCode int

func (e ErrCode) Error() string {
	return e.String()
}

func (e ErrCode) Code() int {
	return int(e)
}

// go get golang.org/x/tools/cmd/stringer
//
//go:generate stringer -type=ErrCode -linecomment
const (
	ErrRecordNotFound ErrCode = iota + 1002 // record not found
	// ErrNeedWait 需要等待 重复请求
	ErrNeedWait // need wait
	// ErrSendVerificationCodeFailed 验证码发送失败
	ErrSendVerificationCodeFailed // send verification code failed
	// ErrVerificationCodeWrong 验证码发送失败
	ErrVerificationCodeWrong //verification code wrong

	ErrRecaptchaNotFound // recaptcha not found
	ErrRecaptchaTimeout  // recaptcha timeout
	ErrRecaptchaFailed   // recaptcha failed
)

// 登录错误
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

// 支付错误
const (
	// ErrPaymentFailed 支付失败
	ErrPaymentFailed ErrCode = iota + 1201 //payment failed
	// ErrAlreadyPaid 已支付，请勿重复支付
	ErrAlreadyPaid //already paid
	// ErrNeedPay  需要支付
	ErrNeedPay //need pay
)

// 售卖相关错误
const (
	// ErrNotQualify 用户无资格购买
	ErrNotQualify ErrCode = iota + 1301 // not qualify
	// ErrGoods 商品信息错误
	ErrGoods // goods not available
	//ErrPriceNotEnough 出价不够
	ErrPriceNotEnough // price not enough
	//ErrAuctionEnded 拍卖已结束
	ErrAuctionEnded // auction ended
	//ErrBuyItNowPricesChanged 限时抢拍价已更新
	ErrBuyItNowPricesChanged // buy it now prices change
	//ErrNeedPledge 需付押金
	ErrNeedPledge // need pledge
	//ErrRedeemFailed 兑换失败
	ErrRedeemFailed // redeem failed
	//ErrRedeemCodeUsed 兑换码已兑换
	ErrRedeemCodeUsed // redeem code used
	//ErrBuyLimit 限购
	ErrBuyLimit // buy limit reached
	//ErrNotEnoughMoney 余额不足
	ErrNotEnoughMoney // not enough money

)

// 游戏相关错误
const (
	//ErrNoRunningGame  没有正在运行的游戏场次
	ErrNoRunningGame    ErrCode = iota + 1401 //no running game
	ErrRunningGameExist                       //running game exist
)

// 调查问卷相关错误
const (
	ErrQuestionnaireFailed           ErrCode = iota + 1501 //questionnaire running failed
	ErrQuestionnaireInviteCode                             // invited code invalid
	ErrQuestionnaireAddressUsed                            // address used
	ErrQuestionnaireAddressNotExists                       // address not exists
	ErrQuestionnaireCodeExists                             // code exists
)

// 钱包相关错误
const (
	ErrWalletInUse ErrCode = iota + 1601 //wallet is in use
)

// 瓜皮兔相关错误
const (
	ErrGuapituFailed         ErrCode = iota + 1701 //guapitu failed
	ErrGuapituTokenNotEnough                       //guapitu tokken not enough
)
