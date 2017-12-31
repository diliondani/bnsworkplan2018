package guestbook

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"google.golang.org/appengine/log"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// [START greeting_struct]
type Task struct {
	Number      int
	Division    string
	Department  string
	Type        string
	SuperGoal   string
	Goal        string
	Target      string
	Mission     string
	Parameter   string
	Q1          bool
	Q2          bool
	Q3          bool
	Q4          bool
	Responsible string
	Partners    string
	Notes       string
}

var tpl *template.Template

// [END greeting_struct]

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	http.HandleFunc("/", root)
	//http.HandleFunc("/bns-office-outlook-manifest.xml", manifest)
	http.HandleFunc("/sign", sign)
	http.HandleFunc("/read", read)
	http.HandleFunc("/dialog", dialog)
	// http.HandleFunc("/favicon.ico", favicon)
	// http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
}
func dialog(w http.ResponseWriter, r *http.Request) {
	if err := tpl.ExecuteTemplate(w, "dialog.gohtml", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func read(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	xlsx, err := excelize.OpenFile("bnsworkplan.xlsx")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t := Task{
		Division:   "ראש המועצה",
		Department: "ראש המועצה",
	}
	// Get value from cell by given worksheet name and axis.
	//cell := xlsx.GetCellValue("Mayor", "B2")
	//log.Debugf(c, "This is cell content %v", cell)
	//io.WriteString(w, cell)
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Mayor")

	for i, row := range rows {
		log.Debugf(c, "Row: %v", i)
		t.Number = i
		t.SuperGoal = row[0]
		t.Goal = row[1]
		t.Target = row[2]
		t.Mission = row[3]
		t.Parameter = row[4]
		//t.Quarter = row[5]
		t.Q1, _ = strconv.ParseBool(row[6])
		t.Q2, _ = strconv.ParseBool(row[7])
		t.Q3, _ = strconv.ParseBool(row[8])
		t.Q4, _ = strconv.ParseBool(row[9])
		t.Responsible = row[10]
		t.Partners = row[11]
		t.Notes = row[12]
		t.Type = row[13]

		key := datastore.NewIncompleteKey(c, "Task", guestbookKey(c))
		_, err := datastore.Put(c, key, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		io.WriteString(w, "Succes")
	}

}

// TODO: add a task to outlook tasks

// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c context.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "WorkPlan", "mayor_workplan", 0, nil)
}

// [START func_root]
func root(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	// Ancestor queries, as shown here, are strongly consistent with the High
	// Replication Datastore. Queries that span entity groups are eventually
	// consistent. If we omitted the .Ancestor from this query there would be
	// a slight chance that Greeting that had just been written would not
	// show up in a query.
	// [START query]
	//q := datastore.NewQuery("Greeting").Ancestor(guestbookKey(c)).Order("-Date").Limit(10)
	// [END query]
	// [START getall]
	//greetings := make([]Greeting, 0, 10)
	// if _, err := q.GetAll(c, &greetings); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// // [END getall]
	// if err := guestbookTemplate.Execute(w, greetings); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	if err := tpl.ExecuteTemplate(w, "index.gohtml", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// [END func_root]

// [START func_sign]
func sign(w http.ResponseWriter, r *http.Request) {
	// [START new_context]
	c := appengine.NewContext(r)
	// [END new_context]
	t := Task{
		Number:     1,
		Division:   "ראש המועצה",
		Department: "",
		SuperGoal:  "",
		Goal:       "",
		Target:     "",
		Mission:    "",
		Parameter:  "",
		//Quarter:     []int8{1, 2, 3, 4},
		Responsible: "",
		Partners:    "",
		Notes:       "",
	}
	// [START if_user]
	// if u := user.Current(c); u != nil {
	// 	g.Author = u.String()
	// }
	// We set the same parent key on every Greeting entity to ensure each Greeting
	// is in the same entity group. Queries across the single entity group
	// will be consistent. However, the write rate to a single entity group
	// should be limited to ~1/second.
	key := datastore.NewIncompleteKey(c, "Task", guestbookKey(c))
	_, err := datastore.Put(c, key, &t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
	// [END if_user]
}

// [END func_sign]
