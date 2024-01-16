package examples

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/hujinrun-github/go-press-test/constant"
	model "github.com/hujinrun-github/go-press-test/model/interface"
	tool_functions "github.com/hujinrun-github/tool_functions/go/string"
	"github.com/pkg/errors"
)

type RpcModel struct {
	model.BaseModel
}

type JsonConfig struct {
	Source     constant.SOURCE_TYPE // data source; file - from file, origin - the config itself is the json string
	ReadMethod constant.READ_METHOD
	Payload    string // if Source=constant.ST_FILE, it is the file name, else if Source=ST_ORIGIN, it is the string itself
}

// input - JsonConfig
// output - []string
func (j *RpcModel) LoadData(c any) (any, constant.ErrorCode, error) {
	rc, ok := c.(JsonConfig)
	if !ok {
		return nil, constant.ERR_TYPE_CONVERT, errors.New(fmt.Sprintf("type:%+v convert to %+v failed", reflect.TypeOf(c), reflect.TypeOf(JsonConfig{})))
	}

	switch rc.Source {
	case constant.ST_FILE:
		{
			switch rc.ReadMethod {
			case constant.RM_ALL:
				file, err := os.Open(rc.Payload)
				if err != nil {
					return nil, constant.ERR_OPEN_FILE, errors.New(fmt.Sprintf("open file:%+v failed", rc.Payload))
				}

				defer file.Close()

				content, err := io.ReadAll(file)
				if err != nil {
					return nil, constant.ERR_READ_FILE, errors.New(fmt.Sprintf("read file:%+v failed", rc.Payload))
				}
				return tool_functions.Bytes2Str(content), constant.ERR_CODE_SUCCESS, nil
			case constant.RM_LINE:
				file, err := os.Open(rc.Payload)
				if err != nil {
					return nil, constant.ERR_OPEN_FILE, errors.New(fmt.Sprintf("open file:%+v failed", rc.Payload))
				}

				defer file.Close()

				r := bufio.NewReader(file)
				result := []string{}
				for {
					line, err := r.ReadString('\n')
					if err != nil && err != io.EOF {
						return nil, constant.ERR_READ_FILE_LINE, errors.New(fmt.Sprintf("read one line failed, err:%+v", err))
					}
					if err == io.EOF {
						break
					}

					line = strings.TrimSpace(line)
					result = append(result, line)
				}
				return result, constant.ERR_CODE_SUCCESS, nil
			}
		}
	case constant.ST_ORIGIN:
		return []string{rc.Payload}, constant.ERR_CODE_SUCCESS, nil
	default:
		return nil, constant.ERR_INVALID_PARAM, errors.New(fmt.Sprintf("invalid source type:%+v", rc.Source))
	}

	return nil, constant.ERR_UNKNOW, errors.New("unknow error")
}

func (r *RpcModel) Request(rsp chan model.IRequestResult, wg *sync.WaitGroup) constant.ErrorCode {

	defer wg.Done()
	return constant.ERR_CODE_SUCCESS
}
