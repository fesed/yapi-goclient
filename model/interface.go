package model

type InterfaceMethodType string

const (
	InterfaceMethodTypeGet     InterfaceMethodType = "GET"
	InterfaceMethodTypePost    InterfaceMethodType = "POST"
	InterfaceMethodTypePUT     InterfaceMethodType = "PUT"
	InterfaceMethodTypeDelete  InterfaceMethodType = "DELETE"
	InterfaceMethodTypePatch   InterfaceMethodType = "PATCH"
	InterfaceMethodTypeHead    InterfaceMethodType = "HEAD"
	InterfaceMethodTypeOptions InterfaceMethodType = "OPTIONS"
)

type ListInterfaceResp struct {
	CommonResp
	Data struct {
		Count int          `json:"count"`
		Total int          `json:"total"`
		List  []*Interface `json:"list"`
	} `json:"data"`
}

type Interface struct {
	EditUID   int      `json:"edit_uid"`
	Status    string   `json:"status"`
	APIOpened bool     `json:"api_opened"`
	Tag       []string `json:"tag"`
	ID        int      `json:"_id"`
	Method    string   `json:"method"`
	Title     string   `json:"title"`
	Path      string   `json:"path"`
	ProjectID int      `json:"project_id"`
	Catid     int      `json:"catid"`
	UID       int      `json:"uid"`
	AddTime   int      `json:"add_time"`
}

type ListMenuResp struct {
	CommonResp
	Data []*Menu `json:"data"`
}

type Menu struct {
	Index     int    `json:"index"`
	ID        int    `json:"_id"`
	Name      string `json:"name"`
	ProjectID int    `json:"project_id"`
	Desc      string `json:"desc"`
	UID       int    `json:"uid"`
	AddTime   int    `json:"add_time"`
	UpTime    int    `json:"up_time"`
	V         int    `json:"__v"`
	List      []struct {
		EditUID   int      `json:"edit_uid"`
		Status    string   `json:"status"`
		Index     int      `json:"index"`
		Tag       []string `json:"tag"`
		ID        int      `json:"_id"`
		Method    string   `json:"method"`
		Title     string   `json:"title"`
		Path      string   `json:"path"`
		ProjectID int      `json:"project_id"`
		Catid     int      `json:"catid"`
		UID       int      `json:"uid"`
		AddTime   int      `json:"add_time"`
		UpTime    int      `json:"up_time"`
	} `json:"list,omitempty"`
}

type AddInterfaceReq struct {
	Method    InterfaceMethodType `json:"method"`
	Catid     string              `json:"catid"`
	Title     string              `json:"title"`
	Path      string              `json:"path"`
	ProjectID int                 `json:"project_id"`
}

type GetInterfaceResp struct {
	CommonResp
	Data *InterfaceDetail `json:"data"`
}

type UpdateInterfaceReq struct {
	ReqParams           []ReqParams   `json:"req_params,omitempty"`
	ReqQuery            []ReqQuery    `json:"req_query,omitempty"`
	ReqHeaders          []ReqHeaders  `json:"req_headers,omitempty"`
	ReqBodyForm         []ReqBodyForm `json:"req_body_form,omitempty"`
	Title               string        `json:"title,omitempty"`
	Catid               string        `json:"catid,omitempty"`
	Path                string        `json:"path,omitempty"`
	Tag                 []string      `json:"tag,omitempty"`
	Status              string        `json:"status,omitempty"`
	CustomFieldValue    string        `json:"custom_field_value,omitempty"`
	ReqBodyType         string        `json:"req_body_type,omitempty"`
	ReqBodyIsJSONSchema bool          `json:"req_body_is_json_schema,omitempty"`
	ResBodyIsJSONSchema bool          `json:"res_body_is_json_schema,omitempty"`
	ResBodyType         string        `json:"res_body_type,omitempty"`
	ResBody             string        `json:"res_body,omitempty"`
	SwitchNotice        bool          `json:"switch_notice,omitempty"`
	APIOpened           bool          `json:"api_opened,omitempty"`
	Desc                string        `json:"desc,omitempty"`
	Markdown            string        `json:"markdown,omitempty"`
	Method              string        `json:"method,omitempty"`
	ID                  string        `json:"id,omitempty"`
}

type AddCatReq struct {
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	ProjectID string `json:"project_id"`
}

