package v1

type CreateLinkRequest struct {
	Url string `json:"url" binding:"required,url" example:"https://www.google.com"`
}

type CreateLinkResponseData struct {
	ShortLink string `json:"shortLink"`
	Url       string `json:"url"`
}

type CreateLinkResponse struct {
	Response
	Data CreateLinkResponseData
}

type GetLinkResponseData struct {
	Url string `json:"url"`
}

type GetLinkResponse struct {
	Response
	Data GetLinkResponseData
}
