package sqli

import "strings"

// col connect
func sqlColConnect(col []string) string {

	if len(col) == 1 && col[0] == "*" {
		return "*"
	}
	for k, s := range col {
		col[k] = "`" + s + "`"
	}
	return strings.Join(col, ",")
}

// condition connect
func sqlConditionConnect(condCol []string, pad string) string {

	conStr := make([]string, 0, len(condCol))
	if pad == "" {
		pad = " AND "
	}
	for _, val := range condCol {
		conStr = append(conStr, val +" = ?")
	}
	return "(" + strings.Join(conStr, pad) + ")"

}

// update connect
func sqlUpdateConnect(upCol []string) string {

	conStr := make([]string, 0, len(upCol))
	for _, val := range upCol {
		conStr = append(conStr, "`" + val +"` = ?")
	}
	return strings.Join(conStr, " , ")

}

// insert connect
func sqlInsertConnect(insCol []string) (string,string){

	conStrCol := make([]string,0,len(insCol))
	conStrVal := make([]string,0,len(insCol))

	for _,val := range insCol{
		conStrCol = append(conStrCol, "`"+ val +"`")
		conStrVal = append(conStrVal," ? ")
	}

	return "(" + strings.Join(conStrCol," , ") + ")","(" + strings.Join(conStrVal,",") +")"

}

// delete connect
func sqlDeleteConnect(delCol []string,pad string) string{

	conStr := make([]string, 0, len(delCol))
	if pad == "" {
		pad = " AND "
	}
	for _, val := range delCol {
		conStr = append(conStr, val+" = ?")
	}
	return "(" + strings.Join(conStr, pad) + ")"
}