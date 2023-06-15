//go:build networking

package api

//import (
//	"sync"
//	"testing"
//	//"os"
//)

//func TestUsernames(t *testing.T) {
//	wg := &sync.WaitGroup{}
//	username := RandomString(16)
//	for i := 0; i < len(DefaultServices); i++ { // loop over all services
//		wg.Add(1)
//		go func(i int) {
//			// Do something
//			service := DefaultServices[i]                  // current service
//			if service.UserExistsFunc(service, username) { // if service exisits
//				t.Errorf("This no work %s", service.Name)
//			}
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}
