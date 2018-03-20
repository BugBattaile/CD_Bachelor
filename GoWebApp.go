package main

import (
	"html/template"
	"io"
	"net/http"

	"cloud.google.com/go/compute/metadata"
)

type Instance struct {
	Id         string
	Name       string
	Version    string
	Hostname   string
	Zone       string
	Project    string
	InternalIP string
	ExternalIP string
	LBRequest  string
	ClientIP   string
	Error      string
}

const (
	index = `<!DOCTYPE html>
<html>
<head>
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.0/css/materialize.min.css">
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
<title>GoWebApp</title>
</head>
<body>
<h2>Introducing GoWebApp</h2>
<div class="card blue">
	<div class="card-content white-text">
		<div class="card-title">
			<h3>This is a simple Webapp written in Go</h3>
			<p>It has been deployed throw a Continuous Delivery Pipeline in a Kubernetes Cluster.</p>		
		</div>
		<table class="bordered">
  			<tbody>
				<tr>
	  			<td>Projekt</td>
				  <td>{{.Project}}</td>
				  </tr>
				  <tr>
				  <td>InternalIP</td>
				  <td>{{.InternalIP}}</td>
				  </tr>
				  <tr> 
				  <td>ExternalIP</td>
	  				<td>{{.ExternalIP}}</td>
				</tr>
			</tbody>
		</table>
	</div>
</div>
</body>
</html>`
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("out").Parse(index))
	i := &Instance{}
	fmt.Fprintf(w, index)

	tpl.Execute(w, i)
	fmt.Fprintf(w, index)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":80", nil)
}

// HealthCheckHandler as a simple HTTP Health check
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
	//	w.Header().Set("Content-Type", "application/json")
}

type assigner struct {
	err error
}

func (a *assigner) assign(getVal func() (string, error)) string {
	if a.err != nil {
		return ""
	}
	s, err := getVal()
	if err != nil {
		a.err = err
	}
	return s
}

func newInstance() *Instance {
	var i = new(Instance)
	if !metadata.OnGCE() {
		i.Error = "Not running on GCE"
		return i
	}

	a := &assigner{}
	i.Id = a.assign(metadata.InstanceID)
	i.Zone = a.assign(metadata.Zone)
	i.Name = a.assign(metadata.InstanceName)
	i.Hostname = a.assign(metadata.Hostname)
	i.Project = a.assign(metadata.ProjectID)
	i.InternalIP = a.assign(metadata.InternalIP)
	i.ExternalIP = a.assign(metadata.ExternalIP)

	if a.err != nil {
		i.Error = a.err.Error()
	}
	return i
}
