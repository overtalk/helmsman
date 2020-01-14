package consul_test

//import (
//	"testing"
//
//	. "helmsman/source/consul"
//)
//
//type Config struct {
//	Addr string `consul:"addr"`
//	Path string `consul:"data_path"`
//}
//
//func TestConsulConfig_ParseConfig(t *testing.T) {
//	updateChan := make(chan interface{}, 100)
//	errChan := make(chan error, 100)
//
//	cc := ConsulConfig{
//		Prefix:     "/",
//		UpdateChan: updateChan,
//		ErrChan:    errChan,
//	}
//
//	go cc.ParseConfig(&Config{})
//	defer cc.Close()
//
//	select {
//	case c := <-updateChan:
//		t.Log(c)
//	case err := <-errChan:
//		t.Error(err)
//	}
//}
