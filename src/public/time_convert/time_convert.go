package timeconvert

import(
	"strconv"
)

func TimeConvert(timestamp string) (int64,error){
	mtime, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil{
		return 0,err
	}
	return mtime,nil
}