type ListInterfacesByCat struct {
	CommonResp
	Data struct {
		Count int          `json:"count"`
		Total int          `json:"total"`
		List  []*Interface `json:"list"`
	} `json:"data"`
}

type UpdateCatReq struct {
	Catid int    `json:"catid"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
}

type DeleteCatReq struct {
	CatId int `json:"catid"`
}

type QueryPath struct {
	Path   string        `json:"path"`
	Params []interface{} `json:"params"`
}
type ReqParams struct {
	ID      string `json:"_id,omitempty"`
	Name    string `json:"name,omitempty"`
	Example string `json:"example,omitempty"`
	Desc    string `json:"desc,omitempty"`
}
type ReqBodyForm struct {
	Required string `json:"required,omitempty"`
	ID       string `json:"_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	Example  string `json:"example,omitempty"`
	Desc     string `json:"desc,omitempty"`
}
type ReqHeaders struct {
	Required string `json:"required,omitempty"`
	ID       string `json:"_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Value    string `json:"value,omitempty"`
	Example  string `json:"example,omitempty"`
	Desc     string `json:"desc,omitempty"`
}
type ReqQuery struct {
	Required string `json:"required,omitempty"`
	ID       string `json:"_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Example  string `json:"example,omitempty"`
	Desc     string `json:"desc,omitempty"`
}
type InterfaceDetail struct {
	QueryPath           QueryPath     `json:"query_path"`
	EditUID             int           `json:"edit_uid"`
	Status              string        `json:"status"`
	Type                string        `json:"type"`
	ReqBodyIsJSONSchema bool          `json:"req_body_is_json_schema"`
	ResBodyIsJSONSchema bool          `json:"res_body_is_json_schema"`
	APIOpened           bool          `json:"api_opened"`
	Index               int           `json:"index"`
	Tag                 []string      `json:"tag"`
	ID                  int           `json:"_id"`
	Method              string        `json:"method"`
	Title               string        `json:"title"`
	Desc                string        `json:"desc"`
	Path                string        `json:"path"`
	ReqParams           []ReqParams   `json:"req_params"`
	ReqBodyForm         []ReqBodyForm `json:"req_body_form"`
	ReqHeaders          []ReqHeaders  `json:"req_headers"`
	ReqQuery            []ReqQuery    `json:"req_query"`
	ReqBodyType         string        `json:"req_body_type"`
	ResBodyType         string        `json:"res_body_type"`
	ResBody             string        `json:"res_body"`
	ReqBodyOther        string        `json:"req_body_other"`
	ProjectID           int           `json:"project_id"`
	Catid               int           `json:"catid"`
	UID                 int           `json:"uid"`
	AddTime             int           `json:"add_time"`
	UpTime              int           `json:"up_time"`
	V                   int           `json:"__v"`
	CustomFieldValue    string        `json:"custom_field_value"`
	Markdown            string        `json:"markdown"`
	Username            string        `json:"username"`
}

type SaveInterfaceReq struct {
	Method              string        `json:"method,omitempty"`
	Title               string        `json:"title,omitempty"`
	Desc                string        `json:"desc,omitempty"`
	Catname             string        `json:"catname,omitempty"`
	Tag                 []string      `json:"tag,omitempty"`
	Path                string        `json:"path,omitempty"`
	ReqParams           []ReqParams   `json:"req_params,omitempty"`
	ReqBodyForm         []ReqBodyForm `json:"req_body_form,omitempty"`
	ReqHeaders          []ReqHeaders  `json:"req_headers,omitempty"`
	ReqQuery            []ReqQuery    `json:"req_query,omitempty"`
	ReqBodyType         string        `json:"req_body_type,omitempty"`
	ResBodyType         string        `json:"res_body_type,omitempty"`
	ResBody             string        `json:"res_body,omitempty"`
	ResBodyIsJSONSchema bool          `json:"res_body_is_json_schema,omitempty"`
	ProjectID           string        `json:"project_id,omitempty"`
	Catid               int           `json:"catid,omitempty"`
	DataSync            string        `json:"dataSync,omitempty"`
}

type SaveInterfaceResp struct {
	CommonResp
	Data []struct {
		ID      int    `json:"_id"`
		ResBody string `json:"res_body"`
	} `json:"data"`
}
