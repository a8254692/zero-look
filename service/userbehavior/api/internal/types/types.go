// Code generated by goctl. DO NOT EDIT.
package types

type FocusFollowListReq struct {
	AppId     int64  `json:"appId"`
	Type      int64  `json:"type"`
	TargetUin string `json:"targetUin"`
	PageIndex int64  `json:"pageIndex"`
	PageSize  int64  `json:"pageSize"`
}

type FocusFollowListData struct {
	Id          int64  `json:"id"`
	UIn         string `json:"uin"`
	UserName    string `json:"userName"`
	UserHeadUrl string `json:"userHeadUrl"`
	Status      int32  `json:"status"`
}

type FocusFollowListResp struct {
	List  []FocusFollowListData `json:"list"`
	Total int64                 `json:"total"`
}

type AddFocusReq struct {
	AppId    int64  `json:"appId"`
	FocusUin string `json:"focusUin"`
}

type AddFocusResp struct {
}

type AddPraiseReq struct {
	AppId     int64  `json:"appId"`
	TopicType int64  `json:"topicType"`
	TopicId   string `json:"topicId"`
}

type AddPraiseResp struct {
}

type AddShareNumReq struct {
	AppId     int64  `json:"appId"`
	TopicType int64  `json:"topicType"`
	TopicId   string `json:"topicId"`
}

type AddShareNumResp struct {
}

type DelFocusReq struct {
	AppId    int64  `json:"appId"`
	FocusUin string `json:"focusUin"`
}

type DelFocusResp struct {
}

type DelPraiseReq struct {
	AppId     int64  `json:"appId"`
	TopicType int64  `json:"topicType"`
	TopicId   string `json:"topicId"`
}

type DelPraiseResp struct {
}

type AddCommentReq struct {
	AppId     int64  `json:"appId"`
	TopicId   string `json:"topicId"`
	TopicType int64  `json:"topicType"`
	Content   string `json:"content"`
}

type AddCommentResp struct {
}

type GetCommentListReq struct {
	AppId     int64  `json:"appId"`
	TopicId   string `json:"topicId"`
	TopicType int64  `json:"topicType"`
	PageSize  int32  `json:"pageSize"`
	Page      int32  `json:"page"`
	Sort      int32  `json:"sort"`
}

type CommentList struct {
	Id          int64  `json:"id"`
	Uin         string `json:"uin"`
	NickName    string `json:"nickname"`
	UserAvatar  string `json:"userAvatar"`
	Content     string `json:"content"`
	CreatedAt   string `json:"createdAt"`
	PraiseCount int64  `json:"praiseCount"`
	IsAuthor    bool   `json:"isAuthor"`
	IsPraise    bool   `json:"isPraise"`
}

type GetCommentListResp struct {
	Total int64         `json:"total"`
	List  []CommentList `json:"list"`
}