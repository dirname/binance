package spotclient

import "github.com/shopspring/decimal"

// AllCoinInfoResponse all coin info response
type AllCoinInfoResponse []struct {
	Coin             string `json:"coin"`
	DepositAllEnable bool   `json:"depositAllEnable"`
	Free             string `json:"free"`
	Freeze           string `json:"freeze"`
	IPOAble          string `json:"ipoable"`
	IPOIng           string `json:"ipoing"`
	IsLegalMoney     bool   `json:"isLegalMoney"`
	Locked           string `json:"locked"`
	Name             string `json:"name"`
	NetworkList      []struct {
		AddressRegex            string `json:"addressRegex"`
		Coin                    string `json:"coin"`
		DepositDesc             string `json:"depositDesc,omitempty"`
		DepositEnable           bool   `json:"depositEnable"`
		IsDefault               bool   `json:"isDefault"`
		MemoRegex               string `json:"memoRegex"`
		MinConfirm              int    `json:"minConfirm"`
		Name                    string `json:"name"`
		Network                 string `json:"network"`
		ResetAddressStatus      bool   `json:"resetAddressStatus"`
		SpecialTips             string `json:"specialTips"`
		UnLockConfirm           int    `json:"unLockConfirm"`
		WithdrawDesc            string `json:"withdrawDesc,omitempty"`
		WithdrawEnable          bool   `json:"withdrawEnable"`
		WithdrawFee             string `json:"withdrawFee"`
		WithdrawMin             string `json:"withdrawMin"`
		InsertTime              int64  `json:"insertTime,omitempty"`
		UpdateTime              int64  `json:"updateTime,omitempty"`
		WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple,omitempty"`
	} `json:"networkList"`
	Storage           string `json:"storage"`
	Trading           bool   `json:"trading"`
	WithdrawAllEnable bool   `json:"withdrawAllEnable"`
	Withdrawing       string `json:"withdrawing"`
}

// SpotSnapshotResponse spot snapshot response
type SpotSnapshotResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	SnapshotVos []struct {
		Data struct {
			Balances []struct {
				Asset  string `json:"asset"`
				Free   string `json:"free"`
				Locked string `json:"locked"`
			} `json:"balances"`
			TotalAssetOfBtc string `json:"totalAssetOfBtc"`
		} `json:"data"`
		Type       string `json:"type"`
		UpdateTime int64  `json:"updateTime"`
	} `json:"snapshotVos"`
}

// FuturesSnapshotResponse futures snapshot response
type FuturesSnapshotResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	SnapshotVos []struct {
		Data struct {
			Assets []struct {
				Asset         string `json:"asset"`
				MarginBalance string `json:"marginBalance"`
				WalletBalance string `json:"walletBalance"`
			} `json:"assets"`
			Position []struct {
				EntryPrice       string `json:"entryPrice"`
				MarkPrice        string `json:"markPrice"`
				PositionAmt      string `json:"positionAmt"`
				Symbol           string `json:"symbol"`
				UnRealizedProfit string `json:"unRealizedProfit"`
			} `json:"position"`
		} `json:"data"`
		Type       string `json:"type"`
		UpdateTime int64  `json:"updateTime"`
	} `json:"snapshotVos"`
}

