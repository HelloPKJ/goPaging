package libs

import "fmt"

//开始分页
//dataSource, 数据源
//eachPageRecords, 每页记录数
func StartPaging(dataSource *[]interface{}, eachPageRecords int) *[]interface{} {
	totalRecords := len(*dataSource)
	resultSet := []interface{}{}
	if tpCount, surplusRecords := getTotalPagesAndSurplusRecords(totalRecords, eachPageRecords); tpCount != 0 {
		if tpCount == 1 {
			resultSet = append(
				resultSet,
				map[string]interface{}{
					"pageID": 1,
					"data":   (*dataSource)[:],
				},
			)
		}
		if tpCount > 1 {
			startIndex := 0
			endIndex := eachPageRecords
			for i := 0; i < tpCount; i++ {
				if i == 0 {
					startIndex = 0
					endIndex = eachPageRecords
				} else {
					if i == tpCount-1 && surplusRecords > 0 {
						startIndex += eachPageRecords
						endIndex += surplusRecords
					} else {
						startIndex += eachPageRecords
						endIndex += eachPageRecords
					}
				}
				resultSet = append(
					resultSet,
					map[string]interface{}{
						"pageID": i + 1,
						"data":   (*dataSource)[startIndex:endIndex],
					},
				)
			}
		}
		return &resultSet
	}
	return nil
}

//求总页数和剩余记录数
//totalRecords, 总记录数
//eachPageRecords, 每页记录数
//tc, 总页数
//sr, 剩余记录数
func getTotalPagesAndSurplusRecords(totalRecords int, eachPageRecords int) (tc int, sr int) {

	if eachPageRecords > 0 {
		if totalRecords > eachPageRecords {
			if totalRecords%eachPageRecords == 0 {
				tc = totalRecords / eachPageRecords
				sr = 0
				return tc, sr
			}

			if totalRecords%eachPageRecords > 0 {
				tc = totalRecords/eachPageRecords + 1
				sr = totalRecords % eachPageRecords
				return tc, sr
			}
		} else {
			fmt.Println("我在这里")
			tc = 1
			sr = totalRecords
			return tc, sr
		}
	}
	return 0, 0
}
