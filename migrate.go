package main

import (
	sqlite "code.google.com/p/go-sqlite/go1/sqlite3"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var( 
	GET_CHAPTER_REG = regexp.MustCompile(`^\d{0,}[^\.]`)
	GET_VERSE_REG = regexp.MustCompile(`[^\.]\d{0,}$`)
	
	SELECT_VERSES = "SELECT * FROM verses order by id;"
	INSERT_VERSES = "INSERT INTO verses VALUES($id, $chapter, $verse, $book, $unformatted);"
	VERSES_COUNT = "SELECT COUNT(*) as count FROM verses;"
)


func MigrateVerses(dbConn *sqlite.Conn, newDb *sqlite.Conn){
	rowData := make(sqlite.RowMap)
	for s, err := dbConn.Query(SELECT_VERSES); err == nil; err = s.Next(){
		s.Scan(rowData)
		c, v := ExpandVerse(fmt.Sprint(rowData["verse"].(float64)))
		v_int, _ := strconv.ParseInt(FormatVerse(v), 10, 64)
		c_int, _ := strconv.ParseInt(c, 10, 32)
		perc := (float64(rowData["id"].(int64)) /
			float64(versesCount(dbConn))) * 100.0
		fmt.Println("At id:", rowData["id"], " =>", 
			fmt.Sprintf("%.2f", perc), "%")
		// Create a NamedArgs
		args := sqlite.NamedArgs{
			"$id" : rowData["id"],
			"$chapter" : c_int,
			"$verse" : v_int,
			"$book" : rowData["book"],
			"$unformatted" : rowData["unformatted"]}
		InsertIntoVerse(newDb, args)
	}
}

func InsertIntoVerse(c *sqlite.Conn, args sqlite.NamedArgs){
	c.Exec(INSERT_VERSES, args)
}

func ExpandVerse(val string) (chapter string, verse string) {
	chapter = GET_CHAPTER_REG.FindString(val)
	verse = GET_VERSE_REG.FindString(val)
	return
}

func FormatVerse(verse string) string {
	// If the verse is string length is less than three append zeros
	diff := 3 - len(verse)
	zeros := strings.Repeat("0", diff)
	return strings.Join([]string{verse, zeros}, "")
}

func versesCount(conn *sqlite.Conn) int64 {
	stmt, _ := conn.Query(VERSES_COUNT)
	data := make(sqlite.RowMap)
	stmt.Scan(data)
	return data["count"].(int64)
}

func main(){
	dbConn, _ := sqlite.Open("amp.sqlite3")
	newDb, _ := sqlite.Open("amp_bible.sqlite3")
	MigrateVerses(dbConn, newDb)
}
