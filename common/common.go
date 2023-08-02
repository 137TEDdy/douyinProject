/*
   @Author Ted
   @Since 2023/7/25
*/

package common

import "douyinProject/model"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// 功能： user登录的响应，包含 response(code,msg), userId, token
type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// 功能： 获取用户信息时返回的响应
type UserResponse struct {
	Response
	User model.User `json:"user"`
}

type CommentResponse struct {
	Response
	Comment model.Comment `json:"comment"`
}

type CommentListResponse struct {
	Response
	Comments []*model.Comment `json:"comment_list"`
}
