// Code generated by "stringer -type=ErrCode -linecomment"; DO NOT EDIT.

package response

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrRecordNotFound-1002]
	_ = x[ErrNeedWait-1003]
	_ = x[ErrSendVerificationCodeFailed-1004]
	_ = x[ErrVerificationCodeWrong-1005]
	_ = x[ErrRecaptchaNotFound-1006]
	_ = x[ErrRecaptchaTimeout-1007]
	_ = x[ErrRecaptchaFailed-1008]
	_ = x[ErrNeedLogin-1101]
	_ = x[ErrTokenExpire-1102]
	_ = x[ErrTokenVerificationFail-1103]
	_ = x[ErrAccountNotExist-1104]
	_ = x[ErrOtherLogin-1105]
	_ = x[ErrPasswordWrong-1106]
	_ = x[ErrNeedBindInfo-1107]
	_ = x[ErrNicknameExist-1108]
	_ = x[ErrSignFail-1109]
	_ = x[ErrHashVerificationFail-1110]
	_ = x[ErrAuthExpired-1111]
	_ = x[ErrPaymentFailed-1201]
	_ = x[ErrAlreadyPaid-1202]
	_ = x[ErrNeedPay-1203]
	_ = x[ErrNotQualify-1301]
	_ = x[ErrGoods-1302]
	_ = x[ErrPriceNotEnough-1303]
	_ = x[ErrAuctionEnded-1304]
	_ = x[ErrBuyItNowPricesChanged-1305]
	_ = x[ErrNeedPledge-1306]
	_ = x[ErrRedeemFailed-1307]
	_ = x[ErrRedeemCodeUsed-1308]
	_ = x[ErrBuyLimit-1309]
	_ = x[ErrNotEnoughMoney-1310]
	_ = x[ErrNoRunningGame-1401]
	_ = x[ErrRunningGameExist-1402]
	_ = x[ErrQuestionnaireFailed-1501]
	_ = x[ErrQuestionnaireInviteCode-1502]
	_ = x[ErrQuestionnaireAddressUsed-1503]
	_ = x[ErrQuestionnaireAddressNotExists-1504]
	_ = x[ErrQuestionnaireCodeExists-1505]
}

const (
	_ErrCode_name_0 = "record not foundneed waitsend verification code failedverification code wrongrecaptcha not foundrecaptcha timeoutrecaptcha failed"
	_ErrCode_name_1 = "need logintoken expiredtoken verification failedaccount not existother loginpassword wrongneed bind infonickname existsignature failedhash verification failedauthorization expired."
	_ErrCode_name_2 = "payment failedalready paidneed pay"
	_ErrCode_name_3 = "not qualifygoods not availableprice not enoughauction endedbuy it now prices changeneed pledgeredeem failedredeem code usedbuy limit reachednot enough money"
	_ErrCode_name_4 = "no running gamerunning game exist"
	_ErrCode_name_5 = "questionnaire running failedinvited code invalidaddress usedaddress not existscode exists"
)

var (
	_ErrCode_index_0 = [...]uint8{0, 16, 25, 54, 77, 96, 113, 129}
	_ErrCode_index_1 = [...]uint8{0, 10, 23, 48, 65, 76, 90, 104, 118, 134, 158, 180}
	_ErrCode_index_2 = [...]uint8{0, 14, 26, 34}
	_ErrCode_index_3 = [...]uint8{0, 11, 30, 46, 59, 83, 94, 107, 123, 140, 156}
	_ErrCode_index_4 = [...]uint8{0, 15, 33}
	_ErrCode_index_5 = [...]uint8{0, 28, 48, 60, 78, 89}
)

func (i ErrCode) String() string {
	switch {
	case 1002 <= i && i <= 1008:
		i -= 1002
		return _ErrCode_name_0[_ErrCode_index_0[i]:_ErrCode_index_0[i+1]]
	case 1101 <= i && i <= 1111:
		i -= 1101
		return _ErrCode_name_1[_ErrCode_index_1[i]:_ErrCode_index_1[i+1]]
	case 1201 <= i && i <= 1203:
		i -= 1201
		return _ErrCode_name_2[_ErrCode_index_2[i]:_ErrCode_index_2[i+1]]
	case 1301 <= i && i <= 1310:
		i -= 1301
		return _ErrCode_name_3[_ErrCode_index_3[i]:_ErrCode_index_3[i+1]]
	case 1401 <= i && i <= 1402:
		i -= 1401
		return _ErrCode_name_4[_ErrCode_index_4[i]:_ErrCode_index_4[i+1]]
	case 1501 <= i && i <= 1505:
		i -= 1501
		return _ErrCode_name_5[_ErrCode_index_5[i]:_ErrCode_index_5[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
