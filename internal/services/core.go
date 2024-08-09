package services

// TODO: Finish later on

/* import (
	"context"
	"fmt"
	"reflect"
)

type BindingMap struct {
	factoryFunc reflect.Value
	lifecycle
}

var services = make(map[reflect.Type]BindingMap)

func addService(life lifecycle, factoryFunc interface{}) (err error) {
	factoryFuncType := reflect.TypeOf(factoryFunc)
	if factoryFuncType.Kind() == reflect.Func && factoryFuncType.NumOut() == 1 {
		services[factoryFuncType.Out(0)] = BindingMap{
			factoryFunc: reflect.ValueOf(factoryFunc),
			lifecycle:   life,
		}
	} else {
		err = fmt.Errorf("type cannot be used as service: %v", factoryFuncType)
	}

	return
}

var (
	contextReference     = (*context.Context)(nil)
	contextReferenceType = reflect.TypeOf(contextReference).Elem()
)

func resolveServiceFromValue(c context.Context, val reflect.Value) (err error) {
	serviceType := val.Elem().Type()
} */
