package SQLRead

import (
  "database/sql"
  "strings"
  "errors"
  "strconv"
)

var NumCols[100] string

func CreateResultSet(result *sql.Rows) (map[string][]string){
    ResultSet := make(map[string][]string)
    columns,_ := result.Columns()
    recordcount := 0
    for result.Next() {
        err := generic(result)
        if err != nil {
            panic(err)
            break
        }
        i := 0
        for _, column := range columns {
            if len(ResultSet[strings.ToLower(column)]) == 0 {
                ResultSet[strings.ToLower(column)] = []string{}
            }
            ResultSet[strings.ToLower(column)] = append(ResultSet[strings.ToLower(column)],NumCols[i])
            i++
        }
        recordcount = recordcount + 1
    }

    ResultSet["RecordCount"] = append(ResultSet["RecordCount"],strconv.Itoa(recordcount))

    return ResultSet
}

func generic(row *sql.Rows) (error){
    col,_ := row.Columns()
    switch len(col){
    case 1:
        return row.Scan(col1())
    case 2:
        return row.Scan(col2())
    case 3:
        return row.Scan(col3())
    case 4:
        return row.Scan(col4())
    case 5:
        return row.Scan(col5())
    case 6:
        return row.Scan(col6())
    case 7:
        return row.Scan(col7())
    case 8:
        return row.Scan(col8())
    case 9:
        return row.Scan(col9())
    case 10:
        return row.Scan(col10())
    case 11:
        return row.Scan(col11())
    case 12:
        return row.Scan(col12())
    case 13:
        return row.Scan(col13())
    case 14:
        return row.Scan(col14())
    case 15:
        return row.Scan(col15())
    case 16:
        return row.Scan(col16())
    case 17:
        return row.Scan(col17())
    case 18:
        return row.Scan(col18())
    case 19:
        return row.Scan(col19())
    case 20:
        return row.Scan(col20())
    case 21:
        return row.Scan(col21())
    case 22:
        return row.Scan(col22())
    case 23:
        return row.Scan(col23())
    case 24:
        return row.Scan(col24())
    case 25:
        return row.Scan(col25()) 
    default:
        return errors.New("More than 25 columns not allowed to fetch")
    }   
}

func col1() (*string) {return &NumCols[0]}
func col2() (*string,*string) {return &NumCols[0],&NumCols[1]}
func col3() (*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2]}
func col4() (*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3]}
func col5() (*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4]}
func col6() (*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5]}
func col7() (*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6]}
func col8() (*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7]}
func col9() (*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8]}
func col10() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9]}
func col11() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10]}
func col12() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10],&NumCols[11]}
func col13() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10],&NumCols[11],&NumCols[12]}
func col14() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10],&NumCols[11],&NumCols[12],&NumCols[13]}
func col15() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10],&NumCols[11],&NumCols[12],&NumCols[13],&NumCols[14]}
func col16() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10],&NumCols[11],&NumCols[12],&NumCols[13],&NumCols[14],&NumCols[15]}
func col17() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10],&NumCols[11],&NumCols[12],&NumCols[13],&NumCols[14],&NumCols[15],&NumCols[16]}
func col18() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10],&NumCols[11],&NumCols[12],&NumCols[13],&NumCols[14],&NumCols[15],&NumCols[16],&NumCols[17]}
func col19() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10],&NumCols[11],&NumCols[12],&NumCols[13],&NumCols[14],&NumCols[15],&NumCols[16],&NumCols[17],&NumCols[18]}
func col20() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10],&NumCols[11],&NumCols[12],&NumCols[13],&NumCols[14],&NumCols[15],&NumCols[16],&NumCols[17],&NumCols[18],&NumCols[19]}
func col21() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10],&NumCols[11],&NumCols[12],&NumCols[13],&NumCols[14],&NumCols[15],&NumCols[16],&NumCols[17],&NumCols[18],&NumCols[19],&NumCols[20]}
func col22() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10],&NumCols[11],&NumCols[12],&NumCols[13],&NumCols[14],&NumCols[15],&NumCols[16],&NumCols[17],&NumCols[18],&NumCols[19],&NumCols[20],&NumCols[21]}
func col23() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10],&NumCols[11],&NumCols[12],&NumCols[13],&NumCols[14],&NumCols[15],&NumCols[16],&NumCols[17],&NumCols[18],&NumCols[19],&NumCols[20],&NumCols[21],&NumCols[22]}
func col24() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10],&NumCols[11],&NumCols[12],&NumCols[13],&NumCols[14],&NumCols[15],&NumCols[16],&NumCols[17],&NumCols[18],&NumCols[19],&NumCols[20],&NumCols[21],&NumCols[22],&NumCols[23]}
func col25() (*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string,*string) {return &NumCols[0],&NumCols[1],&NumCols[2],&NumCols[3],&NumCols[4],&NumCols[5],&NumCols[6],&NumCols[7],&NumCols[8],&NumCols[9],&NumCols[10],&NumCols[11],&NumCols[12],&NumCols[13],&NumCols[14],&NumCols[15],&NumCols[16],&NumCols[17],&NumCols[18],&NumCols[19],&NumCols[20],&NumCols[21],&NumCols[22],&NumCols[23],&NumCols[24]}