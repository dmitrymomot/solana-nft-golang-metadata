package metadata

type TokenAccountsByOwner struct {
	Context struct {
		Slot int64 `json:"slot"`
	} `json:"context"`
	Value []Value `json:"value"`
}

type Value struct {
	Account struct {
		Data struct {
			Parsed struct {
				Info struct {
					IsNative    bool   `json:"isNative"`
					Mint        string `json:"mint"`
					Owner       string `json:"owner"`
					State       string `json:"state"`
					TokenAmount struct {
						Amount         int64   `json:"amount,string"`
						Decimals       int64   `json:"decimals"`
						UiAmount       float64 `json:"uiAmount"`
						UiAmountString float64 `json:"uiAmountString,string"`
					} `json:"tokenAmount"`
				} `json:"info"`
				Type string `json:"type"`
			} `json:"parsed"`
			Program string `json:"program"`
			Space   int64  `json:"space"`
		} `json:"data"`
		Executable bool   `json:"executable"`
		Lamports   int64  `json:"lamports"`
		Owner      string `json:"owner"`
		RentEpoch  int64  `json:"rentEpoch"`
	} `json:"account"`
	Pubkey string `json:"pubkey"`
}

type AccountInfo struct {
	Context struct {
		Slot int64 `json:"slot"`
	} `json:"context"`
	Value struct {
		Data       []string `json:"data"`
		Executable bool     `json:"executable"`
		Lamports   int64    `json:"lamports"`
		Owner      string   `json:"owner"`
		RentEpoch  int64    `json:"rentEpoch"`
	} `json:"value"`
}
