package guestbook

import (
	"html/template"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// [START greeting_struct]
type Task struct {
	Division  string
	SuperGoal string
	Goal string
	Target string
	Tasks []string
	Parameters []string
	Quarter []int
	Responsible string
	Partners []string
	Notes []string
	Date    time.Time
}

var tpl *template.Template

// [END greeting_struct]

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	http.HandleFunc("/", root)
	//http.HandleFunc("/bns-office-outlook-manifest.xml", manifest)
	http.HandleFunc("/sign", sign)
	// http.HandleFunc("/favicon.ico", favicon)
	// http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
}
func favicon(w http.ResponseWriter, r *http.Request)  {
	http.ServeFile(w,r,"assets/icons/favicon.ico")
}
func manifest(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r, "bns-office-outlook-manifest.xml")
}

// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c context.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
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
		Division: "ראש המועצה",
		SuperGoal: "",
		Goal: "",
		Target: "",
		Tasks: []string{"", ""},
		Parameters: []string{"", ""},
		Quarter: []int{1,2,3,4},
		Responsible: "",
		Partners: []string{"", ""},
		Notes: []string{"", ""},
		Date:    time.Now(),
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
