package framework

import (
	"testing"
	"fmt"
	"time"
	"errors"
	"io/ioutil"
	"net/url"
	"github.com/jglobant/yale/model"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/gambol99/go-marathon"
)
var AppMock *marathon.Application
var ErrorMock error
var CreateAppErrorMock error
var ListAppErrorMock error
var ScaleAppErrorMock error
var ListAppMock []string
type MarathonMockClient struct {
}

func createMarathonMockClient() marathon.Marathon {
	client := new(MarathonMockClient)
	return client
}
func (client *MarathonMockClient) ListApplications(url.Values) ([]string, error){return ListAppMock,ListAppErrorMock}
func (client *MarathonMockClient) ApplicationVersions(name string) (*marathon.ApplicationVersions, error){return nil,nil}
func (client *MarathonMockClient) HasApplicationVersion(name, version string) (bool, error){return false,nil}
func (client *MarathonMockClient) SetApplicationVersion(name string, version *marathon.ApplicationVersion) (*marathon.DeploymentID, error){return nil,nil}
func (client *MarathonMockClient) ApplicationOK(name string) (bool, error){return false,nil}
func (client *MarathonMockClient) CreateApplication(application *marathon.Application) (*marathon.Application, error){return nil,CreateAppErrorMock}
func (client *MarathonMockClient) DeleteApplication(name string) (*marathon.DeploymentID, error){return nil,nil}
func (client *MarathonMockClient) UpdateApplication(application *marathon.Application) (*marathon.DeploymentID, error){return nil,nil}
func (client *MarathonMockClient) ApplicationDeployments(name string) ([]*marathon.DeploymentID, error){return nil,nil}
func (client *MarathonMockClient) ScaleApplicationInstances(name string, instances int, force bool) (*marathon.DeploymentID, error){
	deploymentId := &marathon.DeploymentID{}
	deploymentId.Version = "version-1"
	return deploymentId, ScaleAppErrorMock
}
func (client *MarathonMockClient) RestartApplication(name string, force bool) (*marathon.DeploymentID, error){return nil,nil}
func (client *MarathonMockClient) Applications(url.Values) (*marathon.Applications, error){return nil,nil}
func (client *MarathonMockClient) WaitOnApplication(name string, timeout time.Duration) error {return nil}
func (client *MarathonMockClient) Tasks(application string) (*marathon.Tasks, error){return nil,nil}
func (client *MarathonMockClient) AllTasks(opts *marathon.AllTasksOpts) (*marathon.Tasks, error){return nil,nil}
func (client *MarathonMockClient) TaskEndpoints(name string, port int, healthCheck bool) ([]string, error){return nil,nil}
func (client *MarathonMockClient) KillApplicationTasks(applicationID string, opts *marathon.KillApplicationTasksOpts) (*marathon.Tasks, error){return nil,nil}
func (client *MarathonMockClient) KillTask(taskID string, opts *marathon.KillTaskOpts) (*marathon.Task, error){return nil,nil}
func (client *MarathonMockClient) KillTasks(taskIDs []string, opts *marathon.KillTaskOpts) error {return nil}
func (client *MarathonMockClient) Groups() (*marathon.Groups, error){return nil,nil}
func (client *MarathonMockClient) Group(name string) (*marathon.Group, error){return nil,nil}
func (client *MarathonMockClient) DeleteGroup(name string) (*marathon.DeploymentID, error){return nil,nil}
func (client *MarathonMockClient) UpdateGroup(id string, group *marathon.Group) (*marathon.DeploymentID, error){return nil,nil}
func (client *MarathonMockClient) HasGroup(name string) (bool, error){return false,nil}
func (client *MarathonMockClient) WaitOnGroup(name string, timeout time.Duration) error{return nil}
func (client *MarathonMockClient) Deployments() ([]*marathon.Deployment, error){return nil,nil}
func (client *MarathonMockClient) DeleteDeployment(id string, force bool) (*marathon.DeploymentID, error){return nil,nil}
func (client *MarathonMockClient) HasDeployment(id string) (bool, error){return false,nil}
func (client *MarathonMockClient) Subscriptions() (*marathon.Subscriptions, error){return nil,nil}
func (client *MarathonMockClient) AddEventsListener(channel marathon.EventsChannel, filter int) error {return nil}
func (client *MarathonMockClient) RemoveEventsListener(channel marathon.EventsChannel){}
func (client *MarathonMockClient) Unsubscribe(string) error {return nil}
func (client *MarathonMockClient) GetMarathonURL() string {return ""}
func (client *MarathonMockClient) Ping() (bool, error){return false,nil}
func (client *MarathonMockClient) Info() (*marathon.Info, error){return nil,nil}
func (client *MarathonMockClient) Leader() (string, error){return "",nil}
func (client *MarathonMockClient) AbdicateLeader() (string, error){return "", nil}
func (client *MarathonMockClient) CreateGroup(group *marathon.Group) error {return nil}
func (client *MarathonMockClient) WaitOnDeployment(id string, timeout time.Duration) error {return nil}
func (client *MarathonMockClient) Application(name string) (*marathon.Application, error) {
	return AppMock, ErrorMock
}

func TestConstructorError(t *testing.T) {
	_, err := NewMarathon("malformed url")
	assert.True(t, err != nil, "Malformed Url should throw error")
}

