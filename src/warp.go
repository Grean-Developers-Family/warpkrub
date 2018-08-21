package main

type Warp struct {
	Warp []struct {
		Name string `json:"name"`
		Ig   string `json:"ig"`
		Fb   string `json:"fb"`
		Ref  string `json:"ref"`
	} `json:"warps"`
}
