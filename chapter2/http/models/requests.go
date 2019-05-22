package models

type GeekBang struct {
	Host           string `json:"host"`
	Connection     string `json:"connection"`
	ContentLength  int    `json:"content-length"`
	Pragma         string `json:"pragma"`
	CacheControl   string `json:"cache-control"`
	Accept         string `json:"accept"`
	Origin         string `json:"origin"`
	UserAgent      string `json:"user-agent"`
	DNT            int    `json:"DNT"`
	ContentType    string `json:"content-type"`
	Referer        string `json:"referer"`
	AcceptEncoding string `json:"accept-encoding"`
	AcceptLanguage string `json:"accept-language"`
	Cookie         string `json:"cookie"`
}

type DiDiRecruitRequestParams struct {
	JobType     int    `json:"jobType"`     //工作类型，技术、设计、运营
	Page        int    `json:"page"`        //页码
	RecruitType int    `json:"recruitType"` //招聘类型，社招，校招
	Size        int    `json:"size"`        //分页大小
	WorkArea    string `json:"workArea"`    //工作地区
}

func (requestParams *DiDiRecruitRequestParams) Format() string {

}
