package webapi

const UrlBase = "https://localhost:5000/v1/portal"

// session endpoints
const (
	UrlStatus   = "/iserver/auth/status"
	UrlReauth   = "/iserver/reauthenticate"
	UrlValidate = "/sso/validate"
	UrlTickle   = "/tickle"
	UrlLogout   = "/logout"
)

//  trading history endpoint
const UrlTrade = "/iserver/account/trades"

// account endpoints
const (
	UrlAccounts  = "/iserver/accounts"
	UrlAccount   = "/iserver/account"
	UrlPnl       = "/iserver/account/pnl/partitioned"
	UrlPortfolio = "/portfolio/accounts"
)

// order endpoints
const (
	UrlLiveOrder          = "/iserver/account/orders"
	UrlPlaceOrder         = "/iserver/account/{accountId}/order"
	UrlPlaceOrdersBracket = "/iserver/account/{accountId}/orders"
	UrlPlaceOrderReply    = "/iserver/reply/{replyid}"
	UrlPreviewOrder       = "/iserver/account/{accountId}/order/whatif"
	UrlModifyOrder        = "/iserver/account/{accountId}/order/{origCustomerOrderId}"
	UrlDeleteOrder        = "/iserver/account/{accountId}/order/{origCustomerOrderId}"
)

// market data endpoints
const UrlMarketData = "/iserver/marketdata/snapshot"
const UrlMarketDataHistory = "/iserver/marketdata/history"

// contract endpoints
const (
	UrlCtx       = "/iserver/contract"
	UrlCtxInfo   = "/{conid}/info"
	UrlCtxSearch = "/iserver/secdef/search"
	UrlCtxConId  = "/trsrv/secdef"
)