func TestFindServiceInformation(t *testing.T) {
	
	m := createMarathonHelper(t)
	createBasicMockApp()
	
	services, _ := m.FindServiceInformation("nginx")
	assert.Equal(t, 2, len(services), "Should have found two services")
}

func TestFindServiceInformationError(t *testing.T) {

        m := createMarathonHelper(t)
	ErrorMock = errors.New("no app found")
        _, err := m.FindServiceInformation("nginx")
        assert.NotNil(t, err, "Should throw error on recieving application")
}

func TestFindServiceInformationNil(t *testing.T) {

        m := createMarathonHelper(t)
        ErrorMock = nil
	AppMock = nil
        service, _ := m.FindServiceInformation("nginx")
        assert.NotNil(t, service, "Should return service")
	assert.Equal(t, 0, len(service), "instances should be empty")
}

func TestDeployService(t *testing.T) {
	m := createMarathonHelper(t)
	createBasicMockApp()
	
	instances, err := m.DeployService(model.ServiceConfig{})
	assert.Nil(t, err, "Deploy should work")
	assert.NotNil(t, instances, "Should return new instances")
}

func TestErrorDeployService(t *testing.T) {
	ListAppErrorMock = errors.New("ListApplications throws error")
	helper := createMarathonHelper(t)
	_, err := helper.DeployService(model.ServiceConfig{})
	assert.NotNil(t, err, "Deploy should throw error")
}

func TestDeployServiceErrorOnCreate(t *testing.T) {
        m := createMarathonHelper(t)
        createBasicMockApp()
	CreateAppErrorMock = errors.New("Error on create app")
        _, err := m.DeployService(model.ServiceConfig{})
        assert.NotNil(t, err, "Deploy should not work")
}

func TestDeployServiceErrorGetNewApp(t *testing.T) {
        m := createMarathonHelper(t)
        createBasicMockApp()
        ErrorMock = errors.New("Error on create app")
        _, err := m.DeployService(model.ServiceConfig{})
        assert.NotNil(t, err, "Deploy should not work")
}

func TestScaleService(t *testing.T) {
	helper := createMarathonHelper(t)
	ListAppMock = []string{"nginx2","nginx"}
	createBasicMockApp()
	AppMock.Tasks[0].Version = "version-0"
	AppMock.Tasks[1].Version = "version-1"
        instances, err := helper.DeployService(model.ServiceConfig{ServiceId : "nginx", Instances : 2,})
        assert.Nil(t, err, "Scale should work")
	assert.NotNil(t, instances, "Scale should return instances")
}
func TestScaleServiceError(t *testing.T) {
        helper := createMarathonHelper(t)
        ListAppMock = []string{"nginx2","nginx"}
        createBasicMockApp()
	ScaleAppErrorMock = errors.New("Error on scaling App")
        _, err := helper.DeployService(model.ServiceConfig{ServiceId : "nginx", Instances : 2,})
	assert.NotNil(t, err, "Scale should throw error")
}

func TestScaleServiceErrorGetApp(t *testing.T) {
        helper := createMarathonHelper(t)
        ListAppMock = []string{"nginx2","nginx"}
        createBasicMockApp()
	ErrorMock = errors.New("Error on getting app")
        _, err := helper.DeployService(model.ServiceConfig{ServiceId : "nginx", Instances : 2,})
        assert.NotNil(t, err, "Scale should throw error")
}

func TestDeleteService(t *testing.T) {
        ts := setup("../test/resources/marathon_delete_app_response.json")
	defer ts.Close()
	helper, error := NewMarathon(ts.URL)

        if error != nil {
                t.Errorf("Error: " + error.Error())
        }

        err := helper.DeleteService("nginx")
        assert.True(t, err == nil, "Delete should work")
}

func TestErrorDeleteService(t *testing.T) {
	ts := setupFailure()
	defer ts.Close()
        helper, error := NewMarathon(ts.URL)
        
        if error != nil {
                t.Errorf("Error: " + error.Error())
        }       
        err := helper.DeleteService("error-id")
        assert.True(t, err != nil, "Delete should throw error")
}

func setupFailure() (*httptest.Server){
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintln(w, "failed response")
        }))
	return ts
}

func setup(url string) (*httptest.Server) {
        content, _ := ioutil.ReadFile(url)
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintln(w, string(content))
        }))
	return ts
}

func TestUndeployInstance(t *testing.T) {
        m, error := NewMarathon("http://localhost:8081")

        if error != nil {
                t.Errorf("Error: " + error.Error())
        }
	i := model.Instance{}
        err := m.UndeployInstance(&i)
        assert.NotNil(t, err, "UndeployInstance should not work")
}

func createMarathonHelper(t *testing.T) *Marathon {
        m, error := NewMarathon("http://localhost:8081")
        if error != nil {
                t.Errorf("Error: " + error.Error())
        }
        client := createMarathonMockClient()
        m.SetClient(client)
	return m
}
func createBasicMockApp(){
	ErrorMock = nil
	CreateAppErrorMock = nil
	ListAppErrorMock = nil
	ScaleAppErrorMock = nil
        tasks := make([]*marathon.Task, 2)
        tasks[0] = &marathon.Task{}
        tasks[1] = &marathon.Task{}
        AppMock = marathon.NewDockerApplication()
        AppMock.Tasks = tasks
}
