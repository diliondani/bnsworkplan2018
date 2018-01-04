package guestbook

import (
	"encoding/json"
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

// [START User_struct]
type User struct {
	Name       string `datastore:"Name"`
	Email      string `datastore:"Email"`
	ID        string `datastore:"ID"`
	Department string `datastore:"Department"`
	Division   string `datastore:"Division"`
}

// [START Uuid_struct]
type Uuid struct {
	Email string `datastore:"Email"`
	ID   string `datastore:"ID"`
}

// [START greeting_struct]
type Task struct {
	Number      int    `json:"number" | datastore:"number"`
	Division    string `json:"division"`
	Department  string `json:"department"`
	Type        string `json: "type"`
	SuperGoal   string `json: "supergoal"`
	Goal        string `json: "Goal"`
	Target      string `json: "target"`
	Mission     string `json: "mission"`
	Parameter   string `json: "parameter"`
	Q1          bool   `json: "q1"`
	Q2          bool   `json: "q2"`
	Q3          bool   `json: "q3"`
	Q4          bool   `json: "q4"`
	Responsible string `json: "responsible"`
	Partners    string `json: "partners"`
	Notes       string `json: "notes"`
}

var tpl *template.Template

// [END greeting_struct]

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	http.HandleFunc("/", root)
	//http.HandleFunc("/bns-office-outlook-manifest.xml", manifest)

	http.HandleFunc("/dialog", dialog)
	http.HandleFunc("/gettasks", getTasks)
	http.HandleFunc("/setcookie", setCookie)
	//TODO: delete this handlers on production
	http.HandleFunc("/sign", sign)
	http.HandleFunc("/read", read)
	http.HandleFunc("/createuser", createUser)
	// http.HandleFunc("/favicon.ico", favicon)
	// http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	user := User{
		Name:       "Dani Dilion",
		Email:      "dilion.dani@gmail.com",
		ID:         "cb311cc0-5748-4ef5-8513-94dc1fc0dde1",
		Department: "mayor_workplan",
		Division:   "ראש המועצה",
	}
	u := Uuid{
		Email: user.Email,
		ID:    user.ID,
	}
	k := datastore.NewKey(c, "Users", user.ID, 0, nil)
	_, error := datastore.Put(c, k, &user)
	//log.Debugf(c, "User was added: %v", key)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	k = datastore.NewKey(c, "Uuids", u.ID, 0, nil)
	_, err := datastore.Put(c, k, &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getUser(c context.Context, email string) (User, error) {
	var users []User
	q := datastore.NewQuery("Users").Filter("Email =", email).Limit(1)
	_, err := q.GetAll(c, &users)
	// user is not authorized
	if err != nil {
		return User{}, err
	}
	return users[0], nil
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		// get cookie
		ck, err := r.Cookie("session")
		// cookie is not present
		if err != nil {
			// check if user email is on db
			if email != "" {
				// var users []User
				// q := datastore.NewQuery("Users").Ancestor(getParentKey(c, "Users", "default_users")).Filter("email =", email).Limit(1)
				// _, err := q.GetAll(c, &users)
				// // user is not authorized
				// if err != nil {
				// 	http.Error(w, err.Error(), http.StatusInternalServerError)
				// 	log.Errorf(c, "DS Error 1: %v", err)
				// 	return
				// }
				user, err := getUser(c, email)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					log.Errorf(c, "DS Error 1: %v", err)
					return
				}
				// user is authorized but there is no cookie
				ck = &http.Cookie {
					Name:   "session",
					Value:  user.ID,
					MaxAge: 60 * 60 * 24 * 7,
					//TODO: uncomment on production
					Secure:   true,
					HttpOnly: true,
					Path: "/",
				}
				//add user uuid to db
				u := Uuid {
					Email: email,
					ID:  user.ID,
				}
				k := datastore.NewKey(c, "Uuids", u.ID, 0, nil)
				_, error := datastore.Put(c, k, &u)
				if error != nil {
					http.Error(w, error.Error(), http.StatusInternalServerError)
					log.Errorf(c, "Ds Error 2: v%", error.Error())
					return
				}

				// sID, _ := uuid.NewRandom()
				// ck = &http.Cookie{
				// 	Name:     "session",
				// 	Value:    sID.String(),
				// 	MaxAge:   60 * 60 * 24 * 7,
				// 	Secure:   true,
				// 	HttpOnly: true,
				// 	//TODO: add user uuid to db
				// }
				http.SetCookie(w, ck)
				//http.Redirect(w, r, "/", http.StatusAccepted)
				io.WriteString(w, "success")
				return
			}
			// email field is empty
			http.Error(w, "No user email", http.StatusForbidden)
			return
		}
		// cookie is present
		io.WriteString(w, "Cookie is present")
	}
	// not a post request
	// Do nothing
}

func getUserByID(c context.Context, id string) (User, error) {
	var users []User
	q := datastore.NewQuery("Users").Filter("ID =", id).Limit(1)
	if keys, err := q.GetAll(c, &users) ; err != nil || keys == nil {
		return User{}, err
	}
	return users[0], nil
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodPost {
		// get cookie
		ck, err := r.Cookie("session")
		if err != nil {
			// if no cookie is set the user is not unauthorized to access the route
			http.Error(w, err.Error(), http.StatusForbidden)
			log.Errorf(c, "error: %v", err)
			return
		}
		// cookie is present
		log.Debugf(c, "Cookie value %v", ck.Value)
		user, err := getUserByID(c, ck.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			log.Errorf(c, "error: %v", err)
			return
		}
		//TODO: check in users db if a user exists, only then return tasks
		// body
		// bs := make([]byte, r.ContentLength)
		// r.Body.Read(bs)
		// body := string(bs)
		var tasks []*Task
		q := datastore.NewQuery("Tasks").Ancestor(getParentKey(c, "WorkPlan", user.Department)).Filter("Division =",user.Division).Order(" Number")
		//email := r.FormValue("email")
		//subject := r.FormValue("subject")
		//w.Write()
		//[START getall]
		if _, err := q.GetAll(c, &tasks); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Errorf(c, "error: %v", err)
			return
		}
		log.Debugf(c, "Tasks returned %v", len(tasks))
		// [END getall]

		j, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Errorf(c, "error: %v", err)
			return
		}

		//log.Debugf(c, "This is request content %v", j)

		io.WriteString(w, string(j))
	}
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

		key := datastore.NewIncompleteKey(c, "Tasks", getParentKey(c, "WorkPlan", "mayor_workplan"))
		_, err := datastore.Put(c, key, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		io.WriteString(w, "Succes")
	}

}

// TODO: add a task to outlook tasks

// getParentKey returns the key used for all task entries.
func getParentKey(c context.Context, kind string, id string) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, kind, id, 0, nil)
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
	//q := datastore.NewQuery("Greeting").Ancestor(getParentKey(c)).Order("-Date").Limit(10)
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
	key := datastore.NewIncompleteKey(c, "Tasks", getParentKey(c, "Tasks", "mayor_workplan"))
	_, err := datastore.Put(c, key, &t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
	// [END if_user]
}

// [END func_sign]