// MarginSnapshotResponse margin snapshot response
type MarginSnapshotResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	SnapshotVos []struct {
		Data struct {
			MarginLevel         string `json:"marginLevel"`
			TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
			TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
			TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`
			UserAssets          []struct {
				Asset    string `json:"asset"`
				Borrowed string `json:"borrowed"`
				Free     string `json:"free"`
				Interest string `json:"interest"`
				Locked   string `json:"locked"`
				NetAsset string `json:"netAsset"`
			} `json:"userAssets"`
		} `json:"data"`
		Type       string `json:"type"`
		UpdateTime int64  `json:"updateTime"`
	} `json:"snapshotVos"`
}

// SAPIWithdrawResponse sapi withdraw response
type SAPIWithdrawResponse struct {
	ID string `json:"id"`
}

// WAPIWithdrawResponse wapi withdraw response
type WAPIWithdrawResponse struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	ID      string `json:"id"`
}

// DepositHistoryNetworkResponse DepositHistoryNetwork response
type DepositHistoryNetworkResponse []struct {
	Amount       string `json:"amount"`
	Coin         string `json:"coin"`
	Network      string `json:"network"`
	Status       int    `json:"status"`
	Address      string `json:"address"`
	AddressTag   string `json:"addressTag"`
	TxID         string `json:"txId"`
	InsertTime   int64  `json:"insertTime"`
	TransferType int    `json:"transferType"`
	ConfirmTimes string `json:"confirmTimes"`
}

// DepositHistoryResponse DepositHistory response
type DepositHistoryResponse struct {
	DepositList []struct {
		InsertTime int64           `json:"insertTime"`
		Amount     decimal.Decimal `json:"amount"`
		Asset      string          `json:"asset"`
		Address    string          `json:"address"`
		TxID       string          `json:"txId"`
		Status     int             `json:"status"`
		AddressTag string          `json:"addressTag,omitempty"`
	} `json:"depositList"`
	Success bool `json:"success"`
}

// WithdrawHistoryNetworkResponse withdraw network response
type WithdrawHistoryNetworkResponse []struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ApplyTime       string `json:"applyTime"`
	Coin            string `json:"coin"`
	ID              string `json:"id"`
	WithdrawOrderID string `json:"withdrawOrderId,omitempty"`
	Network         string `json:"network"`
	TransferType    int    `json:"transferType,omitempty"`
	Status          int    `json:"status"`
	TxID            string `json:"txId"`
}

// WithdrawHistoryResponse withdraw response
type WithdrawHistoryResponse struct {
	WithdrawList []struct {
		ID              string          `json:"id"`
		WithdrawOrderID interface{}     `json:"withdrawOrderId"`
		Amount          decimal.Decimal `json:"amount"`
		TransactionFee  decimal.Decimal `json:"transactionFee"`
		Address         string          `json:"address"`
		Asset           string          `json:"asset"`
		TxID            string          `json:"txId"`
		ApplyTime       int64           `json:"applyTime"`
		Status          int             `json:"status"`
		AddressTag      string          `json:"addressTag,omitempty"`
	} `json:"withdrawList"`
	Success bool `json:"success"`
}

// DepositAddressNetworkResponse deposit address response
type DepositAddressNetworkResponse struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}

// DepositAddressResponse deposit address response
type DepositAddressResponse struct {
	Address    string `json:"address"`
	Success    bool   `json:"success"`
	AddressTag string `json:"addressTag"`
	Asset      string `json:"asset"`
}

// WAPIAccountResponse account response
type WAPIAccountResponse struct {
	Msg     string   `json:"msg"`
	Success bool     `json:"success"`
	Objs    []string `json:"objs"`
}

// SAPIAccountResponse account response
type SAPIAccountResponse struct {
	Data string `json:"data"`
}

// AccountWAPIAPIStatusResponse account API trading status response
type AccountWAPIAPIStatusResponse struct {
	Success bool `json:"success"`
	Status  struct { // API trading status detail
		IsLocked           bool `json:"isLocked"`
		PlannedRecoverTime int  `json:"plannedRecoverTime"` // If API trading function is locked, this is the planned recover time
		TriggerCondition   struct {
			GCR  int `json:"gcr"`  // Number of GTC orders
			IFER int `json:"ifer"` // Number of FOK/IOC orders
			UFR  int `json:"ufr"`  // Number of orders
		} `json:"triggerCondition"`
		Indicators map[string][]struct {
			Unfilled     string          `json:"i"` // Unfilled Ratio (UFR)
			Count        int64           `json:"c"` // Count of all order
			Value        decimal.Decimal `json:"v"` // Current value
			TriggerValue decimal.Decimal `json:"t"` // Trigger value
		} `json:"indicators"`
		UpdateTime int64 `json:"updateTime"`
	} `json:"status"`
}

// AccountSAPIStatusResponse account API trading status response
type AccountSAPIStatusResponse struct {
	Data struct {
		IsLocked           bool `json:"isLocked"`
		PlannedRecoverTime int  `json:"plannedRecoverTime"` // If API trading function is locked, this is the planned recover time
		TriggerCondition   struct {
			GCR  int `json:"GCR"`  // Number of GTC orders
			IFER int `json:"IFER"` // Number of FOK/IOC orders
			UFR  int `json:"UFR"`  // Number of orders
		} `json:"triggerCondition"`
		Indicators map[string][]struct {
			Unfilled     string          `json:"i"` // Unfilled Ratio (UFR)
			Count        int64           `json:"c"` // Count of all order
			Value        decimal.Decimal `json:"v"` // Current value
			TriggerValue decimal.Decimal `json:"t"` // Trigger value
		} `json:"indicators"`
		UpdateTime int64 `json:"updateTime"`
	} `json:"data"`
}

// DustLogWAPIResponse Fetch small amounts of assets exchanged BNB records
type DustLogWAPIResponse struct {
	Success bool `json:"success"`
	Results struct {
		Total int `json:"total"`
		Rows  []struct {
			TransferredTotal   string `json:"transfered_total"`
			ServiceChargeTotal string `json:"service_charge_total"`
			TranID             int    `json:"tran_id"`
			Logs               []struct {
				TranID              int    `json:"tranId"`
				ServiceChargeAmount string `json:"serviceChargeAmount"`
				Uid                 string `json:"uid"`
				Amount              string `json:"amount"`
				OperateTime         string `json:"operateTime"`
				TransferredAmount   string `json:"transferedAmount"`
				FromAsset           string `json:"fromAsset"`
			} `json:"logs"`
			OperateTime string `json:"operate_time"`
		} `json:"rows"`
	} `json:"results"`
}

// DustLogSAPIResponse Fetch small amounts of assets exchanged BNB records
type DustLogSAPIResponse struct {
	Total             int `json:"total"`
	UserAssetDriblets []struct {
		TotalTransferredAmount   string `json:"totalTransferedAmount"`
		TotalServiceChargeAmount string `json:"totalServiceChargeAmount"`
		TransId                  int64  `json:"transId"`
		UserAssetDribletDetails  []struct {
			TransId             int         `json:"transId"`
			ServiceChargeAmount string      `json:"serviceChargeAmount"`
			Amount              string      `json:"amount"`
			OperateTime         interface{} `json:"operateTime"`
			TransferredAmount   string      `json:"transferedAmount"`
			FromAsset           string      `json:"fromAsset"`
		} `json:"userAssetDribbletDetails"`
		OperateTime int64 `json:"operateTime,omitempty"`
	} `json:"userAssetDribblets"`
}

// DustTransferResponse Convert dust assets to BNB.
type DustTransferResponse struct {
	TotalServiceCharge string `json:"totalServiceCharge"`
	TotalTransferred   string `json:"totalTransfered"`
	TransferResult     []struct {
		Amount              string `json:"amount"`
		FromAsset           string `json:"fromAsset"`
		OperateTime         int64  `json:"operateTime"`
		ServiceChargeAmount string `json:"serviceChargeAmount"`
		TranId              int64  `json:"tranId"`
		TransferredAmount   string `json:"transferedAmount"`
	} `json:"transferResult"`
}

// AccountDividendRecordResponse Query asset dividend record.
type AccountDividendRecordResponse struct {
	Rows []struct {
		Amount  string `json:"amount"`
		Asset   string `json:"asset"`
		DivTime int64  `json:"divTime"`
		EnInfo  string `json:"enInfo"`
		TranID  int64  `json:"tranId"`
	} `json:"rows"`
	Total int `json:"total"`
}

// WAPIAssetDetailResponse Fetch details of assets supported on Binance.
type WAPIAssetDetailResponse struct {
	Success     bool `json:"success"`
	AssetDetail map[string]struct {
		MinWithdrawAmount interface{} `json:"minWithdrawAmount"`
		DepositStatus     bool        `json:"depositStatus"`
		WithdrawFee       interface{} `json:"withdrawFee"`
		WithdrawStatus    bool        `json:"withdrawStatus"`
		DepositTip        string      `json:"depositTip"`
	} `json:"assetDetail"`
}

// SAPIAssetDetailResponse Fetch details of assets supported on Binance.
type SAPIAssetDetailResponse map[string]struct {
	MinWithdrawAmount string `json:"minWithdrawAmount"`
	DepositStatus     bool   `json:"depositStatus"`
	WithdrawFee       string `json:"withdrawFee"`
	WithdrawStatus    bool   `json:"withdrawStatus"`
	DepositTip        string `json:"depositTip"`
}

// WAPITradeFeeResponse Fetch trade fee, values in percentage.
type WAPITradeFeeResponse struct {
	TradeFee []struct {
		Symbol string          `json:"symbol"`
		Maker  decimal.Decimal `json:"maker"`
		Taker  decimal.Decimal `json:"taker"`
	} `json:"tradeFee"`
	Success bool `json:"success"`
}

// SAPITradeFeeResponse Fetch trade fee, values in percentage.
type SAPITradeFeeResponse []struct {
	Symbol          string          `json:"symbol"`
	MakerCommission decimal.Decimal `json:"makerCommission"`
	TakerCommission decimal.Decimal `json:"takerCommission"`
}

// UniversalTransferResponse universal transfer
type UniversalTransferResponse struct {
	TranID int64 `json:"tranId"`
}

// UniversalTransferRecordResponse universal transfer record
type UniversalTransferRecordResponse struct {
	Total int `json:"total"`
	Rows  []struct {
		Asset     string `json:"asset"`
		Amount    string `json:"amount"`
		Type      string `json:"type"`
		Status    string `json:"status"`
		TranID    int64  `json:"tranId"`
		Timestamp int64  `json:"timestamp"`
	} `json:"rows"`
}
