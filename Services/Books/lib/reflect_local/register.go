package reflect_local

import (
	"DigitalPayment/Services/Books/Models"
	"reflect"
)

var registeredMessages map[string]registeredRequest
var registeredRequestType map[string]string
var requestInterfaceType = reflect.TypeOf((*Models.Logic)(nil)).Elem()

type registeredRequest struct {
	request reflect.Type
}

func init() {
	registeredRequestType = make(map[string]string, 10)
	registeredMessages = make(map[string]registeredRequest, 10)
}

func Register(request string, requestType Models.Logic) {

	obj := registeredRequest{}

	tmp := reflect.TypeOf(requestType)

	if tmp.Kind() != reflect.Ptr || tmp.Elem().Kind() != reflect.Struct || !tmp.Implements(requestInterfaceType) {
		panic("Unable to register Request argument. Must be pointer to struct which implements Request interface.")
	}

	obj.request = tmp.Elem()

	registeredMessages[request] = obj
	registeredRequestType[tmp.String()] = request
}
func (obj *registeredRequest) CreateRequest() Models.Logic {
	return reflect.New(obj.request).Interface().(Models.Logic)
}
