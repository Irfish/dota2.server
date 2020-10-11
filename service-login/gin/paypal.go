package gin

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Irfish/component/log"
	"github.com/Irfish/dota2.server/service-login/base"
	"io/ioutil"
	"net/http"
	"time"
)

var payPalManager = NewPayPalManager()

type  PayPalManager struct{
	BaseURl string
	ClientID string
	Secret string
	Version string
	Access *PayPalAccess
	ReqTokenTime  int64
}

func NewPayPalManager() *PayPalManager {
	p := &PayPalManager{}
	p.ClientID = base.Server.PayPayClientID
	p.Secret = base.Server.PayPaySecret
	p.BaseURl = base.Server.PayPayUrl
	return p
}

//检查token是否过期
func (m *PayPalManager)CheckExpired() bool  {
	realTime:= m.ReqTokenTime+m.Access.ExpiresIn
	now:= time.Now().Unix()
	return  realTime<now
}

func (m *PayPalManager)GetAuthorization() string  {
	_,e:= m.PayPalOauth()
	if e!=nil{
		log.Debug("PayPalManager GetAuthor error: %s", e.Error())
		return ""
	}
	return "Bearer "+m.Access.AccessToken
}

//认证
func (m *PayPalManager)PayPalOauth() (*PayPalAccess,error) {
	if m.CheckExpired() {
		url:= m.BaseURl+"/v1/oauth2/token"
		b:= []byte(m.ClientID+":"+m.Secret)
		authorization:="Basic "+base64.URLEncoding.EncodeToString(b)
		d := struct {
			grant_type string
		}{
			grant_type:"client_credentials",
		}
		body,e:=doPost(url,authorization,d)
		if e!=nil {
			return nil, e
		}
		result := &PayPalAccess{}
		e1 := json.Unmarshal(body, result)
		if e1!=nil {
			return nil, e1
		}
		m.Access = result
		return result,e
	}
	return m.Access,nil
}

//创建订单
func (m *PayPalManager)PayPalOrderCreate(purchaseUnits []PurchaseUnits) (*OrderResp,error) {
	url:= m.BaseURl+"/v2/checkout/orders/"
	d:= OrderReq{
		Intent:INTENT_CAPTURE,
		PurchaseUnits:purchaseUnits,
	}
	body,e:=doPost(url,m.GetAuthorization(),d)
	if e!=nil {
		return nil, e
	}
	result := &OrderResp{}
	e1 := json.Unmarshal(body, result)
	if e1!=nil {
		return nil, e1
	}
	switch result.Status {
		case ORDER_RESPONSE_STATUS_CREATED://创建订单
		case ORDER_RESPONSE_STATUS_SAVED:
		case ORDER_RESPONSE_STATUS_APPROVED:
		case ORDER_RESPONSE_STATUS_VOIDED:
		case ORDER_RESPONSE_STATUS_COMPLETED:
			return  result,nil
		default:
			return  nil,nil
	}
	return  nil,nil
}

//
func (m *PayPalManager)PayPalOrderCapture(orderID string) (*CaptureOrderResp,error) {
	url:= m.BaseURl+"/v2/checkout/orders/"+orderID+"/capture"
	d := struct {}{}
	body,e:=doPost(url,m.GetAuthorization(), d)
	if e!=nil {
		return nil, e
	}
	result := &CaptureOrderResp{}
	e1 := json.Unmarshal(body, result)
	if e1!=nil {
		return nil, e1
	}
	switch result.Status {
	case ORDER_RESPONSE_STATUS_CREATED://创建订单
	case ORDER_RESPONSE_STATUS_SAVED:
	case ORDER_RESPONSE_STATUS_APPROVED:
	case ORDER_RESPONSE_STATUS_VOIDED:
	case ORDER_RESPONSE_STATUS_COMPLETED:
		return  result,nil
	default:
		return  nil,nil
	}
	return nil, nil
}


func doPost(url string,authorization string,data interface{}) ([]byte,error) {
	var e error
	bodyStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyStr))
	if err != nil {
		e = fmt.Errorf(err.Error())
		return nil,e
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err1 := client.Do(req)
	if err1 != nil {
		e = fmt.Errorf(err1.Error())
		return nil,e
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2!=nil {
		e = fmt.Errorf(err2.Error())
		return nil,e
	}
	return body,e
}


