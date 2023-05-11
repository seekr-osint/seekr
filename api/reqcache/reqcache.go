package reqcache

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

// var (
//
//	EmptyResp = http.Response{}
//
// )
var (
	GlobalCacheAge = 5
	ErrNotCached   = errors.New("not cached")
)

var GlobalCache Cache = Cache{}

type Cache map[string]CachedReq
type CachedReq struct {
	ReqResp
}
type ReqResp struct {
	Rid     string
	Pending bool
	Resp    http.Response
	Time    time.Time
	Body    []byte
	Header  http.Header
}

type Req http.Request

func ReqcacheReq(method string, url string, body io.Reader) (http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	return *req, err
}

func (req Req) NewReqResp(resp http.Response) ReqResp {
	reqresp := ReqResp{
		Rid:  req.GetRID(),
		Resp: resp,
		Time: time.Now(),
	}
	if resp.Body == nil {
		log.Printf("pending req: %s\n", reqresp)
		reqresp.Pending = true
	}
	return reqresp
}

func (reqresp ReqResp) ValidAge(CacheAge int) bool {
	return time.Since(reqresp.Time) < time.Duration(CacheAge)*time.Minute
}

func (req Req) GetRID() string {
	return fmt.Sprintf("%s-%s-%s", req.Method, req.URL, req.Host)
}

func (reqresp ReqResp) String() string {
	return reqresp.Rid
}

func Reqcache(req http.Request) (ReqResp, error) {
	reqresp := Req(req).NewReqResp(http.Response{})
	reqresp, err := reqresp.ReadCache()
	if err == ErrNotCached {
		resp, err := http.DefaultClient.Do(&req)
		if err != nil {
			return ReqResp{}, err
		}
		fmt.Println(len(resp.Header))
		reqresp = Req(req).NewReqResp(*resp)
		log.Printf("Request (rid): %s\n", reqresp)

		err = reqresp.WriteCache()
		if err != nil {
			return ReqResp{}, err
		}
		reqresp, err = reqresp.ReadCache()
		if err != nil {
			return ReqResp{}, err
		}
	}
	//defer resp.Body.Close()
	return reqresp, nil
}
func (reqresp ReqResp) ReadCache() (ReqResp, error) {
	if cachedReq, ok := GlobalCache[reqresp.Rid]; ok {
		if cachedReq.ValidAge(GlobalCacheAge) {
			return cachedReq.ReqResp, nil
		} else {
			delete(GlobalCache, reqresp.Rid)
			return reqresp, ErrNotCached
		}
	}
	return reqresp, ErrNotCached
}

func (reqresp ReqResp) WriteCache() error {
	var err error
	if reqresp.Resp.Body != nil {
		reqresp.Body, err = ioutil.ReadAll(reqresp.Resp.Body)
		if err != nil {
			return err
		}
		cloneHeaders := make(http.Header, len(reqresp.Resp.Header))
		for k, vv := range reqresp.Resp.Header {
			for _, v := range vv {
				cloneHeaders.Add(k, v)
			}
		}
		reqresp.Header = cloneHeaders
		fmt.Printf("length: %d,%d", len(reqresp.Header), len(reqresp.Resp.Header))
		reqresp.Resp.Body.Close()
		if !reqresp.Pending {
			GlobalCache[reqresp.Rid] = CachedReq{reqresp}
		} else {
			fmt.Printf("err still pending")
		}
	}
	return nil
}

func (reqresp ReqResp) GetCacheFile() (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println("Error getting cache directory:", err)
		return cacheDir, err
	}
	return path.Join(cacheDir, "seekr-reqcache", reqresp.Rid), err
}
