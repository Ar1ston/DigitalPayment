package register_requests

import (
	"DigitalPayment/lib/Models"
	"reflect"
)

var registeredRequests map[string]registeredRequest
var requestInterfaceType = reflect.TypeOf((*Models.Logic)(nil)).Elem()

type registeredRequest struct {
	request reflect.Type
}

func init() {
	registeredRequests = make(map[string]registeredRequest, 10)
}

func Register(request string, requestType Models.Logic) {

	obj := registeredRequest{}

	tmp := reflect.TypeOf(requestType)

	if tmp.Kind() != reflect.Ptr || tmp.Elem().Kind() != reflect.Struct || !tmp.Implements(requestInterfaceType) {
		panic("Not found requests")
	}
	obj.request = tmp.Elem()

	registeredRequests[request] = obj
}
func (obj *registeredRequest) CreateRequest() Models.Logic {
	return reflect.New(obj.request).Interface().(Models.Logic)
}
