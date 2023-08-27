/*
   @Author Ted
   @Since 2023/8/17 19:16
*/

package common

const (
	CodeUserNotExist int = 400 + iota
	CodeUserIdError
	CodeServerError
	CodeTokenNotexist
	CodeInvalidParams
	CodeUserExist
	CodeTokenError
	CodeCommentError
	CodeFavoriteError
	CodeFollowError
	CodeGetVideoListError
	CodePasswordError
	CodeVideoPublishError
	CodeMessageError
)
const CodeSuccess = 200

var codeMap = map[int]string{
	CodeSuccess:           "响应成功",
	CodeServerError:       "服务错误",
	CodeUserNotExist:      "用户不存在",
	CodeTokenNotexist:     "token不存在",
	CodeInvalidParams:     "参数错误",
	CodeUserIdError:       "用户id错误",
	CodeUserExist:         "用户已存在",
	CodeTokenError:        "token相关错误",
	CodeCommentError:      "评论相关错误",
	CodeFavoriteError:     "点赞错误",
	CodeFollowError:       "关注错误",
	CodePasswordError:     "密码错误",
	CodeGetVideoListError: "获取视频列表错误",
	CodeVideoPublishError: "视频发布错误",
	CodeMessageError:      "消息发送出错",
}

func Msg(code int) string {
	msg, isOk := codeMap[code]
	if !isOk {
		msg = codeMap[CodeServerError]
	}
	return msg
}
